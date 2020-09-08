// This file has messages related to read & write rawdata

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: alameda_api/v1alpha1/datahub/rawdata/rawdata.proto

package rawdata

import (
	common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	proto "github.com/golang/protobuf/proto"
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

//*
// Represents a rawdata whcih is read from datahub.
type ReadRawdata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query   *Query          `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Columns []string        `protobuf:"bytes,2,rep,name=columns,proto3" json:"columns,omitempty"`
	Groups  []*common.Group `protobuf:"bytes,3,rep,name=groups,proto3" json:"groups,omitempty"`
	Rawdata string          `protobuf:"bytes,4,opt,name=rawdata,proto3" json:"rawdata,omitempty"`
}

func (x *ReadRawdata) Reset() {
	*x = ReadRawdata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRawdata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRawdata) ProtoMessage() {}

func (x *ReadRawdata) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRawdata.ProtoReflect.Descriptor instead.
func (*ReadRawdata) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_rawDescGZIP(), []int{0}
}

func (x *ReadRawdata) GetQuery() *Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *ReadRawdata) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *ReadRawdata) GetGroups() []*common.Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *ReadRawdata) GetRawdata() string {
	if x != nil {
		return x.Rawdata
	}
	return ""
}

//*
// Represents a rawdata which will be written to datahub.
type WriteRawdata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database    string              `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Table       string              `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	Columns     []string            `protobuf:"bytes,3,rep,name=columns,proto3" json:"columns,omitempty"`
	Rows        []*common.Row       `protobuf:"bytes,4,rep,name=rows,proto3" json:"rows,omitempty"`
	ColumnTypes []common.ColumnType `protobuf:"varint,5,rep,packed,name=column_types,json=columnTypes,proto3,enum=containersai.alameda.v1alpha1.datahub.common.ColumnType" json:"column_types,omitempty"`
	DataTypes   []common.DataType   `protobuf:"varint,6,rep,packed,name=data_types,json=dataTypes,proto3,enum=containersai.alameda.v1alpha1.datahub.common.DataType" json:"data_types,omitempty"`
}

func (x *WriteRawdata) Reset() {
	*x = WriteRawdata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRawdata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRawdata) ProtoMessage() {}

func (x *WriteRawdata) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRawdata.ProtoReflect.Descriptor instead.
func (*WriteRawdata) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_rawDescGZIP(), []int{1}
}

func (x *WriteRawdata) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *WriteRawdata) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *WriteRawdata) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *WriteRawdata) GetRows() []*common.Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

func (x *WriteRawdata) GetColumnTypes() []common.ColumnType {
	if x != nil {
		return x.ColumnTypes
	}
	return nil
}

func (x *WriteRawdata) GetDataTypes() []common.DataType {
	if x != nil {
		return x.DataTypes
	}
	return nil
}

var File_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto protoreflect.FileDescriptor

var file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_rawDesc = []byte{
	0x0a, 0x32, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x72,
	0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x61, 0x77, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x31, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75,
	0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61,
	0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x01, 0x0a, 0x0b, 0x52, 0x65,
	0x61, 0x64, 0x52, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x12, 0x4a, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62,
	0x2e, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12,
	0x4b, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61,
	0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x22, 0xd5, 0x02, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x73, 0x12, 0x45, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69,
	0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x12, 0x5b, 0x0a, 0x0c, 0x63, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e,
	0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65,
	0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x73, 0x42, 0x43,
	0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2d, 0x61, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x61, 0x77, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_rawDescOnce sync.Once
	file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_rawDescData = file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_rawDesc
)

func file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_rawDescGZIP() []byte {
	file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_rawDescOnce.Do(func() {
		file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_rawDescData)
	})
	return file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_rawDescData
}

var file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_goTypes = []interface{}{
	(*ReadRawdata)(nil),    // 0: containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdata
	(*WriteRawdata)(nil),   // 1: containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdata
	(*Query)(nil),          // 2: containersai.alameda.v1alpha1.datahub.rawdata.Query
	(*common.Group)(nil),   // 3: containersai.alameda.v1alpha1.datahub.common.Group
	(*common.Row)(nil),     // 4: containersai.alameda.v1alpha1.datahub.common.Row
	(common.ColumnType)(0), // 5: containersai.alameda.v1alpha1.datahub.common.ColumnType
	(common.DataType)(0),   // 6: containersai.alameda.v1alpha1.datahub.common.DataType
}
var file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_depIdxs = []int32{
	2, // 0: containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdata.query:type_name -> containersai.alameda.v1alpha1.datahub.rawdata.Query
	3, // 1: containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdata.groups:type_name -> containersai.alameda.v1alpha1.datahub.common.Group
	4, // 2: containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdata.rows:type_name -> containersai.alameda.v1alpha1.datahub.common.Row
	5, // 3: containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdata.column_types:type_name -> containersai.alameda.v1alpha1.datahub.common.ColumnType
	6, // 4: containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdata.data_types:type_name -> containersai.alameda.v1alpha1.datahub.common.DataType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_init() }
func file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_init() {
	if File_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto != nil {
		return
	}
	file_alameda_api_v1alpha1_datahub_rawdata_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRawdata); i {
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
		file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRawdata); i {
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
			RawDescriptor: file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_goTypes,
		DependencyIndexes: file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_depIdxs,
		MessageInfos:      file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_msgTypes,
	}.Build()
	File_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto = out.File
	file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_rawDesc = nil
	file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_goTypes = nil
	file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_depIdxs = nil
}
