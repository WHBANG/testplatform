// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: distributed.proto

package api

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type DistributedFramework int32

const (
	UnknownFramework DistributedFramework = 0
	Caffe            DistributedFramework = 1
	MXNet            DistributedFramework = 2
	TensorFlow       DistributedFramework = 3
	PyTorch          DistributedFramework = 4
)

var DistributedFramework_name = map[int32]string{
	0: "UnknownFramework",
	1: "Caffe",
	2: "MXNet",
	3: "TensorFlow",
	4: "PyTorch",
}

var DistributedFramework_value = map[string]int32{
	"UnknownFramework": 0,
	"Caffe":            1,
	"MXNet":            2,
	"TensorFlow":       3,
	"PyTorch":          4,
}

func (x DistributedFramework) String() string {
	return proto.EnumName(DistributedFramework_name, int32(x))
}

func (DistributedFramework) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_31930b1b278a4150, []int{0}
}

type DistributedReplicaType int32

const (
	UnknownReplicaType   DistributedReplicaType = 0
	ReplicaTypeServer    DistributedReplicaType = 1
	ReplicaTypeWorker    DistributedReplicaType = 2
	ReplicaTypeMaster    DistributedReplicaType = 3
	ReplicaTypeScheduler DistributedReplicaType = 4
)

var DistributedReplicaType_name = map[int32]string{
	0: "UnknownReplicaType",
	1: "ReplicaTypeServer",
	2: "ReplicaTypeWorker",
	3: "ReplicaTypeMaster",
	4: "ReplicaTypeScheduler",
}

var DistributedReplicaType_value = map[string]int32{
	"UnknownReplicaType":   0,
	"ReplicaTypeServer":    1,
	"ReplicaTypeWorker":    2,
	"ReplicaTypeMaster":    3,
	"ReplicaTypeScheduler": 4,
}

func (x DistributedReplicaType) String() string {
	return proto.EnumName(DistributedReplicaType_name, int32(x))
}

func (DistributedReplicaType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_31930b1b278a4150, []int{1}
}

type DistributedSpec struct {
	Framework    DistributedFramework  `protobuf:"varint,1,opt,name=framework,proto3,enum=apiserver.v2.DistributedFramework" json:"framework,omitempty" bson:"framework"`
	Replicas     []*DistributedReplica `protobuf:"bytes,2,rep,name=replicas,proto3" json:"replicas,omitempty" bson:"replicas,omitempty"`
	EnableLogger bool                  `protobuf:"varint,3,opt,name=enableLogger,proto3" json:"enableLogger,omitempty" bson:"enableLogger,omitempty"`
}

func (m *DistributedSpec) Reset()         { *m = DistributedSpec{} }
func (m *DistributedSpec) String() string { return proto.CompactTextString(m) }
func (*DistributedSpec) ProtoMessage()    {}
func (*DistributedSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_31930b1b278a4150, []int{0}
}
func (m *DistributedSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributedSpec.Unmarshal(m, b)
}
func (m *DistributedSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributedSpec.Marshal(b, m, deterministic)
}
func (m *DistributedSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributedSpec.Merge(m, src)
}
func (m *DistributedSpec) XXX_Size() int {
	return xxx_messageInfo_DistributedSpec.Size(m)
}
func (m *DistributedSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributedSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DistributedSpec proto.InternalMessageInfo

func (m *DistributedSpec) GetFramework() DistributedFramework {
	if m != nil {
		return m.Framework
	}
	return UnknownFramework
}

func (m *DistributedSpec) GetReplicas() []*DistributedReplica {
	if m != nil {
		return m.Replicas
	}
	return nil
}

func (m *DistributedSpec) GetEnableLogger() bool {
	if m != nil {
		return m.EnableLogger
	}
	return false
}

type DistributedReplica struct {
	Type     DistributedReplicaType `protobuf:"varint,1,opt,name=type,proto3,enum=apiserver.v2.DistributedReplicaType" json:"type,omitempty" bson:"type"`
	Replicas int32                  `protobuf:"varint,2,opt,name=replicas,proto3" json:"replicas,omitempty" bson:"replicas"`
	Package  *ResourceReference     `protobuf:"bytes,3,opt,name=package,proto3" json:"package,omitempty" bson:"package,omitempty"`
}

func (m *DistributedReplica) Reset()         { *m = DistributedReplica{} }
func (m *DistributedReplica) String() string { return proto.CompactTextString(m) }
func (*DistributedReplica) ProtoMessage()    {}
func (*DistributedReplica) Descriptor() ([]byte, []int) {
	return fileDescriptor_31930b1b278a4150, []int{1}
}
func (m *DistributedReplica) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributedReplica.Unmarshal(m, b)
}
func (m *DistributedReplica) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributedReplica.Marshal(b, m, deterministic)
}
func (m *DistributedReplica) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributedReplica.Merge(m, src)
}
func (m *DistributedReplica) XXX_Size() int {
	return xxx_messageInfo_DistributedReplica.Size(m)
}
func (m *DistributedReplica) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributedReplica.DiscardUnknown(m)
}

var xxx_messageInfo_DistributedReplica proto.InternalMessageInfo

func (m *DistributedReplica) GetType() DistributedReplicaType {
	if m != nil {
		return m.Type
	}
	return UnknownReplicaType
}

func (m *DistributedReplica) GetReplicas() int32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

func (m *DistributedReplica) GetPackage() *ResourceReference {
	if m != nil {
		return m.Package
	}
	return nil
}

type DistributedStatus struct {
	Replicas []*ReplicaStatus `protobuf:"bytes,4,rep,name=replicas,proto3" json:"replicas,omitempty" bson:"replicas,omitempty"`
}

func (m *DistributedStatus) Reset()         { *m = DistributedStatus{} }
func (m *DistributedStatus) String() string { return proto.CompactTextString(m) }
func (*DistributedStatus) ProtoMessage()    {}
func (*DistributedStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_31930b1b278a4150, []int{2}
}
func (m *DistributedStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributedStatus.Unmarshal(m, b)
}
func (m *DistributedStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributedStatus.Marshal(b, m, deterministic)
}
func (m *DistributedStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributedStatus.Merge(m, src)
}
func (m *DistributedStatus) XXX_Size() int {
	return xxx_messageInfo_DistributedStatus.Size(m)
}
func (m *DistributedStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributedStatus.DiscardUnknown(m)
}

var xxx_messageInfo_DistributedStatus proto.InternalMessageInfo

func (m *DistributedStatus) GetReplicas() []*ReplicaStatus {
	if m != nil {
		return m.Replicas
	}
	return nil
}

type ReplicaStatus struct {
	Type     DistributedReplicaType `protobuf:"varint,1,opt,name=type,proto3,enum=apiserver.v2.DistributedReplicaType" json:"type,omitempty" bson:"type,omitempty"`
	Replicas int32                  `protobuf:"varint,2,opt,name=replicas,proto3" json:"replicas,omitempty" bson:"replicas,omitempty"`
	Resource *PackageSpec           `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty" bson:"resource,omitempty"`
}

func (m *ReplicaStatus) Reset()         { *m = ReplicaStatus{} }
func (m *ReplicaStatus) String() string { return proto.CompactTextString(m) }
func (*ReplicaStatus) ProtoMessage()    {}
func (*ReplicaStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_31930b1b278a4150, []int{3}
}
func (m *ReplicaStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplicaStatus.Unmarshal(m, b)
}
func (m *ReplicaStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplicaStatus.Marshal(b, m, deterministic)
}
func (m *ReplicaStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplicaStatus.Merge(m, src)
}
func (m *ReplicaStatus) XXX_Size() int {
	return xxx_messageInfo_ReplicaStatus.Size(m)
}
func (m *ReplicaStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplicaStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ReplicaStatus proto.InternalMessageInfo

func (m *ReplicaStatus) GetType() DistributedReplicaType {
	if m != nil {
		return m.Type
	}
	return UnknownReplicaType
}

func (m *ReplicaStatus) GetReplicas() int32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

func (m *ReplicaStatus) GetResource() *PackageSpec {
	if m != nil {
		return m.Resource
	}
	return nil
}

func init() {
	proto.RegisterEnum("apiserver.v2.DistributedFramework", DistributedFramework_name, DistributedFramework_value)
	proto.RegisterEnum("apiserver.v2.DistributedReplicaType", DistributedReplicaType_name, DistributedReplicaType_value)
	proto.RegisterType((*DistributedSpec)(nil), "apiserver.v2.DistributedSpec")
	proto.RegisterType((*DistributedReplica)(nil), "apiserver.v2.DistributedReplica")
	proto.RegisterType((*DistributedStatus)(nil), "apiserver.v2.DistributedStatus")
	proto.RegisterType((*ReplicaStatus)(nil), "apiserver.v2.ReplicaStatus")
}

func init() { proto.RegisterFile("distributed.proto", fileDescriptor_31930b1b278a4150) }

var fileDescriptor_31930b1b278a4150 = []byte{
	// 644 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xe3, 0x24, 0xfd, 0xb5, 0xdd, 0xf4, 0xcf, 0x76, 0x7f, 0x6d, 0x49, 0x4b, 0x1b, 0x07,
	0x8b, 0x43, 0x54, 0x51, 0x5b, 0x0a, 0x27, 0xb8, 0xe1, 0x42, 0x4f, 0x14, 0x95, 0x34, 0x50, 0x04,
	0x12, 0x68, 0xe3, 0x4c, 0x5c, 0x2b, 0xb1, 0xd7, 0x5a, 0xaf, 0x53, 0x72, 0xe5, 0x05, 0xe0, 0xc2,
	0x3b, 0xf0, 0x38, 0x3d, 0x72, 0x8a, 0x04, 0x7d, 0x02, 0x22, 0x71, 0xe1, 0x84, 0xe2, 0x75, 0xda,
	0x75, 0x5a, 0x44, 0xc5, 0x6d, 0x76, 0x76, 0xbe, 0x9f, 0x9d, 0xf9, 0xae, 0xbd, 0x68, 0xa5, 0xed,
	0x45, 0x82, 0x7b, 0xad, 0x58, 0x40, 0xdb, 0x0c, 0x39, 0x13, 0x8c, 0x2c, 0xd0, 0xd0, 0x8b, 0x80,
	0xf7, 0x81, 0x9b, 0xfd, 0xfa, 0xe6, 0xae, 0xeb, 0x89, 0x93, 0xb8, 0x65, 0x3a, 0xcc, 0xb7, 0x5c,
	0xe6, 0x32, 0x2b, 0x29, 0x6a, 0xc5, 0x9d, 0x64, 0x95, 0x2c, 0x92, 0x48, 0x8a, 0x37, 0x1f, 0x29,
	0xe5, 0x10, 0xf4, 0xd9, 0x20, 0xe4, 0xec, 0xfd, 0x40, 0x8a, 0x9c, 0x5d, 0x17, 0x82, 0xdd, 0x3e,
	0xed, 0x79, 0x6d, 0x2a, 0xc0, 0xba, 0x12, 0xa4, 0x08, 0xe4, 0x83, 0xa0, 0x69, 0xbc, 0x18, 0x52,
	0xa7, 0x4b, 0xdd, 0x74, 0xcb, 0xf8, 0x9c, 0x47, 0xcb, 0x8f, 0x2f, 0x1b, 0x3e, 0x0a, 0xc1, 0x21,
	0x6f, 0xd1, 0x7c, 0x87, 0x53, 0x1f, 0x4e, 0x19, 0xef, 0x96, 0xb5, 0xaa, 0x56, 0x5b, 0xaa, 0x1b,
	0xa6, 0x3a, 0x82, 0xa9, 0x28, 0xf6, 0x27, 0x95, 0xf6, 0xd6, 0x2f, 0x7b, 0xe6, 0x83, 0x96, 0xc7,
	0xda, 0x68, 0xa8, 0xe3, 0x56, 0xc4, 0x82, 0x87, 0xc6, 0x05, 0xc6, 0x68, 0x5c, 0x22, 0xc9, 0x1b,
	0x34, 0xc7, 0x21, 0xec, 0x79, 0x0e, 0x8d, 0xca, 0xf9, 0x6a, 0xa1, 0x56, 0xaa, 0x57, 0xff, 0x88,
	0x6f, 0xc8, 0x42, 0x7b, 0x7b, 0x34, 0xd4, 0x37, 0x24, 0x74, 0xa2, 0xbd, 0xc7, 0x7c, 0x4f, 0x80,
	0x1f, 0x8a, 0x81, 0xd1, 0xb8, 0x00, 0x92, 0x27, 0x68, 0x01, 0x02, 0xda, 0xea, 0xc1, 0x53, 0xe6,
	0xba, 0xc0, 0xcb, 0x85, 0xaa, 0x56, 0x9b, 0xb3, 0xef, 0x8c, 0x86, 0xfa, 0xb6, 0x94, 0xab, 0xbb,
	0x2a, 0x22, 0x23, 0x33, 0x7e, 0x68, 0x88, 0x5c, 0x6d, 0x83, 0x3c, 0x47, 0x45, 0x31, 0x08, 0x21,
	0x75, 0xe5, 0xee, 0xdf, 0xda, 0x6e, 0x0e, 0x42, 0xb0, 0x6f, 0x29, 0xbe, 0x94, 0x64, 0x0f, 0x63,
	0x86, 0xd1, 0x48, 0x50, 0xc4, 0xca, 0xb8, 0xa1, 0xd5, 0x66, 0xec, 0xff, 0x47, 0x43, 0x7d, 0x39,
	0x3b, 0xab, 0x3a, 0xe1, 0x31, 0x9a, 0x4d, 0xef, 0x30, 0x19, 0xae, 0x54, 0xd7, 0xb3, 0x6d, 0x34,
	0x20, 0x62, 0x31, 0x77, 0xa0, 0x01, 0x1d, 0xe0, 0x10, 0x38, 0x60, 0x6f, 0x8d, 0x86, 0x7a, 0x59,
	0x02, 0x53, 0xa5, 0x3a, 0xf8, 0x84, 0x66, 0x74, 0xd1, 0x8a, 0xfa, 0x29, 0x08, 0x2a, 0xe2, 0x88,
	0xbc, 0x54, 0xda, 0x2b, 0x26, 0x97, 0x75, 0x7b, 0xfa, 0xb8, 0x64, 0x57, 0x96, 0xdf, 0xf8, 0x9e,
	0x8c, 0x9f, 0x1a, 0x5a, 0xcc, 0x48, 0xc9, 0xd1, 0x3f, 0x78, 0xbb, 0x31, 0x1a, 0xea, 0x6b, 0x97,
	0x9e, 0xaa, 0x47, 0x49, 0x77, 0x1f, 0x5c, 0x71, 0xf7, 0xc6, 0x5f, 0x52, 0x73, 0x2c, 0x95, 0x56,
	0xa6, 0x46, 0x6f, 0x64, 0x7b, 0x3a, 0x94, 0xbe, 0x8d, 0xff, 0x99, 0x2c, 0x55, 0x8a, 0xa6, 0xa8,
	0x32, 0xb9, 0xf3, 0x0e, 0xad, 0x5e, 0xf7, 0xf7, 0x90, 0x55, 0x84, 0x5f, 0x04, 0xdd, 0x80, 0x9d,
	0x06, 0x17, 0x39, 0x9c, 0x23, 0xf3, 0x68, 0x66, 0x8f, 0x76, 0x3a, 0x80, 0xb5, 0x71, 0x78, 0xf0,
	0xea, 0x19, 0x08, 0x9c, 0x27, 0x4b, 0x08, 0x35, 0x21, 0x88, 0x18, 0xdf, 0xef, 0xb1, 0x53, 0x5c,
	0x20, 0x25, 0x34, 0x7b, 0x38, 0x68, 0x32, 0xee, 0x9c, 0xe0, 0xe2, 0xce, 0x47, 0x0d, 0xad, 0x5f,
	0xef, 0x16, 0x59, 0x47, 0x24, 0x3d, 0x43, 0xc9, 0xe2, 0x1c, 0x59, 0x43, 0x2b, 0x4a, 0xe2, 0x28,
	0x19, 0x10, 0x6b, 0x53, 0xe9, 0x63, 0xc6, 0xbb, 0xc0, 0x71, 0x7e, 0x2a, 0x7d, 0x40, 0x23, 0x01,
	0x1c, 0x17, 0x48, 0x19, 0xad, 0xaa, 0x10, 0xe7, 0x04, 0xda, 0x71, 0x0f, 0x38, 0x2e, 0xda, 0x7b,
	0x67, 0xdf, 0x2a, 0xb9, 0x2f, 0xdf, 0x2b, 0xda, 0xa7, 0xf3, 0x4a, 0xee, 0xec, 0xbc, 0x92, 0xfb,
	0x7a, 0x5e, 0xc9, 0xbd, 0x1e, 0x3f, 0x83, 0x66, 0x14, 0x87, 0x1c, 0x7c, 0x2f, 0x68, 0x9b, 0x5e,
	0xd0, 0x49, 0x5e, 0xc2, 0x76, 0xec, 0x88, 0xc8, 0xa2, 0x82, 0xf9, 0xd6, 0xe4, 0x29, 0xb4, 0x68,
	0xe8, 0xb5, 0xfe, 0x4b, 0x56, 0xf7, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x96, 0xb5, 0xde, 0xc6,
	0x5f, 0x05, 0x00, 0x00,
}

func (this *DistributedSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DistributedSpec)
	if !ok {
		that2, ok := that.(DistributedSpec)
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
	if this.Framework != that1.Framework {
		return false
	}
	if len(this.Replicas) != len(that1.Replicas) {
		return false
	}
	for i := range this.Replicas {
		if !this.Replicas[i].Equal(that1.Replicas[i]) {
			return false
		}
	}
	if this.EnableLogger != that1.EnableLogger {
		return false
	}
	return true
}
func (this *DistributedReplica) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DistributedReplica)
	if !ok {
		that2, ok := that.(DistributedReplica)
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
	if this.Type != that1.Type {
		return false
	}
	if this.Replicas != that1.Replicas {
		return false
	}
	if !this.Package.Equal(that1.Package) {
		return false
	}
	return true
}
func (this *DistributedStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DistributedStatus)
	if !ok {
		that2, ok := that.(DistributedStatus)
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
	if len(this.Replicas) != len(that1.Replicas) {
		return false
	}
	for i := range this.Replicas {
		if !this.Replicas[i].Equal(that1.Replicas[i]) {
			return false
		}
	}
	return true
}
func (this *ReplicaStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReplicaStatus)
	if !ok {
		that2, ok := that.(ReplicaStatus)
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
	if this.Type != that1.Type {
		return false
	}
	if this.Replicas != that1.Replicas {
		return false
	}
	if !this.Resource.Equal(that1.Resource) {
		return false
	}
	return true
}
