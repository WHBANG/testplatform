package oss

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

const (
	FormKeyKey           = "key"
	FormKeyPrefix        = "prefix"
	FormKeyPolicy        = "policy"
	FormkeySignature     = "signature"
	FormKeyAccessID      = "OSSAccessKeyId"
	FormKeySuccessStatus = "success_action_status"
	FormKeyCallback      = "callback"

	defaultTTL = 7 * 24 * time.Hour
)

type PostPolicy struct {
	Expiration time.Time  `json:"expiration,omitempty"`
	Conditions [][]string `json:"conditions,omitempty"`
}

type PostOption func(*PostPolicy)

type CallbackParam struct {
	CallbackUrl      string `json:"callbackUrl"`
	CallbackBody     string `json:"callbackBody"`
	CallbackBodyType string `json:"callbackBodyType"`
}
type Signer struct {
	AccessKeyID     string
	AccessKeySecret string
	Endpoint        string
}

func (s *Signer) SignPostPolicy(bucketName string, callback *CallbackParam, opts ...PostOption) (*url.URL, map[string]string, error) {
	host := fmt.Sprintf("http://%s.%s", bucketName, s.Endpoint)
	u, e := url.Parse(host)
	if e != nil {
		return nil, nil, errors.Wrap(e, "failed to parse endpoint")
	}

	policy := &PostPolicy{}
	for _, opt := range opts {
		opt(policy)
	}
	encodedPolicy, signedPolicy, e := policy.sign(s.AccessKeySecret)
	if e != nil {
		return nil, nil, e
	}

	forms := map[string]string{
		FormKeyAccessID:      s.AccessKeyID,
		FormKeyPolicy:        encodedPolicy,
		FormkeySignature:     signedPolicy,
		FormKeySuccessStatus: "200",
	}

	if callback != nil {
		c, e := callback.format()
		if e != nil {
			return nil, nil, e
		}
		forms[FormKeyCallback] = c
	}

	return u, forms, nil
}

func (c *CallbackParam) format() (string, error) {
	marshalled, e := json.Marshal(c)
	if e != nil {
		return "", errors.Wrap(e, "failed to marshal callback params")
	}

	return base64.StdEncoding.EncodeToString(marshalled), nil
}

func (p *PostPolicy) sign(accessSecret string) (string, string, error) {
	if p.Expiration.IsZero() {
		SetTTL(defaultTTL)(p)
	}

	marshalled, e := json.Marshal(p)
	if e != nil {
		return "", "", errors.Wrap(e, "failed to marshal post policy")
	}
	encoded := base64.StdEncoding.EncodeToString(marshalled)

	h := hmac.New(sha1.New, []byte(accessSecret))
	if _, e := h.Write([]byte(encoded)); e != nil {
		return "", "", errors.Wrap(e, "sign post policy failed")
	}
	signed := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return encoded, signed, nil
}

func SetKey(k string) PostOption {
	return func(p *PostPolicy) {
		p.Conditions = append(p.Conditions, []string{"eq", "$key", k})
	}
}

func SetPrefix(prefix string) PostOption {
	return func(p *PostPolicy) {
		p.Conditions = append(p.Conditions, []string{"starts-with", "$key", prefix})
	}
}

func SetTTL(d time.Duration) PostOption {
	return func(p *PostPolicy) {
		p.Expiration = time.Now().Add(d).UTC()
	}
}

func SetDeadline(t time.Time) PostOption {
	return func(p *PostPolicy) {
		p.Expiration = t.UTC()
	}
}
