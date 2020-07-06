package grpc

import (
	"context"

	"git.supremind.info/products/atom/com/identity"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
)

type AVAKey string

const (
	EmailKey         string = "auth.email"
	UsernameKey      string = "auth.username"
	UpstreamTokenKey string = "auth.upstream_token"
	AuthorizationKey string = "authorization"
	PrivateTokenKey  string = "private"
	IDTokenKey       string = "id_token"
)

// GetUserIdentity get user credential from request context
func GetUserIdentity(ctx context.Context) (*identity.UserIdentity, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("bad metadata")
	}
	emails := md.Get(EmailKey)
	usernames := md.Get(UsernameKey)
	tok := md.Get(UpstreamTokenKey)

	if len(emails) < 1 || len(usernames) < 1 {
		return nil, errors.New("no user info in metadata")
	}

	user := &identity.UserIdentity{Email: emails[0], Username: usernames[0]}
	if len(tok) > 0 {
		user.UpstreamToken = tok[0]
	}

	return user, nil
}

func GetIDToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("bad metadata")
	}
	res := md.Get(IDTokenKey)
	if len(res) < 1 {
		return "", errors.New("no id token in metadata")
	}

	return res[0], nil
}
