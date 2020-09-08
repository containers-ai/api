// This file has messages related general definitions

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: alameda_api/v1alpha1/datahub/common/queries.proto

package common

import (
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TimeRange_AggregateFunction int32

const (
	TimeRange_NONE TimeRange_AggregateFunction = 0
	TimeRange_MAX  TimeRange_AggregateFunction = 1
	TimeRange_AVG  TimeRange_AggregateFunction = 2
)

// Enum value maps for TimeRange_AggregateFunction.
var (
	TimeRange_AggregateFunction_name = map[int32]string{
		0: "NONE",
		1: "MAX",
		2: "AVG",
	}
	TimeRange_AggregateFunction_value = map[string]int32{
		"NONE": 0,
		"MAX":  1,
		"AVG":  2,
	}
)

func (x TimeRange_AggregateFunction) Enum() *TimeRange_AggregateFunction {
	p := new(TimeRange_AggregateFunction)
	*p = x
	return p
}

func (x TimeRange_AggregateFunction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TimeRange_AggregateFunction) Descriptor() protoreflect.EnumDescriptor {
	return file_alameda_api_v1alpha1_datahub_common_queries_proto_enumTypes[0].Descriptor()
}

func (TimeRange_AggregateFunction) Type() protoreflect.EnumType {
	return &file_alameda_api_v1alpha1_datahub_common_queries_proto_enumTypes[0]
}

func (x TimeRange_AggregateFunction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TimeRange_AggregateFunction.Descriptor instead.
func (TimeRange_AggregateFunction) EnumDescriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_common_queries_proto_rawDescGZIP(), []int{0, 0}
}

type QueryCondition_Order int32

const (
	QueryCondition_NONE QueryCondition_Order = 0
	QueryCondition_ASC  QueryCondition_Order = 1
	QueryCondition_DESC QueryCondition_Order = 2
)

// Enum value maps for QueryCondition_Order.
var (
	QueryCondition_Order_name = map[int32]string{
		0: "NONE",
		1: "ASC",
		2: "DESC",
	}
	QueryCondition_Order_value = map[string]int32{
		"NONE": 0,
		"ASC":  1,
		"DESC": 2,
	}
)

func (x QueryCondition_Order) Enum() *QueryCondition_Order {
	p := new(QueryCondition_Order)
	*p = x
	return p
}

func (x QueryCondition_Order) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryCondition_Order) Descriptor() protoreflect.EnumDescriptor {
	return file_alameda_api_v1alpha1_datahub_common_queries_proto_enumTypes[1].Descriptor()
}

func (QueryCondition_Order) Type() protoreflect.EnumType {
	return &file_alameda_api_v1alpha1_datahub_common_queries_proto_enumTypes[1]
}

func (x QueryCondition_Order) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryCondition_Order.Descriptor instead.
func (QueryCondition_Order) EnumDescriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_common_queries_proto_rawDescGZIP(), []int{4, 0}
}

//*
// Represents a time range definition.
type TimeRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime         *timestamp.Timestamp        `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime           *timestamp.Timestamp        `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Timeout           *timestamp.Timestamp        `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Step              *duration.Duration          `protobuf:"bytes,4,opt,name=step,proto3" json:"step,omitempty"`
	AggregateFunction TimeRange_AggregateFunction `protobuf:"varint,5,opt,name=aggregateFunction,proto3,enum=containersai.alameda.v1alpha1.datahub.common.TimeRange_AggregateFunction" json:"aggregateFunction,omitempty"`
	ApplyTime         *timestamp.Timestamp        `protobuf:"bytes,6,opt,name=apply_time,json=applyTime,proto3" json:"apply_time,omitempty"`
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_common_queries_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_common_queries_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRange.ProtoReflect.Descriptor instead.
func (*TimeRange) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_common_queries_proto_rawDescGZIP(), []int{0}
}

func (x *TimeRange) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *TimeRange) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *TimeRange) GetTimeout() *timestamp.Timestamp {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *TimeRange) GetStep() *duration.Duration {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *TimeRange) GetAggregateFunction() TimeRange_AggregateFunction {
	if x != nil {
		return x.AggregateFunction
	}
	return TimeRange_NONE
}

func (x *TimeRange) GetApplyTime() *timestamp.Timestamp {
	if x != nil {
		return x.ApplyTime
	}
	return nil
}

//*
// Represents a query condition.
type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys      []string   `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Values    []string   `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	Operators []string   `protobuf:"bytes,3,rep,name=operators,proto3" json:"operators,omitempty"`
	Types     []DataType `protobuf:"varint,4,rep,packed,name=types,proto3,enum=containersai.alameda.v1alpha1.datahub.common.DataType" json:"types,omitempty"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_common_queries_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_common_queries_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_common_queries_proto_rawDescGZIP(), []int{1}
}

func (x *Condition) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *Condition) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *Condition) GetOperators() []string {
	if x != nil {
		return x.Operators
	}
	return nil
}

func (x *Condition) GetTypes() []DataType {
	if x != nil {
		return x.Types
	}
	return nil
}

//*
// Represents a datahub functional query, includes aggregation, selector and transformation.
type Function struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type              FunctionType `protobuf:"varint,1,opt,name=type,proto3,enum=containersai.alameda.v1alpha1.datahub.common.FunctionType" json:"type,omitempty"`
	Fields            []string     `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	Tags              []string     `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	Target            string       `protobuf:"bytes,4,opt,name=target,proto3" json:"target,omitempty"`
	RegularExpression string       `protobuf:"bytes,5,opt,name=regular_expression,json=regularExpression,proto3" json:"regular_expression,omitempty"`
	Unit              string       `protobuf:"bytes,6,opt,name=unit,proto3" json:"unit,omitempty"`
	Number            int64        `protobuf:"varint,7,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Function) Reset() {
	*x = Function{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_common_queries_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Function) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Function) ProtoMessage() {}

func (x *Function) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_common_queries_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Function.ProtoReflect.Descriptor instead.
func (*Function) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_common_queries_proto_rawDescGZIP(), []int{2}
}

func (x *Function) GetType() FunctionType {
	if x != nil {
		return x.Type
	}
	return FunctionType_FUNCTIONTYPE_UNDEFINED
}

func (x *Function) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Function) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Function) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *Function) GetRegularExpression() string {
	if x != nil {
		return x.RegularExpression
	}
	return ""
}

func (x *Function) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Function) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

//*
// Represents a query to a user-specified measurement.
type Into struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database                 string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	RetentionPolicy          string `protobuf:"bytes,2,opt,name=retention_policy,json=retentionPolicy,proto3" json:"retention_policy,omitempty"`
	Measurement              string `protobuf:"bytes,3,opt,name=measurement,proto3" json:"measurement,omitempty"`
	IsDefaultRetentionPolicy bool   `protobuf:"varint,4,opt,name=is_default_retention_policy,json=isDefaultRetentionPolicy,proto3" json:"is_default_retention_policy,omitempty"`
	IsAllMeasurements        bool   `protobuf:"varint,5,opt,name=is_all_measurements,json=isAllMeasurements,proto3" json:"is_all_measurements,omitempty"`
}

func (x *Into) Reset() {
	*x = Into{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_common_queries_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Into) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Into) ProtoMessage() {}

func (x *Into) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_common_queries_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Into.ProtoReflect.Descriptor instead.
func (*Into) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_common_queries_proto_rawDescGZIP(), []int{3}
}

func (x *Into) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *Into) GetRetentionPolicy() string {
	if x != nil {
		return x.RetentionPolicy
	}
	return ""
}

func (x *Into) GetMeasurement() string {
	if x != nil {
		return x.Measurement
	}
	return ""
}

func (x *Into) GetIsDefaultRetentionPolicy() bool {
	if x != nil {
		return x.IsDefaultRetentionPolicy
	}
	return false
}

func (x *Into) GetIsAllMeasurements() bool {
	if x != nil {
		return x.IsAllMeasurements
	}
	return false
}

//*
// Represents a datahub query request.
type QueryCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeRange      *TimeRange           `protobuf:"bytes,1,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	Order          QueryCondition_Order `protobuf:"varint,2,opt,name=order,proto3,enum=containersai.alameda.v1alpha1.datahub.common.QueryCondition_Order" json:"order,omitempty"`
	Function       *Function            `protobuf:"bytes,3,opt,name=function,proto3" json:"function,omitempty"`
	Into           *Into                `protobuf:"bytes,4,opt,name=into,proto3" json:"into,omitempty"`
	WhereClause    string               `protobuf:"bytes,5,opt,name=where_clause,json=whereClause,proto3" json:"where_clause,omitempty"`
	WhereCondition []*Condition         `protobuf:"bytes,6,rep,name=where_condition,json=whereCondition,proto3" json:"where_condition,omitempty"`
	Selects        []string             `protobuf:"bytes,7,rep,name=selects,proto3" json:"selects,omitempty"`
	Groups         []string             `protobuf:"bytes,8,rep,name=groups,proto3" json:"groups,omitempty"`
	Limit          uint64               `protobuf:"varint,9,opt,name=limit,proto3" json:"limit,omitempty"`
	SubQuery       *QueryCondition      `protobuf:"bytes,10,opt,name=sub_query,json=subQuery,proto3" json:"sub_query,omitempty"`
}

func (x *QueryCondition) Reset() {
	*x = QueryCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_common_queries_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryCondition) ProtoMessage() {}

func (x *QueryCondition) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_common_queries_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryCondition.ProtoReflect.Descriptor instead.
func (*QueryCondition) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_common_queries_proto_rawDescGZIP(), []int{4}
}

func (x *QueryCondition) GetTimeRange() *TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

func (x *QueryCondition) GetOrder() QueryCondition_Order {
	if x != nil {
		return x.Order
	}
	return QueryCondition_NONE
}

func (x *QueryCondition) GetFunction() *Function {
	if x != nil {
		return x.Function
	}
	return nil
}

func (x *QueryCondition) GetInto() *Into {
	if x != nil {
		return x.Into
	}
	return nil
}

func (x *QueryCondition) GetWhereClause() string {
	if x != nil {
		return x.WhereClause
	}
	return ""
}

func (x *QueryCondition) GetWhereCondition() []*Condition {
	if x != nil {
		return x.WhereCondition
	}
	return nil
}

func (x *QueryCondition) GetSelects() []string {
	if x != nil {
		return x.Selects
	}
	return nil
}

func (x *QueryCondition) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *QueryCondition) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *QueryCondition) GetSubQuery() *QueryCondition {
	if x != nil {
		return x.SubQuery
	}
	return nil
}

var File_alameda_api_v1alpha1_datahub_common_queries_proto protoreflect.FileDescriptor

var file_alameda_api_v1alpha1_datahub_common_queries_proto_rawDesc = []byte{
	0x0a, 0x31, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x2c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61,
	0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x1a, 0x2f, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x03, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x73, 0x74, 0x65,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x77, 0x0a, 0x11, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x49, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x11,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4d,
	0x41, 0x58, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x56, 0x47, 0x10, 0x02, 0x22, 0xa3, 0x01,
	0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b,
	0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x4c, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x22, 0xf9, 0x01, 0x0a, 0x08, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x4e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c,
	0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f,
	0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0xde, 0x01, 0x0a, 0x04, 0x49, 0x6e, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x3d, 0x0a, 0x1b, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x2e, 0x0a, 0x13, 0x69, 0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x6d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69,
	0x73, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0xac, 0x05, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x56, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x58, 0x0a, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x42, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64,
	0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x04, 0x69, 0x6e, 0x74,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x74, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x74,
	0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x68, 0x65, 0x72, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x75, 0x73,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x68, 0x65, 0x72, 0x65, 0x43, 0x6c,
	0x61, 0x75, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x0f, 0x77, 0x68, 0x65, 0x72, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61,
	0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x77, 0x68, 0x65, 0x72, 0x65, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x59,
	0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69,
	0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x73, 0x75, 0x62, 0x51, 0x75, 0x65, 0x72, 0x79, 0x22, 0x24, 0x0a, 0x05, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x41, 0x53, 0x43, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x02, 0x42,
	0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2d, 0x61, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alameda_api_v1alpha1_datahub_common_queries_proto_rawDescOnce sync.Once
	file_alameda_api_v1alpha1_datahub_common_queries_proto_rawDescData = file_alameda_api_v1alpha1_datahub_common_queries_proto_rawDesc
)

func file_alameda_api_v1alpha1_datahub_common_queries_proto_rawDescGZIP() []byte {
	file_alameda_api_v1alpha1_datahub_common_queries_proto_rawDescOnce.Do(func() {
		file_alameda_api_v1alpha1_datahub_common_queries_proto_rawDescData = protoimpl.X.CompressGZIP(file_alameda_api_v1alpha1_datahub_common_queries_proto_rawDescData)
	})
	return file_alameda_api_v1alpha1_datahub_common_queries_proto_rawDescData
}

var file_alameda_api_v1alpha1_datahub_common_queries_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_alameda_api_v1alpha1_datahub_common_queries_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_alameda_api_v1alpha1_datahub_common_queries_proto_goTypes = []interface{}{
	(TimeRange_AggregateFunction)(0), // 0: containersai.alameda.v1alpha1.datahub.common.TimeRange.AggregateFunction
	(QueryCondition_Order)(0),        // 1: containersai.alameda.v1alpha1.datahub.common.QueryCondition.Order
	(*TimeRange)(nil),                // 2: containersai.alameda.v1alpha1.datahub.common.TimeRange
	(*Condition)(nil),                // 3: containersai.alameda.v1alpha1.datahub.common.Condition
	(*Function)(nil),                 // 4: containersai.alameda.v1alpha1.datahub.common.Function
	(*Into)(nil),                     // 5: containersai.alameda.v1alpha1.datahub.common.Into
	(*QueryCondition)(nil),           // 6: containersai.alameda.v1alpha1.datahub.common.QueryCondition
	(*timestamp.Timestamp)(nil),      // 7: google.protobuf.Timestamp
	(*duration.Duration)(nil),        // 8: google.protobuf.Duration
	(DataType)(0),                    // 9: containersai.alameda.v1alpha1.datahub.common.DataType
	(FunctionType)(0),                // 10: containersai.alameda.v1alpha1.datahub.common.FunctionType
}
var file_alameda_api_v1alpha1_datahub_common_queries_proto_depIdxs = []int32{
	7,  // 0: containersai.alameda.v1alpha1.datahub.common.TimeRange.start_time:type_name -> google.protobuf.Timestamp
	7,  // 1: containersai.alameda.v1alpha1.datahub.common.TimeRange.end_time:type_name -> google.protobuf.Timestamp
	7,  // 2: containersai.alameda.v1alpha1.datahub.common.TimeRange.timeout:type_name -> google.protobuf.Timestamp
	8,  // 3: containersai.alameda.v1alpha1.datahub.common.TimeRange.step:type_name -> google.protobuf.Duration
	0,  // 4: containersai.alameda.v1alpha1.datahub.common.TimeRange.aggregateFunction:type_name -> containersai.alameda.v1alpha1.datahub.common.TimeRange.AggregateFunction
	7,  // 5: containersai.alameda.v1alpha1.datahub.common.TimeRange.apply_time:type_name -> google.protobuf.Timestamp
	9,  // 6: containersai.alameda.v1alpha1.datahub.common.Condition.types:type_name -> containersai.alameda.v1alpha1.datahub.common.DataType
	10, // 7: containersai.alameda.v1alpha1.datahub.common.Function.type:type_name -> containersai.alameda.v1alpha1.datahub.common.FunctionType
	2,  // 8: containersai.alameda.v1alpha1.datahub.common.QueryCondition.time_range:type_name -> containersai.alameda.v1alpha1.datahub.common.TimeRange
	1,  // 9: containersai.alameda.v1alpha1.datahub.common.QueryCondition.order:type_name -> containersai.alameda.v1alpha1.datahub.common.QueryCondition.Order
	4,  // 10: containersai.alameda.v1alpha1.datahub.common.QueryCondition.function:type_name -> containersai.alameda.v1alpha1.datahub.common.Function
	5,  // 11: containersai.alameda.v1alpha1.datahub.common.QueryCondition.into:type_name -> containersai.alameda.v1alpha1.datahub.common.Into
	3,  // 12: containersai.alameda.v1alpha1.datahub.common.QueryCondition.where_condition:type_name -> containersai.alameda.v1alpha1.datahub.common.Condition
	6,  // 13: containersai.alameda.v1alpha1.datahub.common.QueryCondition.sub_query:type_name -> containersai.alameda.v1alpha1.datahub.common.QueryCondition
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_alameda_api_v1alpha1_datahub_common_queries_proto_init() }
func file_alameda_api_v1alpha1_datahub_common_queries_proto_init() {
	if File_alameda_api_v1alpha1_datahub_common_queries_proto != nil {
		return
	}
	file_alameda_api_v1alpha1_datahub_common_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_alameda_api_v1alpha1_datahub_common_queries_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeRange); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_alameda_api_v1alpha1_datahub_common_queries_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_alameda_api_v1alpha1_datahub_common_queries_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Function); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_alameda_api_v1alpha1_datahub_common_queries_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Into); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_alameda_api_v1alpha1_datahub_common_queries_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryCondition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_alameda_api_v1alpha1_datahub_common_queries_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_alameda_api_v1alpha1_datahub_common_queries_proto_goTypes,
		DependencyIndexes: file_alameda_api_v1alpha1_datahub_common_queries_proto_depIdxs,
		EnumInfos:         file_alameda_api_v1alpha1_datahub_common_queries_proto_enumTypes,
		MessageInfos:      file_alameda_api_v1alpha1_datahub_common_queries_proto_msgTypes,
	}.Build()
	File_alameda_api_v1alpha1_datahub_common_queries_proto = out.File
	file_alameda_api_v1alpha1_datahub_common_queries_proto_rawDesc = nil
	file_alameda_api_v1alpha1_datahub_common_queries_proto_goTypes = nil
	file_alameda_api_v1alpha1_datahub_common_queries_proto_depIdxs = nil
}
