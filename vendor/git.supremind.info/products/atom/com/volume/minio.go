package volume

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/minio/minio-go/v6"
)

type Minio struct {
	Spec

	cli *minio.Client
}

var _ StorageService = (*Minio)(nil)

const (
	didiDefaultEndpoint = "https://s3.didiyunapi.com"
)

func NewMinio(spec Spec, secretID, secretKey string) (*Minio, error) {
	if spec.Endpoint == "" {
		spec.Endpoint = didiDefaultEndpoint
	}
	secure := false
	spec.Endpoint = strings.TrimPrefix(spec.Endpoint, "http://")
	if strings.HasPrefix(spec.Endpoint, "https://") {
		secure = true
		spec.Endpoint = strings.TrimPrefix(spec.Endpoint, "https://")
	}

	cli, e := minio.New(spec.Endpoint, secretID, secretKey, secure)
	if e != nil {
		return nil, fmt.Errorf("initialize minio client: %w", e)
	}

	return &Minio{
		Spec: spec,
		cli:  cli,
	}, nil
}

func (mv *Minio) GetObjectInfo(ctx context.Context, key string) (*FileMeta, error) {
	info, e := mv.cli.StatObject(mv.Bucket, filepath.Join(mv.Path, key), minio.StatObjectOptions{})

	if e != nil {
		if me, ok := e.(minio.ErrorResponse); ok {
			if me.StatusCode == http.StatusNotFound {
				return nil, os.ErrNotExist
			}
		}
		return nil, fmt.Errorf("stat minio object: %w", e)
	}

	return &FileMeta{
		Key:         key,
		Size:        info.Size,
		ModTime:     info.LastModified,
		ContentType: info.ContentType,
		ETag:        info.ETag,
	}, nil
}

// GetDownloadURL returns download url which could not be HEAD-ed
func (mv *Minio) GetDownloadURL(key string, ttl time.Duration) (string, error) {
	if ttl <= 0 {
		ttl = defaultDownloadTTL
	}
	presigned, e := mv.cli.PresignedGetObject(mv.Bucket, filepath.Join(mv.Path, key), ttl, nil)
	if e != nil {
		return "", fmt.Errorf("presign download url: %w", e)
	}
	return presigned.String(), nil
}

// ListDirContents blocks until all contents listed, or a (optional) limit exceeded, it only returns that folder's files (no subfolder's files)
func (mv *Minio) ListDirContents(ctx context.Context, prefix string, limit int, files chan<- *FileMeta) error {
	return mv.listObjects(ctx, prefix, limit, files, true)
}

// ListContents blocks until all contents listed, or a (optional) limit exceeded, including the files under its subfolders
func (mv *Minio) ListContents(ctx context.Context, prefix string, limit int, files chan<- *FileMeta) error {
	return mv.listObjects(ctx, prefix, limit, files, false)
}

// ObjectReader reads remote object content
func (mv *Minio) ObjectReader(ctx context.Context, key string, size int64) (io.Reader, error) {
	return mv.cli.GetObjectWithContext(ctx, mv.Bucket, filepath.Join(mv.Path, key), minio.GetObjectOptions{})
}

// ObjectWriter writes content to remote object, it may also be an io.WriterAt
func (mv *Minio) ObjectWriter(ctx context.Context, key string, size int64, overwrite bool) (io.Writer, error) {
	opts := minio.PutObjectOptions{
		NumThreads: workers,
	}
	if size > blockSize {
		opts.PartSize = uint64(blockSize)
	}
	return newUploadWriter(func(r io.Reader) error {
		_, e := mv.cli.PutObjectWithContext(ctx, mv.Bucket, filepath.Join(mv.Path, key), r, size, opts)
		return e
	}, size), nil
}

func (mv *Minio) listObjects(ctx context.Context, prefix string, limit int, files chan<- *FileMeta, dir bool) error {
	done := make(chan struct{})
	defer close(done)
	count := 0

	prefix = filepath.Join(mv.Path, prefix)
	if dir {
		prefix += "/"
	}
	if strings.HasPrefix(prefix, "/") {
		prefix = strings.TrimPrefix(prefix, "/")
	}

	objects := mv.cli.ListObjectsV2(mv.Bucket, prefix, !dir, done)
	for obj := range objects {
		key := obj.Key
		if !dir {
			// filter out problematic files
			if obj.Size <= 0 {
				continue
			}
			if key == "." || key[len(key)-1] == '/' {
				continue
			}
		}

		rel, e := filepath.Rel(mv.Path, key)
		if e != nil {
			return fmt.Errorf("retrive relative key: %w", e)
		}

		meta := &FileMeta{
			Key:         rel,
			Size:        obj.Size,
			ModTime:     obj.LastModified,
			ContentType: obj.ContentType,
			ETag:        obj.ETag,
			Dir:         false,
		}

		if dir && obj.Size == 0 {
			meta = &FileMeta{
				Key: rel,
				Dir: true,
			}
		}

		select {
		case files <- meta:
			count++
			if limit > 0 && count >= limit {
				return nil
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}
