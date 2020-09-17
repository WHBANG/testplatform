package flow

import (
	"context"
	"time"
)

type AnalyzeSummary struct {
	Index     int         `json:"index"`     // 配置编号
	On        bool        `json:"on"`        // 是否开启
	Code      string      `json:"code"`      // 代码
	Name      string      `json:"name"`      // 名称
	Lane      string      `json:"lane"`      // 车道
	Direction string      `json:"direction"` // 方向
	RawData   interface{} `json:"-"`
}

type Task struct {
	ID            string                 `json:"id" bson:"id"`                         // 任务id, 必须全局唯一
	Namespace     string                 `json:"ns" bson:"ns"`                         // 任务所属的 VMR namespace
	Name          string                 `json:"name" bson:"name"`                     // 名称
	Type          string                 `json:"type" bson:"type"`                     // 任务类型
	Region        string                 `json:"region" bson:"region"`                 // 区域
	AnalyzeConfig map[string]interface{} `json:"analyze_config" bson:"analyze_config"` // 模型参数

	AnalyzeConfigSummaryMap map[string]AnalyzeSummary //摘要map

	StreamSetting `bson:",inline"`
	StreamList    []StreamSetting `json:"stream_list" bson:"stream_list"` // 多个摄像头

	WorkField WorkField `bson:"work_field" json:"work_field"` //任务运行字段

	// lease
	Status     string             `json:"status" bson:"status"`           // 任务状态
	ErrorMsg   string             `json:"error_msg" bson:"error_msg"`     // 错误信息
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`   // 创建时间
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`   //修改时间
	SwitchedAt time.Time          `json:"switched_at" bson:"switched_at"` //status转换时间
	Ver        int64              `json:"ver" bson:"ver"`                 // 版本
	Holder     interface{}        `json:"holder" bson:"holder"`
	CancelFunc context.CancelFunc `json:"-" bson:"-"` // for cancel

	Extra map[string]interface{} `json:"extra" bson:"extra"` // 为业务方保留的自定义字段

	Weight    int  `json:"weight" bson:"weight"`         // 权重
	IsDeleted bool `json:"is_deleted" bson:"is_deleted"` // 业务上是否被删除，和 status 字段区分开

	Device *Device `json:"device,omitempty" bson:"-"` // 摄像头信息
}

type StreamSetting struct {
	StreamID       string `json:"stream_id" bson:"stream_id"`          // StreamID来自device-api, 前端可以匹配到摄像头名称
	Snapshot       string `json:"snapshot" bson:"snapshot"`            // 布控截图
	StreamON       string `json:"stream_on" bson:"stream_on"`          // 是否需要推流
	ReadStreamURL  string `json:"read_stream_url,omitempty" bson:"-"`  // 读流地址, 提供前端辅助使用
	WriteStreamURL string `json:"write_stream_url,omitempty" bson:"-"` // 推流地址, 提供前端辅助使用
}

type WorkField struct {
	IsWorkTimeEnabled bool       `json:"is_work_time_enabled" bson:"is_work_time_enabled"`
	WorkTimes         []WorkTime `bson:"work_times" json:"work_times"` //布控生效的时间段
}

type WorkTime struct {
	Weekday     int          `bson:"weekday" json:"weekday"`           // 周日~周六 取值(0~6)
	TimeBuckets []TimeBucket `bson:"time_buckets" json:"time_buckets"` // 一天的生效时间段
}

type TimeBucket struct {
	StartTime string `bson:"start_time" json:"start_time"` // 开始时间
	EndTime   string `bson:"end_time" json:"end_time"`     // 结束时间
}

type CreateTaskReq struct {
	ID            string                 `json:"id" bson:"id"`                         // 任务id, 必须全局唯一
	Name          string                 `json:"name" bson:"name"`                     //名称
	Type          string                 `json:"type" bson:"type"`                     // 任务类型
	Region        string                 `json:"region" bson:"region"`                 // 区域
	AnalyzeConfig map[string]interface{} `json:"analyze_config" bson:"analyze_config"` // 模型参数

	StreamSetting
	StreamList []StreamSetting `json:"stream_list" bson:"stream_list"` // 多个摄像头

	WorkField WorkField              `bson:"work_field" json:"work_field"` //任务运行字段
	Status    string                 `json:"status" bson:"status"`         // 任务状态
	Extra     map[string]interface{} `json:"extra" bson:"extra"`           // 为业务方保留的自定义字段
}

type SearchReqeust struct {
	Type                  string   `json:"type,omitempty" query:"type"`
	ViolationType         string   `json:"violation_type,omitempty" query:"violation_type"`
	StreamID              string   `json:"stream_id,omitempty" query:"stream_id"`
	StreamON              string   `json:"stream_on" query:"stream_on"` // 是否推流
	Status                string   `json:"status,omitempty" query:"status"`
	CreatedTimeRangeBegin string   `json:"create_time_begin,omitempty" query:"create_time_begin"`
	CreatedTimeRangeEnd   string   `json:"create_time_end,omitempty" query:"create_time_end"`
	ID                    string   `json:"id,omitempty" query:"id"`
	Name                  string   `json:"name,omitempty" query:"name"`                   // 支持模糊匹配
	Simple                bool     `json:"simple,omitempty" query:"simple"`               // 是否只返回简单信息, 如忽略流播放地址等
	IDs                   []string `json:"ids,omitempty" query:"ids"`                     // ID 数组
	OrgCodes              []string `json:"org_codes,omitempty" query:"org_codes"`         // 组织代码列表
	DeviceDetail          bool     `json:"device_detail,omitempty" query:"device_detail"` // 是否返回设备详情
}

type UpdateTaskReq struct {
	Name          string                 `json:"name" query:"name"`           // 名称
	StreamON      string                 `json:"stream_on" query:"stream_on"` // 是否推流
	Snapshot      string                 `json:"snapshot,omitempty"`          // 布控截图, base64编码
	AnalyzeConfig map[string]interface{} `json:"analyze_config,omitempty"`    // 模型参数
	Extra         map[string]interface{} `json:"extra,omitempty"`             // 为业务方保留的自定义字段
}

type AcquireTaskReq struct {
	Type   string `json:"type,omitempty"`
	Region string `json:"region,omitempty"`
}
