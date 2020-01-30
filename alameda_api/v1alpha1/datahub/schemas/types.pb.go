// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/schemas/types.proto

package schemas

import (
	fmt "fmt"
	common "github.com/containers-ai/api/common"
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
	Table_TABLE_RECOMMENDATION Table = 5
	Table_TABLE_RESOURCE       Table = 6
)

var Table_name = map[int32]string{
	0: "TABLE_UNDEFINED",
	1: "TABLE_APPLICATION",
	2: "TABLE_METRIC",
	3: "TABLE_PLANNING",
	4: "TABLE_PREDICTION",
	5: "TABLE_RECOMMENDATION",
	6: "TABLE_RESOURCE",
}

var Table_value = map[string]int32{
	"TABLE_UNDEFINED":      0,
	"TABLE_APPLICATION":    1,
	"TABLE_METRIC":         2,
	"TABLE_PLANNING":       3,
	"TABLE_PREDICTION":     4,
	"TABLE_RECOMMENDATION": 5,
	"TABLE_RESOURCE":       6,
}

func (x Table) String() string {
	return proto.EnumName(Table_name, int32(x))
}

func (Table) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_01194bf23370964f, []int{0}
}

type SchemaMeta struct {
	Table                Table    `protobuf:"varint,1,opt,name=table,proto3,enum=containersai.alameda.v1alpha1.datahub.schemas.Table" json:"table,omitempty"`
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

func (m *SchemaMeta) GetTable() Table {
	if m != nil {
		return m.Table
	}
	return Table_TABLE_UNDEFINED
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

type Measurement struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Columns              []*Column `protobuf:"bytes,2,rep,name=columns,proto3" json:"columns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
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

func (m *Measurement) GetColumns() []*Column {
	if m != nil {
		return m.Columns
	}
	return nil
}

type Column struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Required             bool              `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
	ColumnTypes          common.ColumnType `protobuf:"varint,3,opt,name=column_types,json=columnTypes,proto3,enum=containersai.common.ColumnType" json:"column_types,omitempty"`
	DataTypes            common.DataType   `protobuf:"varint,4,opt,name=data_types,json=dataTypes,proto3,enum=containersai.common.DataType" json:"data_types,omitempty"`
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

func (m *Column) GetColumnTypes() common.ColumnType {
	if m != nil {
		return m.ColumnTypes
	}
	return common.ColumnType_COLUMNTYPE_UDEFINED
}

func (m *Column) GetDataTypes() common.DataType {
	if m != nil {
		return m.DataTypes
	}
	return common.DataType_DATATYPE_UNDEFINED
}

func init() {
	proto.RegisterEnum("containersai.alameda.v1alpha1.datahub.schemas.Table", Table_name, Table_value)
	proto.RegisterType((*SchemaMeta)(nil), "containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta")
	proto.RegisterType((*Measurement)(nil), "containersai.alameda.v1alpha1.datahub.schemas.Measurement")
	proto.RegisterType((*Column)(nil), "containersai.alameda.v1alpha1.datahub.schemas.Column")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/schemas/types.proto", fileDescriptor_01194bf23370964f)
}

var fileDescriptor_01194bf23370964f = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x71, 0xfe, 0x91, 0x4c, 0xaa, 0x60, 0x96, 0x22, 0x59, 0x91, 0x10, 0x51, 0x4e, 0x11,
	0x52, 0x6d, 0x1a, 0xe0, 0xc6, 0xc5, 0xb1, 0x17, 0x64, 0x14, 0x3b, 0xd1, 0xd6, 0xbd, 0x70, 0x89,
	0x26, 0xce, 0xaa, 0xb1, 0x14, 0xff, 0xc1, 0x5e, 0x23, 0xe5, 0x0d, 0x78, 0x09, 0x5e, 0x84, 0xa7,
	0x43, 0xde, 0x75, 0x53, 0x90, 0x7a, 0x68, 0x6e, 0xb3, 0x9f, 0xfd, 0xfd, 0xe6, 0x9b, 0xd1, 0xc0,
	0x7b, 0x3c, 0x60, 0xc2, 0x77, 0xb8, 0xc1, 0x3c, 0xb6, 0x7e, 0x5e, 0xe3, 0x21, 0xdf, 0xe3, 0xb5,
	0xb5, 0x43, 0x81, 0xfb, 0x6a, 0x6b, 0x95, 0xd1, 0x9e, 0x27, 0x58, 0x5a, 0xe2, 0x98, 0xf3, 0xd2,
	0xcc, 0x8b, 0x4c, 0x64, 0xe4, 0x2a, 0xca, 0x52, 0x81, 0x71, 0xca, 0x8b, 0x12, 0x63, 0xb3, 0xb1,
	0x9b, 0xf7, 0x56, 0xb3, 0xb1, 0x9a, 0x8d, 0x75, 0x4c, 0xa2, 0x2c, 0x49, 0xb2, 0xf4, 0x5f, 0xc4,
	0xf4, 0x97, 0x06, 0x70, 0x23, 0xbf, 0xfb, 0x5c, 0x20, 0xf9, 0x06, 0x5d, 0x81, 0xdb, 0x03, 0x37,
	0xb4, 0x89, 0x36, 0x1b, 0xcd, 0x3f, 0x9a, 0x67, 0x75, 0x30, 0xc3, 0xda, 0xcb, 0x14, 0x82, 0x8c,
	0xa1, 0x1f, 0xa1, 0xe0, 0x77, 0x59, 0x71, 0x34, 0x5a, 0x13, 0x6d, 0x36, 0x60, 0xa7, 0x37, 0x21,
	0xd0, 0xa9, 0x53, 0x18, 0x6d, 0xa9, 0xcb, 0x7a, 0x5a, 0xc0, 0xd0, 0xe7, 0x58, 0x56, 0x05, 0x4f,
	0x78, 0x2a, 0xea, 0x5f, 0x52, 0x4c, 0x54, 0x92, 0x01, 0x93, 0x35, 0x59, 0xc1, 0xf3, 0x28, 0x3b,
	0x54, 0x49, 0x5a, 0x1a, 0xad, 0x49, 0x7b, 0x36, 0x9c, 0x7f, 0x3a, 0x33, 0xa0, 0x23, 0xdd, 0xec,
	0x9e, 0x32, 0xfd, 0xa3, 0x41, 0x4f, 0x69, 0x8f, 0xf6, 0x1b, 0x43, 0xbf, 0xe0, 0x3f, 0xaa, 0xb8,
	0xe0, 0x3b, 0x39, 0x42, 0x9f, 0x9d, 0xde, 0x64, 0x01, 0x17, 0x8a, 0xb2, 0x91, 0xfb, 0x94, 0xa3,
	0x8c, 0xe6, 0x6f, 0xff, 0x0f, 0xa4, 0x36, 0xde, 0xb4, 0x0d, 0x8f, 0x39, 0x67, 0xc3, 0xe8, 0x54,
	0x97, 0xe4, 0x33, 0x40, 0x9d, 0xb0, 0x21, 0x74, 0x24, 0xe1, 0xcd, 0xa3, 0x04, 0x17, 0x05, 0x4a,
	0xff, 0x60, 0xd7, 0x54, 0xe5, 0xbb, 0xdf, 0x1a, 0x74, 0xe5, 0xc6, 0xc9, 0x2b, 0x78, 0x11, 0xda,
	0x8b, 0x25, 0xdd, 0xdc, 0x06, 0x2e, 0xfd, 0xe2, 0x05, 0xd4, 0xd5, 0x9f, 0x91, 0xd7, 0xf0, 0x52,
	0x89, 0xf6, 0x7a, 0xbd, 0xf4, 0x1c, 0x3b, 0xf4, 0x56, 0x81, 0xae, 0x11, 0x1d, 0x2e, 0x94, 0xec,
	0xd3, 0x90, 0x79, 0x8e, 0xde, 0x22, 0x04, 0x46, 0x4a, 0x59, 0x2f, 0xed, 0x20, 0xf0, 0x82, 0xaf,
	0x7a, 0x9b, 0x5c, 0x82, 0xde, 0x68, 0x8c, 0xba, 0x9e, 0x23, 0xbd, 0x1d, 0x62, 0xc0, 0xa5, 0x52,
	0x19, 0x75, 0x56, 0xbe, 0x4f, 0x03, 0x57, 0x51, 0xbb, 0x0f, 0x0c, 0x46, 0x6f, 0x56, 0xb7, 0xcc,
	0xa1, 0x7a, 0x6f, 0xe1, 0x7c, 0xb7, 0xef, 0x62, 0x51, 0xaf, 0x3f, 0xca, 0x12, 0xeb, 0x61, 0xaa,
	0x2b, 0x8c, 0xad, 0xfa, 0xc6, 0x9f, 0x72, 0xef, 0xdb, 0x9e, 0xbc, 0xd3, 0x0f, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xb7, 0xcb, 0x64, 0x8f, 0x1e, 0x03, 0x00, 0x00,
}
