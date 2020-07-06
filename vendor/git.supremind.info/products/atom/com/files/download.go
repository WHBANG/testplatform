// deprecated: use HTTPRangeReader and stream.CopyInBlocks
package files

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"git.supremind.info/products/atom/com/stream"
	"golang.org/x/sync/errgroup"
)

const (
	DownloadBlockSize          = 64 << 20
	defaultDownloadConcurrency = 8
)

type HTTPDownloader struct {
	URL string

	status *DownloadStatus
	client *http.Client
	wa     io.WriterAt
	writer io.Writer
	closer io.Closer
}

type DownloadStatus struct {
	TotalSize  int64
	DoneSize   int64
	BlockCount int

	DoneBlocks []bool
}

func ChunkedHTTPDownloader(u string, size int64, wa io.WriterAt, opts ...DownloadOption) *HTTPDownloader {
	d := &HTTPDownloader{
		URL: u,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
		wa: wa,
	}

	for _, opt := range opts {
		opt(d.client)
	}

	status := &DownloadStatus{TotalSize: size}
	blks := status.TotalSize / DownloadBlockSize
	if blks*DownloadBlockSize < status.TotalSize {
		blks++
	}
	status.BlockCount = int(blks)
	status.DoneBlocks = make([]bool, status.BlockCount)

	d.status = status
	return d
}

func SimpleHTTPDownloader(u string, size int64, w io.Writer, opts ...DownloadOption) *HTTPDownloader {
	d := &HTTPDownloader{
		URL:    u,
		status: &DownloadStatus{TotalSize: size},
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
		writer: w,
	}

	for _, opt := range opts {
		opt(d.client)
	}

	return d
}

func NewHTTPDownloader(u, path string, size int64) (*HTTPDownloader, error) {
	d := &HTTPDownloader{
		URL:    u,
		status: &DownloadStatus{TotalSize: size},
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}

	info, err := os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("stat local file failed: %w", err)
		}
	}
	if info != nil && info.IsDir() {
		return nil, fmt.Errorf("filename is a directory")
	}
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return nil, fmt.Errorf("mkdir failed: %w", err)
	}

	if size <= 0 || size > DownloadBlockSize {
		if e := d.tryRanges(); e != nil {
			return nil, e
		}
	}

	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("open local file failed: %w", err)
	}
	if err := fp.Truncate(d.status.TotalSize); err != nil {
		return nil, fmt.Errorf("truncate local file failed: %w", err)
	}

	d.wa = fp
	d.closer = fp

	return d, nil
}

func (d *HTTPDownloader) tryRanges() error {
	resp, e := d.client.Head(d.URL)
	if e != nil {
		return fmt.Errorf("head before download: %w", e)
	}
	if resp.StatusCode/100 != 2 {
		// not an error, some server does not allow head method
		return nil
	}

	if resp.ContentLength > 0 {
		d.status.TotalSize = resp.ContentLength
	}

	status := d.status
	if resp.Header.Get("Accept-Ranges") == "bytes" {
		blks := status.TotalSize / DownloadBlockSize
		if blks*DownloadBlockSize < status.TotalSize {
			blks++
		}
		status.BlockCount = int(blks)
		status.DoneBlocks = make([]bool, status.BlockCount)
	} else {
		status.BlockCount = 1
		status.DoneBlocks = []bool{false}
	}

	return nil
}

func (d *HTTPDownloader) TotalSize() int64 {
	return d.status.TotalSize
}

func (d *HTTPDownloader) StartWithoutProgress(ctx context.Context) error {
	if d.closer != nil {
		defer d.closer.Close()
	}

	if d.wa == nil || d.status.BlockCount <= 1 {
		return d.downloadInOneshot(ctx)
	}
	return d.downloadInBlocks(ctx)
}

func (d *HTTPDownloader) Start(ctx context.Context, progress chan<- int64) error {
	if d.closer != nil {
		defer d.closer.Close()
	}

	wc := stream.NewCountingWriterAt(d.wa)
	go func() {
		defer close(progress)
		for cnt := range wc.Count() {
			progress <- cnt
		}
	}()
	d.wa = wc

	if d.status.BlockCount <= 1 {
		return d.downloadInOneshot(ctx)
	}
	return d.downloadInBlocks(ctx)
}

// download file in blocks, async
func (d *HTTPDownloader) downloadInBlocks(ctx context.Context) error {
	type blockInfo struct {
		idx      int
		from, to int64
	}
	// file into blocks
	blocks := make(chan *blockInfo)

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		defer close(blocks)

		if len(d.status.DoneBlocks) == 0 {
			d.status.DoneBlocks = make([]bool, d.status.BlockCount)
		}

		for i, ok := range d.status.DoneBlocks {
			if !ok {
				from := DownloadBlockSize * int64(i)
				var to int64
				if i == len(d.status.DoneBlocks)-1 {
					to = d.status.TotalSize
				} else {
					to = from + DownloadBlockSize
				}

				select {
				case blocks <- &blockInfo{idx: i, from: from, to: to}:
				case <-ctx.Done():
					return ctx.Err()
				}
			}
		}

		return nil
	})

	for i := 0; i < defaultDownloadConcurrency; i++ {
		eg.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return ctx.Err()

				case blk, ok := <-blocks:
					if !ok {
						return nil
					}

					// todo: retry
					if err := d.downloadOneBlock(ctx, blk.from, blk.to); err != nil {
						return fmt.Errorf("downloading block %d: %w", blk.idx, err)
					}

					d.status.DoneBlocks[blk.idx] = true
				}
			}
		})
	}

	return eg.Wait()
}

// sync download block
func (d *HTTPDownloader) downloadOneBlock(ctx context.Context, from, to int64) error {
	req, err := http.NewRequest("GET", d.URL, nil)
	if err != nil {
		return fmt.Errorf("create block download request failed: %w", err)
	}
	if d.status.BlockCount > 1 {
		req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", from, to-1))
	} else {
		req.Header.Add("Range", fmt.Sprintf("bytes=%d-", from))
	}
	req = req.WithContext(ctx)

	resp, err := d.client.Do(req)
	if err != nil {
		return fmt.Errorf("request to get failed: %w", err)
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	switch resp.StatusCode {
	case http.StatusOK:
		// server returns entire file, drop contents we do not need
		if resp.ContentLength > 0 && resp.ContentLength < to {
			return errors.New("incomplete block in 200 response")
		}
		io.CopyN(ioutil.Discard, resp.Body, from)

	case http.StatusPartialContent:
		if resp.ContentLength > 0 && resp.ContentLength < to-from {
			return errors.New("incomplete block in 206 response")
		}
		if ranges := resp.Header.Get("Content-Range"); ranges != "" {
			ranges = strings.TrimPrefix(ranges, "bytes ")
			parts := strings.SplitN(ranges, "/", 2)
			if parts[0] != "*" {
				parts = strings.SplitN(parts[0], "-", 2)
				left, e := strconv.ParseInt(parts[0], 10, 64)
				if e != nil {
					return fmt.Errorf("parse content range [%s]: %w", ranges, e)
				}
				right, e := strconv.ParseInt(parts[1], 10, 64)
				if e != nil {
					return fmt.Errorf("parse content range [%s]: %w", ranges, e)
				}
				if left != from || right != to-1 {
					return fmt.Errorf("got mismatch ranges %s, expecting %d-%d/%d", ranges, from, to-1, to-from)
				}
			}
		}

	default:
		return fmt.Errorf("download block failed, got %d", resp.StatusCode)
	}

	w := stream.NewSegmentWriter(d.wa, from, to-from)
	if _, e := io.CopyN(w, resp.Body, to-from); e != nil {
		return fmt.Errorf("copy block: %w", e)
	}

	return nil
}

func (d *HTTPDownloader) downloadInOneshot(ctx context.Context) error {
	req, e := http.NewRequest("GET", d.URL, nil)
	if e != nil {
		return fmt.Errorf("create get request: %w", e)
	}
	req = req.WithContext(ctx)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, e := client.Do(req)
	if e != nil {
		return fmt.Errorf("request to download: %w", e)
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	switch resp.StatusCode {
	case http.StatusOK, http.StatusPartialContent:
		if resp.ContentLength > 0 && resp.ContentLength != d.status.TotalSize {
			return fmt.Errorf("incomplete download, expecting %d, got %d", d.status.TotalSize, resp.ContentLength)
		}
	default:
		return fmt.Errorf("download failed, got %d", resp.StatusCode)
	}

	w := d.writer
	if w == nil {
		w = stream.NewSegmentWriter(d.wa, 0, d.status.TotalSize)
	}
	_, e = io.CopyN(w, resp.Body, d.status.TotalSize)
	if e != nil {
		return fmt.Errorf("copy content in oneshot: %w", e)
	}

	return nil
}

// direct http downloading
func DownloadFile(u string, opts ...DownloadOption) (io.ReadCloser, error) {
	c := http.DefaultClient
	for _, opt := range opts {
		opt(c)
	}

	resp, e := c.Get(u)
	if e != nil {
		return nil, e
	}
	if resp.StatusCode/100 != 2 {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("http request failed, got [%d]: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return &RespReader{resp.Body}, nil
}

// todo: options on downloader not client
// todo: options on tls, block size...
type DownloadOption func(*http.Client)

func SetTimeout(d time.Duration) DownloadOption {
	return func(c *http.Client) {
		c.Timeout = d
	}
}

type RespReader struct {
	io.ReadCloser
}

func (r *RespReader) Close() error {
	io.Copy(ioutil.Discard, r.ReadCloser)
	return r.ReadCloser.Close()
}
