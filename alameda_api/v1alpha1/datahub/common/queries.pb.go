// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/common/queries.proto

package common

import (
	fmt "fmt"
	common "github.com/containers-ai/api/common"
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
	return fileDescriptor_d602763cab07305c, []int{0, 0}
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
	return fileDescriptor_d602763cab07305c, []int{2, 0}
}

//*
// Represents a time range definition
//
type TimeRange struct {
	StartTime            *timestamp.Timestamp        `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp        `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Step                 *duration.Duration          `protobuf:"bytes,3,opt,name=step,proto3" json:"step,omitempty"`
	AggregateFunction    TimeRange_AggregateFunction `protobuf:"varint,4,opt,name=aggregateFunction,proto3,enum=containersai.alameda.v1alpha1.datahub.common.TimeRange_AggregateFunction" json:"aggregateFunction,omitempty"`
	ApplyTime            *timestamp.Timestamp        `protobuf:"bytes,5,opt,name=apply_time,json=applyTime,proto3" json:"apply_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TimeRange) Reset()         { *m = TimeRange{} }
func (m *TimeRange) String() string { return proto.CompactTextString(m) }
func (*TimeRange) ProtoMessage()    {}
func (*TimeRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_d602763cab07305c, []int{0}
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

type Condition struct {
	Keys                 []string          `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Values               []string          `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	Operators            []string          `protobuf:"bytes,3,rep,name=operators,proto3" json:"operators,omitempty"`
	Types                []common.DataType `protobuf:"varint,4,rep,packed,name=types,proto3,enum=containersai.common.DataType" json:"types,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Condition) Reset()         { *m = Condition{} }
func (m *Condition) String() string { return proto.CompactTextString(m) }
func (*Condition) ProtoMessage()    {}
func (*Condition) Descriptor() ([]byte, []int) {
	return fileDescriptor_d602763cab07305c, []int{1}
}

func (m *Condition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Condition.Unmarshal(m, b)
}
func (m *Condition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Condition.Marshal(b, m, deterministic)
}
func (m *Condition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Condition.Merge(m, src)
}
func (m *Condition) XXX_Size() int {
	return xxx_messageInfo_Condition.Size(m)
}
func (m *Condition) XXX_DiscardUnknown() {
	xxx_messageInfo_Condition.DiscardUnknown(m)
}

var xxx_messageInfo_Condition proto.InternalMessageInfo

func (m *Condition) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Condition) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Condition) GetOperators() []string {
	if m != nil {
		return m.Operators
	}
	return nil
}

func (m *Condition) GetTypes() []common.DataType {
	if m != nil {
		return m.Types
	}
	return nil
}

type QueryCondition struct {
	TimeRange            *TimeRange           `protobuf:"bytes,1,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	Order                QueryCondition_Order `protobuf:"varint,2,opt,name=order,proto3,enum=containersai.alameda.v1alpha1.datahub.common.QueryCondition_Order" json:"order,omitempty"`
	WhereClause          string               `protobuf:"bytes,3,opt,name=where_clause,json=whereClause,proto3" json:"where_clause,omitempty"`
	WhereCondition       []*Condition         `protobuf:"bytes,4,rep,name=where_condition,json=whereCondition,proto3" json:"where_condition,omitempty"`
	Selects              []string             `protobuf:"bytes,5,rep,name=selects,proto3" json:"selects,omitempty"`
	Groups               []string             `protobuf:"bytes,6,rep,name=groups,proto3" json:"groups,omitempty"`
	Limit                uint64               `protobuf:"varint,7,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *QueryCondition) Reset()         { *m = QueryCondition{} }
func (m *QueryCondition) String() string { return proto.CompactTextString(m) }
func (*QueryCondition) ProtoMessage()    {}
func (*QueryCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_d602763cab07305c, []int{2}
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

func (m *QueryCondition) GetWhereCondition() []*Condition {
	if m != nil {
		return m.WhereCondition
	}
	return nil
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

func init() {
	proto.RegisterEnum("containersai.alameda.v1alpha1.datahub.common.TimeRange_AggregateFunction", TimeRange_AggregateFunction_name, TimeRange_AggregateFunction_value)
	proto.RegisterEnum("containersai.alameda.v1alpha1.datahub.common.QueryCondition_Order", QueryCondition_Order_name, QueryCondition_Order_value)
	proto.RegisterType((*TimeRange)(nil), "containersai.alameda.v1alpha1.datahub.common.TimeRange")
	proto.RegisterType((*Condition)(nil), "containersai.alameda.v1alpha1.datahub.common.Condition")
	proto.RegisterType((*QueryCondition)(nil), "containersai.alameda.v1alpha1.datahub.common.QueryCondition")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/common/queries.proto", fileDescriptor_d602763cab07305c)
}

var fileDescriptor_d602763cab07305c = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xfd, 0x39, 0xb6, 0x9b, 0x7a, 0xfa, 0x53, 0x48, 0x57, 0x08, 0x99, 0x8a, 0x3f, 0x21, 0xe2,
	0x90, 0x03, 0x5d, 0xab, 0xa9, 0x10, 0xe2, 0x46, 0x9a, 0x16, 0xc4, 0x81, 0x56, 0xb8, 0x55, 0x55,
	0x71, 0x09, 0xdb, 0x78, 0x70, 0x57, 0xd8, 0x5e, 0xb3, 0xbb, 0x6e, 0x95, 0xaf, 0xc0, 0x8d, 0x0f,
	0xc3, 0xf7, 0x43, 0xbb, 0xb6, 0x53, 0xb5, 0x3d, 0x94, 0x70, 0xdb, 0x99, 0x79, 0xef, 0xed, 0xcc,
	0xdb, 0xd1, 0xc2, 0x0e, 0xcb, 0x58, 0x8e, 0x09, 0x9b, 0xb1, 0x92, 0x47, 0x97, 0x3b, 0x2c, 0x2b,
	0x2f, 0xd8, 0x4e, 0x94, 0x30, 0xcd, 0x2e, 0xaa, 0xf3, 0x68, 0x2e, 0xf2, 0x5c, 0x14, 0xd1, 0x8f,
	0x0a, 0x25, 0x47, 0x45, 0x4b, 0x29, 0xb4, 0x20, 0xaf, 0xe6, 0xa2, 0xd0, 0x8c, 0x17, 0x28, 0x15,
	0xe3, 0xb4, 0xe1, 0xd3, 0x96, 0x4b, 0x1b, 0x2e, 0xad, 0xb9, 0x5b, 0xa4, 0xd1, 0xd0, 0x8b, 0xb2,
	0x55, 0xd8, 0x7a, 0x96, 0x0a, 0x91, 0x66, 0x18, 0xd9, 0xe8, 0xbc, 0xfa, 0x16, 0x25, 0x95, 0x64,
	0x9a, 0x8b, 0xa2, 0xa9, 0x3f, 0xbf, 0x5d, 0xd7, 0x3c, 0x47, 0xa5, 0x59, 0x5e, 0xd6, 0x80, 0xe1,
	0x2f, 0x17, 0x82, 0x13, 0x9e, 0x63, 0xcc, 0x8a, 0x14, 0xc9, 0x5b, 0x00, 0xa5, 0x99, 0xd4, 0x33,
	0x03, 0x0b, 0x9d, 0x81, 0x33, 0xda, 0x18, 0x6f, 0xd1, 0x5a, 0x83, 0xb6, 0x1a, 0xf4, 0xa4, 0xd5,
	0x88, 0x03, 0x8b, 0x36, 0x31, 0x79, 0x0d, 0xeb, 0x58, 0x24, 0x35, 0xb1, 0x73, 0x2f, 0xb1, 0x8b,
	0x45, 0x62, 0x69, 0xdb, 0xe0, 0x29, 0x8d, 0x65, 0xe8, 0x5a, 0xca, 0xe3, 0x3b, 0x94, 0xfd, 0x66,
	0x9e, 0xd8, 0xc2, 0xc8, 0x15, 0x6c, 0xb2, 0x34, 0x95, 0x98, 0x32, 0x8d, 0xef, 0xab, 0x62, 0x6e,
	0x4a, 0xa1, 0x37, 0x70, 0x46, 0xbd, 0xf1, 0x47, 0xba, 0x8a, 0x9b, 0x74, 0x39, 0x34, 0x9d, 0xdc,
	0x16, 0x8c, 0xef, 0xde, 0x61, 0x9c, 0x61, 0x65, 0x99, 0x2d, 0xea, 0x01, 0xfd, 0xfb, 0x9d, 0xb1,
	0x68, 0x13, 0x0f, 0x23, 0xd8, 0xbc, 0x73, 0x05, 0x59, 0x07, 0xef, 0xf0, 0xe8, 0xf0, 0xa0, 0xff,
	0x1f, 0xe9, 0x82, 0xfb, 0x69, 0x72, 0xd6, 0x77, 0xcc, 0x61, 0x72, 0xfa, 0xa1, 0xdf, 0x19, 0xfe,
	0x74, 0x20, 0x98, 0x8a, 0x22, 0xe1, 0x16, 0x49, 0xc0, 0xfb, 0x8e, 0x0b, 0x15, 0x3a, 0x03, 0x77,
	0x14, 0xc4, 0xf6, 0x4c, 0x1e, 0xc1, 0xda, 0x25, 0xcb, 0x2a, 0x54, 0x61, 0xc7, 0x66, 0x9b, 0x88,
	0x3c, 0x81, 0x40, 0x94, 0x28, 0x99, 0x16, 0x52, 0x85, 0xae, 0x2d, 0x5d, 0x27, 0xc8, 0x2e, 0xf8,
	0x76, 0x77, 0x42, 0x6f, 0xe0, 0x8e, 0x7a, 0xe3, 0xa7, 0x37, 0x0d, 0x6b, 0x7c, 0xd9, 0x67, 0x9a,
	0x9d, 0x2c, 0x4a, 0x8c, 0x6b, 0xec, 0xf0, 0xb7, 0x0b, 0xbd, 0xcf, 0x15, 0xca, 0xc5, 0x75, 0x47,
	0xa7, 0x00, 0xc6, 0x85, 0x99, 0x34, 0xf6, 0x35, 0x5b, 0xf2, 0xe6, 0x1f, 0xdd, 0x8f, 0x03, 0xbd,
	0xdc, 0xbe, 0x33, 0xf0, 0x85, 0x4c, 0x50, 0xda, 0xfd, 0xe9, 0x8d, 0xf7, 0x56, 0x93, 0xbc, 0xd9,
	0x24, 0x3d, 0x32, 0x4a, 0x71, 0x2d, 0x48, 0x5e, 0xc0, 0xff, 0x57, 0x17, 0x28, 0x71, 0x36, 0xcf,
	0x58, 0xa5, 0xd0, 0x6e, 0x5b, 0x10, 0x6f, 0xd8, 0xdc, 0xd4, 0xa6, 0xc8, 0x57, 0x78, 0xd0, 0x40,
	0x5a, 0x09, 0x6b, 0xd3, 0xca, 0x93, 0x2d, 0x3b, 0x88, 0x7b, 0xb5, 0xfc, 0xd2, 0xb6, 0x10, 0xba,
	0x0a, 0x33, 0x9c, 0x6b, 0x15, 0xfa, 0xf6, 0x69, 0xda, 0xd0, 0x3c, 0x67, 0x2a, 0x45, 0x55, 0xaa,
	0x70, 0xad, 0x7e, 0xce, 0x3a, 0x22, 0x0f, 0xc1, 0xcf, 0x78, 0xce, 0x75, 0xd8, 0x1d, 0x38, 0x23,
	0x2f, 0xae, 0x83, 0xe1, 0x4b, 0xf0, 0xed, 0x70, 0x37, 0x77, 0x68, 0x72, 0x3c, 0xed, 0x3b, 0x26,
	0xb5, 0x7f, 0x70, 0x3c, 0xed, 0x77, 0xf6, 0xf6, 0xbe, 0xbc, 0x4b, 0xb9, 0x6e, 0x1a, 0x8b, 0xae,
	0x47, 0xd8, 0x66, 0x3c, 0x32, 0x3f, 0xd4, 0x5f, 0xfc, 0x56, 0xe7, 0x6b, 0x76, 0xb1, 0x77, 0xff,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x48, 0xab, 0xc5, 0x75, 0xdb, 0x04, 0x00, 0x00,
}
