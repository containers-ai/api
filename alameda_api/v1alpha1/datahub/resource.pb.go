// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/resource.proto

package containers_ai_alameda_v1alpha1_datahub

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NamespacedName struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NamespacedName) Reset()         { *m = NamespacedName{} }
func (m *NamespacedName) String() string { return proto.CompactTextString(m) }
func (*NamespacedName) ProtoMessage()    {}
func (*NamespacedName) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_42c7462b98f43aa8, []int{0}
}
func (m *NamespacedName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamespacedName.Unmarshal(m, b)
}
func (m *NamespacedName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamespacedName.Marshal(b, m, deterministic)
}
func (dst *NamespacedName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespacedName.Merge(dst, src)
}
func (m *NamespacedName) XXX_Size() int {
	return xxx_messageInfo_NamespacedName.Size(m)
}
func (m *NamespacedName) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespacedName.DiscardUnknown(m)
}

var xxx_messageInfo_NamespacedName proto.InternalMessageInfo

func (m *NamespacedName) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *NamespacedName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Container struct {
	Name                                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LimitResource                        []*MetricData `protobuf:"bytes,2,rep,name=limit_resource,json=limitResource,proto3" json:"limit_resource,omitempty"`
	RequestResource                      []*MetricData `protobuf:"bytes,3,rep,name=request_resource,json=requestResource,proto3" json:"request_resource,omitempty"`
	LimitResourceRecommendation          []*MetricData `protobuf:"bytes,4,rep,name=limit_resource_recommendation,json=limitResourceRecommendation,proto3" json:"limit_resource_recommendation,omitempty"`
	RequestResourceRecommendation        []*MetricData `protobuf:"bytes,5,rep,name=request_resource_recommendation,json=requestResourceRecommendation,proto3" json:"request_resource_recommendation,omitempty"`
	InitialLimitResourceRecommendation   []*MetricData `protobuf:"bytes,6,rep,name=initial_limit_resource_recommendation,json=initialLimitResourceRecommendation,proto3" json:"initial_limit_resource_recommendation,omitempty"`
	InitialRequestResourceRecommendation []*MetricData `protobuf:"bytes,7,rep,name=initial_request_resource_recommendation,json=initialRequestResourceRecommendation,proto3" json:"initial_request_resource_recommendation,omitempty"`
	XXX_NoUnkeyedLiteral                 struct{}      `json:"-"`
	XXX_unrecognized                     []byte        `json:"-"`
	XXX_sizecache                        int32         `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_42c7462b98f43aa8, []int{1}
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

func (m *Container) GetLimitResourceRecommendation() []*MetricData {
	if m != nil {
		return m.LimitResourceRecommendation
	}
	return nil
}

func (m *Container) GetRequestResourceRecommendation() []*MetricData {
	if m != nil {
		return m.RequestResourceRecommendation
	}
	return nil
}

func (m *Container) GetInitialLimitResourceRecommendation() []*MetricData {
	if m != nil {
		return m.InitialLimitResourceRecommendation
	}
	return nil
}

func (m *Container) GetInitialRequestResourceRecommendation() []*MetricData {
	if m != nil {
		return m.InitialRequestResourceRecommendation
	}
	return nil
}

type Pod struct {
	NamespacedName       *NamespacedName `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	ResourceLink         string          `protobuf:"bytes,2,opt,name=resource_link,json=resourceLink,proto3" json:"resource_link,omitempty"`
	Containers           []*Container    `protobuf:"bytes,3,rep,name=containers,proto3" json:"containers,omitempty"`
	IsAlameda            bool            `protobuf:"varint,4,opt,name=is_alameda,json=isAlameda,proto3" json:"is_alameda,omitempty"`
	AlamedaResource      *NamespacedName `protobuf:"bytes,5,opt,name=alameda_resource,json=alamedaResource,proto3" json:"alameda_resource,omitempty"`
	NodeName             string          `protobuf:"bytes,6,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Pod) Reset()         { *m = Pod{} }
func (m *Pod) String() string { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()    {}
func (*Pod) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_42c7462b98f43aa8, []int{2}
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

func (m *Pod) GetAlamedaResource() *NamespacedName {
	if m != nil {
		return m.AlamedaResource
	}
	return nil
}

func (m *Pod) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type Node struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_42c7462b98f43aa8, []int{3}
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

func init() {
	proto.RegisterType((*NamespacedName)(nil), "containers_ai.alameda.v1alpha1.datahub.NamespacedName")
	proto.RegisterType((*Container)(nil), "containers_ai.alameda.v1alpha1.datahub.Container")
	proto.RegisterType((*Pod)(nil), "containers_ai.alameda.v1alpha1.datahub.Pod")
	proto.RegisterType((*Node)(nil), "containers_ai.alameda.v1alpha1.datahub.Node")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/resource.proto", fileDescriptor_resource_42c7462b98f43aa8)
}

var fileDescriptor_resource_42c7462b98f43aa8 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4b, 0x8f, 0xd3, 0x30,
	0x10, 0x56, 0xfa, 0x62, 0x33, 0xcb, 0xb6, 0x2b, 0x9f, 0xa2, 0x5d, 0x2a, 0xaa, 0xf0, 0x2a, 0x42,
	0x4a, 0xd5, 0x22, 0x71, 0xe7, 0x71, 0x5c, 0x56, 0xe0, 0x1b, 0x07, 0x64, 0xcd, 0xc6, 0x96, 0xd6,
	0xda, 0xc4, 0x0e, 0x89, 0xbb, 0x07, 0xee, 0x1c, 0xf8, 0x99, 0xfc, 0x10, 0x24, 0x94, 0xd4, 0x71,
	0x9b, 0x08, 0x85, 0x2a, 0xdc, 0x9c, 0xc9, 0xcc, 0xf7, 0x18, 0xf9, 0x33, 0xbc, 0xc2, 0x04, 0x53,
	0xc1, 0x91, 0x61, 0x26, 0x57, 0xf7, 0x6b, 0x4c, 0xb2, 0x5b, 0x5c, 0xaf, 0x38, 0x1a, 0xbc, 0xdd,
	0xde, 0xac, 0x72, 0x51, 0xe8, 0x6d, 0x1e, 0x8b, 0x28, 0xcb, 0xb5, 0xd1, 0xe4, 0x79, 0xac, 0x95,
	0x41, 0xa9, 0x44, 0x5e, 0x30, 0x94, 0x91, 0x1d, 0x8d, 0xea, 0xb1, 0xc8, 0x8e, 0x5d, 0xbc, 0xec,
	0x04, 0x4d, 0x85, 0xc9, 0x65, 0xbc, 0x83, 0x0c, 0xdf, 0xc1, 0xf4, 0x1a, 0x53, 0x51, 0x64, 0x18,
	0x0b, 0x5e, 0x9e, 0xc8, 0x23, 0xf0, 0x55, 0x5d, 0x09, 0xbc, 0x85, 0xb7, 0xf4, 0xe9, 0xbe, 0x40,
	0x08, 0x8c, 0xca, 0x8f, 0x60, 0x50, 0xfd, 0xa8, 0xce, 0xe1, 0xaf, 0x31, 0xf8, 0xef, 0x6b, 0x65,
	0xae, 0xc3, 0xdb, 0x77, 0x90, 0x2f, 0x30, 0x4d, 0x64, 0x2a, 0x0d, 0xab, 0x0d, 0x05, 0x83, 0xc5,
	0x70, 0x79, 0xba, 0xd9, 0x44, 0xc7, 0x39, 0x8a, 0x3e, 0x56, 0x9a, 0x3f, 0xa0, 0x41, 0x7a, 0x56,
	0x21, 0x51, 0x0b, 0x44, 0xbe, 0xc2, 0x79, 0x2e, 0xbe, 0x6d, 0x45, 0x71, 0x00, 0x3e, 0xec, 0x0d,
	0x3e, 0xb3, 0x58, 0x0e, 0xfe, 0x1e, 0xe6, 0x4d, 0xe5, 0x2c, 0x17, 0xb1, 0x4e, 0x53, 0xa1, 0x38,
	0x1a, 0xa9, 0x55, 0x30, 0xea, 0xcd, 0x75, 0xd9, 0x30, 0x42, 0x1b, 0xb0, 0xe4, 0x3b, 0x3c, 0x6e,
	0xdb, 0x6a, 0x33, 0x8f, 0x7b, 0x33, 0xcf, 0x5b, 0x2e, 0x5b, 0xdc, 0x3f, 0x3c, 0x78, 0x26, 0x95,
	0x34, 0x12, 0x13, 0xd6, 0x6d, 0x7e, 0xd2, 0x5b, 0x42, 0x68, 0x09, 0xae, 0x3a, 0x76, 0xf0, 0xd3,
	0x83, 0x17, 0xb5, 0x8e, 0x7f, 0x2d, 0xe3, 0x41, 0x6f, 0x25, 0x4f, 0x2d, 0x05, 0xed, 0xda, 0x49,
	0xf8, 0x7b, 0x00, 0xc3, 0x4f, 0x9a, 0x13, 0x06, 0x33, 0x17, 0x06, 0xce, 0xdc, 0x45, 0x3f, 0xdd,
	0xbc, 0x39, 0x96, 0xba, 0x19, 0x37, 0x3a, 0x55, 0xcd, 0xf8, 0x3d, 0x81, 0x33, 0xe7, 0x31, 0x91,
	0xea, 0xce, 0x26, 0xed, 0x61, 0x5d, 0xbc, 0x92, 0xea, 0x8e, 0x7c, 0x06, 0xd8, 0xb3, 0xd9, 0xeb,
	0xbe, 0x3e, 0x56, 0x80, 0x8b, 0x2a, 0x3d, 0x00, 0x21, 0x73, 0x00, 0x59, 0x30, 0x3b, 0x14, 0x8c,
	0x16, 0xde, 0xf2, 0x84, 0xfa, 0xb2, 0x78, 0xbb, 0x2b, 0x10, 0x84, 0xf3, 0xfa, 0x51, 0x71, 0x31,
	0x1b, 0xff, 0x97, 0xf1, 0x99, 0x6d, 0x74, 0x51, 0xbb, 0x04, 0x5f, 0x69, 0x2e, 0x76, 0x4b, 0x9d,
	0x54, 0xae, 0x4f, 0xca, 0x42, 0xd9, 0x1d, 0x5e, 0xc0, 0xe8, 0x5a, 0x73, 0xf1, 0xb7, 0xd7, 0xe5,
	0x66, 0x52, 0x3d, 0x65, 0xaf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x7e, 0x54, 0xe1, 0x4c,
	0x05, 0x00, 0x00,
}
