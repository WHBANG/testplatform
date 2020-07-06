package volume

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"git.supremind.info/products/atom/com/files"
	"git.supremind.info/products/atom/com/identity"
)

const (
	formFieldKey = "key"

	queryFieldPrefix = "prefix"
	queryFieldKey    = "key"
	queryFieldToken  = "token"
	queryFieldLimit  = "limit"
)

type Datakeeper struct {
	Spec
	signer DatakeeperPolicySigner
}

type DatakeeperPolicySigner func(policy *identity.FilePolicy, ttl time.Duration) (string, error)

var _ StorageService = (*Datakeeper)(nil)

func NewDatakeeper(spec Spec, signer DatakeeperPolicySigner) *Datakeeper {
	return &Datakeeper{Spec: spec, signer: signer}
}

func (dv *Datakeeper) GetObjectInfo(ctx context.Context, key string) (*FileMeta, error) {
	policy := &identity.FilePolicy{
		Op:     identity.FileOpDownload,
		Bucket: dv.Bucket,
		Prefix: filepath.Join(dv.Path, key),
	}
	resp, e := dv.doSimpleRequest(ctx, "HEAD", policy, shortRequestTTL, map[string]string{
		queryFieldKey: filepath.Join(dv.Path, key),
	})
	if e != nil {
		if he, ok := e.(*httpError); ok {
			if he.code == http.StatusNotFound {
				return nil, os.ErrNotExist
			}
		}
		return nil, fmt.Errorf("datakeeper head request: %w", e)
	}

	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	return &FileMeta{
		Key:         key,
		Size:        resp.ContentLength,
		ContentType: resp.Header.Get("Content-Type"),
	}, nil
}

func (dv *Datakeeper) GetDownloadURL(key string, ttl time.Duration) (string, error) {
	key = filepath.Join(dv.Path, key)
	token, e := dv.signer(&identity.FilePolicy{
		Op:     identity.FileOpDownload,
		Bucket: dv.Bucket,
		Prefix: key,
	}, ttl)
	if e != nil {
		return "", fmt.Errorf("sign datakeeper download url: %w", e)
	}

	values := make(url.Values)
	values.Set(queryFieldKey, key)
	values.Set(queryFieldToken, token)
	u := fmt.Sprintf("%s?%s", dv.Endpoint, values.Encode())
	return u, nil
}

func (dv *Datakeeper) listObjects(ctx context.Context, prefix string, limit int, files chan<- *FileMeta) error {
	policy := &identity.FilePolicy{
		Op:     identity.FileOpList,
		Bucket: dv.Bucket,
		Prefix: filepath.Join(dv.Path, prefix),
	}

	resp, e := dv.doSimpleRequest(ctx, "GET", policy, longRequestTTL, map[string]string{
		queryFieldPrefix: filepath.Join(dv.Path, prefix),
		queryFieldLimit:  strconv.Itoa(limit),
	})
	if e != nil {
		return fmt.Errorf("listing datakeeper objects: %w", e)
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	count := 0
	s := bufio.NewScanner(resp.Body)
	for s.Scan() {
		if limit > 0 && count >= limit {
			return errors.New("limit exceeded")
		}

		rel, e := filepath.Rel(dv.Path, s.Text())
		if e != nil {
			return fmt.Errorf("retrieve relative key: %w", e)
		}

		select {
		case files <- &FileMeta{
			Key: rel,
			Dir: false,
		}:
			count++
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}

// ListDirContents blocks until all contents listed, or a (optional) limit exceeded, it only returns that folder's files (no subfolder's files)
func (dv *Datakeeper) ListDirContents(ctx context.Context, prefix string, limit int, files chan<- *FileMeta) error {
	panic("unimplemented")
}

// ListContents blocks until all contents listed, or a (optional) limit exceeded, including the files under its subfolders
func (dv *Datakeeper) ListContents(ctx context.Context, prefix string, limit int, files chan<- *FileMeta) error {
	return dv.listObjects(ctx, prefix, limit, files)
}

// ObjectReader reads remote object content
func (dv *Datakeeper) ObjectReader(ctx context.Context, key string, size int64) (io.Reader, error) {
	r, e := newHTTPReader(ctx, dv, key, size)
	return r, e
}

// ObjectWriter writes content to remote object
func (dv *Datakeeper) ObjectWriter(ctx context.Context, key string, size int64, overwrite bool) (io.Writer, error) {
	key = filepath.Join(dv.Path, key)
	cli, e := dv.newClient(&identity.FilePolicy{
		Op:        identity.FileOpUpload,
		Bucket:    dv.Bucket,
		Prefix:    key,
		Overwrite: overwrite,
	}, longRequestTTL)
	if e != nil {
		return nil, e
	}

	return files.NewHTTPFormWriter(ctx, &files.FormUploadReq{
		Endpoint: dv.Endpoint,
		Forms:    map[string]string{formFieldKey: key},
		Filename: filepath.Base(key),
		Client:   cli,
	})
}

type httpError struct {
	code    int
	message string
}

func (e *httpError) Error() string {
	return fmt.Sprintf("unexpected response %d: %s", e.code, e.message)
}

func (dv *Datakeeper) doSimpleRequest(ctx context.Context, method string, policy *identity.FilePolicy, ttl time.Duration, params map[string]string) (*http.Response, error) {
	req, e := dv.newRequest(ctx, method, params)
	if e != nil {
		return nil, e
	}
	cli, e := dv.newClient(policy, ttl)
	if e != nil {
		return nil, e
	}

	resp, e := cli.Do(req)
	if e != nil {
		return nil, fmt.Errorf("request datakeeper: %w", e)
	}
	if code := resp.StatusCode; code/100 != 2 {
		buf, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, &httpError{code: resp.StatusCode, message: string(buf)}
	}

	return resp, nil
}

func (dv *Datakeeper) newRequest(ctx context.Context, method string, params map[string]string) (*http.Request, error) {
	values := make(url.Values)
	for k, v := range params {
		values.Set(k, v)
	}
	u := fmt.Sprintf("%s?%s", dv.Endpoint, values.Encode())
	return http.NewRequestWithContext(ctx, method, u, nil)
}

func (dv *Datakeeper) newClient(policy *identity.FilePolicy, ttl time.Duration) (*http.Client, error) {
	t, e := dv.signer(policy, ttl)
	if e != nil {
		return nil, fmt.Errorf("sign datakeeper token: %w", e)
	}

	return &http.Client{
		Timeout: ttl,
		Transport: &datakeeperTransport{
			rt:    http.DefaultTransport,
			token: t,
		},
	}, nil
}

type datakeeperTransport struct {
	rt    http.RoundTripper
	token string
}

func (dt *datakeeperTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+dt.token)
	return dt.rt.RoundTrip(req)
}
