package files

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"mime/multipart"
)

type HTTPFormWriter struct {
	req    *FormUploadReq
	closer func() error
	w      io.Writer
	done   chan struct{}
	e      error
}

func NewHTTPFormWriter(ctx context.Context, req *FormUploadReq) (*HTTPFormWriter, error) {
	pr, pw := io.Pipe()

	// hope it's big enough to avoid deadlock before writing file content
	buf := bufio.NewWriterSize(pw, 1<<20)

	mw := multipart.NewWriter(buf)
	for key, value := range req.Forms {
		if e := mw.WriteField(key, value); e != nil {
			return nil, fmt.Errorf("write form values failed: %w", e)
		}
	}
	fw, e := mw.CreateFormFile("file", req.Filename)
	if e != nil {
		return nil, fmt.Errorf("create form file writer: %w", e)
	}

	if req.Size > 0 {
		req.Size += int64(buf.Size())
	}

	w := &HTTPFormWriter{
		req:  req,
		w:    fw,
		done: make(chan struct{}),
		closer: func() error {
			e1 := mw.Close()
			e2 := buf.Flush()
			e3 := pw.Close()
			if e1 != nil {
				return e1
			}
			if e2 != nil {
				return e2
			}
			return e3
		},
	}

	go func() {
		defer close(w.done)
		w.e = req.do(ctx, mw.FormDataContentType(), pr)
	}()

	return w, nil
}

func (w *HTTPFormWriter) Write(p []byte) (n int, e error) {
	select {
	case <-w.done:
		if w.e != nil {
			return 0, fmt.Errorf("error in request: %s", e)
		}
		return 0, io.ErrClosedPipe
	default:
		return w.w.Write(p)
	}
}

func (w *HTTPFormWriter) Close() error {
	e := w.closer()
	<-w.done
	if e != nil {
		return e
	}
	return w.e
}
