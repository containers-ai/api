// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/predictions/predictions.proto

package predictions

import (
	fmt "fmt"
	common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	schemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
	proto "github.com/golang/protobuf/proto"
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

// Represents a list of predicted metric data of a container
type ContainerPrediction struct {
	Name                    string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PredictedRawData        []*MetricData `protobuf:"bytes,2,rep,name=predicted_raw_data,json=predictedRawData,proto3" json:"predicted_raw_data,omitempty"`
	PredictedUpperboundData []*MetricData `protobuf:"bytes,3,rep,name=predicted_upperbound_data,json=predictedUpperboundData,proto3" json:"predicted_upperbound_data,omitempty"`
	PredictedLowerboundData []*MetricData `protobuf:"bytes,4,rep,name=predicted_lowerbound_data,json=predictedLowerboundData,proto3" json:"predicted_lowerbound_data,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}      `json:"-"`
	XXX_unrecognized        []byte        `json:"-"`
	XXX_sizecache           int32         `json:"-"`
}

func (m *ContainerPrediction) Reset()         { *m = ContainerPrediction{} }
func (m *ContainerPrediction) String() string { return proto.CompactTextString(m) }
func (*ContainerPrediction) ProtoMessage()    {}
func (*ContainerPrediction) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6f58e3dfb6d18af, []int{0}
}

func (m *ContainerPrediction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerPrediction.Unmarshal(m, b)
}
func (m *ContainerPrediction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerPrediction.Marshal(b, m, deterministic)
}
func (m *ContainerPrediction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerPrediction.Merge(m, src)
}
func (m *ContainerPrediction) XXX_Size() int {
	return xxx_messageInfo_ContainerPrediction.Size(m)
}
func (m *ContainerPrediction) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerPrediction.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerPrediction proto.InternalMessageInfo

func (m *ContainerPrediction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerPrediction) GetPredictedRawData() []*MetricData {
	if m != nil {
		return m.PredictedRawData
	}
	return nil
}

func (m *ContainerPrediction) GetPredictedUpperboundData() []*MetricData {
	if m != nil {
		return m.PredictedUpperboundData
	}
	return nil
}

func (m *ContainerPrediction) GetPredictedLowerboundData() []*MetricData {
	if m != nil {
		return m.PredictedLowerboundData
	}
	return nil
}

// Represents a list of predicted metrics data of a pod
type PodPrediction struct {
	ObjectMeta           *resources.ObjectMeta  `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	ContainerPredictions []*ContainerPrediction `protobuf:"bytes,2,rep,name=container_predictions,json=containerPredictions,proto3" json:"container_predictions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *PodPrediction) Reset()         { *m = PodPrediction{} }
func (m *PodPrediction) String() string { return proto.CompactTextString(m) }
func (*PodPrediction) ProtoMessage()    {}
func (*PodPrediction) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6f58e3dfb6d18af, []int{1}
}

func (m *PodPrediction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodPrediction.Unmarshal(m, b)
}
func (m *PodPrediction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodPrediction.Marshal(b, m, deterministic)
}
func (m *PodPrediction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodPrediction.Merge(m, src)
}
func (m *PodPrediction) XXX_Size() int {
	return xxx_messageInfo_PodPrediction.Size(m)
}
func (m *PodPrediction) XXX_DiscardUnknown() {
	xxx_messageInfo_PodPrediction.DiscardUnknown(m)
}

var xxx_messageInfo_PodPrediction proto.InternalMessageInfo

func (m *PodPrediction) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *PodPrediction) GetContainerPredictions() []*ContainerPrediction {
	if m != nil {
		return m.ContainerPredictions
	}
	return nil
}

type ControllerPrediction struct {
	ObjectMeta              *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	Kind                    resources.Kind        `protobuf:"varint,2,opt,name=kind,proto3,enum=containersai.alameda.v1alpha1.datahub.resources.Kind" json:"kind,omitempty"`
	PredictedRawData        []*MetricData         `protobuf:"bytes,3,rep,name=predicted_raw_data,json=predictedRawData,proto3" json:"predicted_raw_data,omitempty"`
	PredictedUpperboundData []*MetricData         `protobuf:"bytes,4,rep,name=predicted_upperbound_data,json=predictedUpperboundData,proto3" json:"predicted_upperbound_data,omitempty"`
	PredictedLowerboundData []*MetricData         `protobuf:"bytes,5,rep,name=predicted_lowerbound_data,json=predictedLowerboundData,proto3" json:"predicted_lowerbound_data,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *ControllerPrediction) Reset()         { *m = ControllerPrediction{} }
func (m *ControllerPrediction) String() string { return proto.CompactTextString(m) }
func (*ControllerPrediction) ProtoMessage()    {}
func (*ControllerPrediction) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6f58e3dfb6d18af, []int{2}
}

func (m *ControllerPrediction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControllerPrediction.Unmarshal(m, b)
}
func (m *ControllerPrediction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControllerPrediction.Marshal(b, m, deterministic)
}
func (m *ControllerPrediction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControllerPrediction.Merge(m, src)
}
func (m *ControllerPrediction) XXX_Size() int {
	return xxx_messageInfo_ControllerPrediction.Size(m)
}
func (m *ControllerPrediction) XXX_DiscardUnknown() {
	xxx_messageInfo_ControllerPrediction.DiscardUnknown(m)
}

var xxx_messageInfo_ControllerPrediction proto.InternalMessageInfo

func (m *ControllerPrediction) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *ControllerPrediction) GetKind() resources.Kind {
	if m != nil {
		return m.Kind
	}
	return resources.Kind_KIND_UNDEFINED
}

func (m *ControllerPrediction) GetPredictedRawData() []*MetricData {
	if m != nil {
		return m.PredictedRawData
	}
	return nil
}

func (m *ControllerPrediction) GetPredictedUpperboundData() []*MetricData {
	if m != nil {
		return m.PredictedUpperboundData
	}
	return nil
}

func (m *ControllerPrediction) GetPredictedLowerboundData() []*MetricData {
	if m != nil {
		return m.PredictedLowerboundData
	}
	return nil
}

type ApplicationPrediction struct {
	ObjectMeta              *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	PredictedRawData        []*MetricData         `protobuf:"bytes,2,rep,name=predicted_raw_data,json=predictedRawData,proto3" json:"predicted_raw_data,omitempty"`
	PredictedUpperboundData []*MetricData         `protobuf:"bytes,3,rep,name=predicted_upperbound_data,json=predictedUpperboundData,proto3" json:"predicted_upperbound_data,omitempty"`
	PredictedLowerboundData []*MetricData         `protobuf:"bytes,4,rep,name=predicted_lowerbound_data,json=predictedLowerboundData,proto3" json:"predicted_lowerbound_data,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *ApplicationPrediction) Reset()         { *m = ApplicationPrediction{} }
func (m *ApplicationPrediction) String() string { return proto.CompactTextString(m) }
func (*ApplicationPrediction) ProtoMessage()    {}
func (*ApplicationPrediction) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6f58e3dfb6d18af, []int{3}
}

func (m *ApplicationPrediction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationPrediction.Unmarshal(m, b)
}
func (m *ApplicationPrediction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationPrediction.Marshal(b, m, deterministic)
}
func (m *ApplicationPrediction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationPrediction.Merge(m, src)
}
func (m *ApplicationPrediction) XXX_Size() int {
	return xxx_messageInfo_ApplicationPrediction.Size(m)
}
func (m *ApplicationPrediction) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationPrediction.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationPrediction proto.InternalMessageInfo

func (m *ApplicationPrediction) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *ApplicationPrediction) GetPredictedRawData() []*MetricData {
	if m != nil {
		return m.PredictedRawData
	}
	return nil
}

func (m *ApplicationPrediction) GetPredictedUpperboundData() []*MetricData {
	if m != nil {
		return m.PredictedUpperboundData
	}
	return nil
}

func (m *ApplicationPrediction) GetPredictedLowerboundData() []*MetricData {
	if m != nil {
		return m.PredictedLowerboundData
	}
	return nil
}

type NamespacePrediction struct {
	ObjectMeta              *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	PredictedRawData        []*MetricData         `protobuf:"bytes,2,rep,name=predicted_raw_data,json=predictedRawData,proto3" json:"predicted_raw_data,omitempty"`
	PredictedUpperboundData []*MetricData         `protobuf:"bytes,3,rep,name=predicted_upperbound_data,json=predictedUpperboundData,proto3" json:"predicted_upperbound_data,omitempty"`
	PredictedLowerboundData []*MetricData         `protobuf:"bytes,4,rep,name=predicted_lowerbound_data,json=predictedLowerboundData,proto3" json:"predicted_lowerbound_data,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *NamespacePrediction) Reset()         { *m = NamespacePrediction{} }
func (m *NamespacePrediction) String() string { return proto.CompactTextString(m) }
func (*NamespacePrediction) ProtoMessage()    {}
func (*NamespacePrediction) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6f58e3dfb6d18af, []int{4}
}

func (m *NamespacePrediction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamespacePrediction.Unmarshal(m, b)
}
func (m *NamespacePrediction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamespacePrediction.Marshal(b, m, deterministic)
}
func (m *NamespacePrediction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespacePrediction.Merge(m, src)
}
func (m *NamespacePrediction) XXX_Size() int {
	return xxx_messageInfo_NamespacePrediction.Size(m)
}
func (m *NamespacePrediction) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespacePrediction.DiscardUnknown(m)
}

var xxx_messageInfo_NamespacePrediction proto.InternalMessageInfo

func (m *NamespacePrediction) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *NamespacePrediction) GetPredictedRawData() []*MetricData {
	if m != nil {
		return m.PredictedRawData
	}
	return nil
}

func (m *NamespacePrediction) GetPredictedUpperboundData() []*MetricData {
	if m != nil {
		return m.PredictedUpperboundData
	}
	return nil
}

func (m *NamespacePrediction) GetPredictedLowerboundData() []*MetricData {
	if m != nil {
		return m.PredictedLowerboundData
	}
	return nil
}

// Represents a list of predicted metric data of a node
type NodePrediction struct {
	ObjectMeta              *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	IsScheduled             bool                  `protobuf:"varint,2,opt,name=is_scheduled,json=isScheduled,proto3" json:"is_scheduled,omitempty"`
	PredictedRawData        []*MetricData         `protobuf:"bytes,3,rep,name=predicted_raw_data,json=predictedRawData,proto3" json:"predicted_raw_data,omitempty"`
	PredictedUpperboundData []*MetricData         `protobuf:"bytes,4,rep,name=predicted_upperbound_data,json=predictedUpperboundData,proto3" json:"predicted_upperbound_data,omitempty"`
	PredictedLowerboundData []*MetricData         `protobuf:"bytes,5,rep,name=predicted_lowerbound_data,json=predictedLowerboundData,proto3" json:"predicted_lowerbound_data,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *NodePrediction) Reset()         { *m = NodePrediction{} }
func (m *NodePrediction) String() string { return proto.CompactTextString(m) }
func (*NodePrediction) ProtoMessage()    {}
func (*NodePrediction) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6f58e3dfb6d18af, []int{5}
}

func (m *NodePrediction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodePrediction.Unmarshal(m, b)
}
func (m *NodePrediction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodePrediction.Marshal(b, m, deterministic)
}
func (m *NodePrediction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodePrediction.Merge(m, src)
}
func (m *NodePrediction) XXX_Size() int {
	return xxx_messageInfo_NodePrediction.Size(m)
}
func (m *NodePrediction) XXX_DiscardUnknown() {
	xxx_messageInfo_NodePrediction.DiscardUnknown(m)
}

var xxx_messageInfo_NodePrediction proto.InternalMessageInfo

func (m *NodePrediction) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *NodePrediction) GetIsScheduled() bool {
	if m != nil {
		return m.IsScheduled
	}
	return false
}

func (m *NodePrediction) GetPredictedRawData() []*MetricData {
	if m != nil {
		return m.PredictedRawData
	}
	return nil
}

func (m *NodePrediction) GetPredictedUpperboundData() []*MetricData {
	if m != nil {
		return m.PredictedUpperboundData
	}
	return nil
}

func (m *NodePrediction) GetPredictedLowerboundData() []*MetricData {
	if m != nil {
		return m.PredictedLowerboundData
	}
	return nil
}

type ClusterPrediction struct {
	ObjectMeta              *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	PredictedRawData        []*MetricData         `protobuf:"bytes,2,rep,name=predicted_raw_data,json=predictedRawData,proto3" json:"predicted_raw_data,omitempty"`
	PredictedUpperboundData []*MetricData         `protobuf:"bytes,3,rep,name=predicted_upperbound_data,json=predictedUpperboundData,proto3" json:"predicted_upperbound_data,omitempty"`
	PredictedLowerboundData []*MetricData         `protobuf:"bytes,4,rep,name=predicted_lowerbound_data,json=predictedLowerboundData,proto3" json:"predicted_lowerbound_data,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *ClusterPrediction) Reset()         { *m = ClusterPrediction{} }
func (m *ClusterPrediction) String() string { return proto.CompactTextString(m) }
func (*ClusterPrediction) ProtoMessage()    {}
func (*ClusterPrediction) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6f58e3dfb6d18af, []int{6}
}

func (m *ClusterPrediction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterPrediction.Unmarshal(m, b)
}
func (m *ClusterPrediction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterPrediction.Marshal(b, m, deterministic)
}
func (m *ClusterPrediction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterPrediction.Merge(m, src)
}
func (m *ClusterPrediction) XXX_Size() int {
	return xxx_messageInfo_ClusterPrediction.Size(m)
}
func (m *ClusterPrediction) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterPrediction.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterPrediction proto.InternalMessageInfo

func (m *ClusterPrediction) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *ClusterPrediction) GetPredictedRawData() []*MetricData {
	if m != nil {
		return m.PredictedRawData
	}
	return nil
}

func (m *ClusterPrediction) GetPredictedUpperboundData() []*MetricData {
	if m != nil {
		return m.PredictedUpperboundData
	}
	return nil
}

func (m *ClusterPrediction) GetPredictedLowerboundData() []*MetricData {
	if m != nil {
		return m.PredictedLowerboundData
	}
	return nil
}

type ReadPrediction struct {
	SchemaMeta              *schemas.SchemaMeta      `protobuf:"bytes,1,opt,name=schema_meta,json=schemaMeta,proto3" json:"schema_meta,omitempty"`
	PredictedRawData        []*common.ReadMetricData `protobuf:"bytes,2,rep,name=predicted_raw_data,json=predictedRawData,proto3" json:"predicted_raw_data,omitempty"`
	PredictedUpperboundData []*common.ReadMetricData `protobuf:"bytes,3,rep,name=predicted_upperbound_data,json=predictedUpperboundData,proto3" json:"predicted_upperbound_data,omitempty"`
	PredictedLowerboundData []*common.ReadMetricData `protobuf:"bytes,4,rep,name=predicted_lowerbound_data,json=predictedLowerboundData,proto3" json:"predicted_lowerbound_data,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                 `json:"-"`
	XXX_unrecognized        []byte                   `json:"-"`
	XXX_sizecache           int32                    `json:"-"`
}

func (m *ReadPrediction) Reset()         { *m = ReadPrediction{} }
func (m *ReadPrediction) String() string { return proto.CompactTextString(m) }
func (*ReadPrediction) ProtoMessage()    {}
func (*ReadPrediction) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6f58e3dfb6d18af, []int{7}
}

func (m *ReadPrediction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadPrediction.Unmarshal(m, b)
}
func (m *ReadPrediction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadPrediction.Marshal(b, m, deterministic)
}
func (m *ReadPrediction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadPrediction.Merge(m, src)
}
func (m *ReadPrediction) XXX_Size() int {
	return xxx_messageInfo_ReadPrediction.Size(m)
}
func (m *ReadPrediction) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadPrediction.DiscardUnknown(m)
}

var xxx_messageInfo_ReadPrediction proto.InternalMessageInfo

func (m *ReadPrediction) GetSchemaMeta() *schemas.SchemaMeta {
	if m != nil {
		return m.SchemaMeta
	}
	return nil
}

func (m *ReadPrediction) GetPredictedRawData() []*common.ReadMetricData {
	if m != nil {
		return m.PredictedRawData
	}
	return nil
}

func (m *ReadPrediction) GetPredictedUpperboundData() []*common.ReadMetricData {
	if m != nil {
		return m.PredictedUpperboundData
	}
	return nil
}

func (m *ReadPrediction) GetPredictedLowerboundData() []*common.ReadMetricData {
	if m != nil {
		return m.PredictedLowerboundData
	}
	return nil
}

type WritePrediction struct {
	SchemaMeta              *schemas.SchemaMeta       `protobuf:"bytes,1,opt,name=schema_meta,json=schemaMeta,proto3" json:"schema_meta,omitempty"`
	PredictedRawData        []*common.WriteMetricData `protobuf:"bytes,2,rep,name=predicted_raw_data,json=predictedRawData,proto3" json:"predicted_raw_data,omitempty"`
	PredictedUpperboundData []*common.WriteMetricData `protobuf:"bytes,3,rep,name=predicted_upperbound_data,json=predictedUpperboundData,proto3" json:"predicted_upperbound_data,omitempty"`
	PredictedLowerboundData []*common.WriteMetricData `protobuf:"bytes,4,rep,name=predicted_lowerbound_data,json=predictedLowerboundData,proto3" json:"predicted_lowerbound_data,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                  `json:"-"`
	XXX_unrecognized        []byte                    `json:"-"`
	XXX_sizecache           int32                     `json:"-"`
}

func (m *WritePrediction) Reset()         { *m = WritePrediction{} }
func (m *WritePrediction) String() string { return proto.CompactTextString(m) }
func (*WritePrediction) ProtoMessage()    {}
func (*WritePrediction) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6f58e3dfb6d18af, []int{8}
}

func (m *WritePrediction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WritePrediction.Unmarshal(m, b)
}
func (m *WritePrediction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WritePrediction.Marshal(b, m, deterministic)
}
func (m *WritePrediction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WritePrediction.Merge(m, src)
}
func (m *WritePrediction) XXX_Size() int {
	return xxx_messageInfo_WritePrediction.Size(m)
}
func (m *WritePrediction) XXX_DiscardUnknown() {
	xxx_messageInfo_WritePrediction.DiscardUnknown(m)
}

var xxx_messageInfo_WritePrediction proto.InternalMessageInfo

func (m *WritePrediction) GetSchemaMeta() *schemas.SchemaMeta {
	if m != nil {
		return m.SchemaMeta
	}
	return nil
}

func (m *WritePrediction) GetPredictedRawData() []*common.WriteMetricData {
	if m != nil {
		return m.PredictedRawData
	}
	return nil
}

func (m *WritePrediction) GetPredictedUpperboundData() []*common.WriteMetricData {
	if m != nil {
		return m.PredictedUpperboundData
	}
	return nil
}

func (m *WritePrediction) GetPredictedLowerboundData() []*common.WriteMetricData {
	if m != nil {
		return m.PredictedLowerboundData
	}
	return nil
}

func init() {
	proto.RegisterType((*ContainerPrediction)(nil), "containersai.alameda.v1alpha1.datahub.predictions.ContainerPrediction")
	proto.RegisterType((*PodPrediction)(nil), "containersai.alameda.v1alpha1.datahub.predictions.PodPrediction")
	proto.RegisterType((*ControllerPrediction)(nil), "containersai.alameda.v1alpha1.datahub.predictions.ControllerPrediction")
	proto.RegisterType((*ApplicationPrediction)(nil), "containersai.alameda.v1alpha1.datahub.predictions.ApplicationPrediction")
	proto.RegisterType((*NamespacePrediction)(nil), "containersai.alameda.v1alpha1.datahub.predictions.NamespacePrediction")
	proto.RegisterType((*NodePrediction)(nil), "containersai.alameda.v1alpha1.datahub.predictions.NodePrediction")
	proto.RegisterType((*ClusterPrediction)(nil), "containersai.alameda.v1alpha1.datahub.predictions.ClusterPrediction")
	proto.RegisterType((*ReadPrediction)(nil), "containersai.alameda.v1alpha1.datahub.predictions.ReadPrediction")
	proto.RegisterType((*WritePrediction)(nil), "containersai.alameda.v1alpha1.datahub.predictions.WritePrediction")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/predictions/predictions.proto", fileDescriptor_a6f58e3dfb6d18af)
}

var fileDescriptor_a6f58e3dfb6d18af = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x98, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0xd5, 0xa5, 0x20, 0x70, 0xa1, 0x40, 0xb6, 0x89, 0xb2, 0xab, 0xd2, 0xab, 0xde, 0x90,
	0xd0, 0xc2, 0x2e, 0xf8, 0xba, 0x80, 0xf1, 0x21, 0x04, 0x1b, 0x53, 0x26, 0x84, 0x34, 0x21, 0x45,
	0xa7, 0xb6, 0x45, 0xbd, 0x25, 0xb1, 0x65, 0x3b, 0x74, 0x13, 0xaf, 0xc0, 0x53, 0xf0, 0x3c, 0x5c,
	0x70, 0x81, 0x10, 0xcf, 0xc0, 0x35, 0x0f, 0x80, 0x9c, 0xa6, 0x59, 0xca, 0xe8, 0xc7, 0xa0, 0xcd,
	0x24, 0xd4, 0xab, 0x5a, 0x6e, 0xcf, 0xff, 0x77, 0xec, 0xbf, 0xcf, 0x69, 0x1c, 0x74, 0x17, 0x02,
	0x08, 0x29, 0x01, 0x1f, 0x04, 0x73, 0xdf, 0xb7, 0x20, 0x10, 0x5d, 0x68, 0xb9, 0x04, 0x34, 0x74,
	0xe3, 0x8e, 0x2b, 0x24, 0x25, 0x0c, 0x6b, 0xc6, 0x23, 0x95, 0x1f, 0x3b, 0x42, 0x72, 0xcd, 0xed,
	0x16, 0xe6, 0x91, 0x06, 0x16, 0x51, 0xa9, 0x80, 0x39, 0xa9, 0x90, 0x33, 0x10, 0x71, 0x52, 0x11,
	0x27, 0x17, 0xb8, 0xd6, 0x1a, 0x8b, 0xc3, 0x3c, 0x0c, 0x79, 0xe4, 0x86, 0x54, 0x4b, 0x86, 0x53,
	0xca, 0xda, 0xed, 0xa9, 0x33, 0xd4, 0x87, 0x82, 0x0e, 0xa2, 0xd6, 0xc7, 0x46, 0x49, 0xaa, 0x78,
	0x2c, 0x31, 0x55, 0x86, 0x05, 0x66, 0x36, 0x0d, 0xbb, 0x39, 0x36, 0x4c, 0xe1, 0x2e, 0x0d, 0x61,
	0x08, 0xd4, 0xf8, 0x68, 0xa1, 0xe5, 0x8d, 0xc1, 0x3e, 0x6c, 0x67, 0xd9, 0xd8, 0x36, 0x2a, 0x47,
	0x10, 0xd2, 0x5a, 0xa9, 0x5e, 0x6a, 0x9e, 0xf7, 0x92, 0xb1, 0xbd, 0x8f, 0xec, 0x34, 0x5f, 0x4a,
	0x7c, 0x09, 0x3d, 0xdf, 0x08, 0xd7, 0x96, 0xea, 0x56, 0xb3, 0xd2, 0x7e, 0xe0, 0x9c, 0x78, 0x37,
	0x9d, 0xcd, 0x64, 0xa3, 0x1e, 0x83, 0x06, 0xef, 0x72, 0x26, 0xec, 0x41, 0xcf, 0xcc, 0xd8, 0x87,
	0xe8, 0xda, 0x11, 0x2c, 0x16, 0x82, 0xca, 0x0e, 0x8f, 0x23, 0xd2, 0x67, 0x5a, 0xb3, 0x60, 0x5e,
	0xcd, 0xf4, 0x5f, 0x67, 0xf2, 0xc7, 0xd1, 0x01, 0xef, 0x0d, 0xa1, 0xcb, 0xb3, 0x45, 0xbf, 0xcc,
	0xe4, 0xcd, 0x17, 0x8d, 0x1f, 0x25, 0x74, 0x71, 0x9b, 0x93, 0x9c, 0x11, 0x6f, 0x51, 0x85, 0x77,
	0xf6, 0x28, 0xd6, 0xbe, 0xf1, 0x3a, 0xf1, 0xa3, 0xd2, 0xbe, 0x37, 0x25, 0x3e, 0x3b, 0x28, 0xce,
	0xab, 0x44, 0x63, 0x93, 0x6a, 0xf0, 0x10, 0xcf, 0xc6, 0xf6, 0x07, 0xb4, 0x9a, 0x29, 0xf9, 0xb9,
	0x64, 0x53, 0x57, 0x9f, 0xfe, 0xc5, 0x32, 0xff, 0x70, 0x9a, 0xbc, 0x15, 0x7c, 0x7c, 0x52, 0x35,
	0x3e, 0x95, 0xd1, 0x8a, 0xf9, 0xb5, 0xe4, 0x41, 0x30, 0x74, 0xf8, 0xe6, 0xbb, 0xe6, 0xe7, 0xa8,
	0xbc, 0xcf, 0x22, 0x52, 0x5b, 0xaa, 0x97, 0x9a, 0xd5, 0xf6, 0xfa, 0x89, 0x65, 0x5f, 0xb0, 0x88,
	0x78, 0x89, 0xc4, 0x88, 0x8a, 0xb0, 0x4e, 0xa1, 0x22, 0xca, 0xa7, 0x57, 0x11, 0x67, 0xe6, 0x5a,
	0x11, 0xdf, 0x2d, 0xb4, 0xfa, 0x50, 0x88, 0x80, 0x61, 0x30, 0x41, 0x85, 0x9d, 0x92, 0x45, 0xb3,
	0x9b, 0x7b, 0xb3, 0xfb, 0x66, 0xa1, 0xe5, 0x2d, 0x08, 0xa9, 0x12, 0x80, 0xe9, 0xc2, 0xd8, 0xff,
	0xc7, 0xd8, 0x9f, 0x16, 0xaa, 0x6e, 0x71, 0x52, 0x9c, 0xa7, 0xd7, 0xd1, 0x05, 0xa6, 0x7c, 0xf3,
	0x7c, 0x43, 0xe2, 0x80, 0xf6, 0x5b, 0xfb, 0x39, 0xaf, 0xc2, 0xd4, 0xce, 0x60, 0x6a, 0xd1, 0xaa,
	0xe7, 0xdf, 0xaa, 0xbf, 0x5a, 0xe8, 0xca, 0x46, 0x10, 0x2b, 0x5d, 0xe0, 0x9f, 0xf9, 0xa2, 0x9a,
	0xe7, 0x5e, 0xcd, 0x9f, 0x2d, 0x54, 0xf5, 0x28, 0xe4, 0x1f, 0x4a, 0x77, 0x51, 0xa5, 0x7f, 0x99,
	0xc8, 0x7b, 0x7a, 0x67, 0x4a, 0x7e, 0x7a, 0x0d, 0x71, 0x76, 0x92, 0xcf, 0xbe, 0xa3, 0x2a, 0x1b,
	0xdb, 0x7b, 0x63, 0x1c, 0xbd, 0x3f, 0x25, 0xa2, 0x7f, 0x13, 0x73, 0x4c, 0xd6, 0x63, 0x0d, 0x3d,
	0x98, 0x6c, 0xe8, 0xbf, 0x21, 0x47, 0xfa, 0x79, 0x30, 0xd9, 0xcf, 0x59, 0x91, 0x7f, 0xb3, 0xf3,
	0x8b, 0x85, 0x2e, 0xbd, 0x91, 0x4c, 0xd3, 0x82, 0xfc, 0x9c, 0x45, 0x85, 0xa6, 0x4b, 0x4c, 0xd2,
	0x2e, 0xaa, 0x42, 0x47, 0x30, 0x8b, 0xa8, 0xd0, 0x89, 0xe8, 0x61, 0x4b, 0x1f, 0x3d, 0xdb, 0x7d,
	0xf2, 0x8e, 0xe9, 0x34, 0xd2, 0x3d, 0x62, 0xdc, 0x00, 0xe6, 0x82, 0x60, 0xee, 0xb4, 0xaf, 0x20,
	0x3a, 0x67, 0x93, 0x97, 0x02, 0xb7, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x67, 0x64, 0xd6,
	0x57, 0x11, 0x00, 0x00,
}
