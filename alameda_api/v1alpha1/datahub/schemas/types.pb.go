// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/schemas/types.proto

package schemas

import (
	fmt "fmt"
	common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	common1 "github.com/containers-ai/api/common"
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

type Table int32

const (
	Table_TABLE_UNDEFINED      Table = 0
	Table_TABLE_APPLICATION    Table = 1
	Table_TABLE_METRIC         Table = 2
	Table_TABLE_PLANNING       Table = 3
	Table_TABLE_PREDICTION     Table = 4
	Table_TABLE_RESOURCE       Table = 5
	Table_TABLE_RECOMMENDATION Table = 6
)

var Table_name = map[int32]string{
	0: "TABLE_UNDEFINED",
	1: "TABLE_APPLICATION",
	2: "TABLE_METRIC",
	3: "TABLE_PLANNING",
	4: "TABLE_PREDICTION",
	5: "TABLE_RESOURCE",
	6: "TABLE_RECOMMENDATION",
}

var Table_value = map[string]int32{
	"TABLE_UNDEFINED":      0,
	"TABLE_APPLICATION":    1,
	"TABLE_METRIC":         2,
	"TABLE_PLANNING":       3,
	"TABLE_PREDICTION":     4,
	"TABLE_RESOURCE":       5,
	"TABLE_RECOMMENDATION": 6,
}

func (x Table) String() string {
	return proto.EnumName(Table_name, int32(x))
}

func (Table) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_01194bf23370964f, []int{0}
}

type Scope int32

const (
	Scope_SCOPE_UNDEFINED   Scope = 0
	Scope_SCOPE_RESOURCE    Scope = 1
	Scope_SCOPE_APPLICATION Scope = 2
)

var Scope_name = map[int32]string{
	0: "SCOPE_UNDEFINED",
	1: "SCOPE_RESOURCE",
	2: "SCOPE_APPLICATION",
}

var Scope_value = map[string]int32{
	"SCOPE_UNDEFINED":   0,
	"SCOPE_RESOURCE":    1,
	"SCOPE_APPLICATION": 2,
}

func (x Scope) String() string {
	return proto.EnumName(Scope_name, int32(x))
}

func (Scope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_01194bf23370964f, []int{1}
}

type SchemaMeta struct {
	Table                Table    `protobuf:"varint,1,opt,name=table,proto3,enum=containersai.alameda.v1alpha1.datahub.schemas.Table" json:"table,omitempty"`
	Scope                Scope    `protobuf:"varint,2,opt,name=scope,proto3,enum=containersai.alameda.v1alpha1.datahub.schemas.Scope" json:"scope,omitempty"`
	Category             string   `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
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

func (m *SchemaMeta) GetTable() Table {
	if m != nil {
		return m.Table
	}
	return Table_TABLE_UNDEFINED
}

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

type Mesaurement struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MetricType           common.MetricType `protobuf:"varint,2,opt,name=metric_type,json=metricType,proto3,enum=containersai.alameda.v1alpha1.datahub.common.MetricType" json:"metric_type,omitempty"`
	Columns              []*Column         `protobuf:"bytes,3,rep,name=columns,proto3" json:"columns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Mesaurement) Reset()         { *m = Mesaurement{} }
func (m *Mesaurement) String() string { return proto.CompactTextString(m) }
func (*Mesaurement) ProtoMessage()    {}
func (*Mesaurement) Descriptor() ([]byte, []int) {
	return fileDescriptor_01194bf23370964f, []int{1}
}

func (m *Mesaurement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mesaurement.Unmarshal(m, b)
}
func (m *Mesaurement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mesaurement.Marshal(b, m, deterministic)
}
func (m *Mesaurement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mesaurement.Merge(m, src)
}
func (m *Mesaurement) XXX_Size() int {
	return xxx_messageInfo_Mesaurement.Size(m)
}
func (m *Mesaurement) XXX_DiscardUnknown() {
	xxx_messageInfo_Mesaurement.DiscardUnknown(m)
}

var xxx_messageInfo_Mesaurement proto.InternalMessageInfo

func (m *Mesaurement) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Mesaurement) GetMetricType() common.MetricType {
	if m != nil {
		return m.MetricType
	}
	return common.MetricType_METRICS_TYPE_UNDEFINED
}

func (m *Mesaurement) GetColumns() []*Column {
	if m != nil {
		return m.Columns
	}
	return nil
}

type Column struct {
	Name                 string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Required             bool               `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
	ColumnTypes          common1.ColumnType `protobuf:"varint,3,opt,name=column_types,json=columnTypes,proto3,enum=containersai.common.ColumnType" json:"column_types,omitempty"`
	DataTypes            common1.DataType   `protobuf:"varint,4,opt,name=data_types,json=dataTypes,proto3,enum=containersai.common.DataType" json:"data_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
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

func (m *Column) GetColumnTypes() common1.ColumnType {
	if m != nil {
		return m.ColumnTypes
	}
	return common1.ColumnType_COLUMNTYPE_UDEFINED
}

func (m *Column) GetDataTypes() common1.DataType {
	if m != nil {
		return m.DataTypes
	}
	return common1.DataType_DATATYPE_UNDEFINED
}

func init() {
	proto.RegisterEnum("containersai.alameda.v1alpha1.datahub.schemas.Table", Table_name, Table_value)
	proto.RegisterEnum("containersai.alameda.v1alpha1.datahub.schemas.Scope", Scope_name, Scope_value)
	proto.RegisterType((*SchemaMeta)(nil), "containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta")
	proto.RegisterType((*Mesaurement)(nil), "containersai.alameda.v1alpha1.datahub.schemas.Mesaurement")
	proto.RegisterType((*Column)(nil), "containersai.alameda.v1alpha1.datahub.schemas.Column")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/schemas/types.proto", fileDescriptor_01194bf23370964f)
}

var fileDescriptor_01194bf23370964f = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6a, 0xdb, 0x4c,
	0x14, 0xfd, 0xe4, 0xbf, 0xcf, 0xbe, 0x0e, 0xae, 0x3a, 0x4d, 0x41, 0x18, 0x4a, 0x8d, 0x57, 0x26,
	0x10, 0xa9, 0x76, 0x5b, 0xe8, 0xa2, 0x1b, 0x5b, 0x56, 0x83, 0x8a, 0x25, 0x99, 0xb1, 0xb2, 0x68,
	0x37, 0x66, 0x2c, 0x0f, 0xb1, 0xc0, 0xfa, 0xa9, 0x66, 0x5c, 0xf0, 0xc3, 0xf4, 0x45, 0xfa, 0x0e,
	0xdd, 0xf5, 0x81, 0x8a, 0x66, 0xc6, 0x4a, 0x02, 0xa1, 0x24, 0xdd, 0xdd, 0x39, 0xcc, 0x39, 0xf7,
	0x9c, 0xcb, 0x9d, 0x81, 0x37, 0x64, 0x4f, 0x12, 0xba, 0x25, 0x6b, 0x92, 0xc7, 0xd6, 0xf7, 0x31,
	0xd9, 0xe7, 0x3b, 0x32, 0xb6, 0xb6, 0x84, 0x93, 0xdd, 0x61, 0x63, 0xb1, 0x68, 0x47, 0x13, 0xc2,
	0x2c, 0x7e, 0xcc, 0x29, 0x33, 0xf3, 0x22, 0xe3, 0x19, 0xba, 0x8c, 0xb2, 0x94, 0x93, 0x38, 0xa5,
	0x05, 0x23, 0xb1, 0xa9, 0xe8, 0xe6, 0x89, 0x6a, 0x2a, 0xaa, 0xa9, 0xa8, 0xfd, 0xf1, 0x5f, 0x1b,
	0x44, 0x59, 0x92, 0x64, 0xa9, 0x95, 0x50, 0x5e, 0xc4, 0x91, 0xea, 0xd0, 0x47, 0x0a, 0xbd, 0xd3,
	0x75, 0xf8, 0x5b, 0x03, 0x58, 0x09, 0x49, 0x8f, 0x72, 0x82, 0x3e, 0x43, 0x93, 0x93, 0xcd, 0x9e,
	0x1a, 0xda, 0x40, 0x1b, 0xf5, 0x26, 0xef, 0xcc, 0x27, 0x99, 0x32, 0xc3, 0x92, 0x8b, 0xa5, 0x44,
	0xa9, 0xc5, 0xa2, 0x2c, 0xa7, 0x46, 0xed, 0x9f, 0xb4, 0x56, 0x25, 0x17, 0x4b, 0x09, 0xd4, 0x87,
	0x76, 0x44, 0x38, 0xbd, 0xc9, 0x8a, 0xa3, 0x51, 0x1f, 0x68, 0xa3, 0x0e, 0xae, 0xce, 0x08, 0x41,
	0xa3, 0x4c, 0x64, 0x34, 0x04, 0x2e, 0xea, 0xe1, 0x2f, 0x0d, 0xba, 0x1e, 0x65, 0xe4, 0x50, 0xd0,
	0x84, 0xa6, 0xbc, 0xbc, 0x93, 0x92, 0x44, 0xc6, 0xea, 0x60, 0x51, 0xa3, 0x2f, 0xd0, 0x95, 0xf3,
	0x59, 0x0b, 0xba, 0x74, 0xf9, 0xe1, 0x91, 0x2e, 0xe5, 0x28, 0x4d, 0x4f, 0x08, 0x84, 0xc7, 0x9c,
	0x62, 0x48, 0xaa, 0x1a, 0x05, 0xf0, 0x7f, 0x94, 0xed, 0x0f, 0x49, 0xca, 0x8c, 0xfa, 0xa0, 0x3e,
	0xea, 0x4e, 0xde, 0x3f, 0x31, 0xbc, 0x2d, 0xd8, 0xf8, 0xa4, 0x32, 0xfc, 0xa9, 0x41, 0x4b, 0x62,
	0x0f, 0x46, 0xe9, 0x43, 0xbb, 0xa0, 0xdf, 0x0e, 0x71, 0x41, 0xb7, 0x22, 0x47, 0x1b, 0x57, 0x67,
	0x34, 0x83, 0x33, 0xa9, 0x22, 0x62, 0x32, 0x31, 0xbe, 0xde, 0xe4, 0xf5, 0x7d, 0x43, 0x2a, 0x8e,
	0x6c, 0x21, 0xe2, 0x74, 0xa3, 0xaa, 0x66, 0xe8, 0x23, 0x40, 0xe9, 0x50, 0x29, 0x34, 0x84, 0xc2,
	0xab, 0x07, 0x15, 0xe6, 0x84, 0x13, 0xc1, 0xef, 0x6c, 0x55, 0xc5, 0x2e, 0x7e, 0x68, 0xd0, 0x14,
	0x9b, 0x81, 0x5e, 0xc0, 0xb3, 0x70, 0x3a, 0x5b, 0x38, 0xeb, 0x6b, 0x7f, 0xee, 0x7c, 0x72, 0x7d,
	0x67, 0xae, 0xff, 0x87, 0x5e, 0xc2, 0x73, 0x09, 0x4e, 0x97, 0xcb, 0x85, 0x6b, 0x4f, 0x43, 0x37,
	0xf0, 0x75, 0x0d, 0xe9, 0x70, 0x26, 0x61, 0xcf, 0x09, 0xb1, 0x6b, 0xeb, 0x35, 0x84, 0xa0, 0x27,
	0x91, 0xe5, 0x62, 0xea, 0xfb, 0xae, 0x7f, 0xa5, 0xd7, 0xd1, 0x39, 0xe8, 0x0a, 0xc3, 0xce, 0xdc,
	0xb5, 0x05, 0xb7, 0x71, 0x7b, 0x13, 0x3b, 0xab, 0xe0, 0x1a, 0xdb, 0x8e, 0xde, 0x44, 0x06, 0x9c,
	0x9f, 0x30, 0x3b, 0xf0, 0x3c, 0xc7, 0x9f, 0xcb, 0x4e, 0xad, 0x8b, 0x2b, 0x68, 0x8a, 0x65, 0x2b,
	0xed, 0xad, 0xec, 0x60, 0x79, 0xdf, 0x1e, 0x82, 0x9e, 0x04, 0x2b, 0x2d, 0xad, 0xb4, 0x2c, 0xb1,
	0xbb, 0x96, 0x6b, 0x33, 0xfb, 0xeb, 0xf4, 0x26, 0xe6, 0x6a, 0x3d, 0xac, 0xdb, 0xf1, 0x5c, 0x92,
	0xd8, 0x2a, 0x9f, 0xe9, 0x63, 0xfe, 0x84, 0x4d, 0x4b, 0x3c, 0xcc, 0xb7, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xcd, 0x69, 0x6a, 0x34, 0x42, 0x04, 0x00, 0x00,
}
