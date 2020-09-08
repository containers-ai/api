// This file has messages related to metric data

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: alameda_api/v1alpha1/datahub/metrics/types.proto

package metrics

import (
	common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	schemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
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

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaMeta *schemas.SchemaMeta `protobuf:"bytes,1,opt,name=schema_meta,json=schemaMeta,proto3" json:"schema_meta,omitempty"`
	MetricData []*MetricData       `protobuf:"bytes,2,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_metrics_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_metrics_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_metrics_types_proto_rawDescGZIP(), []int{0}
}

func (x *Metric) GetSchemaMeta() *schemas.SchemaMeta {
	if x != nil {
		return x.SchemaMeta
	}
	return nil
}

func (x *Metric) GetMetricData() []*MetricData {
	if x != nil {
		return x.MetricData
	}
	return nil
}

//*
// Represents a piece of metreic data.
type MetricData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricType common.MetricType `protobuf:"varint,1,opt,name=metric_type,json=metricType,proto3,enum=containersai.alameda.v1alpha1.datahub.common.MetricType" json:"metric_type,omitempty"`
	ReadData   *common.ReadData  `protobuf:"bytes,2,opt,name=read_data,json=readData,proto3" json:"read_data,omitempty"`
}

func (x *MetricData) Reset() {
	*x = MetricData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_metrics_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricData) ProtoMessage() {}

func (x *MetricData) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_metrics_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricData.ProtoReflect.Descriptor instead.
func (*MetricData) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_metrics_types_proto_rawDescGZIP(), []int{1}
}

func (x *MetricData) GetMetricType() common.MetricType {
	if x != nil {
		return x.MetricType
	}
	return common.MetricType_METRICS_TYPE_UNDEFINED
}

func (x *MetricData) GetReadData() *common.ReadData {
	if x != nil {
		return x.ReadData
	}
	return nil
}

var File_alameda_api_v1alpha1_datahub_metrics_types_proto protoreflect.FileDescriptor

var file_alameda_api_v1alpha1_datahub_metrics_types_proto_rawDesc = []byte{
	0x0a, 0x30, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69,
	0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x1a, 0x31, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68,
	0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61,
	0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x06, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x12, 0x5a, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75,
	0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d, 0x65, 0x74, 0x61,
	0x12, 0x5a, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x22, 0xbc, 0x01, 0x0a,
	0x0a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x12, 0x59, 0x0a, 0x0b, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e,
	0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x53, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x42, 0x43, 0x5a, 0x41, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x73, 0x2d, 0x61, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x61, 0x6d,
	0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alameda_api_v1alpha1_datahub_metrics_types_proto_rawDescOnce sync.Once
	file_alameda_api_v1alpha1_datahub_metrics_types_proto_rawDescData = file_alameda_api_v1alpha1_datahub_metrics_types_proto_rawDesc
)

func file_alameda_api_v1alpha1_datahub_metrics_types_proto_rawDescGZIP() []byte {
	file_alameda_api_v1alpha1_datahub_metrics_types_proto_rawDescOnce.Do(func() {
		file_alameda_api_v1alpha1_datahub_metrics_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_alameda_api_v1alpha1_datahub_metrics_types_proto_rawDescData)
	})
	return file_alameda_api_v1alpha1_datahub_metrics_types_proto_rawDescData
}

var file_alameda_api_v1alpha1_datahub_metrics_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_alameda_api_v1alpha1_datahub_metrics_types_proto_goTypes = []interface{}{
	(*Metric)(nil),             // 0: containersai.alameda.v1alpha1.datahub.metrics.Metric
	(*MetricData)(nil),         // 1: containersai.alameda.v1alpha1.datahub.metrics.MetricData
	(*schemas.SchemaMeta)(nil), // 2: containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta
	(common.MetricType)(0),     // 3: containersai.alameda.v1alpha1.datahub.common.MetricType
	(*common.ReadData)(nil),    // 4: containersai.alameda.v1alpha1.datahub.common.ReadData
}
var file_alameda_api_v1alpha1_datahub_metrics_types_proto_depIdxs = []int32{
	2, // 0: containersai.alameda.v1alpha1.datahub.metrics.Metric.schema_meta:type_name -> containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta
	1, // 1: containersai.alameda.v1alpha1.datahub.metrics.Metric.metric_data:type_name -> containersai.alameda.v1alpha1.datahub.metrics.MetricData
	3, // 2: containersai.alameda.v1alpha1.datahub.metrics.MetricData.metric_type:type_name -> containersai.alameda.v1alpha1.datahub.common.MetricType
	4, // 3: containersai.alameda.v1alpha1.datahub.metrics.MetricData.read_data:type_name -> containersai.alameda.v1alpha1.datahub.common.ReadData
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_alameda_api_v1alpha1_datahub_metrics_types_proto_init() }
func file_alameda_api_v1alpha1_datahub_metrics_types_proto_init() {
	if File_alameda_api_v1alpha1_datahub_metrics_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alameda_api_v1alpha1_datahub_metrics_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
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
		file_alameda_api_v1alpha1_datahub_metrics_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricData); i {
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
			RawDescriptor: file_alameda_api_v1alpha1_datahub_metrics_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_alameda_api_v1alpha1_datahub_metrics_types_proto_goTypes,
		DependencyIndexes: file_alameda_api_v1alpha1_datahub_metrics_types_proto_depIdxs,
		MessageInfos:      file_alameda_api_v1alpha1_datahub_metrics_types_proto_msgTypes,
	}.Build()
	File_alameda_api_v1alpha1_datahub_metrics_types_proto = out.File
	file_alameda_api_v1alpha1_datahub_metrics_types_proto_rawDesc = nil
	file_alameda_api_v1alpha1_datahub_metrics_types_proto_goTypes = nil
	file_alameda_api_v1alpha1_datahub_metrics_types_proto_depIdxs = nil
}
