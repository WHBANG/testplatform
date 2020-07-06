package apiserver

import (
	"context"
	"crypto/x509"
	"time"

	ava_grpc "git.supremind.info/products/atom/com/grpc"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
)

// Option is a client option for connecting to apiserver
type Option func(*Clientset) error

// WithCert setups connection using given tls cert file
func WithCert(certFile string, serverName string) Option {
	return func(cs *Clientset) error {
		creds, e := credentials.NewClientTLSFromFile(certFile, serverName)
		if e != nil {
			return errors.Wrap(e, "create client tls from file failed")
		}
		cs.dialOpts = append(cs.dialOpts, grpc.WithTransportCredentials(creds))
		return nil
	}
}

func WithInsecure() Option {
	return func(cs *Clientset) error {
		cs.dialOpts = append(cs.dialOpts, grpc.WithInsecure())
		return nil
	}
}

func WithDefaultCertPool() Option {
	return func(cs *Clientset) error {
		pool, err := x509.SystemCertPool()
		if err != nil {
			return errors.Wrap(err, "get system cert pool failed")
		}
		cs.dialOpts = append(cs.dialOpts,
			grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(pool, "")))
		return nil
	}
}

func WithLogger(l *zap.Logger) Option {
	return func(cs *Clientset) error {
		logOpts := []grpc_zap.Option{grpc_zap.WithLevels(ava_grpc.CodeToLevel)}
		cs.streamInterceptors = append(cs.streamInterceptors,
			grpc_zap.StreamClientInterceptor(l, logOpts...))
		cs.unaryInterceptors = append(cs.unaryInterceptors,
			grpc_zap.UnaryClientInterceptor(l, logOpts...))
		return nil
	}
}

func WithConstantIDToken(t string) Option {
	return func(cs *Clientset) error {
		cs.dialOpts = append(cs.dialOpts, grpc.WithPerRPCCredentials(&IDTokenSource{raw: t}))
		return nil
	}
}

type IDTokenSource struct {
	raw string
}

func (t *IDTokenSource) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		ava_grpc.AuthorizationKey: ava_grpc.IDTokenKey + " " + t.raw,
	}, nil
}

func (t *IDTokenSource) RequireTransportSecurity() bool {
	return false
}

func WithRetry(initBackoff time.Duration) Option {
	return func(cs *Clientset) error {
		opts := []grpc_retry.CallOption{
			grpc_retry.WithBackoff(grpc_retry.BackoffExponential(initBackoff)),
			grpc_retry.WithCodes(codes.Unavailable),
		}
		cs.streamInterceptors = append(cs.streamInterceptors,
			grpc_retry.StreamClientInterceptor(opts...))
		cs.unaryInterceptors = append(cs.unaryInterceptors,
			grpc_retry.UnaryClientInterceptor(opts...))
		return nil
	}
}

func SetMaxRecvMsgSize(bytes int) Option {
	return func(cs *Clientset) error {
		cs.streamInterceptors = append(cs.streamInterceptors,
			callOptionStreamInterceptor(grpc.MaxCallRecvMsgSize(bytes)),
		)
		cs.unaryInterceptors = append(cs.unaryInterceptors,
			callOptionUnaryInterceptor(grpc.MaxCallRecvMsgSize(bytes)),
		)
		return nil
	}
}

func SetMaxSendMsgSize(bytes int) Option {
	return func(cs *Clientset) error {
		cs.streamInterceptors = append(cs.streamInterceptors,
			callOptionStreamInterceptor(grpc.MaxCallSendMsgSize(bytes)),
		)
		cs.unaryInterceptors = append(cs.unaryInterceptors,
			callOptionUnaryInterceptor(grpc.MaxCallSendMsgSize(bytes)),
		)
		return nil
	}
}

func callOptionStreamInterceptor(opt grpc.CallOption) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		return streamer(ctx, desc, cc, method, append(opts, opt)...)
	}
}

func callOptionUnaryInterceptor(opt grpc.CallOption) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return invoker(ctx, method, req, reply, cc, append(opts, opt)...)
	}
}
