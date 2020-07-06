// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: meta.proto

package api

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SpecVersion int32

const (
	SpecVersionUnset SpecVersion = 0
	SpecVersionV2    SpecVersion = 2
)

var SpecVersion_name = map[int32]string{
	0: "SpecVersionUnset",
	2: "SpecVersionV2",
}

var SpecVersion_value = map[string]int32{
	"SpecVersionUnset": 0,
	"SpecVersionV2":    2,
}

func (x SpecVersion) String() string {
	return proto.EnumName(SpecVersion_name, int32(x))
}

func (SpecVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{0}
}

type ResourceKind int32

const (
	ResourceKindUnknown        ResourceKind = 0
	ResourceKindVolume         ResourceKind = 1
	ResourceKindDataset        ResourceKind = 2
	ResourceKindDatasetVersion ResourceKind = 3
	ResourceKindSecret         ResourceKind = 4
	ResourceKindJob            ResourceKind = 5
	ResourceKindUserProfile    ResourceKind = 6
	ResourceKindPackage        ResourceKind = 7
	ResourceKindDeviceCategory ResourceKind = 8
	ResourceKindOre            ResourceKind = 9
	ResourceKindStorage        ResourceKind = 10
	ResourceKindQuota          ResourceKind = 11
)

var ResourceKind_name = map[int32]string{
	0:  "ResourceKindUnknown",
	1:  "ResourceKindVolume",
	2:  "ResourceKindDataset",
	3:  "ResourceKindDatasetVersion",
	4:  "ResourceKindSecret",
	5:  "ResourceKindJob",
	6:  "ResourceKindUserProfile",
	7:  "ResourceKindPackage",
	8:  "ResourceKindDeviceCategory",
	9:  "ResourceKindOre",
	10: "ResourceKindStorage",
	11: "ResourceKindQuota",
}

var ResourceKind_value = map[string]int32{
	"ResourceKindUnknown":        0,
	"ResourceKindVolume":         1,
	"ResourceKindDataset":        2,
	"ResourceKindDatasetVersion": 3,
	"ResourceKindSecret":         4,
	"ResourceKindJob":            5,
	"ResourceKindUserProfile":    6,
	"ResourceKindPackage":        7,
	"ResourceKindDeviceCategory": 8,
	"ResourceKindOre":            9,
	"ResourceKindStorage":        10,
	"ResourceKindQuota":          11,
}

func (x ResourceKind) String() string {
	return proto.EnumName(ResourceKind_name, int32(x))
}

func (ResourceKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{1}
}

type ResourceReference struct {
	SpecVersion SpecVersion  `protobuf:"varint,1,opt,name=specVersion,proto3,enum=apiserver.v2.SpecVersion" json:"specVersion,omitempty" bson:"specVersion"`
	Kind        ResourceKind `protobuf:"varint,2,opt,name=kind,proto3,enum=apiserver.v2.ResourceKind" json:"kind,omitempty" bson:"kind"`
	Name        string       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	Creator     string       `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty" bson:"creator"`
}

func (m *ResourceReference) Reset()         { *m = ResourceReference{} }
func (m *ResourceReference) String() string { return proto.CompactTextString(m) }
func (*ResourceReference) ProtoMessage()    {}
func (*ResourceReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{0}
}
func (m *ResourceReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceReference.Unmarshal(m, b)
}
func (m *ResourceReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceReference.Marshal(b, m, deterministic)
}
func (m *ResourceReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceReference.Merge(m, src)
}
func (m *ResourceReference) XXX_Size() int {
	return xxx_messageInfo_ResourceReference.Size(m)
}
func (m *ResourceReference) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceReference.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceReference proto.InternalMessageInfo

func (m *ResourceReference) GetSpecVersion() SpecVersion {
	if m != nil {
		return m.SpecVersion
	}
	return SpecVersionUnset
}

func (m *ResourceReference) GetKind() ResourceKind {
	if m != nil {
		return m.Kind
	}
	return ResourceKindUnknown
}

func (m *ResourceReference) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceReference) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type ID struct {
	Hex string `protobuf:"bytes,1,opt,name=hex,proto3" json:"hex,omitempty"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{1}
}
func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetHex() string {
	if m != nil {
		return m.Hex
	}
	return ""
}

// generic metadata for api objects
type Metadata struct {
	SpecVersion SpecVersion        `protobuf:"varint,1,opt,name=specVersion,proto3,enum=apiserver.v2.SpecVersion" json:"specVersion,omitempty" bson:"specVersion,omitempty"`
	Kind        ResourceKind       `protobuf:"varint,2,opt,name=kind,proto3,enum=apiserver.v2.ResourceKind" json:"kind,omitempty" bson:"kind,omitempty"`
	Name        string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	Desc        string             `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty" bson:"desc,omitempty"`
	Creator     string             `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty" bson:"creator"`
	OwnerRef    *ResourceReference `protobuf:"bytes,6,opt,name=ownerRef,proto3" json:"ownerRef,omitempty" bson:"ownerRef,omitempty"`
	CreateTime  *time.Time         `protobuf:"bytes,7,opt,name=createTime,proto3,stdtime" json:"createTime,omitempty" bson:"createTime"`
	UpdateTime  *time.Time         `protobuf:"bytes,8,opt,name=updateTime,proto3,stdtime" json:"updateTime,omitempty" bson:"updateTime"`
	Tags        map[string]string  `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty" bson:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Id          *ID                `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty" bson:"_id,omitempty"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{2}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetSpecVersion() SpecVersion {
	if m != nil {
		return m.SpecVersion
	}
	return SpecVersionUnset
}

func (m *Metadata) GetKind() ResourceKind {
	if m != nil {
		return m.Kind
	}
	return ResourceKindUnknown
}

func (m *Metadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metadata) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Metadata) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Metadata) GetOwnerRef() *ResourceReference {
	if m != nil {
		return m.OwnerRef
	}
	return nil
}

func (m *Metadata) GetCreateTime() *time.Time {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Metadata) GetUpdateTime() *time.Time {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Metadata) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Metadata) GetId() *ID {
	if m != nil {
		return m.Id
	}
	return nil
}

// PagerReq and PagerRes helps for list* methods
type PagerReq struct {
	Marker string `protobuf:"bytes,1,opt,name=marker,proto3" json:"marker,omitempty"`
	Limit  int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (m *PagerReq) Reset()         { *m = PagerReq{} }
func (m *PagerReq) String() string { return proto.CompactTextString(m) }
func (*PagerReq) ProtoMessage()    {}
func (*PagerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{3}
}
func (m *PagerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PagerReq.Unmarshal(m, b)
}
func (m *PagerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PagerReq.Marshal(b, m, deterministic)
}
func (m *PagerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PagerReq.Merge(m, src)
}
func (m *PagerReq) XXX_Size() int {
	return xxx_messageInfo_PagerReq.Size(m)
}
func (m *PagerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PagerReq.DiscardUnknown(m)
}

var xxx_messageInfo_PagerReq proto.InternalMessageInfo

func (m *PagerReq) GetMarker() string {
	if m != nil {
		return m.Marker
	}
	return ""
}

func (m *PagerReq) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type PagerRes struct {
	Total    int32  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Marker   string `protobuf:"bytes,2,opt,name=marker,proto3" json:"marker,omitempty"`
	PageSize int32  `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (m *PagerRes) Reset()         { *m = PagerRes{} }
func (m *PagerRes) String() string { return proto.CompactTextString(m) }
func (*PagerRes) ProtoMessage()    {}
func (*PagerRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{4}
}
func (m *PagerRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PagerRes.Unmarshal(m, b)
}
func (m *PagerRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PagerRes.Marshal(b, m, deterministic)
}
func (m *PagerRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PagerRes.Merge(m, src)
}
func (m *PagerRes) XXX_Size() int {
	return xxx_messageInfo_PagerRes.Size(m)
}
func (m *PagerRes) XXX_DiscardUnknown() {
	xxx_messageInfo_PagerRes.DiscardUnknown(m)
}

var xxx_messageInfo_PagerRes proto.InternalMessageInfo

func (m *PagerRes) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *PagerRes) GetMarker() string {
	if m != nil {
		return m.Marker
	}
	return ""
}

func (m *PagerRes) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func init() {
	proto.RegisterEnum("apiserver.v2.SpecVersion", SpecVersion_name, SpecVersion_value)
	proto.RegisterEnum("apiserver.v2.ResourceKind", ResourceKind_name, ResourceKind_value)
	proto.RegisterType((*ResourceReference)(nil), "apiserver.v2.ResourceReference")
	proto.RegisterType((*ID)(nil), "apiserver.v2.ID")
	proto.RegisterType((*Metadata)(nil), "apiserver.v2.Metadata")
	proto.RegisterMapType((map[string]string)(nil), "apiserver.v2.Metadata.TagsEntry")
	proto.RegisterType((*PagerReq)(nil), "apiserver.v2.PagerReq")
	proto.RegisterType((*PagerRes)(nil), "apiserver.v2.PagerRes")
}

func init() { proto.RegisterFile("meta.proto", fileDescriptor_3b5ea8fe65782bcc) }

var fileDescriptor_3b5ea8fe65782bcc = []byte{
	// 901 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xcf, 0x73, 0xdb, 0x44,
	0x14, 0xc7, 0x2d, 0xf9, 0x47, 0xec, 0x75, 0x29, 0xca, 0x36, 0x4d, 0x14, 0x03, 0xb6, 0xd1, 0x70,
	0xf0, 0x74, 0x62, 0x99, 0x71, 0x67, 0xa0, 0xe4, 0x44, 0xdd, 0x00, 0x13, 0x18, 0xa6, 0xa9, 0xf2,
	0xe3, 0x90, 0x0e, 0x74, 0xd6, 0xd2, 0xb3, 0xba, 0x63, 0x4b, 0x2b, 0x56, 0x6b, 0xb7, 0x2e, 0xc3,
	0x85, 0xbf, 0xa0, 0x7f, 0x06, 0x77, 0xfe, 0x01, 0x4e, 0x9c, 0x7b, 0xe4, 0x14, 0x06, 0xf2, 0x1f,
	0xe4, 0xd8, 0x4b, 0x18, 0xad, 0x6c, 0x6b, 0xed, 0x74, 0x80, 0xcc, 0xf4, 0xa6, 0xf7, 0xb4, 0xdf,
	0xcf, 0x7b, 0xfb, 0xdd, 0x7d, 0xb3, 0x08, 0x05, 0x20, 0x88, 0x1d, 0x71, 0x26, 0x18, 0xbe, 0x41,
	0x22, 0x1a, 0x03, 0x9f, 0x00, 0xb7, 0x27, 0xdd, 0x5a, 0xdb, 0xa7, 0xe2, 0xe9, 0xb8, 0x6f, 0xbb,
	0x2c, 0xe8, 0xf8, 0xcc, 0x67, 0x1d, 0xb9, 0xa8, 0x3f, 0x1e, 0xc8, 0x48, 0x06, 0xf2, 0x2b, 0x15,
	0xd7, 0xee, 0x2b, 0xcb, 0x21, 0x9c, 0xb0, 0x69, 0xc4, 0xd9, 0xf3, 0x69, 0x2a, 0x72, 0xdb, 0x3e,
	0x84, 0xed, 0x09, 0x19, 0x51, 0x8f, 0x08, 0xe8, 0x5c, 0xf9, 0x98, 0x21, 0x1a, 0x3e, 0x63, 0xfe,
	0x08, 0xb2, 0x42, 0x82, 0x06, 0x10, 0x0b, 0x12, 0x44, 0xe9, 0x02, 0xeb, 0x77, 0x1d, 0xad, 0x3b,
	0x10, 0xb3, 0x31, 0x77, 0xc1, 0x81, 0x01, 0x70, 0x08, 0x5d, 0xc0, 0x8f, 0x51, 0x35, 0x8e, 0xc0,
	0x3d, 0x01, 0x1e, 0x53, 0x16, 0x9a, 0x5a, 0x53, 0x6b, 0xdd, 0xec, 0x6e, 0xdb, 0xea, 0x66, 0xec,
	0xc3, 0x6c, 0x41, 0xaf, 0xfe, 0xba, 0x57, 0xfc, 0x59, 0xd3, 0x0d, 0xed, 0xe2, 0xac, 0x81, 0xfb,
	0x31, 0x0b, 0x77, 0x2d, 0x45, 0x6f, 0x39, 0x2a, 0x0d, 0x7f, 0x85, 0x0a, 0x43, 0x1a, 0x7a, 0xa6,
	0x2e, 0xa9, 0xb5, 0x65, 0xea, 0xbc, 0x97, 0x6f, 0x68, 0xe8, 0xf5, 0xb6, 0x14, 0x6c, 0x35, 0xc5,
	0x26, 0x4a, 0xcb, 0x91, 0x00, 0xfc, 0x39, 0x2a, 0x84, 0x24, 0x00, 0x33, 0xdf, 0xd4, 0x5a, 0x95,
	0xde, 0xce, 0xeb, 0x5e, 0x8d, 0x9b, 0xdd, 0xcd, 0xef, 0x1f, 0x93, 0xf6, 0x8b, 0xfb, 0xed, 0xd3,
	0x8f, 0xdb, 0x9f, 0xb5, 0x9f, 0xd8, 0xdf, 0xfd, 0xd8, 0xdd, 0xf9, 0xe4, 0xee, 0x4f, 0x1f, 0x65,
	0x84, 0x44, 0x62, 0x39, 0x52, 0x89, 0xf7, 0xd1, 0x9a, 0xcb, 0x81, 0x08, 0xc6, 0xcd, 0x82, 0x84,
	0x74, 0xfe, 0x0b, 0x72, 0x33, 0x85, 0xcc, 0x54, 0x96, 0x33, 0xd7, 0x5b, 0x9b, 0x48, 0xdf, 0xdf,
	0xc3, 0x06, 0xca, 0x3f, 0x85, 0xe7, 0xd2, 0xb0, 0x8a, 0x93, 0x7c, 0x5a, 0xbf, 0x96, 0x50, 0xf9,
	0x5b, 0x10, 0xc4, 0x23, 0x82, 0x60, 0xb8, 0xa6, 0xaf, 0x2d, 0xc5, 0x80, 0xf7, 0xaf, 0xf8, 0xba,
	0xc3, 0x02, 0x2a, 0x20, 0x88, 0xc4, 0x74, 0xc5, 0xe1, 0x47, 0xff, 0xdb, 0xe1, 0x0f, 0x95, 0x02,
	0xb7, 0x33, 0x87, 0x55, 0xf2, 0xdb, 0xf2, 0xba, 0x8d, 0x0a, 0x1e, 0xc4, 0xee, 0xcc, 0xe8, 0xed,
	0xac, 0x60, 0x92, 0x5d, 0x2a, 0x98, 0x24, 0xf0, 0x4e, 0x76, 0x34, 0x45, 0xa9, 0xc0, 0xff, 0xe2,
	0x3e, 0x3e, 0x45, 0x65, 0xf6, 0x2c, 0x04, 0xee, 0xc0, 0xc0, 0x2c, 0x35, 0xb5, 0x56, 0xb5, 0xdb,
	0x78, 0xf3, 0xae, 0x17, 0x77, 0xbc, 0xf7, 0xc1, 0xc5, 0x59, 0x63, 0x3b, 0xe5, 0xcd, 0xa5, 0x6a,
	0x17, 0x0b, 0x1e, 0x3e, 0x46, 0x48, 0x96, 0x81, 0x23, 0x1a, 0x80, 0xb9, 0x26, 0xe9, 0x35, 0x3b,
	0x1d, 0x2c, 0x7b, 0x3e, 0x58, 0xf6, 0xd1, 0x7c, 0xb0, 0xe4, 0xd6, 0xd6, 0x95, 0x46, 0xa5, 0xce,
	0x7a, 0xf9, 0x67, 0x43, 0x73, 0x14, 0x50, 0x82, 0x1d, 0x47, 0xde, 0x1c, 0x5b, 0xbe, 0x0e, 0x36,
	0xd3, 0xcd, 0xb0, 0x59, 0x02, 0x3f, 0x44, 0x05, 0x41, 0xfc, 0xd8, 0xac, 0x34, 0xf3, 0xad, 0x6a,
	0xb7, 0xb9, 0xec, 0xc2, 0xfc, 0x22, 0xda, 0x47, 0xc4, 0x8f, 0xbf, 0x08, 0x05, 0x9f, 0xaa, 0x07,
	0x91, 0xe8, 0x96, 0x0e, 0x22, 0x49, 0xe0, 0x5d, 0xa4, 0x53, 0xcf, 0x44, 0xb2, 0x3f, 0x63, 0x19,
	0xb7, 0xbf, 0xd7, 0x33, 0x2f, 0xce, 0x1a, 0x1b, 0xa9, 0xfc, 0x09, 0x5d, 0xba, 0x37, 0x3a, 0xf5,
	0x6a, 0x9f, 0xa2, 0xca, 0xa2, 0x52, 0x32, 0x1b, 0x43, 0x98, 0xce, 0x67, 0x63, 0x08, 0x53, 0xbc,
	0x81, 0x8a, 0x13, 0x32, 0x1a, 0x83, 0xbc, 0xa8, 0x15, 0x27, 0x0d, 0x76, 0xf5, 0x7b, 0x9a, 0x75,
	0x0f, 0x95, 0x0f, 0x88, 0x9f, 0xf8, 0xff, 0x03, 0xde, 0x44, 0xa5, 0x80, 0xf0, 0x21, 0xf0, 0x99,
	0x74, 0x16, 0x25, 0xea, 0x11, 0x0d, 0xa8, 0x90, 0xea, 0xa2, 0x93, 0x06, 0xd6, 0xd1, 0x42, 0x19,
	0x27, 0x2b, 0x04, 0x13, 0x64, 0x24, 0x85, 0x45, 0x27, 0x0d, 0x14, 0x9e, 0xbe, 0xc4, 0xab, 0xa1,
	0x72, 0x44, 0x7c, 0x38, 0xa4, 0x2f, 0xd2, 0x6b, 0x5e, 0x74, 0x16, 0xf1, 0x9d, 0x2f, 0x51, 0x55,
	0x99, 0x4b, 0xbc, 0x81, 0x0c, 0x25, 0x3c, 0x0e, 0x63, 0x10, 0x46, 0x0e, 0xaf, 0xa3, 0x77, 0x94,
	0xec, 0x49, 0xd7, 0xd0, 0xad, 0x42, 0x59, 0x33, 0x34, 0xab, 0x5c, 0xce, 0x1b, 0x97, 0x97, 0x97,
	0x97, 0x6b, 0x77, 0x7e, 0xd3, 0xd1, 0x0d, 0x75, 0x00, 0xf1, 0x16, 0xba, 0xa5, 0xc6, 0xc7, 0xe1,
	0x30, 0x64, 0xcf, 0x42, 0x23, 0x87, 0x37, 0x11, 0x56, 0x7f, 0x9c, 0xb0, 0xd1, 0x38, 0x00, 0x43,
	0x5b, 0x15, 0xec, 0x11, 0x41, 0x92, 0xea, 0x3a, 0xae, 0xa3, 0xda, 0x1b, 0x7e, 0xcc, 0x9a, 0x31,
	0xf2, 0xab, 0xc0, 0x43, 0x70, 0x39, 0x08, 0xa3, 0x80, 0x6f, 0xa1, 0x77, 0xd5, 0xfc, 0xd7, 0xac,
	0x6f, 0x14, 0xf1, 0x7b, 0x68, 0x6b, 0xa9, 0xad, 0x18, 0xf8, 0x01, 0x67, 0x03, 0x3a, 0x02, 0xa3,
	0xb4, 0xda, 0xc2, 0x01, 0x71, 0x87, 0xc4, 0x07, 0x63, 0xed, 0x4a, 0x0b, 0x30, 0xa1, 0x2e, 0x3c,
	0x20, 0x02, 0x7c, 0xc6, 0xa7, 0x46, 0x79, 0xb5, 0xd4, 0x43, 0x0e, 0x46, 0x65, 0x95, 0x76, 0x28,
	0x18, 0x4f, 0x68, 0x08, 0xdf, 0xce, 0x5e, 0xa6, 0xe4, 0xc7, 0xa3, 0x31, 0x13, 0xc4, 0xa8, 0xf6,
	0x1e, 0xbc, 0xfa, 0xab, 0x9e, 0xfb, 0xe5, 0xef, 0xba, 0xf6, 0xf2, 0xbc, 0x9e, 0x7b, 0x75, 0x5e,
	0xcf, 0xfd, 0x71, 0x5e, 0xcf, 0x9d, 0x26, 0x4f, 0xab, 0x1d, 0x8f, 0x23, 0x0e, 0x01, 0x0d, 0x3d,
	0x9b, 0x86, 0x03, 0xf9, 0xba, 0x7a, 0x63, 0x57, 0xc4, 0x1d, 0x22, 0x58, 0xd0, 0x99, 0x3f, 0xaf,
	0x1d, 0x12, 0xd1, 0x7e, 0x49, 0x46, 0x77, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x36, 0x5d,
	0xd0, 0xac, 0x07, 0x00, 0x00,
}

func (this *ResourceReference) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResourceReference)
	if !ok {
		that2, ok := that.(ResourceReference)
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
	if this.SpecVersion != that1.SpecVersion {
		return false
	}
	if this.Kind != that1.Kind {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Creator != that1.Creator {
		return false
	}
	return true
}
func (this *ID) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ID)
	if !ok {
		that2, ok := that.(ID)
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
	if this.Hex != that1.Hex {
		return false
	}
	return true
}
func (this *Metadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Metadata)
	if !ok {
		that2, ok := that.(Metadata)
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
	if this.SpecVersion != that1.SpecVersion {
		return false
	}
	if this.Kind != that1.Kind {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Desc != that1.Desc {
		return false
	}
	if this.Creator != that1.Creator {
		return false
	}
	if !this.OwnerRef.Equal(that1.OwnerRef) {
		return false
	}
	if that1.CreateTime == nil {
		if this.CreateTime != nil {
			return false
		}
	} else if !this.CreateTime.Equal(*that1.CreateTime) {
		return false
	}
	if that1.UpdateTime == nil {
		if this.UpdateTime != nil {
			return false
		}
	} else if !this.UpdateTime.Equal(*that1.UpdateTime) {
		return false
	}
	if len(this.Tags) != len(that1.Tags) {
		return false
	}
	for i := range this.Tags {
		if this.Tags[i] != that1.Tags[i] {
			return false
		}
	}
	if !this.Id.Equal(that1.Id) {
		return false
	}
	return true
}
func (this *PagerReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PagerReq)
	if !ok {
		that2, ok := that.(PagerReq)
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
	if this.Marker != that1.Marker {
		return false
	}
	if this.Limit != that1.Limit {
		return false
	}
	return true
}
func (this *PagerRes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PagerRes)
	if !ok {
		that2, ok := that.(PagerRes)
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
	if this.Total != that1.Total {
		return false
	}
	if this.Marker != that1.Marker {
		return false
	}
	if this.PageSize != that1.PageSize {
		return false
	}
	return true
}
