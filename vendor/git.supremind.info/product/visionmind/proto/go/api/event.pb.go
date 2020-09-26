// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: event.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CodecType int32

const (
	CodecType_CODEC_TYPE_H264   CodecType = 0
	CodecType_CODEC_TYPE_H265   CodecType = 1
	CodecType_CODEC_TYPE_UNKNOW CodecType = 2
)

var CodecType_name = map[int32]string{
	0: "CODEC_TYPE_H264",
	1: "CODEC_TYPE_H265",
	2: "CODEC_TYPE_UNKNOW",
}

var CodecType_value = map[string]int32{
	"CODEC_TYPE_H264":   0,
	"CODEC_TYPE_H265":   1,
	"CODEC_TYPE_UNKNOW": 2,
}

func (x CodecType) String() string {
	return proto.EnumName(CodecType_name, int32(x))
}

func (CodecType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

type EventType int32

const (
	EventType_EVENT_TYPE_STREAM_OPEN      EventType = 0
	EventType_EVENT_TYPE_STREAM_CLOSE     EventType = 1
	EventType_EVENT_TYPE_STREAM_EXCEPTION EventType = 2
)

var EventType_name = map[int32]string{
	0: "EVENT_TYPE_STREAM_OPEN",
	1: "EVENT_TYPE_STREAM_CLOSE",
	2: "EVENT_TYPE_STREAM_EXCEPTION",
}

var EventType_value = map[string]int32{
	"EVENT_TYPE_STREAM_OPEN":      0,
	"EVENT_TYPE_STREAM_CLOSE":     1,
	"EVENT_TYPE_STREAM_EXCEPTION": 2,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{1}
}

type ReportEventRequest struct {
	DeviceId             string    `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	EventType            EventType `protobuf:"varint,2,opt,name=event_type,json=eventType,proto3,enum=vmr_proto.EventType" json:"event_type,omitempty"`
	Time                 string    `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	EventStatus          string    `protobuf:"bytes,4,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	Comment              string    `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ReportEventRequest) Reset()         { *m = ReportEventRequest{} }
func (m *ReportEventRequest) String() string { return proto.CompactTextString(m) }
func (*ReportEventRequest) ProtoMessage()    {}
func (*ReportEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}
func (m *ReportEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportEventRequest.Unmarshal(m, b)
}
func (m *ReportEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportEventRequest.Marshal(b, m, deterministic)
}
func (m *ReportEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportEventRequest.Merge(m, src)
}
func (m *ReportEventRequest) XXX_Size() int {
	return xxx_messageInfo_ReportEventRequest.Size(m)
}
func (m *ReportEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportEventRequest proto.InternalMessageInfo

func (m *ReportEventRequest) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *ReportEventRequest) GetEventType() EventType {
	if m != nil {
		return m.EventType
	}
	return EventType_EVENT_TYPE_STREAM_OPEN
}

func (m *ReportEventRequest) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *ReportEventRequest) GetEventStatus() string {
	if m != nil {
		return m.EventStatus
	}
	return ""
}

func (m *ReportEventRequest) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type Video struct {
	Height               int32     `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Width                int32     `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Codec                CodecType `protobuf:"varint,3,opt,name=codec,proto3,enum=vmr_proto.CodecType" json:"codec,omitempty"`
	Rate                 int32     `protobuf:"varint,4,opt,name=rate,proto3" json:"rate,omitempty"`
	FrameRate            int32     `protobuf:"varint,5,opt,name=frame_rate,json=frameRate,proto3" json:"frame_rate,omitempty"`
	DropFrameRate        int32     `protobuf:"varint,6,opt,name=drop_frame_rate,json=dropFrameRate,proto3" json:"drop_frame_rate,omitempty"`
	RecvFrameNum         int32     `protobuf:"varint,7,opt,name=recv_frame_num,json=recvFrameNum,proto3" json:"recv_frame_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Video) Reset()         { *m = Video{} }
func (m *Video) String() string { return proto.CompactTextString(m) }
func (*Video) ProtoMessage()    {}
func (*Video) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{1}
}
func (m *Video) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Video.Unmarshal(m, b)
}
func (m *Video) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Video.Marshal(b, m, deterministic)
}
func (m *Video) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Video.Merge(m, src)
}
func (m *Video) XXX_Size() int {
	return xxx_messageInfo_Video.Size(m)
}
func (m *Video) XXX_DiscardUnknown() {
	xxx_messageInfo_Video.DiscardUnknown(m)
}

var xxx_messageInfo_Video proto.InternalMessageInfo

func (m *Video) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Video) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Video) GetCodec() CodecType {
	if m != nil {
		return m.Codec
	}
	return CodecType_CODEC_TYPE_H264
}

func (m *Video) GetRate() int32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *Video) GetFrameRate() int32 {
	if m != nil {
		return m.FrameRate
	}
	return 0
}

func (m *Video) GetDropFrameRate() int32 {
	if m != nil {
		return m.DropFrameRate
	}
	return 0
}

func (m *Video) GetRecvFrameNum() int32 {
	if m != nil {
		return m.RecvFrameNum
	}
	return 0
}

type StreamInfo struct {
	DeviceId             string   `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Duration             string   `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	Video                *Video   `protobuf:"bytes,3,opt,name=video,proto3" json:"video,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamInfo) Reset()         { *m = StreamInfo{} }
func (m *StreamInfo) String() string { return proto.CompactTextString(m) }
func (*StreamInfo) ProtoMessage()    {}
func (*StreamInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{2}
}
func (m *StreamInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamInfo.Unmarshal(m, b)
}
func (m *StreamInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamInfo.Marshal(b, m, deterministic)
}
func (m *StreamInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamInfo.Merge(m, src)
}
func (m *StreamInfo) XXX_Size() int {
	return xxx_messageInfo_StreamInfo.Size(m)
}
func (m *StreamInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StreamInfo proto.InternalMessageInfo

func (m *StreamInfo) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *StreamInfo) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *StreamInfo) GetVideo() *Video {
	if m != nil {
		return m.Video
	}
	return nil
}

type ReportStreamInfoRequest struct {
	Host                 string        `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Infos                []*StreamInfo `protobuf:"bytes,2,rep,name=infos,proto3" json:"infos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReportStreamInfoRequest) Reset()         { *m = ReportStreamInfoRequest{} }
func (m *ReportStreamInfoRequest) String() string { return proto.CompactTextString(m) }
func (*ReportStreamInfoRequest) ProtoMessage()    {}
func (*ReportStreamInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{3}
}
func (m *ReportStreamInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportStreamInfoRequest.Unmarshal(m, b)
}
func (m *ReportStreamInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportStreamInfoRequest.Marshal(b, m, deterministic)
}
func (m *ReportStreamInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportStreamInfoRequest.Merge(m, src)
}
func (m *ReportStreamInfoRequest) XXX_Size() int {
	return xxx_messageInfo_ReportStreamInfoRequest.Size(m)
}
func (m *ReportStreamInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportStreamInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportStreamInfoRequest proto.InternalMessageInfo

func (m *ReportStreamInfoRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ReportStreamInfoRequest) GetInfos() []*StreamInfo {
	if m != nil {
		return m.Infos
	}
	return nil
}

func init() {
	proto.RegisterEnum("vmr_proto.CodecType", CodecType_name, CodecType_value)
	proto.RegisterEnum("vmr_proto.EventType", EventType_name, EventType_value)
	proto.RegisterType((*ReportEventRequest)(nil), "vmr_proto.ReportEventRequest")
	proto.RegisterType((*Video)(nil), "vmr_proto.Video")
	proto.RegisterType((*StreamInfo)(nil), "vmr_proto.StreamInfo")
	proto.RegisterType((*ReportStreamInfoRequest)(nil), "vmr_proto.ReportStreamInfoRequest")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xad, 0xd3, 0xba, 0xad, 0x27, 0xfd, 0x5a, 0x7f, 0x4b, 0x7f, 0xac, 0x44, 0x15, 0x25, 0x42,
	0x55, 0x55, 0x24, 0x47, 0x4a, 0x0b, 0xf7, 0x90, 0x1a, 0x51, 0x51, 0x9c, 0xca, 0x09, 0x05, 0x7a,
	0x63, 0xa5, 0xf1, 0x24, 0x59, 0x09, 0x7b, 0x8d, 0xbd, 0x0e, 0xca, 0xfb, 0xf0, 0x0e, 0xbc, 0x0f,
	0x4f, 0x82, 0x76, 0x36, 0x09, 0x06, 0x0b, 0xee, 0x66, 0xce, 0x39, 0x99, 0x39, 0x33, 0xde, 0x09,
	0xd4, 0x71, 0x86, 0x89, 0x74, 0xd3, 0x4c, 0x48, 0xc1, 0xac, 0x59, 0x9c, 0x85, 0x14, 0x36, 0x9a,
	0x13, 0x21, 0x26, 0x9f, 0xb1, 0x4d, 0xd9, 0x43, 0x31, 0x6e, 0x63, 0x9c, 0xca, 0xb9, 0xd6, 0xb5,
	0xbe, 0x1b, 0xc0, 0x02, 0x4c, 0x45, 0x26, 0x3d, 0xf5, 0xeb, 0x00, 0xbf, 0x14, 0x98, 0x4b, 0xd6,
	0x04, 0x2b, 0xc2, 0x19, 0x1f, 0x61, 0xc8, 0x23, 0xc7, 0x38, 0x31, 0xce, 0xac, 0x60, 0x5b, 0x03,
	0xd7, 0x11, 0xbb, 0x00, 0xa0, 0x56, 0xa1, 0x9c, 0xa7, 0xe8, 0xd4, 0x4e, 0x8c, 0xb3, 0xdd, 0xce,
	0xbe, 0xbb, 0x6a, 0xe8, 0x52, 0xa5, 0xc1, 0x3c, 0xc5, 0xc0, 0xc2, 0x65, 0xc8, 0x18, 0x6c, 0x48,
	0x1e, 0xa3, 0xb3, 0x4e, 0xc5, 0x28, 0x66, 0x4f, 0x60, 0x47, 0x17, 0xca, 0xe5, 0x50, 0x16, 0xb9,
	0xb3, 0x41, 0x9c, 0x9e, 0xa3, 0x4f, 0x10, 0x73, 0x60, 0x6b, 0x24, 0xe2, 0x18, 0x13, 0xe9, 0x98,
	0xc4, 0x2e, 0xd3, 0xd6, 0x0f, 0x03, 0xcc, 0x3b, 0x1e, 0xa1, 0x60, 0x87, 0xb0, 0x39, 0x45, 0x3e,
	0x99, 0x4a, 0x72, 0x6a, 0x06, 0x8b, 0x8c, 0xed, 0x83, 0xf9, 0x95, 0x47, 0x72, 0x4a, 0x16, 0xcd,
	0x40, 0x27, 0xec, 0x1c, 0xcc, 0x91, 0x88, 0x70, 0x44, 0x4e, 0x7e, 0x37, 0xde, 0x55, 0x38, 0x19,
	0xd7, 0x12, 0x65, 0x3a, 0x1b, 0x4a, 0x24, 0x63, 0x66, 0x40, 0x31, 0x3b, 0x06, 0x18, 0x67, 0xc3,
	0x18, 0x43, 0x62, 0x4c, 0x62, 0x2c, 0x42, 0x02, 0x45, 0x9f, 0xc2, 0x5e, 0x94, 0x89, 0x34, 0x2c,
	0x69, 0x36, 0x49, 0xf3, 0x9f, 0x82, 0x5f, 0xaf, 0x74, 0x4f, 0x61, 0x37, 0xc3, 0xd1, 0x6c, 0xa1,
	0x4b, 0x8a, 0xd8, 0xd9, 0x22, 0xd9, 0x8e, 0x42, 0x49, 0xe6, 0x17, 0x71, 0x2b, 0x06, 0xe8, 0xcb,
	0x0c, 0x87, 0xf1, 0x75, 0x32, 0x16, 0xff, 0xfe, 0x2a, 0x0d, 0xd8, 0x8e, 0x8a, 0x6c, 0x28, 0xb9,
	0x48, 0x68, 0x60, 0xc5, 0x2d, 0x72, 0x76, 0x0a, 0xe6, 0x4c, 0xad, 0x8a, 0x66, 0xae, 0x77, 0xec,
	0xd2, 0xcc, 0xb4, 0xc2, 0x40, 0xd3, 0xad, 0x7b, 0x38, 0xd2, 0x8f, 0xe1, 0x57, 0xd3, 0xe5, 0x8b,
	0x60, 0xb0, 0x31, 0x15, 0xb9, 0x5c, 0xb4, 0xa5, 0x98, 0x3d, 0x03, 0x93, 0x27, 0x63, 0x91, 0x3b,
	0xb5, 0x93, 0xf5, 0xb3, 0x7a, 0xe7, 0xa0, 0x54, 0xb6, 0x54, 0x40, 0x6b, 0xce, 0x6f, 0xc0, 0x5a,
	0xed, 0x97, 0x3d, 0x82, 0xbd, 0x6e, 0xef, 0xca, 0xeb, 0x86, 0x83, 0x4f, 0xb7, 0x5e, 0xf8, 0xa6,
	0xf3, 0xe2, 0xd2, 0x5e, 0xab, 0x82, 0xcf, 0x6d, 0x83, 0x1d, 0xc0, 0xff, 0x25, 0xf0, 0xbd, 0xff,
	0xd6, 0xef, 0x7d, 0xb0, 0x6b, 0xe7, 0x08, 0xd6, 0xea, 0x99, 0xb1, 0x06, 0x1c, 0x7a, 0x77, 0x9e,
	0x3f, 0xd0, 0x9a, 0xfe, 0x20, 0xf0, 0x5e, 0xbe, 0x0b, 0x7b, 0xb7, 0x9e, 0x6f, 0xaf, 0xb1, 0x26,
	0x1c, 0x55, 0xb9, 0xee, 0x4d, 0xaf, 0xef, 0xd9, 0x06, 0x7b, 0x0c, 0xcd, 0x2a, 0xe9, 0x7d, 0xec,
	0x7a, 0xb7, 0x83, 0xeb, 0x9e, 0x6f, 0xd7, 0x3a, 0xdf, 0x0c, 0xd8, 0xa1, 0x3e, 0x7d, 0xcc, 0xd4,
	0x9e, 0xd9, 0x15, 0xd4, 0x4b, 0xe7, 0xc2, 0x8e, 0x4b, 0x23, 0x57, 0xcf, 0xa8, 0x71, 0xe8, 0xea,
	0xdb, 0x73, 0x97, 0xb7, 0xe7, 0x7a, 0xea, 0xf6, 0x98, 0x0f, 0xf6, 0x9f, 0x7b, 0x66, 0xad, 0x4a,
	0xa9, 0xca, 0x47, 0xf8, 0x5b, 0xbd, 0x57, 0x97, 0xf7, 0x9d, 0x09, 0x97, 0x6e, 0x5e, 0xa4, 0x19,
	0xc6, 0x3c, 0x89, 0x5c, 0xb5, 0x72, 0x75, 0xf0, 0x51, 0x31, 0x92, 0xed, 0x19, 0xcf, 0xb9, 0x48,
	0x14, 0xae, 0xff, 0x03, 0xda, 0x13, 0xd1, 0x1e, 0xa6, 0xfc, 0x61, 0x93, 0xb2, 0x8b, 0x9f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xb8, 0x76, 0xa1, 0x43, 0x39, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventServiceClient interface {
	ReportEvent(ctx context.Context, in *ReportEventRequest, opts ...grpc.CallOption) (*types.Empty, error)
	ReportStreamInfo(ctx context.Context, in *ReportStreamInfoRequest, opts ...grpc.CallOption) (*types.Empty, error)
}

type eventServiceClient struct {
	cc *grpc.ClientConn
}

func NewEventServiceClient(cc *grpc.ClientConn) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) ReportEvent(ctx context.Context, in *ReportEventRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/vmr_proto.EventService/ReportEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) ReportStreamInfo(ctx context.Context, in *ReportStreamInfoRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/vmr_proto.EventService/ReportStreamInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
type EventServiceServer interface {
	ReportEvent(context.Context, *ReportEventRequest) (*types.Empty, error)
	ReportStreamInfo(context.Context, *ReportStreamInfoRequest) (*types.Empty, error)
}

// UnimplementedEventServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (*UnimplementedEventServiceServer) ReportEvent(ctx context.Context, req *ReportEventRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportEvent not implemented")
}
func (*UnimplementedEventServiceServer) ReportStreamInfo(ctx context.Context, req *ReportStreamInfoRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportStreamInfo not implemented")
}

func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer) {
	s.RegisterService(&_EventService_serviceDesc, srv)
}

func _EventService_ReportEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).ReportEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmr_proto.EventService/ReportEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).ReportEvent(ctx, req.(*ReportEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_ReportStreamInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportStreamInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).ReportStreamInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmr_proto.EventService/ReportStreamInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).ReportStreamInfo(ctx, req.(*ReportStreamInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vmr_proto.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportEvent",
			Handler:    _EventService_ReportEvent_Handler,
		},
		{
			MethodName: "ReportStreamInfo",
			Handler:    _EventService_ReportStreamInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event.proto",
}