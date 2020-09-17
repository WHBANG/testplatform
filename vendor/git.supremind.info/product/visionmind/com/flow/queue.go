package flow

import (
	"context"
	"time"
)

var (
	DefaultClearConfig = ClearConfig{
		StorageDuration: time.Hour * 168,
		CheckDuration:   time.Hour,
	}
)

// ReadOnly
type Message struct {
	ID   int64       `json:"id" bson:"id"` // 日志唯一ID
	Time time.Time   `json:"time,omitempty" bson:"time"`
	Body interface{} `json:"body,omitempty" bson:"body"`
}

type ClearConfig struct {
	StorageDuration time.Duration `json:"storage_duration"`
	CheckDuration   time.Duration `json:"check_duration"`
}

type IQueue interface {
	Push(ctx context.Context, topic string, body interface{}) (int64, error)
	// pull by ID
	PullByID(ctx context.Context, topic string, id int64) (*Message, error)
	PullBatch(ctx context.Context, topic string, begid, endid int64) ([]Message, error)
	// pull by group
	Pull(ctx context.Context, topic, group string, since int64) (*Message, error)
	FetchTopics(ctx context.Context, filter string) ([]string, error)
}

type IIDs interface {
	// push ID
	NextPushID(ctx context.Context, topic string) (int64, error)
	CurPushID(ctx context.Context, topic string) (int64, error)
	// pull ID
	GetPullID(ctx context.Context, topic, group string) (int64, error)
	SetPullID(ctx context.Context, topic, group string, preid, curid int64) error
}
