package identity

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

type FileOp int

const (
	FileOpUnknown FileOp = iota
	FileOpUpload
	FileOpDownload
	FileOpList
	FileOpDelete
)

func (t FileOp) String() string {
	switch t {
	case FileOpUpload:
		return "Upload"
	case FileOpDownload:
		return "Download"
	case FileOpList:
		return "List"
	case FileOpDelete:
		return "Delete"
	default:
	}
	return "Unknown"
}

type FilePolicy struct {
	Op        FileOp `json:"op"`
	Bucket    string `json:"bucket"`
	Prefix    string `json:"prefix"`
	Overwrite bool   `json:"overwrite"`
}

type HSJWT struct {
	issuer   string
	audience []string
	key      []byte

	signer jose.Signer
}

func NewHSJWT(issuer string, audience []string, key []byte) (*HSJWT, error) {
	sig, e := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: key}, nil)
	if e != nil {
		return nil, e
	}
	return &HSJWT{
		issuer:   issuer,
		audience: audience,
		signer:   sig,
		key:      key,
	}, nil
}

func (t *HSJWT) Issue(sub string, ttl time.Duration, payload interface{}) (string, error) {
	now := time.Now()
	cl := &jwt.Claims{
		Issuer:   t.issuer,
		Subject:  sub,
		Audience: jwt.Audience(t.audience),
		IssuedAt: jwt.NewNumericDate(now),
		ID:       uuid.New().String(),
	}
	if ttl > 0 {
		cl.Expiry = jwt.NewNumericDate(now.Add(ttl))
	}
	return jwt.Signed(t.signer).Claims(cl).Claims(payload).CompactSerialize()
}

func (t *HSJWT) Verify(raw string, payload interface{}) error {
	tok, e := jwt.ParseSigned(raw)
	if e != nil {
		return e
	}

	var cl jwt.Claims
	if e := tok.Claims(t.key, &payload, &cl); e != nil {
		return e
	}

	if e := cl.Validate(jwt.Expected{
		Issuer:   t.issuer,
		Audience: jwt.Audience(t.audience),
		Time:     time.Now(),
	}); e != nil {
		return e
	}

	return nil
}
