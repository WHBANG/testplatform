package datakeeper

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"

	"git.supremind.info/products/atom/com/files"
	"github.com/pkg/errors"
)

type Client struct {
	Token      string
	Endpoint   string
	HTTPClient *http.Client
}

const (
	formFieldFile = "file"
	formFieldKey  = "key"

	queryFieldPrefix = "prefix"
	queryFieldKey    = "key"
	queryFieldToken  = "token"
	queryFieldLimit  = "limit"
)

func (c *Client) Upload(ctx context.Context, key string, r io.Reader) error {
	return files.StreamingFormUpload(ctx, &files.FormUploadReq{
		Endpoint: c.Endpoint,
		Forms:    map[string]string{formFieldKey: key},
		Header:   url.Values{"Authorization": []string{"Bearer " + c.Token}},
		Filename: filepath.Base(key),
		Reader:   r,
		Client:   c.HTTPClient,
	})
}

func (c *Client) Download(ctx context.Context, key string) (io.ReadCloser, error) {
	return c.simpleRequest(ctx, "GET", map[string]string{queryFieldKey: key})
}

func (c *Client) DownloadURL(key string) string {
	values := make(url.Values)
	values.Set(queryFieldKey, key)
	values.Set(queryFieldToken, c.Token)
	return fmt.Sprintf("%s?%s", c.Endpoint, values.Encode())
}

func (c *Client) DeleteDirectory(ctx context.Context, dir string) error {
	r, e := c.simpleRequest(ctx, "DELETE", map[string]string{queryFieldPrefix: dir})
	if e != nil {
		return errors.Wrap(e, "delete directory failed")
	}

	r.Close()
	return nil
}

func (c *Client) ListFiles(ctx context.Context, dir string, limit int, files chan<- string) error {
	r, e := c.simpleRequest(ctx, "GET", map[string]string{queryFieldPrefix: dir, queryFieldLimit: strconv.Itoa(limit)})
	if e != nil {
		return e
	}
	defer r.Close()

	s := bufio.NewScanner(r)
	for s.Scan() {
		select {
		case files <- s.Text():
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return s.Err()
}

func (c *Client) simpleRequest(ctx context.Context, method string, params map[string]string) (io.ReadCloser, error) {
	u, e := url.Parse(c.Endpoint)
	if e != nil {
		return nil, e
	}
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req, e := http.NewRequest(method, u.String(), nil)
	if e != nil {
		return nil, e
	}

	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, e := c.HTTPClient.Do(req)
	if e != nil {
		return nil, e
	}

	r := &responseReadCloser{resp.Body}
	if code := resp.StatusCode; code/100 != 2 {
		body, _ := ioutil.ReadAll(resp.Body)
		r.Close()
		return nil, fmt.Errorf("unsuccessful request, got [%d] %s: %s", code, http.StatusText(code), string(body))
	}

	return r, nil
}

// override close method of http response body
type responseReadCloser struct {
	io.ReadCloser
}

func (r *responseReadCloser) Close() error {
	io.Copy(ioutil.Discard, r.ReadCloser)
	return r.ReadCloser.Close()
}
