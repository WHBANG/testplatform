package apiserver

import (
	"context"
	"io"
	"net/http"
	"path/filepath"

	"git.supremind.info/products/atom/com/files"
	ossutil "git.supremind.info/products/atom/com/oss"
	datakeeper "git.supremind.info/products/atom/datakeeper/client"
	"git.supremind.info/products/atom/proto/go/api"
	"github.com/pkg/errors"
	"github.com/qiniu/api.v7/v7/storage"
)

const (
	kodoUpHost = "http://up.qiniup.com"

	kodoFormUploadThreshold = 64 << 20
	kodoFormUploadMaxSize   = 1 << 30
)

// Uploader uploads anything to a volume
type Uploader interface {
	Upload(ctx context.Context, r io.Reader, key string, size int64) error
}

func NewUploader(vol *api.Volume, cli api.VolumeServiceClient, prefix string, overwrite bool) (Uploader, error) {
	f := &uploaderFactory{
		cli:       cli,
		vol:       vol,
		prefix:    prefix,
		overwrite: overwrite,
	}

	switch vol.GetSpec().GetVendor() {
	case api.ResourceVolume_KODO:
		return f.newKodoUploader()
	case api.ResourceVolume_Local:
		return f.newKeeperUploader()
	case api.ResourceVolume_Minio:
		return f.newMinioUploader()
	case api.ResourceVolume_OSS:
		return f.newOSSUploader()
	default:
		return nil, errors.New("invalid volume vendor")
	}
}

type uploaderFactory struct {
	cli       api.VolumeServiceClient
	vol       *api.Volume
	overwrite bool
	prefix    string
}

type kodoUploader struct {
	cli       api.VolumeServiceClient
	vol       *api.Volume
	overwrite bool
	prefix    string // without path
}

func (f *uploaderFactory) getVolumeCredential(ctx context.Context) (*api.Volume_Credential, error) {
	return f.cli.GetVolumeCredential(ctx, &api.GetVolumeCredentialReq{
		Name:      f.vol.GetName(),
		Creator:   f.vol.GetCreator(),
		Operation: api.Volume_Credential_Upload,
		Prefix:    f.prefix,
		Overwrite: f.overwrite,
	})
}

func (f *uploaderFactory) newKodoUploader() (*kodoUploader, error) {
	up := &kodoUploader{
		cli:       f.cli,
		vol:       f.vol,
		overwrite: f.overwrite,
		prefix:    f.prefix,
	}
	return up, nil
}

func (u *kodoUploader) Upload(ctx context.Context, r io.Reader, key string, size int64) error {
	token, e := u.getToken(ctx, key)
	if e != nil {
		return e
	}

	if size > kodoFormUploadThreshold {
		if ra, ok := r.(io.ReaderAt); ok {
			return u.multiBlockUpload(ctx, ra, token, key, size)
		}
	}

	if size > kodoFormUploadMaxSize {
		return errors.New("file larger than 1GiB could not be uploaded using http formdata to kodo")
	}

	return u.formUpload(ctx, r, token, key, size)
}

func (u *kodoUploader) formUpload(ctx context.Context, r io.Reader, token, key string, size int64) error {
	fu := storage.NewFormUploader(&storage.Config{
		UseHTTPS: false,
		UpHost:   kodoUpHost,
	})
	ret := &storage.PutRet{}
	putExt := &storage.PutExtra{
		UpHost: kodoUpHost,
	}
	if size <= 0 {
		size = -1
	}

	return fu.Put(ctx, ret, token, filepath.Join(u.vol.GetSpec().GetPath(), u.prefix, key), r, size, putExt)
}

func (u *kodoUploader) multiBlockUpload(ctx context.Context, r io.ReaderAt, token, key string, size int64) error {
	ru := storage.NewResumeUploader(&storage.Config{
		UseHTTPS: false,
		UpHost:   kodoUpHost,
	})
	ret := &storage.PutRet{}
	putExt := &storage.RputExtra{
		UpHost: kodoUpHost,
	}

	return ru.Put(ctx, ret, token, filepath.Join(u.vol.GetSpec().GetPath(), u.prefix, key), r, size, putExt)
}

func (u *kodoUploader) getToken(ctx context.Context, key string) (string, error) {
	res, e := u.cli.GetVolumeCredential(ctx, &api.GetVolumeCredentialReq{
		Name:      u.vol.GetName(),
		Creator:   u.vol.GetCreator(),
		Operation: api.Volume_Credential_Upload,
		Key:       filepath.Join(u.prefix, key),
		Overwrite: u.overwrite,
	})
	if e != nil {
		return "", errors.Wrap(e, "failed to get kodo upload token")
	}

	return res.GetKodo().GetUploadToken(), nil
}

type keeperUploader struct {
	client *datakeeper.Client
	prefix string
}

func (f *uploaderFactory) newKeeperUploader() (*keeperUploader, error) {
	cred, e := f.getVolumeCredential(context.TODO())
	if e != nil {
		return nil, errors.Wrap(e, "failed to get volume credential")
	}
	if cred.GetLocal() == nil {
		return nil, errors.New("invailable data keeper token")
	}
	up := &keeperUploader{
		client: &datakeeper.Client{
			Endpoint:   f.vol.GetSpec().GetEndpoint(),
			Token:      cred.GetLocal().GetToken(),
			HTTPClient: http.DefaultClient,
		},
		prefix: filepath.Join(f.vol.GetSpec().GetPath(), f.prefix),
	}
	return up, nil
}

func (u *keeperUploader) Upload(ctx context.Context, r io.Reader, key string, size int64) error {
	if e := u.client.Upload(ctx, filepath.Join(u.prefix, key), r); e != nil {
		return errors.Wrap(e, "upload file failed")
	}

	return nil
}

type minioUploader struct {
	presignedPostURL string
	formData         map[string]string
	prefix           string
}

func (f *uploaderFactory) newMinioUploader() (*minioUploader, error) {
	cred, e := f.getVolumeCredential(context.TODO())
	if e != nil {
		return nil, errors.Wrap(e, "failed to get volume credential")
	}
	if cred.GetMinio() == nil {
		return nil, errors.New("invailabel minio credential")
	}
	up := &minioUploader{
		presignedPostURL: cred.GetMinio().GetPresignedPostURL(),
		formData:         cred.GetMinio().GetFormData(),
		prefix:           filepath.Join(f.vol.GetSpec().GetPath(), f.prefix),
	}
	return up, nil
}

func (u *minioUploader) Upload(ctx context.Context, r io.Reader, key string, size int64) error {
	forms := make(map[string]string, len(u.formData))
	for k, v := range u.formData {
		if k == "key" {
			v = filepath.Join(v, key)
		}
		forms[k] = v
	}

	// 0 size is invalid for minio

	return files.StreamingFormUpload(ctx, &files.FormUploadReq{
		Endpoint: u.presignedPostURL,
		Forms:    forms,
		Filename: filepath.Base(key),
		Reader:   r,
		Size:     size,
	})
}

type ossUploader struct {
	presignedURL string
	formData     map[string]string
	prefix       string
}

func (f *uploaderFactory) newOSSUploader() (*ossUploader, error) {
	cred, e := f.getVolumeCredential(context.TODO())
	if e != nil {
		return nil, errors.Wrap(e, "failed to get volume credential")
	}
	if cred.GetOss() == nil {
		return nil, errors.New("invailable oss credential")
	}
	up := &ossUploader{
		presignedURL: cred.GetOss().GetPresignedURL(),
		formData:     cred.GetOss().GetFormData(),
		prefix:       filepath.Join(f.vol.GetSpec().GetPath(), f.prefix),
	}
	return up, nil
}

func (u *ossUploader) Upload(ctx context.Context, r io.Reader, key string, size int64) error {
	forms := make(map[string]string, len(u.formData)+1)
	for k, v := range u.formData {
		forms[k] = v
	}
	forms[ossutil.FormKeyKey] = filepath.Join(forms[ossutil.FormKeyPrefix], key)

	if size <= 0 {
		size = -1
	}

	return files.StreamingFormUpload(ctx, &files.FormUploadReq{
		Endpoint: u.presignedURL,
		Forms:    forms,
		Filename: filepath.Base(key),
		Reader:   r,
		Size:     size,
	})
}
