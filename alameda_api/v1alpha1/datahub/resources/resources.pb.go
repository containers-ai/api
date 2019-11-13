// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/resources/resources.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Represents a container and its containing limit and requeset configurations
type Container struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Resources            *ResourceRequirements `protobuf:"bytes,2,opt,name=resources,proto3" json:"resources,omitempty"`
	Status               *ContainerStatus      `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_b42f420f5e5b9311, []int{0}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
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

func (m *Container) GetResources() *ResourceRequirements {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Container) GetStatus() *ContainerStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// Represents a Kubernetes pod
type Pod struct {
	ObjectMeta           *ObjectMeta          `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	ResourceLink         string               `protobuf:"bytes,3,opt,name=resource_link,json=resourceLink,proto3" json:"resource_link,omitempty"`
	AppName              string               `protobuf:"bytes,4,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	AppPartOf            string               `protobuf:"bytes,5,opt,name=app_part_of,json=appPartOf,proto3" json:"app_part_of,omitempty"`
	Containers           []*Container         `protobuf:"bytes,6,rep,name=containers,proto3" json:"containers,omitempty"`
	TopController        *Controller          `protobuf:"bytes,7,opt,name=top_controller,json=topController,proto3" json:"top_controller,omitempty"`
	Status               *PodStatus           `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	AlamedaPodSpec       *AlamedaPodSpec      `protobuf:"bytes,9,opt,name=alameda_pod_spec,json=alamedaPodSpec,proto3" json:"alameda_pod_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Pod) Reset()         { *m = Pod{} }
func (m *Pod) String() string { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()    {}
func (*Pod) Descriptor() ([]byte, []int) {
	return fileDescriptor_b42f420f5e5b9311, []int{1}
}

func (m *Pod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pod.Unmarshal(m, b)
}
func (m *Pod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pod.Marshal(b, m, deterministic)
}
func (m *Pod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pod.Merge(m, src)
}
func (m *Pod) XXX_Size() int {
	return xxx_messageInfo_Pod.Size(m)
}
func (m *Pod) XXX_DiscardUnknown() {
	xxx_messageInfo_Pod.DiscardUnknown(m)
}

var xxx_messageInfo_Pod proto.InternalMessageInfo

func (m *Pod) GetObjectMeta() *ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *Pod) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Pod) GetResourceLink() string {
	if m != nil {
		return m.ResourceLink
	}
	return ""
}

func (m *Pod) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *Pod) GetAppPartOf() string {
	if m != nil {
		return m.AppPartOf
	}
	return ""
}

func (m *Pod) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Pod) GetTopController() *Controller {
	if m != nil {
		return m.TopController
	}
	return nil
}

func (m *Pod) GetStatus() *PodStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Pod) GetAlamedaPodSpec() *AlamedaPodSpec {
	if m != nil {
		return m.AlamedaPodSpec
	}
	return nil
}

type Controller struct {
	ObjectMeta            *ObjectMeta            `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	Kind                  Kind                   `protobuf:"varint,2,opt,name=kind,proto3,enum=containersai.alameda.v1alpha1.datahub.resources.Kind" json:"kind,omitempty"`
	Replicas              int32                  `protobuf:"varint,3,opt,name=replicas,proto3" json:"replicas,omitempty"`
	SpecReplicas          int32                  `protobuf:"varint,4,opt,name=spec_replicas,json=specReplicas,proto3" json:"spec_replicas,omitempty"`
	OwnerReferences       []*OwnerReference      `protobuf:"bytes,5,rep,name=owner_references,json=ownerReferences,proto3" json:"owner_references,omitempty"`
	AlamedaControllerSpec *AlamedaControllerSpec `protobuf:"bytes,6,opt,name=alameda_controller_spec,json=alamedaControllerSpec,proto3" json:"alameda_controller_spec,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *Controller) Reset()         { *m = Controller{} }
func (m *Controller) String() string { return proto.CompactTextString(m) }
func (*Controller) ProtoMessage()    {}
func (*Controller) Descriptor() ([]byte, []int) {
	return fileDescriptor_b42f420f5e5b9311, []int{2}
}

func (m *Controller) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Controller.Unmarshal(m, b)
}
func (m *Controller) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Controller.Marshal(b, m, deterministic)
}
func (m *Controller) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Controller.Merge(m, src)
}
func (m *Controller) XXX_Size() int {
	return xxx_messageInfo_Controller.Size(m)
}
func (m *Controller) XXX_DiscardUnknown() {
	xxx_messageInfo_Controller.DiscardUnknown(m)
}

var xxx_messageInfo_Controller proto.InternalMessageInfo

func (m *Controller) GetObjectMeta() *ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *Controller) GetKind() Kind {
	if m != nil {
		return m.Kind
	}
	return Kind_POD
}

func (m *Controller) GetReplicas() int32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

func (m *Controller) GetSpecReplicas() int32 {
	if m != nil {
		return m.SpecReplicas
	}
	return 0
}

func (m *Controller) GetOwnerReferences() []*OwnerReference {
	if m != nil {
		return m.OwnerReferences
	}
	return nil
}

func (m *Controller) GetAlamedaControllerSpec() *AlamedaControllerSpec {
	if m != nil {
		return m.AlamedaControllerSpec
	}
	return nil
}

type Application struct {
	ObjectMeta             *ObjectMeta             `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	AlamedaApplicationSpec *AlamedaApplicationSpec `protobuf:"bytes,2,opt,name=alameda_application_spec,json=alamedaApplicationSpec,proto3" json:"alameda_application_spec,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                `json:"-"`
	XXX_unrecognized       []byte                  `json:"-"`
	XXX_sizecache          int32                   `json:"-"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_b42f420f5e5b9311, []int{3}
}

func (m *Application) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Application.Unmarshal(m, b)
}
func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Application.Marshal(b, m, deterministic)
}
func (m *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(m, src)
}
func (m *Application) XXX_Size() int {
	return xxx_messageInfo_Application.Size(m)
}
func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetObjectMeta() *ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *Application) GetAlamedaApplicationSpec() *AlamedaApplicationSpec {
	if m != nil {
		return m.AlamedaApplicationSpec
	}
	return nil
}

type Namespace struct {
	ObjectMeta           *ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Namespace) Reset()         { *m = Namespace{} }
func (m *Namespace) String() string { return proto.CompactTextString(m) }
func (*Namespace) ProtoMessage()    {}
func (*Namespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_b42f420f5e5b9311, []int{4}
}

func (m *Namespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Namespace.Unmarshal(m, b)
}
func (m *Namespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Namespace.Marshal(b, m, deterministic)
}
func (m *Namespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Namespace.Merge(m, src)
}
func (m *Namespace) XXX_Size() int {
	return xxx_messageInfo_Namespace.Size(m)
}
func (m *Namespace) XXX_DiscardUnknown() {
	xxx_messageInfo_Namespace.DiscardUnknown(m)
}

var xxx_messageInfo_Namespace proto.InternalMessageInfo

func (m *Namespace) GetObjectMeta() *ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

// Represents a Kubernetes node
type Node struct {
	ObjectMeta           *ObjectMeta          `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Capacity             *Capacity            `protobuf:"bytes,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	AlamedaNodeSpec      *AlamedaNodeSpec     `protobuf:"bytes,4,opt,name=alameda_node_spec,json=alamedaNodeSpec,proto3" json:"alameda_node_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_b42f420f5e5b9311, []int{5}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetObjectMeta() *ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *Node) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Node) GetCapacity() *Capacity {
	if m != nil {
		return m.Capacity
	}
	return nil
}

func (m *Node) GetAlamedaNodeSpec() *AlamedaNodeSpec {
	if m != nil {
		return m.AlamedaNodeSpec
	}
	return nil
}

type Cluster struct {
	ObjectMeta           *ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_b42f420f5e5b9311, []int{6}
}

func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetObjectMeta() *ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func init() {
	proto.RegisterType((*Container)(nil), "containersai.alameda.v1alpha1.datahub.resources.Container")
	proto.RegisterType((*Pod)(nil), "containersai.alameda.v1alpha1.datahub.resources.Pod")
	proto.RegisterType((*Controller)(nil), "containersai.alameda.v1alpha1.datahub.resources.Controller")
	proto.RegisterType((*Application)(nil), "containersai.alameda.v1alpha1.datahub.resources.Application")
	proto.RegisterType((*Namespace)(nil), "containersai.alameda.v1alpha1.datahub.resources.Namespace")
	proto.RegisterType((*Node)(nil), "containersai.alameda.v1alpha1.datahub.resources.Node")
	proto.RegisterType((*Cluster)(nil), "containersai.alameda.v1alpha1.datahub.resources.Cluster")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/resources/resources.proto", fileDescriptor_b42f420f5e5b9311)
}

var fileDescriptor_b42f420f5e5b9311 = []byte{
	// 705 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0x9a, 0x34, 0x8d, 0x27, 0xfd, 0x62, 0x25, 0xc0, 0xe4, 0x00, 0x55, 0xb8, 0xf4, 0x82,
	0xa3, 0xa6, 0x2a, 0x52, 0xe1, 0x00, 0x25, 0x2a, 0x08, 0x01, 0x6d, 0xb5, 0x80, 0x84, 0x2a, 0x24,
	0x6b, 0x62, 0x6f, 0xd2, 0x6d, 0x6d, 0xef, 0x62, 0x6f, 0x40, 0xbd, 0x20, 0xf1, 0x47, 0x39, 0x20,
	0xee, 0xfc, 0x02, 0x24, 0xb4, 0xeb, 0xaf, 0x44, 0xea, 0x21, 0x8e, 0x14, 0x89, 0xdb, 0xce, 0x78,
	0xfc, 0xc6, 0xef, 0xbd, 0xf1, 0xd8, 0xf0, 0x18, 0x03, 0x0c, 0x99, 0x8f, 0x2e, 0x4a, 0xde, 0xfb,
	0xba, 0x87, 0x81, 0xbc, 0xc0, 0xbd, 0x9e, 0x8f, 0x0a, 0x2f, 0x26, 0xc3, 0x5e, 0xcc, 0x12, 0x31,
	0x89, 0x3d, 0x96, 0x94, 0x27, 0x47, 0xc6, 0x42, 0x09, 0xd2, 0xf3, 0x44, 0xa4, 0x90, 0x47, 0x2c,
	0x4e, 0x90, 0x3b, 0x19, 0x88, 0x93, 0x03, 0x38, 0x19, 0x80, 0x53, 0xdc, 0xd6, 0x39, 0x98, 0xb3,
	0x51, 0xc8, 0x14, 0xea, 0x6c, 0xda, 0xa7, 0xb3, 0x3f, 0xe7, 0x6d, 0x89, 0x42, 0x35, 0xc9, 0x1e,
	0xae, 0xd3, 0x9f, 0xf3, 0x26, 0x75, 0x2d, 0x73, 0x42, 0x9d, 0x07, 0x63, 0x21, 0xc6, 0x01, 0xeb,
	0x99, 0x68, 0x38, 0x19, 0xf5, 0x14, 0x0f, 0x59, 0xa2, 0x30, 0x94, 0x69, 0x41, 0xf7, 0x67, 0x0d,
	0xac, 0x41, 0x4e, 0x9a, 0x10, 0x68, 0x44, 0x18, 0x32, 0xbb, 0xb6, 0x53, 0xdb, 0xb5, 0xa8, 0x39,
	0x13, 0x0f, 0xac, 0x02, 0xdb, 0x5e, 0xd9, 0xa9, 0xed, 0xb6, 0xfb, 0xc7, 0x4e, 0x45, 0x9d, 0x1c,
	0x9a, 0x9d, 0x28, 0xfb, 0x32, 0xe1, 0x31, 0x0b, 0x59, 0xa4, 0x12, 0x5a, 0xe2, 0x92, 0x4f, 0xd0,
	0x4c, 0xb9, 0xda, 0x75, 0xd3, 0xe1, 0x79, 0xe5, 0x0e, 0x05, 0x89, 0xf7, 0x06, 0x87, 0x66, 0x78,
	0xdd, 0xbf, 0x0d, 0xa8, 0x9f, 0x09, 0x9f, 0x7c, 0x86, 0xb6, 0x18, 0x5e, 0x32, 0x4f, 0xb9, 0xda,
	0x0b, 0xc3, 0xb0, 0xdd, 0x7f, 0x5a, 0xb9, 0xcd, 0xa9, 0xc1, 0x78, 0xc7, 0x14, 0x52, 0x10, 0xc5,
	0x99, 0x1c, 0x02, 0x24, 0x0a, 0x63, 0xe5, 0x6a, 0x7d, 0x33, 0x95, 0x3a, 0x4e, 0x2a, 0xbe, 0x93,
	0x8b, 0xef, 0x7c, 0xc8, 0xc5, 0xa7, 0x96, 0xa9, 0xd6, 0x31, 0x79, 0x08, 0x1b, 0x39, 0xbc, 0x1b,
	0xf0, 0xe8, 0xca, 0x28, 0x60, 0xd1, 0xf5, 0x3c, 0xf9, 0x96, 0x47, 0x57, 0xe4, 0x1e, 0xb4, 0x50,
	0x4a, 0xd7, 0x98, 0xd3, 0x30, 0xd7, 0xd7, 0x50, 0xca, 0x13, 0xed, 0xcf, 0x7d, 0x68, 0xeb, 0x4b,
	0x52, 0x77, 0x17, 0x23, 0x7b, 0xd5, 0x5c, 0xb5, 0x50, 0xca, 0x33, 0x8c, 0xd5, 0xe9, 0x88, 0x9c,
	0x03, 0x94, 0x24, 0xed, 0xe6, 0x4e, 0x7d, 0xb7, 0xdd, 0x7f, 0xb2, 0xb8, 0xbc, 0x74, 0x0a, 0x8d,
	0x0c, 0x61, 0x53, 0x09, 0xe9, 0xea, 0x4c, 0x2c, 0x82, 0x80, 0xc5, 0xf6, 0xda, 0x82, 0xba, 0x0e,
	0x0a, 0x08, 0xba, 0xa1, 0x84, 0x2c, 0x43, 0x42, 0x8b, 0xd1, 0x68, 0x19, 0xec, 0xea, 0xcf, 0x7e,
	0x26, 0xfc, 0xd9, 0xa1, 0x20, 0x1c, 0xb6, 0xf3, 0x97, 0x49, 0x0a, 0xdf, 0x4d, 0x24, 0xf3, 0x6c,
	0xcb, 0xa0, 0x3f, 0xab, 0x8c, 0x7e, 0x94, 0x96, 0xe8, 0x26, 0x92, 0x79, 0x74, 0x13, 0x67, 0xe2,
	0xee, 0xaf, 0x3a, 0xc0, 0x14, 0x9b, 0xe5, 0x8e, 0xe1, 0x6b, 0x68, 0x5c, 0xf1, 0xc8, 0x37, 0x03,
	0xb8, 0xd9, 0x3f, 0xa8, 0x0c, 0xfb, 0x86, 0x47, 0x3e, 0x35, 0x10, 0xa4, 0x03, 0xad, 0x98, 0xc9,
	0x80, 0x7b, 0x98, 0xbe, 0x93, 0xab, 0xb4, 0x88, 0xf5, 0xc8, 0x6a, 0xc9, 0xdc, 0xa2, 0xa0, 0x61,
	0x0a, 0xd6, 0x75, 0x92, 0xe6, 0x45, 0x97, 0xb0, 0x2d, 0xbe, 0x45, 0x2c, 0x76, 0x63, 0x36, 0x62,
	0x31, 0x8b, 0xf4, 0xfa, 0x58, 0x35, 0xd3, 0x57, 0x5d, 0xe3, 0x53, 0x0d, 0x44, 0x73, 0x1c, 0xba,
	0x25, 0x66, 0xe2, 0x84, 0x7c, 0x87, 0xbb, 0xb9, 0x9f, 0xe5, 0x2c, 0xa6, 0xb6, 0x36, 0x8d, 0xc2,
	0x2f, 0x17, 0xb5, 0xb5, 0xb4, 0xce, 0xb8, 0x7b, 0x1b, 0x6f, 0x4a, 0x77, 0xff, 0xd4, 0xa0, 0x7d,
	0x24, 0x0d, 0x71, 0xc5, 0x45, 0xb4, 0x64, 0x97, 0x7f, 0xd4, 0xc0, 0x2e, 0xbf, 0x05, 0x45, 0xd7,
	0x94, 0x6f, 0xba, 0x7b, 0x5e, 0x2d, 0xca, 0x77, 0x8a, 0x85, 0x21, 0x7c, 0x07, 0x6f, 0xcc, 0x77,
	0x39, 0x58, 0x7a, 0xfb, 0x24, 0x12, 0x3d, 0xb6, 0x5c, 0xba, 0xdd, 0xdf, 0x2b, 0xd0, 0x38, 0x11,
	0x3e, 0xfb, 0x7f, 0x57, 0xf8, 0x47, 0x68, 0x79, 0x28, 0xd1, 0xe3, 0xea, 0x3a, 0xfb, 0x7e, 0x1d,
	0x56, 0x5f, 0x80, 0x19, 0x00, 0x2d, 0xa0, 0x48, 0x00, 0xb7, 0x72, 0x9b, 0x23, 0xe1, 0xb3, 0xd4,
	0xdf, 0xc6, 0x82, 0xdf, 0xc7, 0xcc, 0x5f, 0x2d, 0xa4, 0x31, 0x76, 0x0b, 0x67, 0x13, 0xdd, 0x31,
	0xac, 0x0d, 0x82, 0x49, 0xa2, 0x96, 0xbd, 0xa4, 0x5e, 0x1c, 0x9f, 0x0f, 0xc6, 0x5c, 0xe9, 0x52,
	0x4f, 0x84, 0x53, 0x7f, 0x5c, 0x8f, 0x90, 0xf7, 0xf4, 0xaf, 0xcd, 0x7c, 0xbf, 0x39, 0xc3, 0xa6,
	0xf1, 0x64, 0xff, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x57, 0x15, 0xec, 0x9a, 0xec, 0x09, 0x00,
	0x00,
}
