package volume

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/client"
	"github.com/qiniu/api.v7/v7/storage"
)

const (
	kodoAPIEndpoint   = "api.qiniu.com"
	kodoJuewaUpHost   = "http://free-jjh-public-up-z0.qiniup.com"
	kodoDefaultUpHost = "http://up.qiniup.com"
)

const (
	kodoBlockSize = int64(4 << 20)
)

type KODO struct {
	Spec

	mgr  *storage.BucketManager
	info *storage.BucketInfo

	mu     sync.Mutex
	domain string
}

var _ StorageService = (*KODO)(nil)

func NewKODO(spec Spec, ak, sk string) (*KODO, error) {
	mac := auth.New(ak, sk)
	mgr := storage.NewBucketManager(mac, &storage.Config{UseHTTPS: false})
	info, e := mgr.GetBucketInfo(spec.Bucket)
	if e != nil {
		return nil, e
	}
	if !strings.HasPrefix(spec.Endpoint, "http://") && !strings.HasPrefix(spec.Endpoint, "https://") {
		spec.Endpoint = "http://" + spec.Endpoint
	}
	return &KODO{
		Spec: spec,
		mgr:  mgr,
		info: &info,
	}, nil
}

func (kv *KODO) GetObjectInfo(ctx context.Context, key string) (*FileMeta, error) {
	info, e := kv.mgr.Stat(kv.Bucket, filepath.Join(kv.Path, key))
	if e != nil {
		if e == storage.ErrNoSuchFile {
			return nil, os.ErrNotExist
		}
		if ke, ok := e.(*client.ErrorInfo); ok {
			code := ke.Code
			if code == 612 {
				return nil, os.ErrNotExist
			}
		}
		return nil, fmt.Errorf("stat kodo object info: %s, %w", key, e)
	}

	return &FileMeta{
		Key:         key,
		Size:        info.Fsize,
		ModTime:     storage.ParsePutTime(info.PutTime),
		ContentType: info.MimeType,
		ETag:        info.Hash,
	}, nil
}

func (kv *KODO) GetDownloadURL(key string, ttl time.Duration) (string, error) {
	key = filepath.Join(kv.Path, key)
	domain, e := kv.getBucketDomain(context.TODO())
	if e != nil {
		return "", fmt.Errorf("get bucket domain: %w", e)
	}
	if !kv.info.IsPrivate() {
		return storage.MakePublicURL(domain, key), nil
	}

	if ttl <= 0 {
		ttl = defaultDownloadTTL
	}
	return storage.MakePrivateURL(kv.mgr.Mac, domain, key, time.Now().Add(ttl).Unix()), nil
}

func (kv *KODO) listObjects(ctx context.Context, prefix string, limit int, files chan<- *FileMeta, dir bool) error {
	delimiter := ""
	aPrefix := filepath.Join(kv.Path, prefix)
	if dir {
		delimiter = "/"
		aPrefix = aPrefix + "/"
	}
	aPrefix = strings.TrimPrefix(aPrefix, "/")

	marker := ""
	count := 0

	for {
		l := 1000
		if limit > 0 {
			if count >= limit {
				return nil
			}
			if limit-count < l {
				l = limit - count
			}
		}

		entries, commonPrefixes, nextMarker, hashNext, e := kv.mgr.ListFiles(kv.Bucket, aPrefix, delimiter, marker, l)
		if e != nil {
			return fmt.Errorf("list files: %w", e)
		}

		for _, entry := range entries {

			// filter out problematic files
			if entry.Fsize <= 0 {
				continue
			}
			// unexpected file key, should not be uploaded
			if entry.Key[len(entry.Key)-1] == '/' {
				continue
			}

			rel, e := filepath.Rel(kv.Path, entry.Key)
			if e != nil {
				return fmt.Errorf("retrieve relative key: %w", e)
			}

			select {
			case files <- &FileMeta{
				Key:         rel,
				Size:        entry.Fsize,
				ModTime:     storage.ParsePutTime(entry.PutTime),
				ContentType: entry.MimeType,
				ETag:        entry.Hash,
				Dir:         false,
			}:
				count++
			case <-ctx.Done():
				return ctx.Err()
			}
		}

		if dir {
			for _, obj := range commonPrefixes {
				obj = strings.TrimPrefix(obj, "/")
				rel, e := filepath.Rel(kv.Path, obj)
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

		if hashNext {
			marker = nextMarker
		} else {
			break
		}
	}

	return nil
}

// ListContents blocks until all contents listed, or a (optional) limit exceeded, including the files under its subfolders
func (kv *KODO) ListContents(ctx context.Context, prefix string, limit int, files chan<- *FileMeta) error {
	return kv.listObjects(ctx, prefix, limit, files, false)
}

// ListDirContents blocks until all contents listed, or a (optional) limit exceeded, it only returns that folder's files (no subfolder's files)
func (kv *KODO) ListDirContents(ctx context.Context, prefix string, limit int, files chan<- *FileMeta) error {
	return kv.listObjects(ctx, prefix, limit, files, true)
}

// ObjectReader reads remote object content
func (kv *KODO) ObjectReader(ctx context.Context, key string, size int64) (io.Reader, error) {
	r, e := newHTTPReader(ctx, kv, key, size)
	return r, e
}

// ObjectWriter writes content to remote object, it may also be an io.WriterAt
func (kv *KODO) ObjectWriter(ctx context.Context, key string, size int64, overwrite bool) (io.Writer, error) {
	return newUploadWriter(func(r io.Reader) error {
		return kv.chunkedUpload(ctx, key, size, overwrite, r)
	}, size), nil
}

func (kv *KODO) uploadToken(key string, overwrite bool) string {
	putPolicy := storage.PutPolicy{
		Expires: uint64(uploadTTL.Seconds()),
		Scope:   kv.Bucket + ":" + key,
	}
	if !overwrite {
		putPolicy.InsertOnly = 1
	}

	return putPolicy.UploadToken(kv.mgr.Mac)
}

func (kv *KODO) formUpload(ctx context.Context, key string, size int64, overwrite bool, r io.Reader) error {
	uploader := storage.NewFormUploader(&storage.Config{
		UseHTTPS: false,
		UpHost:   kv.upHost(),
	})
	ret := &storage.PutRet{}
	putExt := &storage.PutExtra{
		UpHost: kv.upHost(),
	}

	key = filepath.Join(kv.Path, key)
	e := uploader.Put(ctx, ret, kv.uploadToken(key, overwrite), key, r, size, putExt)
	if e != nil {
		return fmt.Errorf("form upload: %w", e)
	}
	return nil
}

func (kv *KODO) chunkedUpload(ctx context.Context, key string, size int64, overwrite bool, r io.Reader) error {
	uploader := storage.NewResumeUploader(&storage.Config{
		UseHTTPS: false,
		UpHost:   kv.upHost(),
	})
	key = filepath.Join(kv.Path, key)

	token := kv.uploadToken(key, overwrite)
	host := kv.upHost()
	cnt := size / kodoBlockSize
	if kodoBlockSize*cnt < size {
		cnt++
	}
	prg := make([]storage.BlkputRet, int(cnt))

	for i := 0; i < int(cnt); i++ {
		from := int64(i) * kodoBlockSize
		to := from + kodoBlockSize
		if to > size {
			to = size
		}

		lr := io.LimitReader(r, to-from)
		e := uploader.Mkblk(ctx, token, host, &prg[i], int(to-from), lr, int(to-from))
		if e != nil {
			return fmt.Errorf("failed on block %d-%d: %w", from, to, e)
		}
	}

	e := uploader.Mkfile(ctx, token, host, nil, key, true, size, &storage.RputExtra{
		Progresses: prg,
	})
	return e
}

func (kv *KODO) upHost() string {
	if kv.InJuewa {
		return kodoJuewaUpHost
	}
	return kodoDefaultUpHost
}

func (kv *KODO) getBucketDomain(ctx context.Context) (string, error) {
	if kv.Endpoint != "" {
		return kv.Endpoint, nil
	}

	kv.mu.Lock()
	defer kv.mu.Unlock()
	if kv.domain != "" {
		return kv.domain, nil
	}

	domains, e := kv.mgr.ListBucketDomains(kv.Bucket)
	if e != nil {
		return "", fmt.Errorf("list bucket domains: %w", e)
	}
	if len(domains) == 0 {
		return "", errors.New("bucket without domain")
	}

	for _, d := range domains {
		if !strings.HasPrefix(d.Domain, ".") {
			kv.domain = d.Domain
			break
		}
	}
	if kv.domain == "" {
		return "", errors.New("no available domain for bucket")
	}

	return kv.domain, nil
}
