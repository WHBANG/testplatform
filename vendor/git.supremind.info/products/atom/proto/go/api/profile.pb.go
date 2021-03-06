// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: profile.proto

package api

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type UserProfile struct {
	*Metadata `protobuf:"bytes,1,opt,name=meta,proto3,embedded=meta" json:"meta,omitempty" bson:",inline"`
	Spec      *UserProfileSpec   `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty" bson:"spec"`
	Status    *UserProfileStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty" bson:"status,omitempty"`
}

func (m *UserProfile) Reset()         { *m = UserProfile{} }
func (m *UserProfile) String() string { return proto.CompactTextString(m) }
func (*UserProfile) ProtoMessage()    {}
func (*UserProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{0}
}
func (m *UserProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserProfile.Unmarshal(m, b)
}
func (m *UserProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserProfile.Marshal(b, m, deterministic)
}
func (m *UserProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserProfile.Merge(m, src)
}
func (m *UserProfile) XXX_Size() int {
	return xxx_messageInfo_UserProfile.Size(m)
}
func (m *UserProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_UserProfile.DiscardUnknown(m)
}

var xxx_messageInfo_UserProfile proto.InternalMessageInfo

func (m *UserProfile) GetSpec() *UserProfileSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *UserProfile) GetStatus() *UserProfileStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type UserProfileSpec struct {
}

func (m *UserProfileSpec) Reset()         { *m = UserProfileSpec{} }
func (m *UserProfileSpec) String() string { return proto.CompactTextString(m) }
func (*UserProfileSpec) ProtoMessage()    {}
func (*UserProfileSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{1}
}
func (m *UserProfileSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserProfileSpec.Unmarshal(m, b)
}
func (m *UserProfileSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserProfileSpec.Marshal(b, m, deterministic)
}
func (m *UserProfileSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserProfileSpec.Merge(m, src)
}
func (m *UserProfileSpec) XXX_Size() int {
	return xxx_messageInfo_UserProfileSpec.Size(m)
}
func (m *UserProfileSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UserProfileSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UserProfileSpec proto.InternalMessageInfo

type UserProfileStatus struct {
}

func (m *UserProfileStatus) Reset()         { *m = UserProfileStatus{} }
func (m *UserProfileStatus) String() string { return proto.CompactTextString(m) }
func (*UserProfileStatus) ProtoMessage()    {}
func (*UserProfileStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{2}
}
func (m *UserProfileStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserProfileStatus.Unmarshal(m, b)
}
func (m *UserProfileStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserProfileStatus.Marshal(b, m, deterministic)
}
func (m *UserProfileStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserProfileStatus.Merge(m, src)
}
func (m *UserProfileStatus) XXX_Size() int {
	return xxx_messageInfo_UserProfileStatus.Size(m)
}
func (m *UserProfileStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_UserProfileStatus.DiscardUnknown(m)
}

var xxx_messageInfo_UserProfileStatus proto.InternalMessageInfo

type SetContainerRegistrySecretReq struct {
	SecretName string `protobuf:"bytes,1,opt,name=secretName,proto3" json:"secretName,omitempty"`
}

func (m *SetContainerRegistrySecretReq) Reset()         { *m = SetContainerRegistrySecretReq{} }
func (m *SetContainerRegistrySecretReq) String() string { return proto.CompactTextString(m) }
func (*SetContainerRegistrySecretReq) ProtoMessage()    {}
func (*SetContainerRegistrySecretReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{3}
}
func (m *SetContainerRegistrySecretReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetContainerRegistrySecretReq.Unmarshal(m, b)
}
func (m *SetContainerRegistrySecretReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetContainerRegistrySecretReq.Marshal(b, m, deterministic)
}
func (m *SetContainerRegistrySecretReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetContainerRegistrySecretReq.Merge(m, src)
}
func (m *SetContainerRegistrySecretReq) XXX_Size() int {
	return xxx_messageInfo_SetContainerRegistrySecretReq.Size(m)
}
func (m *SetContainerRegistrySecretReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SetContainerRegistrySecretReq.DiscardUnknown(m)
}

var xxx_messageInfo_SetContainerRegistrySecretReq proto.InternalMessageInfo

func (m *SetContainerRegistrySecretReq) GetSecretName() string {
	if m != nil {
		return m.SecretName
	}
	return ""
}

func init() {
	proto.RegisterType((*UserProfile)(nil), "apiserver.v2.UserProfile")
	proto.RegisterType((*UserProfileSpec)(nil), "apiserver.v2.UserProfileSpec")
	proto.RegisterType((*UserProfileStatus)(nil), "apiserver.v2.UserProfileStatus")
	proto.RegisterType((*SetContainerRegistrySecretReq)(nil), "apiserver.v2.SetContainerRegistrySecretReq")
}

func init() { proto.RegisterFile("profile.proto", fileDescriptor_744bf7a47b381504) }

var fileDescriptor_744bf7a47b381504 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xb1, 0xee, 0xd3, 0x30,
	0x10, 0xc6, 0x93, 0x3f, 0xa5, 0x02, 0x17, 0xa8, 0x1a, 0x10, 0x54, 0x45, 0x4d, 0x90, 0x27, 0x86,
	0x36, 0x91, 0xca, 0xc6, 0x82, 0x48, 0x37, 0x24, 0x10, 0x4a, 0x61, 0x61, 0x73, 0x92, 0x6b, 0xb0,
	0xd4, 0xd8, 0xc6, 0xbe, 0x44, 0xe4, 0x15, 0x98, 0x78, 0x0c, 0x1e, 0xa7, 0x23, 0x53, 0x24, 0xe8,
	0x1b, 0x94, 0x8d, 0x09, 0xd5, 0x29, 0x52, 0xa0, 0xfa, 0x6f, 0xf7, 0x9d, 0xef, 0xfb, 0xdd, 0xc9,
	0xfa, 0xc8, 0x5d, 0xa5, 0xe5, 0x96, 0xef, 0x20, 0x54, 0x5a, 0xa2, 0xf4, 0xee, 0x30, 0xc5, 0x0d,
	0xe8, 0x1a, 0x74, 0x58, 0xaf, 0x66, 0xcb, 0x82, 0xe3, 0xc7, 0x2a, 0x0d, 0x33, 0x59, 0x46, 0x85,
	0x2c, 0x64, 0x64, 0x87, 0xd2, 0x6a, 0x6b, 0x95, 0x15, 0xb6, 0xea, 0xcc, 0xb3, 0x97, 0xbd, 0x71,
	0x10, 0xb5, 0x6c, 0x94, 0x96, 0x9f, 0x9b, 0xce, 0x94, 0x2d, 0x0b, 0x10, 0xcb, 0x9a, 0xed, 0x78,
	0xce, 0x10, 0xa2, 0x8b, 0xe2, 0x8c, 0x20, 0x25, 0x20, 0xeb, 0x6a, 0xfa, 0xcb, 0x25, 0xa3, 0xf7,
	0x06, 0xf4, 0xdb, 0xee, 0x42, 0xef, 0x15, 0x19, 0x9c, 0x5e, 0xa7, 0xee, 0x13, 0xf7, 0xe9, 0x68,
	0xf5, 0x30, 0xec, 0x9f, 0x1a, 0xbe, 0x06, 0x64, 0x39, 0x43, 0x16, 0xfb, 0xbf, 0xe3, 0x9b, 0x5f,
	0xdc, 0xab, 0x5b, 0xee, 0xbe, 0x0d, 0xdc, 0x63, 0x1b, 0xdc, 0x4b, 0x8d, 0x14, 0xcf, 0xe9, 0x82,
	0x8b, 0x1d, 0x17, 0x40, 0x13, 0xcb, 0xf0, 0x62, 0x32, 0x30, 0x0a, 0xb2, 0xe9, 0x95, 0x65, 0xcd,
	0xff, 0x65, 0xf5, 0x96, 0x6e, 0x14, 0x64, 0xf1, 0xf8, 0xd8, 0x06, 0xa3, 0x0e, 0x73, 0x32, 0xd1,
	0xc4, 0x7a, 0xbd, 0x77, 0x64, 0x68, 0x90, 0x61, 0x65, 0xa6, 0x37, 0x2c, 0x25, 0xb8, 0x9e, 0x62,
	0xc7, 0xe2, 0xc7, 0xc7, 0x36, 0x78, 0x74, 0xe6, 0xd8, 0xce, 0x42, 0x96, 0x1c, 0xa1, 0x54, 0xd8,
	0xd0, 0xe4, 0xcc, 0xa2, 0x13, 0x32, 0xfe, 0x6f, 0x3f, 0xbd, 0x4f, 0x26, 0x17, 0x30, 0xfa, 0x82,
	0xcc, 0x37, 0x80, 0x6b, 0x29, 0x90, 0x71, 0x01, 0x3a, 0x81, 0x82, 0x1b, 0xd4, 0xcd, 0x06, 0x32,
	0x0d, 0x98, 0xc0, 0x27, 0xcf, 0x27, 0xc4, 0x58, 0xf1, 0x86, 0x95, 0x60, 0x3f, 0xed, 0x76, 0xd2,
	0xeb, 0xac, 0x1e, 0x10, 0xaf, 0x4f, 0x05, 0x5d, 0xf3, 0x0c, 0xe2, 0xf5, 0xfe, 0x87, 0xef, 0x7c,
	0xfb, 0xe9, 0xbb, 0x5f, 0x0f, 0xbe, 0xb3, 0x3f, 0xf8, 0xce, 0xf7, 0x83, 0xef, 0x7c, 0x38, 0x05,
	0x21, 0x34, 0x95, 0xd2, 0x50, 0x72, 0x91, 0x87, 0x5c, 0x6c, 0x6d, 0x16, 0xf2, 0x2a, 0x43, 0x13,
	0x31, 0x94, 0x65, 0xf4, 0x37, 0x0c, 0x11, 0x53, 0x3c, 0x1d, 0x5a, 0xf5, 0xec, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x88, 0x44, 0xd3, 0x77, 0x5d, 0x02, 0x00, 0x00,
}

func (this *UserProfile) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UserProfile)
	if !ok {
		that2, ok := that.(UserProfile)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	if !this.Status.Equal(that1.Status) {
		return false
	}
	return true
}
func (this *UserProfileSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UserProfileSpec)
	if !ok {
		that2, ok := that.(UserProfileSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *UserProfileStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UserProfileStatus)
	if !ok {
		that2, ok := that.(UserProfileStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *SetContainerRegistrySecretReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SetContainerRegistrySecretReq)
	if !ok {
		that2, ok := that.(SetContainerRegistrySecretReq)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.SecretName != that1.SecretName {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserProfileServiceClient is the client API for UserProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserProfileServiceClient interface {
}

type userProfileServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserProfileServiceClient(cc *grpc.ClientConn) UserProfileServiceClient {
	return &userProfileServiceClient{cc}
}

// UserProfileServiceServer is the server API for UserProfileService service.
type UserProfileServiceServer interface {
}

// UnimplementedUserProfileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserProfileServiceServer struct {
}

func RegisterUserProfileServiceServer(s *grpc.Server, srv UserProfileServiceServer) {
	s.RegisterService(&_UserProfileService_serviceDesc, srv)
}

var _UserProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apiserver.v2.UserProfileService",
	HandlerType: (*UserProfileServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "profile.proto",
}
