// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/common/metrics.proto

package common

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
// Metric type. A metric may be CPU, memory and etc.
type MetricType int32

const (
	MetricType_METRICS_TYPE_UNDEFINED   MetricType = 0
	MetricType_CPU_SECONDS_TOTAL        MetricType = 1
	MetricType_CPU_CORES_ALLOCATABLE    MetricType = 2
	MetricType_CPU_MILLICORES_TOTAL     MetricType = 3
	MetricType_CPU_MILLICORES_AVAIL     MetricType = 4
	MetricType_CPU_MILLICORES_USAGE     MetricType = 5
	MetricType_CPU_MILLICORES_USAGE_PCT MetricType = 6
	MetricType_MEMORY_BYTES_ALLOCATABLE MetricType = 7
	MetricType_MEMORY_BYTES_TOTAL       MetricType = 8
	MetricType_MEMORY_BYTES_AVAIL       MetricType = 9
	MetricType_MEMORY_BYTES_USAGE       MetricType = 10
	MetricType_MEMORY_BYTES_USAGE_PCT   MetricType = 11
	MetricType_FS_BYTES_TOTAL           MetricType = 12
	MetricType_FS_BYTES_AVAIL           MetricType = 13
	MetricType_FS_BYTES_USAGE           MetricType = 14
	MetricType_FS_BYTES_USAGE_PCT       MetricType = 15
	MetricType_HTTP_REQUESTS_COUNT      MetricType = 16
	MetricType_HTTP_REQUESTS_TOTAL      MetricType = 17
	MetricType_HTTP_RESPONSE_COUNT      MetricType = 18
	MetricType_HTTP_RESPONSE_TOTAL      MetricType = 19
	MetricType_DISK_IO_SECONDS_TOTAL    MetricType = 20
	MetricType_DISK_IO_UTILIZATION      MetricType = 21
	MetricType_RESTARTS_TOTAL           MetricType = 22
	MetricType_UNSCHEDULABLE            MetricType = 23
	MetricType_HEALTH                   MetricType = 24
	MetricType_POWER_USAGE_WATTS        MetricType = 25
	MetricType_TEMPERATURE_CELSIUS      MetricType = 26
	MetricType_DUTY_CYCLE               MetricType = 27
	MetricType_CURRENT_OFFSET           MetricType = 28
	MetricType_LAG                      MetricType = 29
	MetricType_LATENCY                  MetricType = 30
	MetricType_NUMBER                   MetricType = 31
)

var MetricType_name = map[int32]string{
	0:  "METRICS_TYPE_UNDEFINED",
	1:  "CPU_SECONDS_TOTAL",
	2:  "CPU_CORES_ALLOCATABLE",
	3:  "CPU_MILLICORES_TOTAL",
	4:  "CPU_MILLICORES_AVAIL",
	5:  "CPU_MILLICORES_USAGE",
	6:  "CPU_MILLICORES_USAGE_PCT",
	7:  "MEMORY_BYTES_ALLOCATABLE",
	8:  "MEMORY_BYTES_TOTAL",
	9:  "MEMORY_BYTES_AVAIL",
	10: "MEMORY_BYTES_USAGE",
	11: "MEMORY_BYTES_USAGE_PCT",
	12: "FS_BYTES_TOTAL",
	13: "FS_BYTES_AVAIL",
	14: "FS_BYTES_USAGE",
	15: "FS_BYTES_USAGE_PCT",
	16: "HTTP_REQUESTS_COUNT",
	17: "HTTP_REQUESTS_TOTAL",
	18: "HTTP_RESPONSE_COUNT",
	19: "HTTP_RESPONSE_TOTAL",
	20: "DISK_IO_SECONDS_TOTAL",
	21: "DISK_IO_UTILIZATION",
	22: "RESTARTS_TOTAL",
	23: "UNSCHEDULABLE",
	24: "HEALTH",
	25: "POWER_USAGE_WATTS",
	26: "TEMPERATURE_CELSIUS",
	27: "DUTY_CYCLE",
	28: "CURRENT_OFFSET",
	29: "LAG",
	30: "LATENCY",
	31: "NUMBER",
}

var MetricType_value = map[string]int32{
	"METRICS_TYPE_UNDEFINED":   0,
	"CPU_SECONDS_TOTAL":        1,
	"CPU_CORES_ALLOCATABLE":    2,
	"CPU_MILLICORES_TOTAL":     3,
	"CPU_MILLICORES_AVAIL":     4,
	"CPU_MILLICORES_USAGE":     5,
	"CPU_MILLICORES_USAGE_PCT": 6,
	"MEMORY_BYTES_ALLOCATABLE": 7,
	"MEMORY_BYTES_TOTAL":       8,
	"MEMORY_BYTES_AVAIL":       9,
	"MEMORY_BYTES_USAGE":       10,
	"MEMORY_BYTES_USAGE_PCT":   11,
	"FS_BYTES_TOTAL":           12,
	"FS_BYTES_AVAIL":           13,
	"FS_BYTES_USAGE":           14,
	"FS_BYTES_USAGE_PCT":       15,
	"HTTP_REQUESTS_COUNT":      16,
	"HTTP_REQUESTS_TOTAL":      17,
	"HTTP_RESPONSE_COUNT":      18,
	"HTTP_RESPONSE_TOTAL":      19,
	"DISK_IO_SECONDS_TOTAL":    20,
	"DISK_IO_UTILIZATION":      21,
	"RESTARTS_TOTAL":           22,
	"UNSCHEDULABLE":            23,
	"HEALTH":                   24,
	"POWER_USAGE_WATTS":        25,
	"TEMPERATURE_CELSIUS":      26,
	"DUTY_CYCLE":               27,
	"CURRENT_OFFSET":           28,
	"LAG":                      29,
	"LATENCY":                  30,
	"NUMBER":                   31,
}

func (x MetricType) String() string {
	return proto.EnumName(MetricType_name, int32(x))
}

func (MetricType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0dcb68d43798345b, []int{0}
}

//*
// Represents Kubernetes resources which will be allocated to pods.
type ResourceName int32

const (
	ResourceName_RESOURCE_NAME_UNDEFINED ResourceName = 0
	ResourceName_CPU                     ResourceName = 1
	ResourceName_MEMORY                  ResourceName = 2
)

var ResourceName_name = map[int32]string{
	0: "RESOURCE_NAME_UNDEFINED",
	1: "CPU",
	2: "MEMORY",
}

var ResourceName_value = map[string]int32{
	"RESOURCE_NAME_UNDEFINED": 0,
	"CPU":                     1,
	"MEMORY":                  2,
}

func (x ResourceName) String() string {
	return proto.EnumName(ResourceName_name, int32(x))
}

func (ResourceName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0dcb68d43798345b, []int{1}
}

//*
// Represents a data point of time-series metric data.
type Sample struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	NumValue             string               `protobuf:"bytes,3,opt,name=num_value,json=numValue,proto3" json:"num_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Sample) Reset()         { *m = Sample{} }
func (m *Sample) String() string { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()    {}
func (*Sample) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dcb68d43798345b, []int{0}
}

func (m *Sample) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sample.Unmarshal(m, b)
}
func (m *Sample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sample.Marshal(b, m, deterministic)
}
func (m *Sample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sample.Merge(m, src)
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

func (m *Sample) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Sample) GetNumValue() string {
	if m != nil {
		return m.NumValue
	}
	return ""
}

//*
// Represents a piece of metreic data.
type MetricData struct {
	MetricType           MetricType `protobuf:"varint,1,opt,name=metric_type,json=metricType,proto3,enum=containersai.alameda.v1alpha1.datahub.common.MetricType" json:"metric_type,omitempty"`
	Data                 []*Sample  `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	Granularity          int64      `protobuf:"varint,3,opt,name=granularity,proto3" json:"granularity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MetricData) Reset()         { *m = MetricData{} }
func (m *MetricData) String() string { return proto.CompactTextString(m) }
func (*MetricData) ProtoMessage()    {}
func (*MetricData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dcb68d43798345b, []int{1}
}

func (m *MetricData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricData.Unmarshal(m, b)
}
func (m *MetricData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricData.Marshal(b, m, deterministic)
}
func (m *MetricData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricData.Merge(m, src)
}
func (m *MetricData) XXX_Size() int {
	return xxx_messageInfo_MetricData.Size(m)
}
func (m *MetricData) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricData.DiscardUnknown(m)
}

var xxx_messageInfo_MetricData proto.InternalMessageInfo

func (m *MetricData) GetMetricType() MetricType {
	if m != nil {
		return m.MetricType
	}
	return MetricType_METRICS_TYPE_UNDEFINED
}

func (m *MetricData) GetData() []*Sample {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MetricData) GetGranularity() int64 {
	if m != nil {
		return m.Granularity
	}
	return 0
}

func init() {
	proto.RegisterEnum("containersai.alameda.v1alpha1.datahub.common.MetricType", MetricType_name, MetricType_value)
	proto.RegisterEnum("containersai.alameda.v1alpha1.datahub.common.ResourceName", ResourceName_name, ResourceName_value)
	proto.RegisterType((*Sample)(nil), "containersai.alameda.v1alpha1.datahub.common.Sample")
	proto.RegisterType((*MetricData)(nil), "containersai.alameda.v1alpha1.datahub.common.MetricData")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/common/metrics.proto", fileDescriptor_0dcb68d43798345b)
}

var fileDescriptor_0dcb68d43798345b = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x4e, 0xe3, 0x38,
	0x18, 0xdd, 0xb6, 0xd0, 0x82, 0x0b, 0x5d, 0x63, 0xfe, 0x4a, 0x61, 0x97, 0x8a, 0x2b, 0x84, 0x76,
	0x13, 0xc1, 0xee, 0x4a, 0x7b, 0x49, 0x9a, 0xba, 0x34, 0xda, 0xfc, 0xad, 0xed, 0x80, 0xca, 0x8d,
	0x65, 0xda, 0x4c, 0x89, 0xd4, 0x24, 0x55, 0x9b, 0x22, 0xf1, 0x0e, 0xf3, 0x76, 0x73, 0x3d, 0xef,
	0x32, 0x72, 0xdc, 0x32, 0x94, 0x22, 0xcd, 0x70, 0x97, 0x9c, 0xe3, 0x73, 0xce, 0xf7, 0x7d, 0xfe,
	0x01, 0x97, 0x62, 0x24, 0xe2, 0x70, 0x20, 0xb8, 0x18, 0x47, 0xfa, 0xd3, 0xa5, 0x18, 0x8d, 0x1f,
	0xc5, 0xa5, 0x3e, 0x10, 0x99, 0x78, 0x9c, 0x3d, 0xe8, 0xfd, 0x34, 0x8e, 0xd3, 0x44, 0x8f, 0xc3,
	0x6c, 0x12, 0xf5, 0xa7, 0xda, 0x78, 0x92, 0x66, 0x29, 0xfa, 0xa3, 0x9f, 0x26, 0x99, 0x88, 0x92,
	0x70, 0x32, 0x15, 0x91, 0x36, 0xd7, 0x6b, 0x0b, 0xad, 0x36, 0xd7, 0x6a, 0x4a, 0xdb, 0x38, 0x1d,
	0xa6, 0xe9, 0x70, 0x14, 0xea, 0xb9, 0xf6, 0x61, 0xf6, 0x49, 0xcf, 0xa2, 0x38, 0x9c, 0x66, 0x22,
	0x1e, 0x2b, 0xbb, 0xb3, 0xcf, 0x05, 0x50, 0xa6, 0x22, 0x1e, 0x8f, 0x42, 0xa4, 0x81, 0x35, 0xc9,
	0xd6, 0x0b, 0xcd, 0xc2, 0x79, 0xf5, 0xaa, 0xa1, 0x29, 0xa9, 0xb6, 0x90, 0x6a, 0x6c, 0x21, 0x25,
	0xf9, 0x3a, 0xf4, 0x0f, 0xd8, 0x08, 0x93, 0x01, 0xcf, 0x35, 0xc5, 0x1f, 0x6a, 0x2a, 0x61, 0x32,
	0x90, 0x7f, 0xe8, 0x18, 0x6c, 0x26, 0xb3, 0x98, 0x3f, 0x89, 0xd1, 0x2c, 0xac, 0x97, 0x9a, 0x85,
	0xf3, 0x4d, 0xb2, 0x91, 0xcc, 0xe2, 0x5b, 0xf9, 0x7f, 0xf6, 0xa5, 0x00, 0x80, 0x93, 0xf7, 0xdb,
	0x16, 0x99, 0x40, 0x3d, 0x50, 0x55, 0xdd, 0xf3, 0xec, 0x79, 0xac, 0x2a, 0xab, 0x5d, 0xfd, 0xab,
	0x7d, 0x64, 0x04, 0x9a, 0xb2, 0x63, 0xcf, 0xe3, 0x90, 0x80, 0xf8, 0xe5, 0x1b, 0x75, 0xc1, 0x9a,
	0x5c, 0x58, 0x2f, 0x36, 0x4b, 0xe7, 0xd5, 0xab, 0xbf, 0x3f, 0xe6, 0xa9, 0x26, 0x46, 0x72, 0x07,
	0xd4, 0x04, 0xd5, 0xe1, 0x44, 0x24, 0xb3, 0x91, 0x98, 0x44, 0xd9, 0x73, 0xde, 0x52, 0x89, 0xbc,
	0x86, 0x2e, 0xbe, 0xae, 0x2f, 0xba, 0xca, 0xa3, 0x1b, 0xe0, 0xc0, 0xc1, 0x8c, 0x58, 0x26, 0xe5,
	0xac, 0xe7, 0x63, 0x1e, 0xb8, 0x6d, 0xdc, 0xb1, 0x5c, 0xdc, 0x86, 0xbf, 0xa0, 0x7d, 0xb0, 0x63,
	0xfa, 0x01, 0xa7, 0xd8, 0xf4, 0xdc, 0x36, 0xe5, 0xcc, 0x63, 0x86, 0x0d, 0x0b, 0xe8, 0x08, 0xec,
	0x4b, 0xd8, 0xf4, 0x08, 0xa6, 0xdc, 0xb0, 0x6d, 0xcf, 0x34, 0x98, 0xd1, 0xb2, 0x31, 0x2c, 0xa2,
	0x3a, 0xd8, 0x93, 0x94, 0x63, 0xd9, 0xb6, 0xa5, 0x78, 0x25, 0x2a, 0xbd, 0xc3, 0x18, 0xb7, 0x86,
	0x65, 0xc3, 0xb5, 0x77, 0x98, 0x80, 0x1a, 0x37, 0x18, 0xae, 0xa3, 0x13, 0x50, 0x7f, 0x8f, 0xe1,
	0xbe, 0xc9, 0x60, 0x59, 0xb2, 0x0e, 0x76, 0x3c, 0xd2, 0xe3, 0xad, 0x1e, 0x7b, 0x53, 0x49, 0x05,
	0x1d, 0x00, 0xb4, 0xc4, 0xaa, 0x3a, 0x36, 0x56, 0x70, 0x55, 0xc5, 0xe6, 0x0a, 0xae, 0x6a, 0x00,
	0x6a, 0x3e, 0x6f, 0xf1, 0xbc, 0x82, 0x2a, 0x42, 0xa0, 0xd6, 0xa1, 0x4b, 0xfe, 0x5b, 0x4b, 0x98,
	0xf2, 0xde, 0x5e, 0xc2, 0x94, 0x6f, 0x4d, 0xe6, 0x2d, 0x63, 0xb9, 0xe7, 0xaf, 0xe8, 0x10, 0xec,
	0x76, 0x19, 0xf3, 0x39, 0xc1, 0xff, 0x07, 0x98, 0x32, 0xca, 0x4d, 0x2f, 0x70, 0x19, 0x84, 0xab,
	0x84, 0x4a, 0xdc, 0x79, 0x45, 0x50, 0xdf, 0x73, 0x29, 0x9e, 0x2b, 0xd0, 0x2a, 0xa1, 0x14, 0xbb,
	0x72, 0x03, 0xdb, 0x16, 0xfd, 0x8f, 0x5b, 0xde, 0x9b, 0xbd, 0xdd, 0x93, 0x9a, 0x05, 0x15, 0x30,
	0xcb, 0xb6, 0xee, 0x0d, 0x66, 0x79, 0x2e, 0xdc, 0x97, 0x3d, 0x10, 0x4c, 0x99, 0x41, 0x5e, 0x92,
	0x0f, 0xd0, 0x0e, 0xd8, 0x0e, 0x5c, 0x6a, 0x76, 0x71, 0x3b, 0xb0, 0xf3, 0xb1, 0x1f, 0x22, 0x00,
	0xca, 0x5d, 0x6c, 0xd8, 0xac, 0x0b, 0xeb, 0xf2, 0xf8, 0xf8, 0xde, 0x1d, 0x26, 0xf3, 0xfe, 0xee,
	0x0c, 0xc6, 0x28, 0x3c, 0x92, 0x11, 0x0c, 0x3b, 0x3e, 0x26, 0x06, 0x0b, 0x08, 0xe6, 0x26, 0xb6,
	0xa9, 0x15, 0x50, 0xd8, 0x40, 0x35, 0x00, 0xda, 0x01, 0xeb, 0x71, 0xb3, 0x67, 0xda, 0x18, 0x1e,
	0xcb, 0x48, 0x33, 0x20, 0x04, 0xbb, 0x8c, 0x7b, 0x9d, 0x0e, 0xc5, 0x0c, 0x9e, 0xa0, 0x0a, 0x28,
	0xd9, 0xc6, 0x0d, 0xfc, 0x0d, 0x55, 0x41, 0xc5, 0x36, 0x18, 0x76, 0xcd, 0x1e, 0xfc, 0x5d, 0xa6,
	0xba, 0x81, 0xd3, 0xc2, 0x04, 0x9e, 0x5e, 0x5c, 0x83, 0x2d, 0x12, 0x4e, 0xd3, 0xd9, 0xa4, 0x1f,
	0xba, 0x22, 0xbf, 0xe2, 0x87, 0x04, 0x53, 0x2f, 0x20, 0x26, 0xe6, 0xae, 0xe1, 0x2c, 0x9f, 0xf0,
	0x0a, 0x28, 0x99, 0x7e, 0x00, 0x0b, 0xd2, 0x41, 0x6d, 0x33, 0x2c, 0xb6, 0x5a, 0xf7, 0xd7, 0xc3,
	0x28, 0x9b, 0xdf, 0x30, 0xfd, 0xfb, 0x5d, 0xfc, 0x53, 0x44, 0xba, 0x7c, 0x1b, 0x7f, 0xe2, 0x9d,
	0x7c, 0x28, 0xe7, 0xaf, 0xce, 0x5f, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x63, 0x6b, 0xb7, 0x4b,
	0x55, 0x05, 0x00, 0x00,
}
