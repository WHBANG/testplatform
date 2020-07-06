package apiserver

import (
	"context"
	"io"
	"os"
	"sync/atomic"

	"git.supremind.info/products/atom/proto/go/api"
	"github.com/pkg/errors"
)

// FileUploader upload files to volumes
type FileUploader struct {
	u Uploader
}

func NewFileUploader(vol *api.Volume, cli api.VolumeServiceClient, prefix string, overwrite bool) (*FileUploader, error) {
	u, e := NewUploader(vol, cli, prefix, overwrite)
	if e != nil {
		return nil, e
	}

	return &FileUploader{u: u}, nil
}

func (fp *FileUploader) UploadFile(ctx context.Context, key, path string, doneSize chan<- int64) (int64, error) {
	f, e := os.Open(path)
	if e != nil {
		return 0, errors.Wrap(e, "failed to open file")
	}
	defer f.Close()
	stat, e := f.Stat()
	if e != nil {
		return 0, errors.Wrap(e, "failed to get file stat")
	}

	rc := newReaderAtCounter(f, f)

	go func() {
		for prg := range rc.prg {
			doneSize <- prg
		}
		close(doneSize)
	}()

	return stat.Size(), fp.u.Upload(ctx, rc, key, stat.Size())
}

type readerCounter struct {
	io.Reader

	cnt int64
	prg chan int64
}

func newReaderCounter(r io.Reader) *readerCounter {
	return &readerCounter{
		Reader: r,
		prg:    make(chan int64, 16),
	}
}

func (rc *readerCounter) Read(p []byte) (n int, err error) {
	n, err = rc.Reader.Read(p)
	cnt := atomic.AddInt64(&rc.cnt, int64(n))
	rc.prg <- cnt
	return
}

func (rc *readerCounter) Close() error {
	if c, ok := rc.Reader.(io.Closer); ok {
		c.Close()
	}
	close(rc.prg)
	return nil
}

type readerAtCounter struct {
	*readerCounter
	io.ReaderAt
}

func newReaderAtCounter(r io.Reader, ra io.ReaderAt) *readerAtCounter {
	return &readerAtCounter{
		readerCounter: newReaderCounter(r),
		ReaderAt:      ra,
	}
}

func (rc *readerAtCounter) ReadAt(p []byte, off int64) (n int, err error) {
	n, err = rc.ReaderAt.ReadAt(p, off)
	cnt := atomic.AddInt64(&rc.cnt, int64(n))
	rc.prg <- cnt
	return
}
