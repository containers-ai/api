// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/resource.proto

package containers_ai_alameda_v1alpha1_datahub

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

//*
// Represents kubernetes resource kind
//
type Kind int32

const (
	Kind_POD              Kind = 0
	Kind_DEPLOYMENT       Kind = 1
	Kind_DEPLOYMENTCONFIG Kind = 2
	Kind_ALAMEDASCALER    Kind = 3
)

var Kind_name = map[int32]string{
	0: "POD",
	1: "DEPLOYMENT",
	2: "DEPLOYMENTCONFIG",
	3: "ALAMEDASCALER",
}

var Kind_value = map[string]int32{
	"POD":              0,
	"DEPLOYMENT":       1,
	"DEPLOYMENTCONFIG": 2,
	"ALAMEDASCALER":    3,
}

func (x Kind) String() string {
	return proto.EnumName(Kind_name, int32(x))
}

func (Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e45e09c80210b55a, []int{0}
}

//*
// Represents a container and its containing limit and requeset configurations
//
type Container struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LimitResource        []*MetricData    `protobuf:"bytes,2,rep,name=limit_resource,json=limitResource,proto3" json:"limit_resource,omitempty"`
	RequestResource      []*MetricData    `protobuf:"bytes,3,rep,name=request_resource,json=requestResource,proto3" json:"request_resource,omitempty"`
	Status               *ContainerStatus `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_e45e09c80210b55a, []int{0}
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

func (m *Container) GetStatus() *ContainerStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

//*
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
	UsedRecommendationId string               `protobuf:"bytes,10,opt,name=used_recommendation_id,json=usedRecommendationId,proto3" json:"used_recommendation_id,omitempty"`
	Status               *PodStatus           `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	Enable_VPA           bool                 `protobuf:"varint,12,opt,name=enable_VPA,json=enableVPA,proto3" json:"enable_VPA,omitempty"`
	Enable_HPA           bool                 `protobuf:"varint,13,opt,name=enable_HPA,json=enableHPA,proto3" json:"enable_HPA,omitempty"`
	AppName              string               `protobuf:"bytes,14,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	AppPartOf            string               `protobuf:"bytes,15,opt,name=app_part_of,json=appPartOf,proto3" json:"app_part_of,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Pod) Reset()         { *m = Pod{} }
func (m *Pod) String() string { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()    {}
func (*Pod) Descriptor() ([]byte, []int) {
	return fileDescriptor_e45e09c80210b55a, []int{1}
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

func (m *Pod) GetUsedRecommendationId() string {
	if m != nil {
		return m.UsedRecommendationId
	}
	return ""
}

func (m *Pod) GetStatus() *PodStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Pod) GetEnable_VPA() bool {
	if m != nil {
		return m.Enable_VPA
	}
	return false
}

func (m *Pod) GetEnable_HPA() bool {
	if m != nil {
		return m.Enable_HPA
	}
	return false
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

//*
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
	return fileDescriptor_e45e09c80210b55a, []int{2}
}

func (m *Capacity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Capacity.Unmarshal(m, b)
}
func (m *Capacity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Capacity.Marshal(b, m, deterministic)
}
func (m *Capacity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Capacity.Merge(m, src)
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

//*
// Represents a Kubernetes node
//
type Node struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Capacity             *Capacity            `protobuf:"bytes,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Provider             *Provider            `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_e45e09c80210b55a, []int{3}
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

func (m *Node) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Node) GetProvider() *Provider {
	if m != nil {
		return m.Provider
	}
	return nil
}

//*
// Represents top controller of the pod
//
type TopController struct {
	NamespacedName       *NamespacedName `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	Kind                 Kind            `protobuf:"varint,2,opt,name=kind,proto3,enum=containers_ai.alameda.v1alpha1.datahub.Kind" json:"kind,omitempty"`
	Replicas             int32           `protobuf:"varint,3,opt,name=Replicas,proto3" json:"Replicas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TopController) Reset()         { *m = TopController{} }
func (m *TopController) String() string { return proto.CompactTextString(m) }
func (*TopController) ProtoMessage()    {}
func (*TopController) Descriptor() ([]byte, []int) {
	return fileDescriptor_e45e09c80210b55a, []int{4}
}

func (m *TopController) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopController.Unmarshal(m, b)
}
func (m *TopController) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopController.Marshal(b, m, deterministic)
}
func (m *TopController) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopController.Merge(m, src)
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

func (m *TopController) GetReplicas() int32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

type ResourceInfo struct {
	NamespacedName       *NamespacedName `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	Kind                 Kind            `protobuf:"varint,2,opt,name=kind,proto3,enum=containers_ai.alameda.v1alpha1.datahub.Kind" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ResourceInfo) Reset()         { *m = ResourceInfo{} }
func (m *ResourceInfo) String() string { return proto.CompactTextString(m) }
func (*ResourceInfo) ProtoMessage()    {}
func (*ResourceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e45e09c80210b55a, []int{5}
}

func (m *ResourceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceInfo.Unmarshal(m, b)
}
func (m *ResourceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceInfo.Marshal(b, m, deterministic)
}
func (m *ResourceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceInfo.Merge(m, src)
}
func (m *ResourceInfo) XXX_Size() int {
	return xxx_messageInfo_ResourceInfo.Size(m)
}
func (m *ResourceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceInfo proto.InternalMessageInfo

func (m *ResourceInfo) GetNamespacedName() *NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *ResourceInfo) GetKind() Kind {
	if m != nil {
		return m.Kind
	}
	return Kind_POD
}

type Controller struct {
	ControllerInfo                *ResourceInfo        `protobuf:"bytes,1,opt,name=controller_info,json=controllerInfo,proto3" json:"controller_info,omitempty"`
	OwnerInfo                     []*ResourceInfo      `protobuf:"bytes,2,rep,name=owner_info,json=ownerInfo,proto3" json:"owner_info,omitempty"`
	Replicas                      int32                `protobuf:"varint,3,opt,name=replicas,proto3" json:"replicas,omitempty"`
	EnableRecommendationExecution bool                 `protobuf:"varint,4,opt,name=enable_recommendation_execution,json=enableRecommendationExecution,proto3" json:"enable_recommendation_execution,omitempty"`
	Policy                        RecommendationPolicy `protobuf:"varint,5,opt,name=policy,proto3,enum=containers_ai.alameda.v1alpha1.datahub.RecommendationPolicy" json:"policy,omitempty"`
	SpecReplicas                  int32                `protobuf:"varint,6,opt,name=spec_replicas,json=specReplicas,proto3" json:"spec_replicas,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}             `json:"-"`
	XXX_unrecognized              []byte               `json:"-"`
	XXX_sizecache                 int32                `json:"-"`
}

func (m *Controller) Reset()         { *m = Controller{} }
func (m *Controller) String() string { return proto.CompactTextString(m) }
func (*Controller) ProtoMessage()    {}
func (*Controller) Descriptor() ([]byte, []int) {
	return fileDescriptor_e45e09c80210b55a, []int{6}
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

func (m *Controller) GetControllerInfo() *ResourceInfo {
	if m != nil {
		return m.ControllerInfo
	}
	return nil
}

func (m *Controller) GetOwnerInfo() []*ResourceInfo {
	if m != nil {
		return m.OwnerInfo
	}
	return nil
}

func (m *Controller) GetReplicas() int32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

func (m *Controller) GetEnableRecommendationExecution() bool {
	if m != nil {
		return m.EnableRecommendationExecution
	}
	return false
}

func (m *Controller) GetPolicy() RecommendationPolicy {
	if m != nil {
		return m.Policy
	}
	return RecommendationPolicy_RECOMMENDATIONPOLICY_UNDEFINED
}

func (m *Controller) GetSpecReplicas() int32 {
	if m != nil {
		return m.SpecReplicas
	}
	return 0
}

type Provider struct {
	Provider             string   `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	InstanceType         string   `protobuf:"bytes,2,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`
	Region               string   `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Zone                 string   `protobuf:"bytes,4,opt,name=zone,proto3" json:"zone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Provider) Reset()         { *m = Provider{} }
func (m *Provider) String() string { return proto.CompactTextString(m) }
func (*Provider) ProtoMessage()    {}
func (*Provider) Descriptor() ([]byte, []int) {
	return fileDescriptor_e45e09c80210b55a, []int{7}
}

func (m *Provider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Provider.Unmarshal(m, b)
}
func (m *Provider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Provider.Marshal(b, m, deterministic)
}
func (m *Provider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Provider.Merge(m, src)
}
func (m *Provider) XXX_Size() int {
	return xxx_messageInfo_Provider.Size(m)
}
func (m *Provider) XXX_DiscardUnknown() {
	xxx_messageInfo_Provider.DiscardUnknown(m)
}

var xxx_messageInfo_Provider proto.InternalMessageInfo

func (m *Provider) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *Provider) GetInstanceType() string {
	if m != nil {
		return m.InstanceType
	}
	return ""
}

func (m *Provider) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *Provider) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterEnum("containers_ai.alameda.v1alpha1.datahub.Kind", Kind_name, Kind_value)
	proto.RegisterType((*Container)(nil), "containers_ai.alameda.v1alpha1.datahub.Container")
	proto.RegisterType((*Pod)(nil), "containers_ai.alameda.v1alpha1.datahub.Pod")
	proto.RegisterType((*Capacity)(nil), "containers_ai.alameda.v1alpha1.datahub.Capacity")
	proto.RegisterType((*Node)(nil), "containers_ai.alameda.v1alpha1.datahub.Node")
	proto.RegisterType((*TopController)(nil), "containers_ai.alameda.v1alpha1.datahub.TopController")
	proto.RegisterType((*ResourceInfo)(nil), "containers_ai.alameda.v1alpha1.datahub.ResourceInfo")
	proto.RegisterType((*Controller)(nil), "containers_ai.alameda.v1alpha1.datahub.Controller")
	proto.RegisterType((*Provider)(nil), "containers_ai.alameda.v1alpha1.datahub.Provider")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/resource.proto", fileDescriptor_e45e09c80210b55a)
}

var fileDescriptor_e45e09c80210b55a = []byte{
	// 1006 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xef, 0x6e, 0x1b, 0x45,
	0x10, 0xc7, 0xb9, 0xd4, 0x39, 0x8f, 0x63, 0xc7, 0xac, 0xaa, 0xea, 0x48, 0x55, 0x1a, 0x5c, 0x09,
	0x85, 0x82, 0x1c, 0x62, 0x0a, 0x08, 0x09, 0x24, 0x8c, 0x93, 0x92, 0x08, 0x27, 0x39, 0x36, 0x51,
	0xa5, 0x4a, 0x44, 0xa7, 0xf5, 0xdd, 0x26, 0x5d, 0xc5, 0x77, 0xbb, 0xec, 0xae, 0x1b, 0xcc, 0x1b,
	0xf0, 0x14, 0x7c, 0xe7, 0x1b, 0x2f, 0xc2, 0x43, 0xf0, 0x06, 0xbc, 0x01, 0xda, 0xbd, 0xbd, 0xb3,
	0x1d, 0xa1, 0x70, 0x2d, 0xf0, 0x81, 0x6f, 0x3b, 0xff, 0x7e, 0x37, 0xf3, 0xdb, 0x99, 0xd9, 0x83,
	0xf7, 0xc9, 0x84, 0xa4, 0x34, 0x21, 0x11, 0x11, 0x6c, 0xe7, 0xe5, 0x2e, 0x99, 0x88, 0x17, 0x64,
	0x77, 0x27, 0x21, 0x9a, 0xbc, 0x98, 0x8e, 0x77, 0x24, 0x55, 0x7c, 0x2a, 0x63, 0xda, 0x13, 0x92,
	0x6b, 0x8e, 0xde, 0x8d, 0x79, 0xa6, 0x09, 0xcb, 0xa8, 0x54, 0x11, 0x61, 0x3d, 0x17, 0xda, 0x2b,
	0xc2, 0x7a, 0x2e, 0x6c, 0xf3, 0xf1, 0xad, 0xa0, 0x42, 0xd2, 0x84, 0xc5, 0x3a, 0xc7, 0xdc, 0x7c,
	0x78, 0xc9, 0xf9, 0xe5, 0x84, 0xee, 0x58, 0x69, 0x3c, 0xbd, 0xd8, 0xd1, 0x2c, 0xa5, 0x4a, 0x93,
	0x54, 0x38, 0x87, 0xdb, 0x33, 0x4c, 0xa9, 0x26, 0xe6, 0xec, 0x9c, 0xdf, 0xfb, 0x3b, 0x67, 0xc9,
	0x62, 0xe7, 0xba, 0x7d, 0xab, 0xab, 0x9e, 0x09, 0xaa, 0x72, 0xcf, 0xee, 0x2f, 0x2b, 0xd0, 0x18,
	0x16, 0x95, 0x23, 0x04, 0xab, 0x19, 0x49, 0x69, 0x50, 0xdb, 0xaa, 0x6d, 0x37, 0xb0, 0x3d, 0xa3,
	0xe7, 0xd0, 0x9e, 0xb0, 0x94, 0xe9, 0xa8, 0x20, 0x2c, 0x58, 0xd9, 0xf2, 0xb6, 0x9b, 0xfd, 0x7e,
	0xaf, 0x1a, 0x63, 0xbd, 0x23, 0x9b, 0xd9, 0x1e, 0xd1, 0x04, 0xb7, 0x2c, 0x12, 0x76, 0x40, 0xe8,
	0x1c, 0x3a, 0x92, 0x7e, 0x3f, 0xa5, 0x6a, 0x01, 0xdc, 0x7b, 0x6d, 0xf0, 0x0d, 0x87, 0x55, 0xc2,
	0x9f, 0x40, 0x5d, 0x69, 0xa2, 0xa7, 0x2a, 0x58, 0xdd, 0xaa, 0x6d, 0x37, 0xfb, 0x9f, 0x56, 0x05,
	0x2d, 0x09, 0x39, 0xb5, 0xe1, 0xd8, 0xc1, 0x74, 0x7f, 0xaf, 0x83, 0x17, 0xf2, 0x04, 0x45, 0xb0,
	0x61, 0xa8, 0x51, 0x82, 0xc4, 0x34, 0x89, 0x4a, 0xc6, 0x9a, 0xfd, 0x4f, 0xaa, 0x7e, 0xe1, 0xb8,
	0x0c, 0x37, 0x27, 0xdc, 0xce, 0x96, 0x64, 0xf4, 0x08, 0x5a, 0x05, 0x21, 0xd1, 0x84, 0x65, 0x57,
	0xc1, 0x8a, 0xbd, 0x90, 0xf5, 0x42, 0x39, 0x62, 0xd9, 0x15, 0xfa, 0x16, 0x60, 0xfe, 0x35, 0xc7,
	0xdb, 0xee, 0x2b, 0x97, 0x88, 0x17, 0x40, 0xd0, 0x03, 0x00, 0xa6, 0x22, 0x17, 0x64, 0x59, 0xf3,
	0x71, 0x83, 0xa9, 0x41, 0xae, 0x40, 0xe7, 0xd0, 0x2e, 0x1a, 0x4b, 0xc5, 0x64, 0x42, 0x65, 0x70,
	0xe7, 0x1f, 0x95, 0xdd, 0x72, 0x8e, 0xa7, 0x16, 0x0c, 0xdd, 0x87, 0x46, 0xc6, 0x13, 0x9a, 0x13,
	0x5a, 0xb7, 0x15, 0xfb, 0x46, 0x61, 0x29, 0xf9, 0x0c, 0x40, 0x69, 0x22, 0x75, 0x64, 0x66, 0x28,
	0x58, 0xb3, 0xdf, 0xdd, 0xec, 0xe5, 0x03, 0xd6, 0x2b, 0x06, 0xac, 0x77, 0x56, 0x0c, 0x18, 0x6e,
	0x58, 0x6f, 0x23, 0xa3, 0x33, 0xa8, 0x0b, 0x3e, 0x61, 0xf1, 0x2c, 0xf0, 0xb7, 0x6a, 0xdb, 0xed,
	0xfe, 0xe7, 0x55, 0xd3, 0xc5, 0x34, 0xe6, 0x69, 0x4a, 0xb3, 0x84, 0x68, 0xc6, 0xb3, 0xd0, 0x62,
	0x60, 0x87, 0x85, 0xbe, 0x83, 0xb6, 0xe6, 0x22, 0x32, 0x50, 0x92, 0x4f, 0x0c, 0x19, 0x0d, 0x9b,
	0xd4, 0xc7, 0x55, 0xd1, 0xcf, 0xb8, 0x18, 0x96, 0xc1, 0xb8, 0xa5, 0x17, 0x45, 0xf4, 0x04, 0xee,
	0x4d, 0x15, 0x4d, 0x22, 0xb9, 0x94, 0x42, 0xc4, 0x92, 0x00, 0x2c, 0x31, 0x77, 0x8d, 0x75, 0x39,
	0xbf, 0xc3, 0x04, 0x1d, 0x96, 0x1d, 0xdf, 0xb4, 0xb9, 0x54, 0x6e, 0x87, 0x90, 0x27, 0xcb, 0xbd,
	0x6e, 0x5a, 0x81, 0x66, 0x64, 0x3c, 0xa1, 0xd1, 0xb3, 0x70, 0x10, 0xac, 0xe7, 0xad, 0x90, 0x6b,
	0x9e, 0x85, 0x83, 0x05, 0xf3, 0x41, 0x38, 0x08, 0x5a, 0x8b, 0xe6, 0x83, 0x70, 0x80, 0xde, 0x02,
	0x9f, 0x08, 0x91, 0xdf, 0x64, 0xdb, 0x26, 0xbc, 0x46, 0x84, 0xb0, 0x17, 0xf9, 0x36, 0x34, 0x8d,
	0x49, 0x98, 0xbb, 0xe4, 0x17, 0xc1, 0x86, 0xb5, 0x36, 0x88, 0x10, 0x21, 0x91, 0xfa, 0xe4, 0xa2,
	0xfb, 0x53, 0x0d, 0xfc, 0x21, 0x11, 0x24, 0x66, 0x7a, 0x66, 0x5a, 0x22, 0x16, 0xd3, 0x28, 0xe6,
	0x92, 0x2a, 0x3b, 0x63, 0x1e, 0xf6, 0x63, 0x31, 0x1d, 0x1a, 0x19, 0xbd, 0x03, 0xeb, 0x29, 0x4d,
	0xb9, 0x9c, 0x45, 0xe3, 0x99, 0xa6, 0xca, 0x0e, 0x89, 0x87, 0x9b, 0xb9, 0xee, 0x2b, 0xa3, 0x42,
	0x5f, 0xc0, 0xfd, 0x8c, 0xea, 0x6b, 0xae, 0xaf, 0xa2, 0x94, 0x5e, 0x92, 0x31, 0xd3, 0x2a, 0x12,
	0x54, 0x46, 0x8a, 0xc6, 0x3c, 0x4b, 0x02, 0xcf, 0x46, 0x04, 0xce, 0xe5, 0xc8, 0x79, 0x84, 0x54,
	0x9e, 0x5a, 0x7b, 0xf7, 0x8f, 0x1a, 0xac, 0x1e, 0xf3, 0x84, 0xfe, 0xe5, 0x62, 0x1c, 0x81, 0x1f,
	0xbb, 0x3c, 0xed, 0xa7, 0x9b, 0xfd, 0x0f, 0x2b, 0x4f, 0x9f, 0x8b, 0xc3, 0x25, 0xc2, 0x8d, 0xfe,
	0xf6, 0x5e, 0xa5, 0xbf, 0x47, 0xe0, 0x0b, 0xc9, 0x5f, 0xb2, 0x84, 0x4a, 0xb7, 0xe9, 0x2a, 0x27,
	0x12, 0xba, 0x38, 0x5c, 0x22, 0x74, 0x7f, 0xab, 0x41, 0x6b, 0xa9, 0x35, 0xff, 0xfb, 0x75, 0xf7,
	0x25, 0xac, 0x5e, 0xb1, 0x2c, 0xb1, 0x2c, 0xb6, 0xfb, 0x1f, 0x54, 0x45, 0xfd, 0x86, 0x65, 0x09,
	0xb6, 0x91, 0x68, 0x13, 0x7c, 0x4c, 0xc5, 0x84, 0xc5, 0x44, 0x59, 0xee, 0xee, 0xe0, 0x52, 0xee,
	0xfe, 0x5a, 0x83, 0xf5, 0xe2, 0x4d, 0x38, 0xcc, 0x2e, 0xf8, 0xff, 0xa0, 0x9e, 0xee, 0xcf, 0x1e,
	0xc0, 0xc2, 0x0d, 0x9c, 0xc3, 0xc6, 0x7c, 0xcf, 0x44, 0x2c, 0xbb, 0xe0, 0x2e, 0xe3, 0x27, 0xd5,
	0x57, 0xd9, 0x9c, 0x00, 0xdc, 0x9e, 0x83, 0x59, 0x42, 0x4e, 0x01, 0xf8, 0x75, 0x56, 0x20, 0xe7,
	0xcf, 0xfb, 0xeb, 0x21, 0x37, 0x2c, 0x8e, 0x05, 0xdd, 0x04, 0x5f, 0xde, 0xb8, 0x92, 0x42, 0x46,
	0x4f, 0xe1, 0xa1, 0xdb, 0x1e, 0x37, 0xf6, 0x1b, 0xfd, 0x81, 0xc6, 0x53, 0x73, 0x72, 0x8f, 0xcf,
	0x83, 0xdc, 0x6d, 0x79, 0xd1, 0xed, 0x17, 0x4e, 0x0b, 0x9b, 0xfd, 0xce, 0xbf, 0xb8, 0xd9, 0x1f,
	0x41, 0x4b, 0x09, 0x1a, 0x47, 0x65, 0xfa, 0x75, 0x9b, 0xfe, 0xba, 0x51, 0x96, 0x5d, 0x75, 0x0d,
	0x7e, 0x31, 0x3c, 0xa6, 0xd4, 0x72, 0x00, 0xf3, 0x0d, 0x51, 0xca, 0x06, 0x8c, 0x65, 0x4a, 0x93,
	0x2c, 0xa6, 0x91, 0xf9, 0xf1, 0x2a, 0x9e, 0xf2, 0x42, 0x79, 0x36, 0x13, 0x14, 0xdd, 0x83, 0xba,
	0xa4, 0x97, 0xa6, 0x6c, 0xcf, 0x5a, 0x9d, 0x64, 0xd6, 0xce, 0x8f, 0x3c, 0xa3, 0x96, 0x8c, 0x06,
	0xb6, 0xe7, 0xc7, 0x07, 0xb0, 0x6a, 0x1a, 0x05, 0xad, 0x81, 0x17, 0x9e, 0xec, 0x75, 0xde, 0x40,
	0x6d, 0x80, 0xbd, 0xfd, 0x70, 0x74, 0xf2, 0xfc, 0x68, 0xff, 0xf8, 0xac, 0x53, 0x43, 0x77, 0xa1,
	0x33, 0x97, 0x87, 0x27, 0xc7, 0x4f, 0x0f, 0xbf, 0xee, 0xac, 0xa0, 0x37, 0xa1, 0x35, 0x18, 0x0d,
	0x8e, 0xf6, 0xf7, 0x06, 0xa7, 0xc3, 0xc1, 0x68, 0x1f, 0x77, 0xbc, 0x71, 0xdd, 0xae, 0x95, 0x8f,
	0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x73, 0x4a, 0xee, 0xf4, 0x28, 0x0b, 0x00, 0x00,
}
