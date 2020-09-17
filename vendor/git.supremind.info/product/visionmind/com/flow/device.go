package flow

import (
	"context"
	"time"

	baseProto "git.supremind.info/product/visionmind/com/proto"
)

type Device struct {
	ID         string      `json:"id"`          // 设备id
	Namespace  string      `json:"ns"`          // 设备所属的 VMR namespace
	Name       string      `json:"name"`        // 设备名称
	BoxChannel string      `json:"box_channel"` // 通道号
	Meta       interface{} `json:"meta"`        // meta 信息，对应vms的user_data
	Extra      interface{} `json:"extra"`       // 对应业务信息
	BaseModel  BaseModel   `json:"base_model"`
}

type MainDevice struct {
	ID         string    `json:"id"`         // 主设备id
	Namespace  string    `json:"namespace"`  // 设备所属的 VMR namespace
	Name       string    `json:"name"`       // 主设备名称
	Type       int       `json:"type"`       // 主设备类型
	Permission int       `json:"permission"` // 主设备权限信息
	BaseModel  BaseModel `json:"base_model"`
}

// Base公共属性
type BaseModel struct {
	IsDeleted int        `json:"is_deleted"`
	CreatedAt int        `json:"created_at"`
	UpdatedAt int        `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type CamerasMap map[string][]Camera
type Camera struct {
	Channel         int                    `json:"channel"`             // 通道号
	IndexCode       string                 `json:"index_code"`          // 结点
	Name            string                 `json:"name"`                // 名称
	ParentIndexCode string                 `json:"parent_index_code"`   // 父节点
	CameraType      int                    `json:"camera_type"`         // 摄像头类型
	Description     string                 `json:"description"`         // 描述
	IsOnline        int                    `json:"is_online,omitempty"` // 在线
	ExtraField      map[string]interface{} `json:"extra_field"`         // 多余字段
	RegionInfo      interface{}            `json:"region_info"`         // 区域信息，对应vms的user_data
	Video           baseProto.Video        `json:"video"`               // 视频流信息
	Extra           interface{}            `json:"extra"`               // 对应业务信息
}

type UnitsMap map[string][]UnitInfo
type UnitInfo struct {
	IndexCode       string `json:"index_code"`          // 节点
	Name            string `json:"name"`                // 名称
	ParentIndexCode string `json:"parent_index_code"`   // 父节点
	UnitLevel       int    `json:"unit_level"`          // 目录层级
	UnitType        int    `json:"unit_type,omitempty"` // 目录类型
}

type NodesMap map[string][]NodeInfo
type NodeInfo struct {
	IndexCode       string                 `json:"index_code"`        // 节点
	Name            string                 `json:"name"`              // 名称
	ParentIndexCode string                 `json:"parent_index_code"` // 父节点
	IsLeaf          bool                   `json:"is_leaf"`           // 是否是叶子节点
	ExtraField      map[string]interface{} `json:"extra_field"`       // 冗余字段，包vms的user_data
}

type Snap struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Base64 string `json:"base64"`
}

type IDeviceMgt interface {
	GetDevices(ctx context.Context) ([]Device, error)
	GetDevice(ctx context.Context, deviceID string) (Device, error)
	GetDeviceStreamURL(ctx context.Context, deviceID string, streamID int, streamType string) (url string, err error) // streamType: rtmp, rtsp, hlv
	GetDeviceStreamFile(ctx context.Context, deviceID string, startTime, endTime string) (content []byte, err error)
	GetDeviceVChannelURL(ctx context.Context, deviceID string, target string) (url string, err error)
	CreateDevice(ctx context.Context, params interface{}) (*Device, error)
	CreateMainDevice(ctx context.Context, params interface{}) (*MainDevice, error)
	DeleteDevice(ctx context.Context, deviceID string) (err error)
	GetDeviceSnap(ctx context.Context, deviceID string, streamType string) (Snap, error)
	UpdateDevice(ctx context.Context, deviceID string, params interface{}) (*Device, error)

	GetCameras(ctx context.Context, deletedInclude bool) ([]Camera, error)
	GetCamera(ctx context.Context, deviceID string) (Camera, error)
	GetUnits(ctx context.Context, parentID, name string) ([]UnitInfo, error)
	GetUnit(ctx context.Context, unitID string) (UnitInfo, error)
	GetChildren(ctx context.Context, orgId string) ([]NodeInfo, error)
	SyncCatalog(ctx context.Context, deviceId string) error
}
