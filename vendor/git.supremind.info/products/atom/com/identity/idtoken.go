package identity

import (
	"crypto/rsa"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

type Issuer interface {
	Issue(sub string, u *UserIdentity) (string, error)
	Verifier() Verifier
}

type Verifier interface {
	Verify(raw string) (*UserIdentity, error)
}

// (RSA signed) IDTokenIssuer
type IDTokenIssuer struct {
	Issuer   string
	Audience []string
	TTL      time.Duration
	Key      *rsa.PrivateKey

	signer jose.Signer
}

func (i *IDTokenIssuer) Init() error {
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: i.Key}, nil)
	if err != nil {
		return errors.Wrap(err, "create jws signer failed")
	}

	i.signer = sig
	return nil
}

func (i *IDTokenIssuer) Issue(sub string, u *UserIdentity) (string, error) {
	now := time.Now()
	cl := &jwt.Claims{
		Issuer:   i.Issuer,
		Subject:  sub,
		Audience: jwt.Audience(i.Audience),
		IssuedAt: jwt.NewNumericDate(now),
		Expiry:   jwt.NewNumericDate(now.Add(i.TTL)),
		ID:       uuid.New().String(),
	}
	return jwt.Signed(i.signer).Claims(cl).Claims(u).CompactSerialize()
}

func (i *IDTokenIssuer) Verifier() Verifier {
	return &IDTokenVerifier{
		Key:      &i.Key.PublicKey,
		Issuer:   i.Issuer,
		Audience: i.Audience,
	}
}

// (RSA signed) IDTokenVerifier
type IDTokenVerifier struct {
	Key      *rsa.PublicKey
	Issuer   string
	Audience []string
}

func (i *IDTokenVerifier) Verify(raw string) (*UserIdentity, error) {
	tok, err := jwt.ParseSigned(raw)
	if err != nil {
		return nil, errors.Wrap(err, "parse jws failed")
	}

	var u UserIdentity
	var cl jwt.Claims
	if err := tok.Claims(i.Key, &u, &cl); err != nil {
		return nil, errors.Wrap(err, "extract claim failed")
	}

	err = cl.Validate(jwt.Expected{
		Issuer:   i.Issuer,
		Audience: jwt.Audience(i.Audience),
		Time:     time.Now(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "invalid claims")
	}

	u.Username = trimQuotes(u.Username)
	u.Email = trimQuotes(u.Email)
	u.UpstreamToken = trimQuotes(u.UpstreamToken)

	return &u, nil
}

func trimQuotes(s string) string {
	return strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)
}

// UnsafeVerifyIDToken checks experation only
func UnsafeVerifyIDToken(raw string) error {
	tok, err := jwt.ParseSigned(raw)
	if err != nil {
		return errors.Wrap(err, "parse jws failed")
	}

	var cl jwt.Claims
	if err := tok.UnsafeClaimsWithoutVerification(&cl); err != nil {
		return errors.Wrap(err, "extract claim failed")
	}

	if err := cl.Validate(jwt.Expected{Time: time.Now()}); err != nil {
		return errors.Wrap(err, "invalid claims")
	}

	return nil
}

func UnsafeParseIDToken(raw string) (*UserIdentity, error) {
	tok, err := jwt.ParseSigned(raw)
	if err != nil {
		return nil, errors.Wrap(err, "parse jws failed")
	}

	var user UserIdentity
	if err := tok.UnsafeClaimsWithoutVerification(&user); err != nil {
		return nil, errors.Wrap(err, "extract user info failed")
	}

	return &user, nil
}
