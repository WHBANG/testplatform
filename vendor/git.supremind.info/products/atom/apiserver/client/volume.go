package apiserver

import (
	"context"
	"fmt"
	"time"

	"git.supremind.info/products/atom/com/identity"
	"git.supremind.info/products/atom/com/volume"
	"git.supremind.info/products/atom/proto/go/api"
	"github.com/pkg/errors"
)

func GetStorageService(ctx context.Context, name, creator string,
	volClient api.VolumeServiceClient,
	secClient api.SecretServiceClient,
	signer volume.DatakeeperPolicySigner) (volume.StorageService, error) {

	vol, e := volClient.GetVolume(ctx, &api.GetVolumeReq{Name: name, Creator: creator})
	if e != nil {
		return nil, fmt.Errorf("get volume detail: %w", e)
	}

	spec := volume.Spec{
		Bucket:   vol.GetSpec().GetBucket(),
		Path:     vol.GetSpec().GetPath(),
		Endpoint: vol.GetSpec().GetEndpoint(),
	}

	var sec *api.Secret
	if vol.GetSpec().GetVendor() != api.ResourceVolume_Local {
		ref := vol.GetSpec().GetSecret()
		sec, e = secClient.GetSecret(ctx, &api.GetSecretReq{Name: ref.GetName(), Creator: ref.GetCreator()})
		if e != nil {
			return nil, fmt.Errorf("get secret: %w", e)
		}
	}

	switch vol.GetSpec().GetVendor() {
	case api.ResourceVolume_KODO:
		qiniu := sec.GetSpec().GetQiniu()
		return volume.NewKODO(spec, qiniu.GetAccessKey(), qiniu.GetSecretKey())
	case api.ResourceVolume_OSS:
		aliyun := sec.GetSpec().GetAliyun()
		return volume.NewOSS(spec, aliyun.GetAccessKeyID(), aliyun.GetAccessKeySecret())
	case api.ResourceVolume_Minio:
		didiyun := sec.GetSpec().GetDidiyun()
		return volume.NewMinio(spec, didiyun.GetSecretID(), didiyun.GetSecretKey())
	case api.ResourceVolume_Local:
		if signer == nil {
			signer = serverSideTokenSigner(ctx, name, creator, volClient)
		}
		return volume.NewDatakeeper(spec, signer), nil
	}

	return nil, errors.New("unsupported volume vendor")
}

func serverSideTokenSigner(ctx context.Context, name, creator string, cli api.VolumeServiceClient) volume.DatakeeperPolicySigner {
	return func(policy *identity.FilePolicy, ttl time.Duration) (string, error) {
		var op api.Volume_Credential_Operation
		switch policy.Op {
		case identity.FileOpUpload:
			op = api.Volume_Credential_Upload
		case identity.FileOpDownload:
			op = api.Volume_Credential_Download
		case identity.FileOpList:
			op = api.Volume_Credential_List
		case identity.FileOpDelete:
			op = api.Volume_Credential_Delete
		}

		res, e := cli.GetVolumeCredential(ctx, &api.GetVolumeCredentialReq{
			Name:      name,
			Creator:   creator,
			Operation: op,
			Prefix:    policy.Prefix,
		})
		if e != nil {
			return "", e
		}

		return res.GetLocal().GetToken(), nil
	}
}
