// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/schemas/types.proto

package schemas

import (
	fmt "fmt"
	common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
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

type Scope int32

const (
	Scope_SCOPE_UNDEFINED      Scope = 0
	Scope_SCOPE_APPLICATION    Scope = 1
	Scope_SCOPE_FEDEMETER      Scope = 2
	Scope_SCOPE_METRIC         Scope = 3
	Scope_SCOPE_PLANNING       Scope = 4
	Scope_SCOPE_PREDICTION     Scope = 5
	Scope_SCOPE_RECOMMENDATION Scope = 6
	Scope_SCOPE_RESOURCE       Scope = 7
	Scope_SCOPE_TARGET         Scope = 8
)

var Scope_name = map[int32]string{
	0: "SCOPE_UNDEFINED",
	1: "SCOPE_APPLICATION",
	2: "SCOPE_FEDEMETER",
	3: "SCOPE_METRIC",
	4: "SCOPE_PLANNING",
	5: "SCOPE_PREDICTION",
	6: "SCOPE_RECOMMENDATION",
	7: "SCOPE_RESOURCE",
	8: "SCOPE_TARGET",
}

var Scope_value = map[string]int32{
	"SCOPE_UNDEFINED":      0,
	"SCOPE_APPLICATION":    1,
	"SCOPE_FEDEMETER":      2,
	"SCOPE_METRIC":         3,
	"SCOPE_PLANNING":       4,
	"SCOPE_PREDICTION":     5,
	"SCOPE_RECOMMENDATION": 6,
	"SCOPE_RESOURCE":       7,
	"SCOPE_TARGET":         8,
}

func (x Scope) String() string {
	return proto.EnumName(Scope_name, int32(x))
}

func (Scope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_01194bf23370964f, []int{0}
}

//*
// Represents the private metadata of datahub schema.
type SchemaMeta struct {
	Scope                Scope    `protobuf:"varint,1,opt,name=scope,proto3,enum=containersai.alameda.v1alpha1.datahub.schemas.Scope" json:"scope,omitempty"`
	Category             string   `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SchemaMeta) Reset()         { *m = SchemaMeta{} }
func (m *SchemaMeta) String() string { return proto.CompactTextString(m) }
func (*SchemaMeta) ProtoMessage()    {}
func (*SchemaMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_01194bf23370964f, []int{0}
}

func (m *SchemaMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchemaMeta.Unmarshal(m, b)
}
func (m *SchemaMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchemaMeta.Marshal(b, m, deterministic)
}
func (m *SchemaMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchemaMeta.Merge(m, src)
}
func (m *SchemaMeta) XXX_Size() int {
	return xxx_messageInfo_SchemaMeta.Size(m)
}
func (m *SchemaMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_SchemaMeta.DiscardUnknown(m)
}

var xxx_messageInfo_SchemaMeta proto.InternalMessageInfo

func (m *SchemaMeta) GetScope() Scope {
	if m != nil {
		return m.Scope
	}
	return Scope_SCOPE_UNDEFINED
}

func (m *SchemaMeta) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *SchemaMeta) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

//*
// Represents the information of measurment which datahub will write data in InfluxDB.
type Measurement struct {
	Name                 string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MetricType           common.MetricType       `protobuf:"varint,2,opt,name=metric_type,json=metricType,proto3,enum=containersai.alameda.v1alpha1.datahub.common.MetricType" json:"metric_type,omitempty"`
	ResourceBoundary     common.ResourceBoundary `protobuf:"varint,3,opt,name=resource_boundary,json=resourceBoundary,proto3,enum=containersai.alameda.v1alpha1.datahub.common.ResourceBoundary" json:"resource_boundary,omitempty"`
	ResourceQuota        common.ResourceQuota    `protobuf:"varint,4,opt,name=resource_quota,json=resourceQuota,proto3,enum=containersai.alameda.v1alpha1.datahub.common.ResourceQuota" json:"resource_quota,omitempty"`
	IsTs                 bool                    `protobuf:"varint,5,opt,name=is_ts,json=isTs,proto3" json:"is_ts,omitempty"`
	Columns              []*Column               `protobuf:"bytes,6,rep,name=columns,proto3" json:"columns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Measurement) Reset()         { *m = Measurement{} }
func (m *Measurement) String() string { return proto.CompactTextString(m) }
func (*Measurement) ProtoMessage()    {}
func (*Measurement) Descriptor() ([]byte, []int) {
	return fileDescriptor_01194bf23370964f, []int{1}
}

func (m *Measurement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Measurement.Unmarshal(m, b)
}
func (m *Measurement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Measurement.Marshal(b, m, deterministic)
}
func (m *Measurement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Measurement.Merge(m, src)
}
func (m *Measurement) XXX_Size() int {
	return xxx_messageInfo_Measurement.Size(m)
}
func (m *Measurement) XXX_DiscardUnknown() {
	xxx_messageInfo_Measurement.DiscardUnknown(m)
}

var xxx_messageInfo_Measurement proto.InternalMessageInfo

func (m *Measurement) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Measurement) GetMetricType() common.MetricType {
	if m != nil {
		return m.MetricType
	}
	return common.MetricType_METRICS_TYPE_UNDEFINED
}

func (m *Measurement) GetResourceBoundary() common.ResourceBoundary {
	if m != nil {
		return m.ResourceBoundary
	}
	return common.ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED
}

func (m *Measurement) GetResourceQuota() common.ResourceQuota {
	if m != nil {
		return m.ResourceQuota
	}
	return common.ResourceQuota_RESOURCE_QUOTA_UNDEFINED
}

func (m *Measurement) GetIsTs() bool {
	if m != nil {
		return m.IsTs
	}
	return false
}

func (m *Measurement) GetColumns() []*Column {
	if m != nil {
		return m.Columns
	}
	return nil
}

//*
// Represents a data record.
type Column struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Required             bool              `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
	ColumnType           common.ColumnType `protobuf:"varint,3,opt,name=column_type,json=columnType,proto3,enum=containersai.alameda.v1alpha1.datahub.common.ColumnType" json:"column_type,omitempty"`
	DataType             common.DataType   `protobuf:"varint,4,opt,name=data_type,json=dataType,proto3,enum=containersai.alameda.v1alpha1.datahub.common.DataType" json:"data_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Column) Reset()         { *m = Column{} }
func (m *Column) String() string { return proto.CompactTextString(m) }
func (*Column) ProtoMessage()    {}
func (*Column) Descriptor() ([]byte, []int) {
	return fileDescriptor_01194bf23370964f, []int{2}
}

func (m *Column) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Column.Unmarshal(m, b)
}
func (m *Column) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Column.Marshal(b, m, deterministic)
}
func (m *Column) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Column.Merge(m, src)
}
func (m *Column) XXX_Size() int {
	return xxx_messageInfo_Column.Size(m)
}
func (m *Column) XXX_DiscardUnknown() {
	xxx_messageInfo_Column.DiscardUnknown(m)
}

var xxx_messageInfo_Column proto.InternalMessageInfo

func (m *Column) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Column) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

func (m *Column) GetColumnType() common.ColumnType {
	if m != nil {
		return m.ColumnType
	}
	return common.ColumnType_COLUMNTYPE_UDEFINED
}

func (m *Column) GetDataType() common.DataType {
	if m != nil {
		return m.DataType
	}
	return common.DataType_DATATYPE_UNDEFINED
}

func init() {
	proto.RegisterEnum("containersai.alameda.v1alpha1.datahub.schemas.Scope", Scope_name, Scope_value)
	proto.RegisterType((*SchemaMeta)(nil), "containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta")
	proto.RegisterType((*Measurement)(nil), "containersai.alameda.v1alpha1.datahub.schemas.Measurement")
	proto.RegisterType((*Column)(nil), "containersai.alameda.v1alpha1.datahub.schemas.Column")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/schemas/types.proto", fileDescriptor_01194bf23370964f)
}

var fileDescriptor_01194bf23370964f = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdb, 0x6a, 0xdb, 0x3e,
	0x1c, 0xfe, 0xbb, 0x39, 0xd4, 0x55, 0xff, 0xcb, 0x5c, 0xb5, 0x03, 0x93, 0xab, 0xd0, 0xab, 0x30,
	0xa8, 0xbd, 0x64, 0x07, 0x06, 0x83, 0x41, 0x6a, 0xab, 0xc5, 0xa3, 0x76, 0x32, 0xc5, 0xbd, 0xd8,
	0x6e, 0x8c, 0xe2, 0x88, 0xc6, 0x2c, 0x3e, 0x54, 0x92, 0x07, 0x79, 0x83, 0xbd, 0xcf, 0x5e, 0x61,
	0x0f, 0xb3, 0xc7, 0x18, 0x96, 0x5c, 0xa7, 0x8c, 0x31, 0x92, 0xdd, 0xfd, 0xf4, 0xa1, 0xef, 0x20,
	0x7d, 0x42, 0xe0, 0x05, 0x59, 0x93, 0x94, 0x2e, 0x49, 0x44, 0x8a, 0xc4, 0xfe, 0x3a, 0x22, 0xeb,
	0x62, 0x45, 0x46, 0xf6, 0x92, 0x08, 0xb2, 0x2a, 0x17, 0x36, 0x8f, 0x57, 0x34, 0x25, 0xdc, 0x16,
	0x9b, 0x82, 0x72, 0xab, 0x60, 0xb9, 0xc8, 0xe1, 0x45, 0x9c, 0x67, 0x82, 0x24, 0x19, 0x65, 0x9c,
	0x24, 0x56, 0x4d, 0xb7, 0x1e, 0xa8, 0x56, 0x4d, 0xb5, 0x6a, 0x6a, 0x7f, 0xf4, 0x57, 0x83, 0x38,
	0x4f, 0xd3, 0x3c, 0xb3, 0x53, 0x2a, 0x58, 0x12, 0xd7, 0x0e, 0x7d, 0x7b, 0x17, 0xca, 0xa3, 0x48,
	0xe7, 0xdf, 0x34, 0x00, 0xe6, 0xd2, 0xcf, 0xa7, 0x82, 0xc0, 0x0f, 0xa0, 0xc3, 0xe3, 0xbc, 0xa0,
	0xa6, 0x36, 0xd0, 0x86, 0xbd, 0xf1, 0x2b, 0x6b, 0xaf, 0xc4, 0xd6, 0xbc, 0xe2, 0x62, 0x25, 0x01,
	0xfb, 0x40, 0x8f, 0x89, 0xa0, 0x77, 0x39, 0xdb, 0x98, 0x07, 0x03, 0x6d, 0x78, 0x84, 0x9b, 0x35,
	0x84, 0xa0, 0x5d, 0xa5, 0x30, 0x5b, 0x12, 0x97, 0xf3, 0xf9, 0xf7, 0x16, 0x38, 0xf6, 0x29, 0xe1,
	0x25, 0xa3, 0x29, 0xcd, 0x44, 0xb5, 0x27, 0x23, 0xa9, 0x8a, 0x72, 0x84, 0xe5, 0x0c, 0x3f, 0x81,
	0x63, 0x75, 0xe0, 0x48, 0xd2, 0x0f, 0x64, 0xca, 0xb7, 0x3b, 0xa6, 0x54, 0xc7, 0xb7, 0x7c, 0x29,
	0x10, 0x6e, 0x0a, 0x8a, 0x41, 0xda, 0xcc, 0xf0, 0x0b, 0x38, 0x61, 0x94, 0xe7, 0x25, 0x8b, 0x69,
	0xb4, 0xc8, 0xcb, 0x6c, 0x49, 0xd8, 0x46, 0xe6, 0xeb, 0x8d, 0xdf, 0xef, 0x67, 0x80, 0x6b, 0x99,
	0xcb, 0x5a, 0x05, 0x1b, 0xec, 0x37, 0x04, 0x2e, 0x40, 0xaf, 0x31, 0xbb, 0x2f, 0x73, 0x41, 0xcc,
	0xb6, 0x74, 0x7a, 0xf7, 0x6f, 0x4e, 0x1f, 0x2b, 0x09, 0xfc, 0x84, 0x3d, 0x5e, 0xc2, 0x53, 0xd0,
	0x49, 0x78, 0x24, 0xb8, 0xd9, 0x19, 0x68, 0x43, 0x1d, 0xb7, 0x13, 0x1e, 0x72, 0x38, 0x05, 0x87,
	0x71, 0xbe, 0x2e, 0xd3, 0x8c, 0x9b, 0xdd, 0x41, 0x6b, 0x78, 0x3c, 0x7e, 0xbd, 0x67, 0xc5, 0x8e,
	0x64, 0xe3, 0x07, 0x95, 0xf3, 0x9f, 0x1a, 0xe8, 0x2a, 0xec, 0x8f, 0x85, 0xf5, 0x81, 0xce, 0xe8,
	0x7d, 0x99, 0x30, 0xba, 0x94, 0x6d, 0xe9, 0xb8, 0x59, 0x57, 0x65, 0x2a, 0x95, 0xa8, 0x79, 0x0b,
	0x7b, 0x97, 0xa9, 0xac, 0x55, 0x99, 0x71, 0x33, 0xc3, 0x39, 0x38, 0xaa, 0x36, 0x2a, 0x61, 0x75,
	0xb5, 0x6f, 0xf6, 0x13, 0x76, 0x89, 0x20, 0x52, 0x56, 0x5f, 0xd6, 0xd3, 0xf3, 0x1f, 0x1a, 0xe8,
	0xc8, 0x17, 0x0e, 0x4f, 0xc1, 0xd3, 0xb9, 0x33, 0x9d, 0xa1, 0xe8, 0x36, 0x70, 0xd1, 0x95, 0x17,
	0x20, 0xd7, 0xf8, 0x0f, 0x3e, 0x03, 0x27, 0x0a, 0x9c, 0xcc, 0x66, 0x37, 0x9e, 0x33, 0x09, 0xbd,
	0x69, 0x60, 0x68, 0xdb, 0xbd, 0x57, 0xc8, 0x45, 0x3e, 0x0a, 0x11, 0x36, 0x0e, 0xa0, 0x01, 0xfe,
	0x57, 0xa0, 0x8f, 0x42, 0xec, 0x39, 0x46, 0x0b, 0x42, 0xd0, 0x53, 0xc8, 0xec, 0x66, 0x12, 0x04,
	0x5e, 0x70, 0x6d, 0xb4, 0xe1, 0x19, 0x30, 0x6a, 0x0c, 0x23, 0xd7, 0x73, 0xa4, 0x60, 0x07, 0x9a,
	0xe0, 0x4c, 0xa1, 0x18, 0x39, 0x53, 0xdf, 0x47, 0x81, 0xab, 0xac, 0xba, 0x5b, 0x0d, 0x8c, 0xe6,
	0xd3, 0x5b, 0xec, 0x20, 0xe3, 0x70, 0xeb, 0x14, 0x4e, 0xf0, 0x35, 0x0a, 0x0d, 0xfd, 0xd2, 0xf9,
	0x3c, 0xb9, 0x4b, 0x44, 0x7d, 0x54, 0x7b, 0x7b, 0x29, 0x17, 0x24, 0xb1, 0xab, 0x6f, 0x63, 0x97,
	0x6f, 0x6d, 0xd1, 0x95, 0xdf, 0xc7, 0xcb, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xf7, 0x72,
	0x4a, 0x05, 0x05, 0x00, 0x00,
}
