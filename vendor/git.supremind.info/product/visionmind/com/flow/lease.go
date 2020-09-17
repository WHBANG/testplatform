package flow

import (
	"context"
	"time"
)

type Lease struct {
	ID         string             `json:"id" bson:"id"`
	Status     string             `json:"status" bson:"status"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
	SwitchedAt time.Time          `json:"switched_at" bson:"switched_at"`
	Ver        int64              `json:"ver" bson:"ver"`
	Holder     interface{}        `json:"holder" bson:"holder"`
	CancelFunc context.CancelFunc `json:"-" bson:"-"` // for cancel
}

type IGrantor interface {
	// 设置 / 删除
	Set(ctx context.Context, lease *Lease, info map[string]interface{}) error
	Start(ctx context.Context, id string) error
	Stop(ctx context.Context, id string) error
	Del(ctx context.Context, id string) error
	// 抢占 / 持有 / 释放
	Get(ctx context.Context, selector map[string]interface{}, holder interface{}) (map[string]interface{}, error)
	GetWithWeight(ctx context.Context, selector map[string]interface{}, holder interface{}, weight int) (map[string]interface{}, error)
	Hold(ctx context.Context, id string, ver int64, holder interface{}, extra ...string) (map[string]interface{}, error)
	Release(ctx context.Context, id string, ver int64, extra ...string) error
}
