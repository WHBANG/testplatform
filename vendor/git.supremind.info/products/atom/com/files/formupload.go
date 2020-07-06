package files

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"

	"git.supremind.info/products/atom/com/stream"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

const defaultBlockSize int64 = 32 << 20

type FormUploadReq struct {
	Endpoint  string
	Forms     map[string]string
	Header    url.Values
	Filename  string
	Reader    io.Reader
	Client    *http.Client
	Size      int64
	BlockSize int64
}

func SimpleFormUpload(ctx context.Context, req *FormUploadReq) error {
	if e := req.validate(); e != nil {
		return e
	}

	buf := &bytes.Buffer{}
	mw := multipart.NewWriter(buf)

	if e := req.writeBody(mw); e != nil {
		return e
	}

	contentType := mw.FormDataContentType()
	mw.Close()

	return req.do(ctx, contentType, buf)
}

// StreamingFormUpload uploads file by streaming file content into multiple parts in request body
func StreamingFormUpload(ctx context.Context, req *FormUploadReq) error {
	if e := req.validate(); e != nil {
		return e
	}

	if e := req.fixSize(); e != nil {
		return e
	}

	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		defer pw.Close()
		defer mw.Close()

		return req.writeBody(mw)
	})

	if e := req.do(ctx, mw.FormDataContentType(), pr); e != nil {
		return e
	}

	if e := eg.Wait(); e != nil {
		return errors.Wrap(e, "copy file as request body failed")
	}

	return nil
}

// content length should be file size + size of other metadata
func (req *FormUploadReq) fixSize() error {
	// if we do not know file size in advance, do not include it in header
	if req.Size <= 0 {
		return nil
	}

	buf := &bytes.Buffer{}
	mw := multipart.NewWriter(buf)
	for key, value := range req.Forms {
		if e := mw.WriteField(key, value); e != nil {
			return errors.Wrap(e, "write form values failed")
		}
	}
	// file is not actually written, only filename
	if _, e := mw.CreateFormFile("file", req.Filename); e != nil {
		return errors.Wrap(e, "create form file writer failed")
	}

	// all but file content is written, and the missing content length is known now
	mw.Close()

	n := buf.Len()
	req.Size += int64(n)

	return nil
}

func (req *FormUploadReq) validate() error {
	if req.Client == nil {
		req.Client = http.DefaultClient
	}
	if req.Endpoint == "" {
		return errors.New("empty endpoint")
	}
	if req.Reader == nil {
		return errors.New("nil reader")
	}

	return nil
}

func (req *FormUploadReq) writeBody(mw *multipart.Writer) error {
	for key, value := range req.Forms {
		if e := mw.WriteField(key, value); e != nil {
			return errors.Wrap(e, "write form values failed")
		}
	}
	fw, e := mw.CreateFormFile("file", req.Filename)
	if e != nil {
		return errors.Wrap(e, "create form file writer failed")
	}

	if req.BlockSize <= 0 {
		req.BlockSize = defaultBlockSize
	}
	if _, e := stream.DoubleBufferedCopy(fw, req.Reader, int(req.BlockSize)); e != nil {
		return errors.Wrap(e, "copy content to file writer failed")
	}

	return nil
}

func (req *FormUploadReq) do(ctx context.Context, contentType string, r io.Reader) error {
	postReq, e := http.NewRequest("POST", req.Endpoint, r)
	if e != nil {
		return errors.Wrap(e, "failed to create post request")
	}
	postReq = postReq.WithContext(ctx)
	h := postReq.Header
	for k, vs := range req.Header {
		for _, v := range vs {
			h.Add(k, v)
		}
	}
	h.Set("Content-Type", contentType)
	postReq.Header = h
	if req.Size > 0 {
		postReq.ContentLength = req.Size
	}

	resp, e := req.Client.Do(postReq)
	if e != nil {
		return errors.Wrap(e, "failed to send post request")
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()
	if resp.StatusCode/100 != 2 {
		buf := &bytes.Buffer{}
		io.CopyN(buf, resp.Body, 1024)
		return fmt.Errorf("post file failed, got [%d] %s: %s", resp.StatusCode, http.StatusText(resp.StatusCode), buf.String())
	}

	return nil
}
