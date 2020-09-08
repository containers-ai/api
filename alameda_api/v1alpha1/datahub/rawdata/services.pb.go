// This file has messages related to read & write rawdata

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: alameda_api/v1alpha1/datahub/rawdata/services.proto

package rawdata

import (
	common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	proto "github.com/golang/protobuf/proto"
	status "google.golang.org/genproto/googleapis/rpc/status"
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
// Represents a request for reading rawdata from database.
type ReadRawdataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatabaseType common.DatabaseType `protobuf:"varint,1,opt,name=database_type,json=databaseType,proto3,enum=containersai.alameda.v1alpha1.datahub.common.DatabaseType" json:"database_type,omitempty"`
	Queries      []*Query            `protobuf:"bytes,2,rep,name=queries,proto3" json:"queries,omitempty"`
}

func (x *ReadRawdataRequest) Reset() {
	*x = ReadRawdataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_rawdata_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRawdataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRawdataRequest) ProtoMessage() {}

func (x *ReadRawdataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_rawdata_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRawdataRequest.ProtoReflect.Descriptor instead.
func (*ReadRawdataRequest) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_rawdata_services_proto_rawDescGZIP(), []int{0}
}

func (x *ReadRawdataRequest) GetDatabaseType() common.DatabaseType {
	if x != nil {
		return x.DatabaseType
	}
	return common.DatabaseType_UNDEFINED
}

func (x *ReadRawdataRequest) GetQueries() []*Query {
	if x != nil {
		return x.Queries
	}
	return nil
}

//*
// Represents a response for listing rawdata from database.
type ReadRawdataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Rawdata []*ReadRawdata `protobuf:"bytes,2,rep,name=rawdata,proto3" json:"rawdata,omitempty"`
}

func (x *ReadRawdataResponse) Reset() {
	*x = ReadRawdataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_rawdata_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRawdataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRawdataResponse) ProtoMessage() {}

func (x *ReadRawdataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_rawdata_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRawdataResponse.ProtoReflect.Descriptor instead.
func (*ReadRawdataResponse) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_rawdata_services_proto_rawDescGZIP(), []int{1}
}

func (x *ReadRawdataResponse) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ReadRawdataResponse) GetRawdata() []*ReadRawdata {
	if x != nil {
		return x.Rawdata
	}
	return nil
}

//*
// Represents a request for writing rawdata to database.
type WriteRawdataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatabaseType common.DatabaseType `protobuf:"varint,1,opt,name=database_type,json=databaseType,proto3,enum=containersai.alameda.v1alpha1.datahub.common.DatabaseType" json:"database_type,omitempty"`
	Rawdata      []*WriteRawdata     `protobuf:"bytes,2,rep,name=rawdata,proto3" json:"rawdata,omitempty"`
}

func (x *WriteRawdataRequest) Reset() {
	*x = WriteRawdataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_rawdata_services_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRawdataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRawdataRequest) ProtoMessage() {}

func (x *WriteRawdataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_rawdata_services_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRawdataRequest.ProtoReflect.Descriptor instead.
func (*WriteRawdataRequest) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_rawdata_services_proto_rawDescGZIP(), []int{2}
}

func (x *WriteRawdataRequest) GetDatabaseType() common.DatabaseType {
	if x != nil {
		return x.DatabaseType
	}
	return common.DatabaseType_UNDEFINED
}

func (x *WriteRawdataRequest) GetRawdata() []*WriteRawdata {
	if x != nil {
		return x.Rawdata
	}
	return nil
}

var File_alameda_api_v1alpha1_datahub_rawdata_services_proto protoreflect.FileDescriptor

var file_alameda_api_v1alpha1_datahub_rawdata_services_proto_rawDesc = []byte{
	0x0a, 0x33, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x72,
	0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x61, 0x77,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x2f, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68,
	0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2f, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x72, 0x61, 0x77, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x61, 0x6c, 0x61, 0x6d, 0x65,
	0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x61, 0x64, 0x52, 0x61, 0x77,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5f, 0x0a, 0x0d, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61,
	0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4e, 0x0a, 0x07,
	0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61,
	0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x22, 0x97, 0x01, 0x0a,
	0x13, 0x52, 0x65, 0x61, 0x64, 0x52, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x54, 0x0a, 0x07, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69,
	0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x52, 0x07, 0x72,
	0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x22, 0xcd, 0x01, 0x0a, 0x13, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5f,
	0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x55, 0x0a, 0x07, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e,
	0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x52, 0x07, 0x72,
	0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2d,
	0x61, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2f, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_alameda_api_v1alpha1_datahub_rawdata_services_proto_rawDescOnce sync.Once
	file_alameda_api_v1alpha1_datahub_rawdata_services_proto_rawDescData = file_alameda_api_v1alpha1_datahub_rawdata_services_proto_rawDesc
)

func file_alameda_api_v1alpha1_datahub_rawdata_services_proto_rawDescGZIP() []byte {
	file_alameda_api_v1alpha1_datahub_rawdata_services_proto_rawDescOnce.Do(func() {
		file_alameda_api_v1alpha1_datahub_rawdata_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_alameda_api_v1alpha1_datahub_rawdata_services_proto_rawDescData)
	})
	return file_alameda_api_v1alpha1_datahub_rawdata_services_proto_rawDescData
}

var file_alameda_api_v1alpha1_datahub_rawdata_services_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_alameda_api_v1alpha1_datahub_rawdata_services_proto_goTypes = []interface{}{
	(*ReadRawdataRequest)(nil),  // 0: containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataRequest
	(*ReadRawdataResponse)(nil), // 1: containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataResponse
	(*WriteRawdataRequest)(nil), // 2: containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdataRequest
	(common.DatabaseType)(0),    // 3: containersai.alameda.v1alpha1.datahub.common.DatabaseType
	(*Query)(nil),               // 4: containersai.alameda.v1alpha1.datahub.rawdata.Query
	(*status.Status)(nil),       // 5: google.rpc.Status
	(*ReadRawdata)(nil),         // 6: containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdata
	(*WriteRawdata)(nil),        // 7: containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdata
}
var file_alameda_api_v1alpha1_datahub_rawdata_services_proto_depIdxs = []int32{
	3, // 0: containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataRequest.database_type:type_name -> containersai.alameda.v1alpha1.datahub.common.DatabaseType
	4, // 1: containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataRequest.queries:type_name -> containersai.alameda.v1alpha1.datahub.rawdata.Query
	5, // 2: containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataResponse.status:type_name -> google.rpc.Status
	6, // 3: containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdataResponse.rawdata:type_name -> containersai.alameda.v1alpha1.datahub.rawdata.ReadRawdata
	3, // 4: containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdataRequest.database_type:type_name -> containersai.alameda.v1alpha1.datahub.common.DatabaseType
	7, // 5: containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdataRequest.rawdata:type_name -> containersai.alameda.v1alpha1.datahub.rawdata.WriteRawdata
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_alameda_api_v1alpha1_datahub_rawdata_services_proto_init() }
func file_alameda_api_v1alpha1_datahub_rawdata_services_proto_init() {
	if File_alameda_api_v1alpha1_datahub_rawdata_services_proto != nil {
		return
	}
	file_alameda_api_v1alpha1_datahub_rawdata_rawdata_proto_init()
	file_alameda_api_v1alpha1_datahub_rawdata_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_alameda_api_v1alpha1_datahub_rawdata_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRawdataRequest); i {
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
		file_alameda_api_v1alpha1_datahub_rawdata_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRawdataResponse); i {
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
		file_alameda_api_v1alpha1_datahub_rawdata_services_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRawdataRequest); i {
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
			RawDescriptor: file_alameda_api_v1alpha1_datahub_rawdata_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_alameda_api_v1alpha1_datahub_rawdata_services_proto_goTypes,
		DependencyIndexes: file_alameda_api_v1alpha1_datahub_rawdata_services_proto_depIdxs,
		MessageInfos:      file_alameda_api_v1alpha1_datahub_rawdata_services_proto_msgTypes,
	}.Build()
	File_alameda_api_v1alpha1_datahub_rawdata_services_proto = out.File
	file_alameda_api_v1alpha1_datahub_rawdata_services_proto_rawDesc = nil
	file_alameda_api_v1alpha1_datahub_rawdata_services_proto_goTypes = nil
	file_alameda_api_v1alpha1_datahub_rawdata_services_proto_depIdxs = nil
}
