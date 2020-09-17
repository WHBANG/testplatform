// +build !ca

package client

import (
	"context"

	"git.supremind.info/product/visionmind/com/go_sdk"
	mgoutil "github.com/qiniu/db/mgoutil.v3"
)

type CAClientConfig struct {
	CaServerHost string `json:"host"`
}

type CAClient struct {
}

func NewCAClient(cfg *CAClientConfig, tasksColl *mgoutil.Collection) (*CAClient, error) {
	return &CAClient{}, nil
}

func (c *CAClient) FetchQuotas() (quotas []go_sdk.Quota, err error) {
	return nil, nil
}

func (c *CAClient) Start(ctx context.Context) {
}

func (c *CAClient) CheckQuota(name string) (enough bool) {
	return true
}
