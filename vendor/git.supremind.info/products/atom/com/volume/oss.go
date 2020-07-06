package volume

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const (
	ossMaxPartCount = 10000
	ossMaxPartSize  = 5 << 30
)

type OSS struct {
	Spec

	bkt *oss.Bucket
}

var _ StorageService = (*OSS)(nil)

func NewOSS(spec Spec, accessKeyID, accessKeySecret string) (*OSS, error) {
	cli, e := oss.New(spec.Endpoint, accessKeyID, accessKeySecret)
	if e != nil {
		return nil, fmt.Errorf("create oss client: %w", e)
	}
	bkt, e := cli.Bucket(spec.Bucket)
	if e != nil {
		return nil, fmt.Errorf("get oss bucket manager: %w", e)
	}

	return &OSS{Spec: spec, bkt: bkt}, nil
}

func (ov *OSS) GetObjectInfo(ctx context.Context, key string) (*FileMeta, error) {
	meta, e := ov.bkt.GetObjectMeta(filepath.Join(ov.Path, key))

	if key == "" {
		return nil, os.ErrNotExist
	}
	if e != nil {
		if ke, ok := e.(oss.ServiceError); ok {
			if http.StatusNotFound == ke.StatusCode {
				return nil, os.ErrNotExist
			}
		}
		return nil, fmt.Errorf("get oss object meta: %s, %w", key, e)
	}

	size, e := strconv.ParseInt(meta.Get("Content-Length"), 10, 64)
	if e != nil {
		return nil, fmt.Errorf("parse oss object size: %w", e)
	}

	return &FileMeta{
		Key:         key,
		Size:        size,
		ContentType: meta.Get("Content-Type"),
		ETag:        meta.Get("ETag"),
	}, nil
}

func (ov *OSS) GetDownloadURL(key string, ttl time.Duration) (string, error) {
	if ttl <= 0 {
		ttl = defaultDownloadTTL
	}
	return ov.bkt.SignURL(filepath.Join(ov.Path, key), "GET", ttl.Nanoseconds()/1e9)
}

// ListDirContents blocks until all contents listed, or a (optional) limit exceeded, it only returns that folder's files (no subfolder's files)
func (ov *OSS) ListDirContents(ctx context.Context, prefix string, limit int, files chan<- *FileMeta) error {
	return ov.listObjects(ctx, prefix, limit, files, true)
}

// ListContents blocks until all contents listed, or a (optional) limit exceeded, including the files under its subfolders
// files will be closed inside
func (ov *OSS) ListContents(ctx context.Context, prefix string, limit int, files chan<- *FileMeta) error {
	return ov.listObjects(ctx, prefix, limit, files, false)
}

func (ov *OSS) ObjectReader(ctx context.Context, key string, size int64) (io.Reader, error) {
	r, e := newHTTPReader(ctx, ov, key, size)
	return r, e
}

// ObjectWriter writes content to remote object, it may also be an io.WriterAt
func (ov *OSS) ObjectWriter(ctx context.Context, key string, size int64, overwrite bool) (io.Writer, error) {
	return newUploadWriter(func(r io.Reader) error {
		return ov.bkt.PutObject(filepath.Join(ov.Path, key), r)
	}, size), nil
}

func (ov *OSS) listObjects(ctx context.Context, prefix string, limit int, files chan<- *FileMeta, dir bool) error {
	delimiter := ""
	aPrefix := filepath.Join(ov.Path, prefix)
	if dir {
		delimiter = "/"
		aPrefix = aPrefix + "/"
	}
	aPrefix = strings.TrimPrefix(aPrefix, "/")
	count := 0
	marker := ""
	for {
		size := 1000
		if limit > 0 && count+size > limit {
			size = limit - count
		}
		lor, e := ov.bkt.ListObjects(oss.MaxKeys(size), oss.Marker(marker), oss.Prefix(aPrefix), oss.Delimiter(delimiter))
		if e != nil {
			return fmt.Errorf("list oss objects: %w", e)
		}
		for _, obj := range lor.Objects {
			// filter out problematic files
			if obj.Size <= 0 {
				continue
			}
			rel, e := filepath.Rel(ov.Path, obj.Key)
			if e != nil {
				return fmt.Errorf("retrieve relative key: %w", e)
			}

			select {
			case files <- &FileMeta{
				Key:         rel,
				Size:        obj.Size,
				ContentType: obj.Type,
				ModTime:     obj.LastModified,
				ETag:        obj.ETag,
				Dir:         false,
			}:
				count++
			case <-ctx.Done():
				return ctx.Err()
			}
		}

		if dir {
			for _, obj := range lor.CommonPrefixes {
				obj = strings.TrimPrefix(obj, "/")
				rel, e := filepath.Rel(ov.Path, obj)
				if e != nil {
					return fmt.Errorf("retrieve relative key: %w", e)
				}
				select {
				case files <- &FileMeta{
					Key: rel,
					Dir: true,
				}:
					count++
				case <-ctx.Done():
					return ctx.Err()
				}
			}
		}

		marker = lor.NextMarker
		if !lor.IsTruncated || (limit > 0 && count >= limit) {
			break
		}
	}

	return nil
}
