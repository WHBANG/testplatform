package apiserver

import (
	"context"

	ava_grpc "git.supremind.info/products/atom/com/grpc"
	"git.supremind.info/products/atom/com/identity"
	"google.golang.org/grpc/metadata"
)

func WithUserInfo(ctx context.Context, userinfo *identity.UserIdentity) context.Context {
	return metadata.AppendToOutgoingContext(ctx,
		ava_grpc.EmailKey, userinfo.Email,
		ava_grpc.UsernameKey, userinfo.Username,
	)
}

// deprecated, will be removed in Augest
func WithPrivateToken(ctx context.Context, token string) context.Context {
	return metadata.AppendToOutgoingContext(ctx,
		ava_grpc.AuthorizationKey, ava_grpc.PrivateTokenKey+" "+token,
	)
}

func WithIDToken(ctx context.Context, token string) context.Context {
	return metadata.AppendToOutgoingContext(ctx,
		ava_grpc.AuthorizationKey, ava_grpc.IDTokenKey+" "+token,
	)
}
