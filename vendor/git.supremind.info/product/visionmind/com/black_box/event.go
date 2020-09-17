package black_box

import (
	"context"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	DefaultPage    = 1
	DefaultPerPage = 100
)

type Event struct {
	ID             bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Product        string        `json:"product" bson:"product"`                // 产品名称
	ProductVersion string        `json:"productVersion" bson:"product_version"` // 产品版本
	Service        string        `json:"service" bson:"service"`                // 具体服务名称
	Meta           Meta          `json:"meta" bson:"meta"`                      // 事件的 meta 信息，用于事件的检索
	Time           int64         `json:"time,omitempty" bson:"time"`            // 事件发生的时间
	CreatedAt      time.Time     `json:"createdAt" bson:"created_at"`           // 事件入库的时间，若time不为0则以time为准，否则以入库时间为准
	Message        interface{}   `json:"message" bson:"message"`                // 事件其他的详细信息
	Name           string        `json:"name" bson:"name"`                      // 事件名称
	Level          string        `json:"level" bson:"level"`                    // 事件级别 INFO/ERROR
}

type Meta struct {
	TaskID     string `json:"taskId,omitempty" bson:"task_id,omitempty"`          // 任务 ID
	StreamID   string `json:"streamId,omitempty" bson:"stream_id,omitempty"`      // 视频流 ID
	SessionID  string `json:"sessionId,omitempty" bson:"session_id,omitempty"`    // 流会话标识
	VioEventID string `json:"vioEventId,omitempty" bson:"vio_event_id,omitempty"` // 预警事件 ID
	Resource   string `json:"resource,omitempty" bson:"resource,omitempty"`       // 资源标记
}

type Message struct {
	Info   string `json:"info" bson:"info"`
	Reason string `json:"reason" bson:"reason"`
}

type IEvent interface {
	Set(ctx context.Context, event *Event) error
	Get(ctx context.Context, req *EventRequest) ([]Event, int, error)
}

var (
	DefaultClearConfig = ClearConfig{
		StorageDuration: 72,
		CheckDuration:   24,
	}
)

type ClearConfig struct {
	StorageDuration int `json:"storageDuration"`
	CheckDuration   int `json:"checkDuration"`
}

type CoolDownConfig struct {
	Duration  int `json:"duration"`
	AllowList []struct {
		// 暂定这2个字段
		Product string `json:"product"`
		Service string `json:"service"`
	} `json:"allowList"`
}

type EventRequest struct {
	Page      int `json:"page" query:"page"`
	PerPage   int `json:"perPage" query:"perPage"`
	StartTime int `json:"startTime" query:"startTime"`
	EndTime   int `json:"endTime" query:"endTime"`

	Products       []string `json:"products" query:"products"`
	ProductVersion string   `json:"productVersion" query:"productVersion"`
	Service        string   `json:"service" query:"service"`
	Name           string   `json:"name" query:"name"`
	Levels         []string `json:"levels" query:"levels"`

	TaskID     string `json:"taskId" query:"taskId"`
	StreamID   string `json:"streamId" query:"streamId"`
	VioEventID string `json:"vioEventId" query:"vioEventId"`
	Resource   string `json:"resource"  query:"resource"`

	Marker string `json:"marker"  query:"marker"`
}

type EventResponse struct {
	Page      int     `json:"page"`
	PerPage   int     `json:"per_page"`
	Total     int     `json:"total"`
	TotalPage int     `json:"total_page"`
	Content   []Event `json:"content"`
}

type Alert struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	StartsAt    time.Time         `json:"startsAt,omitempty"`
	EndsAt      time.Time         `json:"endsAt,omitempty"`
}

type Notification struct {
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	Status            string            `json:"status"`
	Receiver          string            `json:"receiver"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Alerts            []Alert           `json:"alerts"`
}
