package go_sdk

import "context"

type Quota struct {
	Name       string `json:"name"`
	Quota      int    `json:"quota"`
	ExpireTime string `json:"expire_time"`
}

type ICAClient interface {
	Start(context.Context)  // 启动循环
	CheckQuota(string) bool // 单次quota检查
	FetchQuotas() (quotas []Quota, err error)
}
