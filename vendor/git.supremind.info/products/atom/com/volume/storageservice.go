package volume

import (
	"context"
	"io"
	"time"

	"git.supremind.info/products/atom/com/files"
	"git.supremind.info/products/atom/com/stream"
)

const (
	defaultDownloadTTL = time.Hour * 24 * 7
	uploadTTL          = time.Hour * 24
	shortRequestTTL    = 1 * time.Minute
	longRequestTTL     = 24 * time.Hour

	blockSize int64 = 128 << 20
	workers         = 8
)

type StorageService interface {
	GetObjectInfo(ctx context.Context, key string) (*FileMeta, error)

	GetDownloadURL(key string, ttl time.Duration) (string, error)

	// ListContents blocks until all contents listed, or a (optional) limit exceeded, including the files under its subfolders
	ListContents(ctx context.Context, prefix string, limit int, files chan<- *FileMeta) error

	// ListDirContents blocks until all contents listed, or a (optional) limit exceeded, it only returns that folder's files (no subfolder's files)
	ListDirContents(ctx context.Context, prefix string, limit int, files chan<- *FileMeta) error

	// ObjectReader reads from remote object content, it may also be an io.ReaderAt
	ObjectReader(ctx context.Context, key string, size int64) (io.Reader, error)

	// ObjectWriter writes content to remote object, it may also be an io.WriterAt or io.ReaderFrom
	ObjectWriter(ctx context.Context, key string, size int64, overwrite bool) (io.Writer, error)
}

type FileMeta struct {
	Key         string // relative to (optional) path
	Size        int64
	ContentType string
	ModTime     time.Time
	ETag        string
	Dir         bool
}

type Spec struct {
	Bucket   string
	Path     string
	Endpoint string

	InJuewa bool
}

func newHTTPReader(ctx context.Context, ss StorageService, key string, size int64) (stream.AutoReader, error) {
	u, e := ss.GetDownloadURL(key, defaultDownloadTTL)
	if e != nil {
		return nil, e
	}

	r := files.NewHTTPRangeReader(ctx, u, size)
	return r, nil
}

type uploadWriter struct {
	size int64

	w    *io.PipeWriter
	r    *io.PipeReader
	done chan struct{}
	e    error
}

func newUploadWriter(upload func(io.Reader) error, size int64) *uploadWriter {
	uw := &uploadWriter{
		size: size,
		done: make(chan struct{}),
	}

	uw.r, uw.w = io.Pipe()
	go func() {
		defer func() {
			close(uw.done)
			uw.w.Close()
		}()
		uw.e = upload(uw.r)
	}()

	return uw
}

func (uw *uploadWriter) Write(p []byte) (int, error) {
	return uw.w.Write(p)
}

func (uw *uploadWriter) Close() error {
	e := uw.w.Close()
	<-uw.done
	if e != nil {
		return e
	}
	return uw.e
}
