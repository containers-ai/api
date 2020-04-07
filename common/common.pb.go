// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type DatabaseType int32

const (
	DatabaseType_UNDEFINED  DatabaseType = 0
	DatabaseType_INFLUXDB   DatabaseType = 1
	DatabaseType_PROMETHEUS DatabaseType = 2
)

var DatabaseType_name = map[int32]string{
	0: "UNDEFINED",
	1: "INFLUXDB",
	2: "PROMETHEUS",
}

var DatabaseType_value = map[string]int32{
	"UNDEFINED":  0,
	"INFLUXDB":   1,
	"PROMETHEUS": 2,
}

func (x DatabaseType) String() string {
	return proto.EnumName(DatabaseType_name, int32(x))
}

func (DatabaseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

type TimeRange_AggregateFunction int32

const (
	TimeRange_NONE TimeRange_AggregateFunction = 0
	TimeRange_MAX  TimeRange_AggregateFunction = 1
	TimeRange_AVG  TimeRange_AggregateFunction = 2
)

var TimeRange_AggregateFunction_name = map[int32]string{
	0: "NONE",
	1: "MAX",
	2: "AVG",
}

var TimeRange_AggregateFunction_value = map[string]int32{
	"NONE": 0,
	"MAX":  1,
	"AVG":  2,
}

func (x TimeRange_AggregateFunction) String() string {
	return proto.EnumName(TimeRange_AggregateFunction_name, int32(x))
}

func (TimeRange_AggregateFunction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0, 0}
}

type QueryCondition_Order int32

const (
	QueryCondition_NONE QueryCondition_Order = 0
	QueryCondition_ASC  QueryCondition_Order = 1
	QueryCondition_DESC QueryCondition_Order = 2
)

var QueryCondition_Order_name = map[int32]string{
	0: "NONE",
	1: "ASC",
	2: "DESC",
}

var QueryCondition_Order_value = map[string]int32{
	"NONE": 0,
	"ASC":  1,
	"DESC": 2,
}

func (x QueryCondition_Order) String() string {
	return proto.EnumName(QueryCondition_Order_name, int32(x))
}

func (QueryCondition_Order) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1, 0}
}

// Represents a time range definition
type TimeRange struct {
	StartTime            *timestamp.Timestamp        `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp        `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Timeout              *timestamp.Timestamp        `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Step                 *duration.Duration          `protobuf:"bytes,4,opt,name=step,proto3" json:"step,omitempty"`
	AggregateFunction    TimeRange_AggregateFunction `protobuf:"varint,5,opt,name=aggregate_function,json=aggregateFunction,proto3,enum=containersai.common.TimeRange_AggregateFunction" json:"aggregate_function,omitempty"`
	ApplyTime            *timestamp.Timestamp        `protobuf:"bytes,6,opt,name=apply_time,json=applyTime,proto3" json:"apply_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TimeRange) Reset()         { *m = TimeRange{} }
func (m *TimeRange) String() string { return proto.CompactTextString(m) }
func (*TimeRange) ProtoMessage()    {}
func (*TimeRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *TimeRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRange.Unmarshal(m, b)
}
func (m *TimeRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRange.Marshal(b, m, deterministic)
}
func (m *TimeRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRange.Merge(m, src)
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

func (m *TimeRange) GetTimeout() *timestamp.Timestamp {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *TimeRange) GetStep() *duration.Duration {
	if m != nil {
		return m.Step
	}
	return nil
}

func (m *TimeRange) GetAggregateFunction() TimeRange_AggregateFunction {
	if m != nil {
		return m.AggregateFunction
	}
	return TimeRange_NONE
}

func (m *TimeRange) GetApplyTime() *timestamp.Timestamp {
	if m != nil {
		return m.ApplyTime
	}
	return nil
}

type QueryCondition struct {
	TimeRange            *TimeRange           `protobuf:"bytes,1,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	Order                QueryCondition_Order `protobuf:"varint,2,opt,name=order,proto3,enum=containersai.common.QueryCondition_Order" json:"order,omitempty"`
	WhereClause          string               `protobuf:"bytes,3,opt,name=where_clause,json=whereClause,proto3" json:"where_clause,omitempty"`
	Selects              []string             `protobuf:"bytes,4,rep,name=selects,proto3" json:"selects,omitempty"`
	Groups               []string             `protobuf:"bytes,5,rep,name=groups,proto3" json:"groups,omitempty"`
	Limit                uint64               `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *QueryCondition) Reset()         { *m = QueryCondition{} }
func (m *QueryCondition) String() string { return proto.CompactTextString(m) }
func (*QueryCondition) ProtoMessage()    {}
func (*QueryCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

func (m *QueryCondition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryCondition.Unmarshal(m, b)
}
func (m *QueryCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryCondition.Marshal(b, m, deterministic)
}
func (m *QueryCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCondition.Merge(m, src)
}
func (m *QueryCondition) XXX_Size() int {
	return xxx_messageInfo_QueryCondition.Size(m)
}
func (m *QueryCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCondition.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCondition proto.InternalMessageInfo

func (m *QueryCondition) GetTimeRange() *TimeRange {
	if m != nil {
		return m.TimeRange
	}
	return nil
}

func (m *QueryCondition) GetOrder() QueryCondition_Order {
	if m != nil {
		return m.Order
	}
	return QueryCondition_NONE
}

func (m *QueryCondition) GetWhereClause() string {
	if m != nil {
		return m.WhereClause
	}
	return ""
}

func (m *QueryCondition) GetSelects() []string {
	if m != nil {
		return m.Selects
	}
	return nil
}

func (m *QueryCondition) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *QueryCondition) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type Query struct {
	Database             string          `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Table                string          `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	Expression           string          `protobuf:"bytes,3,opt,name=expression,proto3" json:"expression,omitempty"`
	Condition            *QueryCondition `protobuf:"bytes,4,opt,name=condition,proto3" json:"condition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{2}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *Query) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *Query) GetExpression() string {
	if m != nil {
		return m.Expression
	}
	return ""
}

func (m *Query) GetCondition() *QueryCondition {
	if m != nil {
		return m.Condition
	}
	return nil
}

type Row struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Values               []string             `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Row) Reset()         { *m = Row{} }
func (m *Row) String() string { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()    {}
func (*Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{3}
}

func (m *Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Row.Unmarshal(m, b)
}
func (m *Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Row.Marshal(b, m, deterministic)
}
func (m *Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Row.Merge(m, src)
}
func (m *Row) XXX_Size() int {
	return xxx_messageInfo_Row.Size(m)
}
func (m *Row) XXX_DiscardUnknown() {
	xxx_messageInfo_Row.DiscardUnknown(m)
}

var xxx_messageInfo_Row proto.InternalMessageInfo

func (m *Row) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Row) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type Group struct {
	Rows                 []*Row   `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{4}
}

func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetRows() []*Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

type ReadRawdata struct {
	Query                *Query   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Columns              []string `protobuf:"bytes,2,rep,name=columns,proto3" json:"columns,omitempty"`
	Groups               []*Group `protobuf:"bytes,3,rep,name=groups,proto3" json:"groups,omitempty"`
	Rawdata              string   `protobuf:"bytes,4,opt,name=rawdata,proto3" json:"rawdata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRawdata) Reset()         { *m = ReadRawdata{} }
func (m *ReadRawdata) String() string { return proto.CompactTextString(m) }
func (*ReadRawdata) ProtoMessage()    {}
func (*ReadRawdata) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{5}
}

func (m *ReadRawdata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRawdata.Unmarshal(m, b)
}
func (m *ReadRawdata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRawdata.Marshal(b, m, deterministic)
}
func (m *ReadRawdata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRawdata.Merge(m, src)
}
func (m *ReadRawdata) XXX_Size() int {
	return xxx_messageInfo_ReadRawdata.Size(m)
}
func (m *ReadRawdata) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRawdata.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRawdata proto.InternalMessageInfo

func (m *ReadRawdata) GetQuery() *Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *ReadRawdata) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *ReadRawdata) GetGroups() []*Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *ReadRawdata) GetRawdata() string {
	if m != nil {
		return m.Rawdata
	}
	return ""
}

type WriteRawdata struct {
	Database             string       `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Table                string       `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	Columns              []string     `protobuf:"bytes,3,rep,name=columns,proto3" json:"columns,omitempty"`
	Rows                 []*Row       `protobuf:"bytes,4,rep,name=rows,proto3" json:"rows,omitempty"`
	ColumnTypes          []ColumnType `protobuf:"varint,5,rep,packed,name=column_types,json=columnTypes,proto3,enum=containersai.common.ColumnType" json:"column_types,omitempty"`
	DataTypes            []DataType   `protobuf:"varint,6,rep,packed,name=data_types,json=dataTypes,proto3,enum=containersai.common.DataType" json:"data_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *WriteRawdata) Reset()         { *m = WriteRawdata{} }
func (m *WriteRawdata) String() string { return proto.CompactTextString(m) }
func (*WriteRawdata) ProtoMessage()    {}
func (*WriteRawdata) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{6}
}

func (m *WriteRawdata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRawdata.Unmarshal(m, b)
}
func (m *WriteRawdata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRawdata.Marshal(b, m, deterministic)
}
func (m *WriteRawdata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRawdata.Merge(m, src)
}
func (m *WriteRawdata) XXX_Size() int {
	return xxx_messageInfo_WriteRawdata.Size(m)
}
func (m *WriteRawdata) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRawdata.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRawdata proto.InternalMessageInfo

func (m *WriteRawdata) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *WriteRawdata) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *WriteRawdata) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *WriteRawdata) GetRows() []*Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

func (m *WriteRawdata) GetColumnTypes() []ColumnType {
	if m != nil {
		return m.ColumnTypes
	}
	return nil
}

func (m *WriteRawdata) GetDataTypes() []DataType {
	if m != nil {
		return m.DataTypes
	}
	return nil
}

func init() {
	proto.RegisterEnum("containersai.common.DatabaseType", DatabaseType_name, DatabaseType_value)
	proto.RegisterEnum("containersai.common.TimeRange_AggregateFunction", TimeRange_AggregateFunction_name, TimeRange_AggregateFunction_value)
	proto.RegisterEnum("containersai.common.QueryCondition_Order", QueryCondition_Order_name, QueryCondition_Order_value)
	proto.RegisterType((*TimeRange)(nil), "containersai.common.TimeRange")
	proto.RegisterType((*QueryCondition)(nil), "containersai.common.QueryCondition")
	proto.RegisterType((*Query)(nil), "containersai.common.Query")
	proto.RegisterType((*Row)(nil), "containersai.common.Row")
	proto.RegisterType((*Group)(nil), "containersai.common.Group")
	proto.RegisterType((*ReadRawdata)(nil), "containersai.common.ReadRawdata")
	proto.RegisterType((*WriteRawdata)(nil), "containersai.common.WriteRawdata")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6) }

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 758 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdf, 0x6f, 0xeb, 0x34,
	0x18, 0xbd, 0x69, 0x92, 0xb6, 0xf9, 0x5a, 0xaa, 0x5c, 0x5f, 0x84, 0x42, 0x25, 0x76, 0x4b, 0x2e,
	0x48, 0x05, 0xb1, 0x74, 0x2a, 0xec, 0x01, 0x01, 0x42, 0x5d, 0xdb, 0x8d, 0x49, 0xac, 0x03, 0xaf,
	0x83, 0x89, 0x97, 0xca, 0x4d, 0xbc, 0x2c, 0x52, 0x9a, 0x84, 0xd8, 0xa1, 0xf4, 0x5f, 0xe0, 0x6f,
	0xe0, 0x85, 0x27, 0x1e, 0xf9, 0x17, 0x91, 0xed, 0xa4, 0xfb, 0x55, 0xd8, 0xee, 0x53, 0x7d, 0xec,
	0x73, 0xbe, 0xef, 0xf8, 0x7c, 0x6e, 0x0b, 0xaf, 0xfc, 0x74, 0xb5, 0x4a, 0x93, 0x81, 0xfa, 0xf0,
	0xb2, 0x3c, 0xe5, 0x29, 0x7a, 0xe5, 0xa7, 0x09, 0x27, 0x51, 0x42, 0x73, 0x46, 0x22, 0x4f, 0x1d,
	0x75, 0x51, 0xc9, 0xe4, 0x9b, 0x8c, 0x32, 0x45, 0xec, 0xee, 0x85, 0x69, 0x1a, 0xc6, 0x74, 0x20,
	0xd1, 0xb2, 0xb8, 0x1e, 0x04, 0x45, 0x4e, 0x78, 0x54, 0x15, 0xea, 0xbe, 0x7e, 0x78, 0xce, 0xa3,
	0x15, 0x65, 0x9c, 0xac, 0x32, 0x45, 0x70, 0xff, 0xd1, 0xc1, 0x9a, 0x47, 0x2b, 0x8a, 0x49, 0x12,
	0x52, 0xf4, 0x25, 0x00, 0xe3, 0x24, 0xe7, 0x0b, 0x41, 0x73, 0xb4, 0x9e, 0xd6, 0x6f, 0x0d, 0xbb,
	0x9e, 0xaa, 0xe1, 0x55, 0x35, 0xbc, 0x79, 0x55, 0x03, 0x5b, 0x92, 0x2d, 0x30, 0x3a, 0x84, 0x26,
	0x4d, 0x02, 0x25, 0xac, 0x3d, 0x29, 0x6c, 0xd0, 0x24, 0x90, 0xb2, 0x2f, 0xa0, 0x21, 0x24, 0x69,
	0xc1, 0x1d, 0xfd, 0x69, 0x55, 0x49, 0x45, 0xfb, 0x60, 0x30, 0x4e, 0x33, 0xc7, 0x90, 0x92, 0xf7,
	0x1f, 0x49, 0x26, 0x65, 0x0a, 0x58, 0xd2, 0xd0, 0x02, 0x10, 0x09, 0xc3, 0x9c, 0x86, 0x84, 0xd3,
	0xc5, 0x75, 0x91, 0xf8, 0xe2, 0xcc, 0x31, 0x7b, 0x5a, 0xbf, 0x33, 0x3c, 0xf0, 0x76, 0x64, 0xed,
	0x6d, 0x23, 0xf1, 0x46, 0x95, 0xf0, 0xb8, 0xd4, 0xe1, 0x97, 0xe4, 0xe1, 0x96, 0xc8, 0x8d, 0x64,
	0x59, 0xbc, 0x51, 0xd7, 0xaf, 0x3f, 0x9d, 0x9b, 0x64, 0x0b, 0xec, 0x0e, 0xe0, 0xe5, 0xa3, 0x16,
	0xa8, 0x09, 0xc6, 0xec, 0x7c, 0x36, 0xb5, 0x5f, 0xa0, 0x06, 0xe8, 0x67, 0xa3, 0x2b, 0x5b, 0x13,
	0x8b, 0xd1, 0x4f, 0x27, 0x76, 0xcd, 0xfd, 0xab, 0x06, 0x9d, 0x1f, 0x0b, 0x9a, 0x6f, 0xc6, 0x69,
	0x12, 0x44, 0x92, 0xfe, 0x0d, 0x80, 0x68, 0xbc, 0xc8, 0x85, 0xe3, 0x72, 0x6c, 0x7b, 0xff, 0x7f,
	0x2f, 0x6c, 0xf1, 0xed, 0xd4, 0xbf, 0x05, 0x33, 0xcd, 0x03, 0x9a, 0xcb, 0xb9, 0x75, 0x86, 0x9f,
	0xec, 0x54, 0xde, 0x6f, 0xe9, 0x9d, 0x0b, 0x01, 0x56, 0x3a, 0xf4, 0x21, 0xb4, 0xd7, 0x37, 0x34,
	0xa7, 0x0b, 0x3f, 0x26, 0x05, 0xa3, 0x72, 0x92, 0x16, 0x6e, 0xc9, 0xbd, 0xb1, 0xdc, 0x42, 0x0e,
	0x34, 0x18, 0x8d, 0xa9, 0xcf, 0x99, 0x63, 0xf4, 0xf4, 0xbe, 0x85, 0x2b, 0x88, 0xde, 0x83, 0x7a,
	0x98, 0xa7, 0x45, 0xc6, 0x1c, 0x53, 0x1e, 0x94, 0x08, 0xbd, 0x0b, 0x66, 0x1c, 0xad, 0x22, 0x2e,
	0xe3, 0x34, 0xb0, 0x02, 0xee, 0x47, 0x60, 0xca, 0xd6, 0xf7, 0x23, 0x1a, 0x5d, 0x8c, 0x6d, 0x4d,
	0x6c, 0x4d, 0xa6, 0x17, 0x63, 0xbb, 0xe6, 0xfe, 0xa9, 0x81, 0x29, 0x0d, 0xa3, 0x2e, 0x34, 0x03,
	0xc2, 0xc9, 0x92, 0x30, 0x15, 0x8c, 0x85, 0xb7, 0x58, 0x74, 0xe0, 0x64, 0x19, 0xab, 0xf7, 0x6a,
	0x61, 0x05, 0xd0, 0x1e, 0x00, 0xfd, 0x3d, 0xcb, 0x29, 0x63, 0xe2, 0x91, 0xa8, 0xab, 0xdc, 0xd9,
	0x41, 0x23, 0xb0, 0xfc, 0x2a, 0x86, 0xf2, 0x01, 0xbe, 0x79, 0x46, 0x62, 0xf8, 0x56, 0xe5, 0x9e,
	0x81, 0x8e, 0xd3, 0x35, 0xf2, 0xc0, 0x78, 0xe6, 0xf7, 0x4c, 0xf2, 0x44, 0x52, 0xbf, 0x91, 0xb8,
	0xa0, 0xcc, 0xa9, 0xa9, 0xa4, 0x14, 0x72, 0x0f, 0xc1, 0x3c, 0x11, 0x99, 0xa1, 0xcf, 0xc0, 0xc8,
	0xd3, 0x35, 0x73, 0xb4, 0x9e, 0xde, 0x6f, 0x0d, 0x9d, 0x9d, 0xae, 0x70, 0xba, 0xc6, 0x92, 0xe5,
	0xfe, 0xad, 0x41, 0x0b, 0x53, 0x12, 0x60, 0xb2, 0x16, 0x91, 0xa0, 0x03, 0x30, 0x7f, 0x15, 0x96,
	0xb7, 0x7e, 0xfe, 0xf3, 0x52, 0x58, 0x11, 0xc5, 0x50, 0xfd, 0x34, 0x2e, 0x56, 0x49, 0xe5, 0xa8,
	0x82, 0x68, 0xb8, 0x1d, 0xaa, 0x2e, 0xbd, 0xec, 0x2e, 0x26, 0x5d, 0x6f, 0x07, 0xee, 0x40, 0x23,
	0x57, 0x56, 0x64, 0xac, 0x16, 0xae, 0xa0, 0xfb, 0x47, 0x0d, 0xda, 0x3f, 0xe7, 0x11, 0xa7, 0x95,
	0xd5, 0xb7, 0x9f, 0xea, 0x1d, 0xab, 0xfa, 0x7d, 0xab, 0x55, 0x68, 0xc6, 0x73, 0x42, 0x43, 0x47,
	0xd0, 0x56, 0xc2, 0x85, 0xfc, 0x19, 0x96, 0x6f, 0xb6, 0x33, 0x7c, 0xbd, 0x53, 0x35, 0x96, 0xc4,
	0xf9, 0x26, 0xa3, 0xb8, 0xe5, 0x6f, 0xd7, 0x0c, 0x7d, 0x0d, 0x20, 0xdc, 0x96, 0x15, 0xea, 0xb2,
	0xc2, 0x07, 0x3b, 0x2b, 0x4c, 0x08, 0x27, 0x52, 0x6f, 0x05, 0xe5, 0x8a, 0x7d, 0xfa, 0x15, 0xb4,
	0x27, 0xe5, 0x5d, 0xc5, 0x06, 0x7a, 0x07, 0xac, 0xcb, 0xd9, 0x64, 0x7a, 0x7c, 0x3a, 0x9b, 0x4e,
	0xec, 0x17, 0xa8, 0x0d, 0xcd, 0xd3, 0xd9, 0xf1, 0xf7, 0x97, 0x57, 0x93, 0x23, 0x5b, 0x43, 0x1d,
	0x80, 0x1f, 0xf0, 0xf9, 0xd9, 0x74, 0xfe, 0xdd, 0xf4, 0xf2, 0xc2, 0xae, 0x1d, 0x7d, 0xfc, 0xcb,
	0x9b, 0x30, 0xe2, 0x37, 0xc5, 0x52, 0x74, 0x18, 0xdc, 0xb6, 0xdc, 0x27, 0xd1, 0x80, 0x64, 0x51,
	0xf9, 0x2f, 0xb4, 0xac, 0xcb, 0x37, 0xf8, 0xf9, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x8c,
	0x14, 0xa3, 0x9d, 0x06, 0x00, 0x00,
}
