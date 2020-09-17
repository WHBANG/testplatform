package proto

import "git.supremind.info/product/visionmind/proto/go/api"

type Video struct {
	Height    int32         `json:"height" bson:"height"`         //视频高度
	Width     int32         `json:"width" bson:"width"`           //视频宽度
	Codec     api.CodecType `json:"codec" bson:"codec"`           //视频编码格式h264 or h265
	Rate      int32         `json:"rate" bson:"rate"`             //视频码率 kbps
	FrameRate int32         `json:"frame_rate" bson:"frame_rate"` //视频帧率
}
