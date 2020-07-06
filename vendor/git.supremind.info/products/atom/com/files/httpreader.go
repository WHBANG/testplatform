package files

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func NewHTTPRangeReader(ctx context.Context, url string, size int64, opts ...DownloadOption) *HTTPRangeReader {
	r := &HTTPRangeReader{
		ctx:    ctx,
		url:    url,
		client: http.DefaultClient,
		size:   size,
	}

	for _, opt := range opts {
		opt(r.client)
	}

	return r
}

type HTTPRangeReader struct {
	ctx    context.Context
	url    string
	client *http.Client
	offset int64
	size   int64
}

func (r *HTTPRangeReader) Read(p []byte) (n int, e error) {
	if r.offset >= r.size {
		return 0, io.EOF
	}

	partSize := len(p)
	if r.offset+int64(partSize) > r.size {
		partSize = int(r.size - r.offset)
	}
	n, e = r.ReadAt(p[:partSize], r.offset)
	r.offset += int64(n)
	if r.offset >= r.size {
		e = io.EOF
	}
	return
}

// Seek sets the offset for the next Read or Write on file to offset
// minio expects readerAt to be a seeker
func (r *HTTPRangeReader) Seek(offset int64, whence int) (ret int64, err error) {
	var pos int64

	switch whence {
	case io.SeekStart:
		pos = offset
	case io.SeekCurrent:
		pos = r.offset + offset
	case io.SeekEnd:
		pos = r.size + offset
	}
	if pos > r.size || pos < 0 {
		return 0, os.ErrInvalid
	}

	r.offset = pos
	return r.offset, nil
}

func (r *HTTPRangeReader) ReadAt(p []byte, off int64) (n int, e error) {
	to := off + int64(len(p))
	if to > r.size {
		to = r.size
	}

	req, e := http.NewRequestWithContext(r.ctx, "GET", r.url, nil)
	if e != nil {
		return 0, e
	}
	req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", off, to-1))
	resp, e := r.client.Do(req)
	if e != nil {
		return 0, e
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	var partSize int64
	switch resp.StatusCode {
	case http.StatusOK:
		// drop unwanted bytes
		_, e := io.CopyN(ioutil.Discard, resp.Body, off)
		if e != nil {
			return 0, fmt.Errorf("drop unwanted heading bytes in a complete response: %w", e)
		}
		if resp.ContentLength > 0 {
			partSize = resp.ContentLength - off
		}
	case http.StatusPartialContent:
		ranges := resp.Header.Get("Content-Range")
		if ranges == "" {
			break
		}
		if strings.HasPrefix(ranges, "bytes") {
			break
		}
		ranges = strings.TrimPrefix(ranges, "bytes ")
		parts := strings.SplitN(ranges, "/", 2)
		if parts[0] != "*" {
			parts = strings.SplitN(parts[0], "-", 2)
			left, e := strconv.ParseInt(parts[0], 10, 64)
			if e != nil {
				return 0, fmt.Errorf("parse content range [%s]: %w", ranges, e)
			}
			right, e := strconv.ParseInt(parts[1], 10, 64)
			if e != nil {
				return 0, fmt.Errorf("parse content range [%s]: %w", ranges, e)
			}
			partSize = right - left + 1
		}
		if parts[1] != "*" {
			size, e := strconv.ParseInt(parts[1], 10, 64)
			if e != nil {
				return 0, fmt.Errorf("parse content range size [%s]: %w", ranges, e)
			}
			if size != r.size {
				return 0, fmt.Errorf("reader size changed, expect %d, new size is %d", r.size, size)
			}
		}
	default:
		e = fmt.Errorf("download failed, got %d", resp.StatusCode)
		return
	}

	cnt := 0
	for e == nil && n < len(p) && (partSize <= 0 || int64(n) < partSize) {
		cnt, e = resp.Body.Read(p[n:])
		n += cnt
	}
	if n >= len(p) {
		e = nil
		n = len(p)
	}
	if e == io.EOF {
		e = nil
	}

	return
}
