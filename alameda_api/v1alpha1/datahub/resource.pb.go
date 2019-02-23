// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/resource.proto

package containers_ai_alameda_v1alpha1_datahub

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// *
// Represents kubernetes resource kind
//
type Kind int32

const (
	Kind_POD              Kind = 0
	Kind_DEPLOYMENT       Kind = 1
	Kind_DEPLOYMENTCONFIG Kind = 2
)

var Kind_name = map[int32]string{
	0: "POD",
	1: "DEPLOYMENT",
	2: "DEPLOYMENTCONFIG",
}
var Kind_value = map[string]int32{
	"POD":              0,
	"DEPLOYMENT":       1,
	"DEPLOYMENTCONFIG": 2,
}

func (x Kind) String() string {
	return proto.EnumName(Kind_name, int32(x))
}
func (Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_resource_9fd046cd277717cb, []int{0}
}

// *
// Represents a container and its containing limit and requeset configurations
//
type Container struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LimitResource        []*MetricData `protobuf:"bytes,2,rep,name=limit_resource,json=limitResource,proto3" json:"limit_resource,omitempty"`
	RequestResource      []*MetricData `protobuf:"bytes,3,rep,name=request_resource,json=requestResource,proto3" json:"request_resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_9fd046cd277717cb, []int{0}
}
func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (dst *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(dst, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Container) GetLimitResource() []*MetricData {
	if m != nil {
		return m.LimitResource
	}
	return nil
}

func (m *Container) GetRequestResource() []*MetricData {
	if m != nil {
		return m.RequestResource
	}
	return nil
}

// *
// Represents a Kubernetes pod
//
type Pod struct {
	NamespacedName       *NamespacedName      `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	ResourceLink         string               `protobuf:"bytes,2,opt,name=resource_link,json=resourceLink,proto3" json:"resource_link,omitempty"`
	Containers           []*Container         `protobuf:"bytes,3,rep,name=containers,proto3" json:"containers,omitempty"`
	IsAlameda            bool                 `protobuf:"varint,4,opt,name=is_alameda,json=isAlameda,proto3" json:"is_alameda,omitempty"`
	AlamedaScaler        *NamespacedName      `protobuf:"bytes,5,opt,name=alameda_scaler,json=alamedaScaler,proto3" json:"alameda_scaler,omitempty"`
	NodeName             string               `protobuf:"bytes,6,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Policy               RecommendationPolicy `protobuf:"varint,8,opt,name=policy,proto3,enum=containers_ai.alameda.v1alpha1.datahub.RecommendationPolicy" json:"policy,omitempty"`
	TopController        *TopController       `protobuf:"bytes,9,opt,name=top_controller,json=topController,proto3" json:"top_controller,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Pod) Reset()         { *m = Pod{} }
func (m *Pod) String() string { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()    {}
func (*Pod) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_9fd046cd277717cb, []int{1}
}
func (m *Pod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pod.Unmarshal(m, b)
}
func (m *Pod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pod.Marshal(b, m, deterministic)
}
func (dst *Pod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pod.Merge(dst, src)
}
func (m *Pod) XXX_Size() int {
	return xxx_messageInfo_Pod.Size(m)
}
func (m *Pod) XXX_DiscardUnknown() {
	xxx_messageInfo_Pod.DiscardUnknown(m)
}

var xxx_messageInfo_Pod proto.InternalMessageInfo

func (m *Pod) GetNamespacedName() *NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *Pod) GetResourceLink() string {
	if m != nil {
		return m.ResourceLink
	}
	return ""
}

func (m *Pod) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Pod) GetIsAlameda() bool {
	if m != nil {
		return m.IsAlameda
	}
	return false
}

func (m *Pod) GetAlamedaScaler() *NamespacedName {
	if m != nil {
		return m.AlamedaScaler
	}
	return nil
}

func (m *Pod) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Pod) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Pod) GetPolicy() RecommendationPolicy {
	if m != nil {
		return m.Policy
	}
	return RecommendationPolicy_RECOMMENDATIONPOLICY_UNDEFINED
}

func (m *Pod) GetTopController() *TopController {
	if m != nil {
		return m.TopController
	}
	return nil
}

// *
// Represents the capacity of a Kubernetes node
//
type Capacity struct {
	CpuCores                 int64    `protobuf:"varint,1,opt,name=cpu_cores,json=cpuCores,proto3" json:"cpu_cores,omitempty"`
	MemoryBytes              int64    `protobuf:"varint,2,opt,name=memory_bytes,json=memoryBytes,proto3" json:"memory_bytes,omitempty"`
	NetwotkMegabitsPerSecond int64    `protobuf:"varint,3,opt,name=netwotk_megabits_per_second,json=netwotkMegabitsPerSecond,proto3" json:"netwotk_megabits_per_second,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *Capacity) Reset()         { *m = Capacity{} }
func (m *Capacity) String() string { return proto.CompactTextString(m) }
func (*Capacity) ProtoMessage()    {}
func (*Capacity) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_9fd046cd277717cb, []int{2}
}
func (m *Capacity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Capacity.Unmarshal(m, b)
}
func (m *Capacity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Capacity.Marshal(b, m, deterministic)
}
func (dst *Capacity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Capacity.Merge(dst, src)
}
func (m *Capacity) XXX_Size() int {
	return xxx_messageInfo_Capacity.Size(m)
}
func (m *Capacity) XXX_DiscardUnknown() {
	xxx_messageInfo_Capacity.DiscardUnknown(m)
}

var xxx_messageInfo_Capacity proto.InternalMessageInfo

func (m *Capacity) GetCpuCores() int64 {
	if m != nil {
		return m.CpuCores
	}
	return 0
}

func (m *Capacity) GetMemoryBytes() int64 {
	if m != nil {
		return m.MemoryBytes
	}
	return 0
}

func (m *Capacity) GetNetwotkMegabitsPerSecond() int64 {
	if m != nil {
		return m.NetwotkMegabitsPerSecond
	}
	return 0
}

// *
// Represents a Kubernetes node
//
type Node struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Capacity             *Capacity `protobuf:"bytes,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_9fd046cd277717cb, []int{3}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Node) GetCapacity() *Capacity {
	if m != nil {
		return m.Capacity
	}
	return nil
}

// *
// Represents top controller of the pod
//
type TopController struct {
	NamespacedName       *NamespacedName `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	Kind                 Kind            `protobuf:"varint,2,opt,name=kind,proto3,enum=containers_ai.alameda.v1alpha1.datahub.Kind" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TopController) Reset()         { *m = TopController{} }
func (m *TopController) String() string { return proto.CompactTextString(m) }
func (*TopController) ProtoMessage()    {}
func (*TopController) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_9fd046cd277717cb, []int{4}
}
func (m *TopController) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopController.Unmarshal(m, b)
}
func (m *TopController) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopController.Marshal(b, m, deterministic)
}
func (dst *TopController) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopController.Merge(dst, src)
}
func (m *TopController) XXX_Size() int {
	return xxx_messageInfo_TopController.Size(m)
}
func (m *TopController) XXX_DiscardUnknown() {
	xxx_messageInfo_TopController.DiscardUnknown(m)
}

var xxx_messageInfo_TopController proto.InternalMessageInfo

func (m *TopController) GetNamespacedName() *NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *TopController) GetKind() Kind {
	if m != nil {
		return m.Kind
	}
	return Kind_POD
}

func init() {
	proto.RegisterType((*Container)(nil), "containers_ai.alameda.v1alpha1.datahub.Container")
	proto.RegisterType((*Pod)(nil), "containers_ai.alameda.v1alpha1.datahub.Pod")
	proto.RegisterType((*Capacity)(nil), "containers_ai.alameda.v1alpha1.datahub.Capacity")
	proto.RegisterType((*Node)(nil), "containers_ai.alameda.v1alpha1.datahub.Node")
	proto.RegisterType((*TopController)(nil), "containers_ai.alameda.v1alpha1.datahub.TopController")
	proto.RegisterEnum("containers_ai.alameda.v1alpha1.datahub.Kind", Kind_name, Kind_value)
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/resource.proto", fileDescriptor_resource_9fd046cd277717cb)
}

var fileDescriptor_resource_9fd046cd277717cb = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xdf, 0x6a, 0x13, 0x41,
	0x14, 0xc6, 0xdd, 0x26, 0xb6, 0xc9, 0x49, 0xb3, 0x0d, 0x83, 0x17, 0x4b, 0x8b, 0x18, 0x23, 0x48,
	0xac, 0xb2, 0xb1, 0x91, 0x0a, 0x82, 0x82, 0x9a, 0x56, 0x11, 0xdb, 0x34, 0x4e, 0x73, 0x53, 0xb0,
	0x0c, 0x93, 0xd9, 0xb1, 0x1d, 0xb2, 0xbb, 0xb3, 0xce, 0x4c, 0x94, 0x3c, 0x82, 0xaf, 0xe3, 0xdb,
	0xf8, 0x06, 0x3e, 0x86, 0xec, 0x64, 0x36, 0x6d, 0x41, 0xea, 0xa2, 0x78, 0x37, 0xe7, 0xe4, 0x7c,
	0xbf, 0xf3, 0x27, 0xdf, 0xc2, 0x43, 0x1a, 0xd3, 0x84, 0x47, 0x94, 0xd0, 0x4c, 0xf4, 0xbe, 0xec,
	0xd0, 0x38, 0x3b, 0xa7, 0x3b, 0xbd, 0x88, 0x1a, 0x7a, 0x3e, 0x9b, 0xf4, 0x14, 0xd7, 0x72, 0xa6,
	0x18, 0x0f, 0x33, 0x25, 0x8d, 0x44, 0xf7, 0x99, 0x4c, 0x0d, 0x15, 0x29, 0x57, 0x9a, 0x50, 0x11,
	0x3a, 0x69, 0x58, 0xc8, 0x42, 0x27, 0xdb, 0xdc, 0xbe, 0x16, 0x9a, 0x29, 0x1e, 0x09, 0x66, 0x16,
	0xcc, 0xcd, 0x3b, 0x67, 0x52, 0x9e, 0xc5, 0xbc, 0x67, 0xa3, 0xc9, 0xec, 0x53, 0xcf, 0x88, 0x84,
	0x6b, 0x43, 0x93, 0xcc, 0x15, 0x5c, 0x3f, 0x61, 0xc2, 0x0d, 0xcd, 0xdf, 0xae, 0xf8, 0xc1, 0x9f,
	0x8a, 0x95, 0x60, 0x8b, 0xd2, 0xce, 0x0f, 0x0f, 0xea, 0x83, 0x62, 0x1f, 0x84, 0xa0, 0x9a, 0xd2,
	0x84, 0x07, 0x5e, 0xdb, 0xeb, 0xd6, 0xb1, 0x7d, 0xa3, 0x13, 0xf0, 0x63, 0x91, 0x08, 0x43, 0x8a,
	0x33, 0x04, 0x2b, 0xed, 0x4a, 0xb7, 0xd1, 0xef, 0x87, 0xe5, 0xee, 0x10, 0x1e, 0xda, 0x7e, 0x7b,
	0xd4, 0x50, 0xdc, 0xb4, 0x24, 0xec, 0x40, 0xe8, 0x14, 0x5a, 0x8a, 0x7f, 0x9e, 0x71, 0x7d, 0x09,
	0x5e, 0xf9, 0x6b, 0xf8, 0x86, 0x63, 0x15, 0xf8, 0xce, 0xcf, 0x2a, 0x54, 0x46, 0x32, 0x42, 0x04,
	0x36, 0xf2, 0x4d, 0x74, 0x46, 0x19, 0x8f, 0xc8, 0x72, 0xc1, 0x46, 0xff, 0x69, 0xd9, 0x2e, 0xc3,
	0xa5, 0x3c, 0x7f, 0x61, 0x3f, 0xbd, 0x12, 0xa3, 0x7b, 0xd0, 0x2c, 0xe6, 0x27, 0xb1, 0x48, 0xa7,
	0xc1, 0x8a, 0xbd, 0xdf, 0x7a, 0x91, 0x3c, 0x10, 0xe9, 0x14, 0x7d, 0x00, 0xb8, 0xe8, 0xe6, 0xd6,
	0xdc, 0x29, 0x3b, 0xc0, 0xf2, 0x2f, 0xc2, 0x97, 0x20, 0xe8, 0x36, 0x80, 0xd0, 0xc4, 0x89, 0x82,
	0x6a, 0xdb, 0xeb, 0xd6, 0x70, 0x5d, 0xe8, 0x57, 0x8b, 0x04, 0x3a, 0x05, 0xbf, 0x30, 0x82, 0x66,
	0x34, 0xe6, 0x2a, 0xb8, 0xf9, 0x4f, 0x6b, 0x37, 0x5d, 0xe1, 0xb1, 0x85, 0xa1, 0x2d, 0xa8, 0xa7,
	0x32, 0xe2, 0x8b, 0x83, 0xae, 0xda, 0x8d, 0x6b, 0x79, 0xc2, 0x9e, 0xe4, 0x19, 0x80, 0x36, 0x54,
	0x19, 0x92, 0x1b, 0x39, 0x58, 0xb3, 0x7d, 0x37, 0xc3, 0x85, 0xcb, 0xc3, 0xc2, 0xe5, 0xe1, 0xb8,
	0x70, 0x39, 0xae, 0xdb, 0xea, 0x3c, 0x46, 0x63, 0x58, 0xcd, 0x64, 0x2c, 0xd8, 0x3c, 0xa8, 0xb5,
	0xbd, 0xae, 0xdf, 0x7f, 0x5e, 0x76, 0x5c, 0xcc, 0x99, 0x4c, 0x12, 0x9e, 0x46, 0xd4, 0x08, 0x99,
	0x8e, 0x2c, 0x03, 0x3b, 0x16, 0xfa, 0x08, 0xbe, 0x91, 0x19, 0xc9, 0x51, 0x4a, 0xc6, 0xf9, 0x31,
	0xea, 0x76, 0xa8, 0xdd, 0xb2, 0xf4, 0xb1, 0xcc, 0x06, 0x4b, 0x31, 0x6e, 0x9a, 0xcb, 0x61, 0xe7,
	0x9b, 0x07, 0xb5, 0x01, 0xcd, 0x28, 0x13, 0x66, 0x9e, 0x1f, 0x86, 0x65, 0x33, 0xc2, 0xa4, 0xe2,
	0xda, 0x3a, 0xad, 0x82, 0x6b, 0x2c, 0x9b, 0x0d, 0xf2, 0x18, 0xdd, 0x85, 0xf5, 0x84, 0x27, 0x52,
	0xcd, 0xc9, 0x64, 0x6e, 0xb8, 0xb6, 0x56, 0xa9, 0xe0, 0xc6, 0x22, 0xf7, 0x3a, 0x4f, 0xa1, 0x17,
	0xb0, 0x95, 0x72, 0xf3, 0x55, 0x9a, 0x29, 0x49, 0xf8, 0x19, 0x9d, 0x08, 0xa3, 0x49, 0xc6, 0x15,
	0xd1, 0x9c, 0xc9, 0x34, 0x0a, 0x2a, 0x56, 0x11, 0xb8, 0x92, 0x43, 0x57, 0x31, 0xe2, 0xea, 0xd8,
	0xfe, 0xde, 0x39, 0x87, 0xea, 0x50, 0x46, 0xfc, 0xb7, 0x1f, 0xf3, 0x01, 0xd4, 0x98, 0x1b, 0xd3,
	0x76, 0x6e, 0xf4, 0x1f, 0x97, 0xb6, 0xa0, 0xd3, 0xe1, 0x25, 0xa1, 0xf3, 0xdd, 0x83, 0xe6, 0x95,
	0xb3, 0xfc, 0xff, 0x4f, 0xed, 0x25, 0x54, 0xa7, 0x22, 0x8d, 0xec, 0xf0, 0x7e, 0xff, 0x51, 0x59,
	0xea, 0x7b, 0x91, 0x46, 0xd8, 0x2a, 0xb7, 0x77, 0xa1, 0x9a, 0x47, 0x68, 0x0d, 0x2a, 0xa3, 0xa3,
	0xbd, 0xd6, 0x0d, 0xe4, 0x03, 0xec, 0xed, 0x8f, 0x0e, 0x8e, 0x4e, 0x0e, 0xf7, 0x87, 0xe3, 0x96,
	0x87, 0x6e, 0x41, 0xeb, 0x22, 0x1e, 0x1c, 0x0d, 0xdf, 0xbc, 0x7b, 0xdb, 0x5a, 0x99, 0xac, 0x5a,
	0xd3, 0x3e, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xf2, 0xe9, 0x6d, 0x2b, 0x06, 0x00, 0x00,
}
