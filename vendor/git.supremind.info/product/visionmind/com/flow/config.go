package flow

import (
	"context"
	"time"
)

type Config struct {
	ID         string             `json:"id" bson:"id"`
	Namespace  string             `json:"ns" bson:"ns"` // Config 所属的 VMR namespace
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
	Ver        int64              `json:"ver" bson:"ver"`
	Info       interface{}        `json:"info" bson:"info"`
	CancelFunc context.CancelFunc `json:"-" bson:"-"` // for cancel
}

type IConfiger interface {
	// 设置 / 删除 / 查询
	Set(ctx context.Context, id string, info interface{}) error
	Del(ctx context.Context, id string) error
	Get(ctx context.Context, id *string) ([]Config, error)
	// 检查配置更新
	Check(ctx context.Context, id string, ver int64) (*Config, error)
}
