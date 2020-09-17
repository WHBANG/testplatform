package go_sdk

import (
	"context"
	"time"

	"git.supremind.info/product/visionmind/com/flow"
)

type CameraOptions struct {
	DeleteInclude bool `json:"delete_include"`
}

type IFlowClient interface {
	// Task
	TaskTry(ctx context.Context, leaseType string, region string, holder string) (*flow.Task, error)
	TaskTryWithWeight(ctx context.Context, leaseType string, region string, holder string, weight int) (*flow.Task, error)
	// taskHold extra切片顺序，status, errorMsg
	TaskHold(ctx context.Context, id string, ver int64, holder string, extra ...string) (*flow.Task, error)
	// TaskRelease extra切片顺序，errorMsg, (if errorMsg != "" status=failed else status=waiting)
	TaskRelease(ctx context.Context, id string, ver int64, extra ...string) error

	// Msg
	Push(ctx context.Context, topic string, body interface{}) (int64, error)
	// pull by ID
	PullByID(ctx context.Context, topic string, id int64) (*flow.Message, error)
	PullBatch(ctx context.Context, topic string, begid, endid int64) ([]flow.Message, error)
	// pull by group
	Pull(ctx context.Context, topic, group string) (*flow.Message, error)
	//pull with since
	PullSince(ctx context.Context, topic, group string, since time.Time) (*flow.Message, error)

	// Config
	ConfigCheck(ctx context.Context, id string) (*flow.Config, error)
	ConfigWatch(ctx context.Context, conf *flow.Config, callback func(context.Context, *flow.Config) bool)

	// Device
	GetDevices(ctx context.Context) ([]flow.Device, error)
	GetDevice(ctx context.Context, deviceID string) (*flow.Device, error)
	GetDeviceStreamURL(ctx context.Context, deviceID string, streamID int, streamType string) (url string, err error) // streamType: rtmp, rtsp, hlv
	GetDeviceVChannelURL(ctx context.Context, deviceID string, target string) (url string, err error)
	GetDeviceSnap(ctx context.Context, deviceID string, streamType string) (*flow.Snap, error)
	GetDownloadUrl(ctx context.Context, deviceID string, startTime, endTime time.Time, timeout time.Duration) (string, error)
	CreateDevice(ctx context.Context, device interface{}) (*flow.Device, error)
	DeleteDevice(ctx context.Context, deviceIDs []string) ([]string, error)
	GetCameras(ctx context.Context, options CameraOptions) ([]flow.Camera, error)

	//Deep.Topics
	FetchTopics(ctx context.Context, filters []string) ([]string, error)
	UpdateConfig(ctx context.Context, configId string, data interface{}) (err error)
}
