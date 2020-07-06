package apiserver

import (
	"git.supremind.info/products/atom/proto/go/api"
)

type Interface interface {
	Close() error

	Secret() api.SecretServiceClient
	Dataset() api.DatasetServiceClient
	Volume() api.VolumeServiceClient
	Watch() api.WatchServiceClient
	Access() api.AccessServiceClient
	Job() api.JobServiceClient
	DeviceCategory() api.DeviceCategoryServiceClient
	Ore() api.OreServiceClient
	Package() api.PackageServiceClient
}
