package apiserver

import (
	"git.supremind.info/products/atom/proto/go/api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var _ Interface = (*Clientset)(nil)

type Clientset struct {
	dialOpts           []grpc.DialOption
	streamInterceptors []grpc.StreamClientInterceptor
	unaryInterceptors  []grpc.UnaryClientInterceptor
	md                 metadata.MD
	conn               *grpc.ClientConn
}

func New(endpoint string, opts ...Option) (*Clientset, error) {
	cs := &Clientset{
		md: metadata.New(nil),
	}

	var e error
	for _, opt := range opts {
		e = opt(cs)
		if e != nil {
			return nil, errors.Wrap(e, "setting up apiserver client options failed")
		}
	}
	if len(cs.streamInterceptors) > 0 {
		cs.dialOpts = append(cs.dialOpts, grpc.WithChainStreamInterceptor(cs.streamInterceptors...))
	}
	if len(cs.unaryInterceptors) > 0 {
		cs.dialOpts = append(cs.dialOpts, grpc.WithChainUnaryInterceptor(cs.unaryInterceptors...))
	}

	conn, e := grpc.Dial(endpoint, cs.dialOpts...)
	if e != nil {
		return nil, errors.Wrap(e, "dial apiserver failed")
	}
	cs.conn = conn

	return cs, nil
}

func (cs *Clientset) Close() error {
	return cs.conn.Close()
}

func (cs *Clientset) Secret() api.SecretServiceClient {
	return api.NewSecretServiceClient(cs.conn)
}

func (cs *Clientset) Dataset() api.DatasetServiceClient {
	return api.NewDatasetServiceClient(cs.conn)
}

func (cs *Clientset) Volume() api.VolumeServiceClient {
	return api.NewVolumeServiceClient(cs.conn)
}

func (cs *Clientset) Watch() api.WatchServiceClient {
	return api.NewWatchServiceClient(cs.conn)
}

func (cs *Clientset) Access() api.AccessServiceClient {
	return api.NewAccessServiceClient(cs.conn)
}

func (cs *Clientset) Job() api.JobServiceClient {
	return api.NewJobServiceClient(cs.conn)
}

func (cs *Clientset) DeviceCategory() api.DeviceCategoryServiceClient {
	return api.NewDeviceCategoryServiceClient(cs.conn)
}

func (cs *Clientset) Ore() api.OreServiceClient {
	return api.NewOreServiceClient(cs.conn)
}

func (cs *Clientset) Package() api.PackageServiceClient {
	return api.NewPackageServiceClient(cs.conn)
}
