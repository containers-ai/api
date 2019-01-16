// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datahub/metric/v1alpha2/metric.proto

package containersai_datahub_metric_v1alpha2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v1alpha2 "datahub/resource/metadata/v1alpha2"
import duration "github.com/golang/protobuf/ptypes/duration"
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
// Metric type. A metric may be either CPU or memory.
type MetricType int32

const (
	MetricType_UNDEFINED                    MetricType = 0
	MetricType_CPU_USAGE_SECONDS_PERCENTAGE MetricType = 1
	MetricType_MEMORY_USAGE_BYTES           MetricType = 2
)

var MetricType_name = map[int32]string{
	0: "UNDEFINED",
	1: "CPU_USAGE_SECONDS_PERCENTAGE",
	2: "MEMORY_USAGE_BYTES",
}
var MetricType_value = map[string]int32{
	"UNDEFINED":                    0,
	"CPU_USAGE_SECONDS_PERCENTAGE": 1,
	"MEMORY_USAGE_BYTES":           2,
}

func (x MetricType) String() string {
	return proto.EnumName(MetricType_name, int32(x))
}
func (MetricType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_metric_6251774d64f426c9, []int{0}
}

// *
// Represents metric data of a container
type ContainerMetric struct {
	Name                 string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MetricData           map[string]*MetricData `protobuf:"bytes,2,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ContainerMetric) Reset()         { *m = ContainerMetric{} }
func (m *ContainerMetric) String() string { return proto.CompactTextString(m) }
func (*ContainerMetric) ProtoMessage()    {}
func (*ContainerMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_6251774d64f426c9, []int{0}
}
func (m *ContainerMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerMetric.Unmarshal(m, b)
}
func (m *ContainerMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerMetric.Marshal(b, m, deterministic)
}
func (dst *ContainerMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerMetric.Merge(dst, src)
}
func (m *ContainerMetric) XXX_Size() int {
	return xxx_messageInfo_ContainerMetric.Size(m)
}
func (m *ContainerMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerMetric.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerMetric proto.InternalMessageInfo

func (m *ContainerMetric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerMetric) GetMetricData() map[string]*MetricData {
	if m != nil {
		return m.MetricData
	}
	return nil
}

// *
// Represents metric data of a pod
type PodMetric struct {
	NamespacedName       *v1alpha2.NamespacedName `protobuf:"bytes,1,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	ContainerMetrics     []*ContainerMetric       `protobuf:"bytes,2,rep,name=container_metrics,json=containerMetrics,proto3" json:"container_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *PodMetric) Reset()         { *m = PodMetric{} }
func (m *PodMetric) String() string { return proto.CompactTextString(m) }
func (*PodMetric) ProtoMessage()    {}
func (*PodMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_6251774d64f426c9, []int{1}
}
func (m *PodMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodMetric.Unmarshal(m, b)
}
func (m *PodMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodMetric.Marshal(b, m, deterministic)
}
func (dst *PodMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodMetric.Merge(dst, src)
}
func (m *PodMetric) XXX_Size() int {
	return xxx_messageInfo_PodMetric.Size(m)
}
func (m *PodMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_PodMetric.DiscardUnknown(m)
}

var xxx_messageInfo_PodMetric proto.InternalMessageInfo

func (m *PodMetric) GetNamespacedName() *v1alpha2.NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *PodMetric) GetContainerMetrics() []*ContainerMetric {
	if m != nil {
		return m.ContainerMetrics
	}
	return nil
}

// *
// Represents metric data of a node
type NodeMetric struct {
	Name                 string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MetricData           map[string]*MetricData `protobuf:"bytes,2,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *NodeMetric) Reset()         { *m = NodeMetric{} }
func (m *NodeMetric) String() string { return proto.CompactTextString(m) }
func (*NodeMetric) ProtoMessage()    {}
func (*NodeMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_6251774d64f426c9, []int{2}
}
func (m *NodeMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeMetric.Unmarshal(m, b)
}
func (m *NodeMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeMetric.Marshal(b, m, deterministic)
}
func (dst *NodeMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeMetric.Merge(dst, src)
}
func (m *NodeMetric) XXX_Size() int {
	return xxx_messageInfo_NodeMetric.Size(m)
}
func (m *NodeMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeMetric.DiscardUnknown(m)
}

var xxx_messageInfo_NodeMetric proto.InternalMessageInfo

func (m *NodeMetric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeMetric) GetMetricData() map[string]*MetricData {
	if m != nil {
		return m.MetricData
	}
	return nil
}

// *
// Represents a data point of time-series metric data
type Sample struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	NumValue             string               `protobuf:"bytes,2,opt,name=num_value,json=numValue,proto3" json:"num_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Sample) Reset()         { *m = Sample{} }
func (m *Sample) String() string { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()    {}
func (*Sample) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_6251774d64f426c9, []int{3}
}
func (m *Sample) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sample.Unmarshal(m, b)
}
func (m *Sample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sample.Marshal(b, m, deterministic)
}
func (dst *Sample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sample.Merge(dst, src)
}
func (m *Sample) XXX_Size() int {
	return xxx_messageInfo_Sample.Size(m)
}
func (m *Sample) XXX_DiscardUnknown() {
	xxx_messageInfo_Sample.DiscardUnknown(m)
}

var xxx_messageInfo_Sample proto.InternalMessageInfo

func (m *Sample) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Sample) GetNumValue() string {
	if m != nil {
		return m.NumValue
	}
	return ""
}

// *
// Represents a time range definition
//
type TimeRange struct {
	StartTime            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Step                 *duration.Duration   `protobuf:"bytes,3,opt,name=step,proto3" json:"step,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimeRange) Reset()         { *m = TimeRange{} }
func (m *TimeRange) String() string { return proto.CompactTextString(m) }
func (*TimeRange) ProtoMessage()    {}
func (*TimeRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_6251774d64f426c9, []int{4}
}
func (m *TimeRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRange.Unmarshal(m, b)
}
func (m *TimeRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRange.Marshal(b, m, deterministic)
}
func (dst *TimeRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRange.Merge(dst, src)
}
func (m *TimeRange) XXX_Size() int {
	return xxx_messageInfo_TimeRange.Size(m)
}
func (m *TimeRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRange.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRange proto.InternalMessageInfo

func (m *TimeRange) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TimeRange) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *TimeRange) GetStep() *duration.Duration {
	if m != nil {
		return m.Step
	}
	return nil
}

// *
// Represents a piece of metreic data
type MetricData struct {
	// data can be time series or non-time series
	Data                 []*Sample `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MetricData) Reset()         { *m = MetricData{} }
func (m *MetricData) String() string { return proto.CompactTextString(m) }
func (*MetricData) ProtoMessage()    {}
func (*MetricData) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_6251774d64f426c9, []int{5}
}
func (m *MetricData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricData.Unmarshal(m, b)
}
func (m *MetricData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricData.Marshal(b, m, deterministic)
}
func (dst *MetricData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricData.Merge(dst, src)
}
func (m *MetricData) XXX_Size() int {
	return xxx_messageInfo_MetricData.Size(m)
}
func (m *MetricData) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricData.DiscardUnknown(m)
}

var xxx_messageInfo_MetricData proto.InternalMessageInfo

func (m *MetricData) GetData() []*Sample {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ContainerMetric)(nil), "containersai.datahub.metric.v1alpha2.ContainerMetric")
	proto.RegisterMapType((map[string]*MetricData)(nil), "containersai.datahub.metric.v1alpha2.ContainerMetric.MetricDataEntry")
	proto.RegisterType((*PodMetric)(nil), "containersai.datahub.metric.v1alpha2.PodMetric")
	proto.RegisterType((*NodeMetric)(nil), "containersai.datahub.metric.v1alpha2.NodeMetric")
	proto.RegisterMapType((map[string]*MetricData)(nil), "containersai.datahub.metric.v1alpha2.NodeMetric.MetricDataEntry")
	proto.RegisterType((*Sample)(nil), "containersai.datahub.metric.v1alpha2.Sample")
	proto.RegisterType((*TimeRange)(nil), "containersai.datahub.metric.v1alpha2.TimeRange")
	proto.RegisterType((*MetricData)(nil), "containersai.datahub.metric.v1alpha2.MetricData")
	proto.RegisterEnum("containersai.datahub.metric.v1alpha2.MetricType", MetricType_name, MetricType_value)
}

func init() {
	proto.RegisterFile("datahub/metric/v1alpha2/metric.proto", fileDescriptor_metric_6251774d64f426c9)
}

var fileDescriptor_metric_6251774d64f426c9 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x41, 0x8f, 0xd2, 0x40,
	0x14, 0xb6, 0x80, 0xeb, 0xf6, 0x11, 0x05, 0xe7, 0x60, 0x10, 0x8d, 0x12, 0xb2, 0x07, 0x62, 0x74,
	0x70, 0x31, 0x9b, 0xa8, 0x17, 0x77, 0x85, 0xd9, 0x8d, 0x07, 0xba, 0x64, 0x00, 0x93, 0x3d, 0x35,
	0x03, 0x9d, 0x05, 0x22, 0x9d, 0x36, 0xed, 0x74, 0x13, 0xfe, 0x93, 0x7f, 0xc8, 0xff, 0xe1, 0xd9,
	0x98, 0x99, 0x4e, 0x4b, 0xc0, 0x8d, 0x36, 0x9e, 0xf6, 0xd4, 0x79, 0x6f, 0xde, 0xf7, 0xcd, 0xf7,
	0xbe, 0x97, 0x57, 0x38, 0xf2, 0x98, 0x64, 0xcb, 0x64, 0xd6, 0xf5, 0xb9, 0x8c, 0x56, 0xf3, 0xee,
	0xcd, 0x31, 0x5b, 0x87, 0x4b, 0xd6, 0x33, 0x31, 0x0e, 0xa3, 0x40, 0x06, 0xe8, 0x68, 0x1e, 0x08,
	0xc9, 0x56, 0x82, 0x47, 0x31, 0x5b, 0x61, 0x03, 0xc1, 0xa6, 0x24, 0x83, 0x34, 0x5f, 0x2e, 0x82,
	0x60, 0xb1, 0xe6, 0x5d, 0x8d, 0x99, 0x25, 0xd7, 0x5d, 0xb9, 0xf2, 0x79, 0x2c, 0x99, 0x1f, 0xa6,
	0x34, 0xcd, 0x17, 0xfb, 0x05, 0x5e, 0x12, 0x31, 0xb9, 0x0a, 0x84, 0xb9, 0x3f, 0xce, 0xc4, 0x44,
	0x3c, 0x0e, 0x92, 0x68, 0xce, 0x95, 0x0a, 0xa6, 0x92, 0x3b, 0xba, 0x74, 0x26, 0x85, 0xb4, 0x7f,
	0x59, 0x50, 0xeb, 0x67, 0xe2, 0x86, 0x5a, 0x10, 0x42, 0x50, 0x11, 0xcc, 0xe7, 0x0d, 0xab, 0x65,
	0x75, 0x6c, 0xaa, 0xcf, 0xe8, 0x1a, 0xaa, 0xa9, 0x5c, 0x57, 0x81, 0x1b, 0xa5, 0x56, 0xb9, 0x53,
	0xed, 0x11, 0x5c, 0xa4, 0x2f, 0xbc, 0xc7, 0x8f, 0xd3, 0xcf, 0x80, 0x49, 0x46, 0x84, 0x8c, 0x36,
	0x14, 0xfc, 0x3c, 0xd1, 0x0c, 0xa0, 0xb6, 0x77, 0x8d, 0xea, 0x50, 0xfe, 0xc6, 0x37, 0x46, 0x8d,
	0x3a, 0xa2, 0x73, 0xb8, 0x7f, 0xc3, 0xd6, 0x09, 0x6f, 0x94, 0x5a, 0x56, 0xa7, 0xda, 0x7b, 0x5b,
	0x4c, 0xc6, 0x96, 0x97, 0xa6, 0xf0, 0x8f, 0xa5, 0xf7, 0x56, 0xfb, 0x87, 0x05, 0xf6, 0x28, 0xf0,
	0x4c, 0xeb, 0x4b, 0xa8, 0xa9, 0x76, 0xe3, 0x90, 0xcd, 0xb9, 0xe7, 0xe6, 0x2e, 0x54, 0x7b, 0x9f,
	0x6e, 0x7f, 0x23, 0x33, 0x1a, 0xe7, 0xb6, 0xe6, 0xcf, 0x39, 0x39, 0x8f, 0x3a, 0xd1, 0x47, 0x62,
	0x27, 0x46, 0x33, 0x78, 0x9c, 0x33, 0xba, 0xa9, 0xd4, 0xd8, 0xd8, 0x7a, 0xf2, 0x5f, 0xb6, 0xd2,
	0xfa, 0x7c, 0x37, 0x11, 0xb7, 0x7f, 0x5a, 0x00, 0x4e, 0xe0, 0xf1, 0xbf, 0xcc, 0x95, 0xdd, 0x36,
	0xd7, 0xd3, 0x62, 0x02, 0xb6, 0xd4, 0x77, 0x6b, 0xa4, 0x53, 0x38, 0x18, 0x33, 0x3f, 0x5c, 0x73,
	0x84, 0xa1, 0xa2, 0x76, 0xc8, 0xcc, 0xb0, 0x89, 0xd3, 0xfd, 0xc1, 0xd9, 0xfe, 0xe0, 0x49, 0xb6,
	0x60, 0x54, 0xd7, 0xa1, 0x67, 0x60, 0x8b, 0xc4, 0x77, 0xb7, 0x4a, 0x6c, 0x7a, 0x28, 0x12, 0xff,
	0xab, 0x8a, 0xdb, 0xdf, 0x2d, 0xb0, 0x15, 0x80, 0x32, 0xb1, 0xe0, 0xe8, 0x03, 0x40, 0x2c, 0x59,
	0x24, 0xdd, 0x82, 0x0f, 0xd8, 0xba, 0x5a, 0xc5, 0xe8, 0x04, 0x0e, 0xb9, 0xf0, 0x52, 0x60, 0xe9,
	0x9f, 0xc0, 0x07, 0x5c, 0x78, 0x1a, 0xf6, 0x06, 0x2a, 0xb1, 0xe4, 0x61, 0xa3, 0xac, 0x21, 0x4f,
	0xff, 0x80, 0x0c, 0xcc, 0xcf, 0x80, 0xea, 0xb2, 0xb6, 0x03, 0xb0, 0xb5, 0x07, 0x9d, 0x42, 0x45,
	0x0f, 0xd8, 0xd2, 0x03, 0x7e, 0x5d, 0xcc, 0xde, 0xd4, 0x45, 0xaa, 0x91, 0xaf, 0xa6, 0x19, 0xdf,
	0x64, 0x13, 0x72, 0xf4, 0x10, 0xec, 0xa9, 0x33, 0x20, 0xe7, 0x5f, 0x1c, 0x32, 0xa8, 0xdf, 0x43,
	0x2d, 0x78, 0xde, 0x1f, 0x4d, 0xdd, 0xe9, 0xf8, 0xec, 0x82, 0xb8, 0x63, 0xd2, 0xbf, 0x74, 0x06,
	0x63, 0x77, 0x44, 0x68, 0x9f, 0x38, 0x93, 0xb3, 0x0b, 0x52, 0xb7, 0xd0, 0x13, 0x40, 0x43, 0x32,
	0xbc, 0xa4, 0x57, 0xa6, 0xe8, 0xf3, 0xd5, 0x84, 0x8c, 0xeb, 0xa5, 0xd9, 0x81, 0xd6, 0xff, 0xee,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x24, 0x66, 0xde, 0x1b, 0x49, 0x05, 0x00, 0x00,
}
