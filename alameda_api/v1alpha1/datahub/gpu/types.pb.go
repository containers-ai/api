// This file has messages related to gpu info, metrics and predictions

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.0.0
// source: alameda_api/v1alpha1/datahub/gpu/types.proto

package gpu

import (
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

type GpuMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host        string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Instance    string `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance,omitempty"`
	Job         string `protobuf:"bytes,3,opt,name=job,proto3" json:"job,omitempty"`
	MinorNumber string `protobuf:"bytes,4,opt,name=minor_number,json=minorNumber,proto3" json:"minor_number,omitempty"`
}

func (x *GpuMetadata) Reset() {
	*x = GpuMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_gpu_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GpuMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GpuMetadata) ProtoMessage() {}

func (x *GpuMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_gpu_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GpuMetadata.ProtoReflect.Descriptor instead.
func (*GpuMetadata) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_gpu_types_proto_rawDescGZIP(), []int{0}
}

func (x *GpuMetadata) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *GpuMetadata) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

func (x *GpuMetadata) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *GpuMetadata) GetMinorNumber() string {
	if x != nil {
		return x.MinorNumber
	}
	return ""
}

type GpuSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemoryTotal float32 `protobuf:"fixed32,1,opt,name=memory_total,json=memoryTotal,proto3" json:"memory_total,omitempty"` // Total memory in bytes
}

func (x *GpuSpec) Reset() {
	*x = GpuSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_gpu_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GpuSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GpuSpec) ProtoMessage() {}

func (x *GpuSpec) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_gpu_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GpuSpec.ProtoReflect.Descriptor instead.
func (*GpuSpec) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_gpu_types_proto_rawDescGZIP(), []int{1}
}

func (x *GpuSpec) GetMemoryTotal() float32 {
	if x != nil {
		return x.MemoryTotal
	}
	return 0
}

var File_alameda_api_v1alpha1_datahub_gpu_types_proto protoreflect.FileDescriptor

var file_alameda_api_v1alpha1_datahub_gpu_types_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x67,
	0x70, 0x75, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61,
	0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x70, 0x75, 0x22, 0x72, 0x0a, 0x0b, 0x47, 0x70, 0x75,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69,
	0x6e, 0x6f, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2c, 0x0a,
	0x07, 0x47, 0x70, 0x75, 0x53, 0x70, 0x65, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x3f, 0x5a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x73, 0x2d, 0x61, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x61, 0x6d,
	0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x67, 0x70, 0x75, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alameda_api_v1alpha1_datahub_gpu_types_proto_rawDescOnce sync.Once
	file_alameda_api_v1alpha1_datahub_gpu_types_proto_rawDescData = file_alameda_api_v1alpha1_datahub_gpu_types_proto_rawDesc
)

func file_alameda_api_v1alpha1_datahub_gpu_types_proto_rawDescGZIP() []byte {
	file_alameda_api_v1alpha1_datahub_gpu_types_proto_rawDescOnce.Do(func() {
		file_alameda_api_v1alpha1_datahub_gpu_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_alameda_api_v1alpha1_datahub_gpu_types_proto_rawDescData)
	})
	return file_alameda_api_v1alpha1_datahub_gpu_types_proto_rawDescData
}

var file_alameda_api_v1alpha1_datahub_gpu_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_alameda_api_v1alpha1_datahub_gpu_types_proto_goTypes = []interface{}{
	(*GpuMetadata)(nil), // 0: containersai.alameda.v1alpha1.datahub.gpu.GpuMetadata
	(*GpuSpec)(nil),     // 1: containersai.alameda.v1alpha1.datahub.gpu.GpuSpec
}
var file_alameda_api_v1alpha1_datahub_gpu_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_alameda_api_v1alpha1_datahub_gpu_types_proto_init() }
func file_alameda_api_v1alpha1_datahub_gpu_types_proto_init() {
	if File_alameda_api_v1alpha1_datahub_gpu_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alameda_api_v1alpha1_datahub_gpu_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GpuMetadata); i {
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
		file_alameda_api_v1alpha1_datahub_gpu_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GpuSpec); i {
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
			RawDescriptor: file_alameda_api_v1alpha1_datahub_gpu_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_alameda_api_v1alpha1_datahub_gpu_types_proto_goTypes,
		DependencyIndexes: file_alameda_api_v1alpha1_datahub_gpu_types_proto_depIdxs,
		MessageInfos:      file_alameda_api_v1alpha1_datahub_gpu_types_proto_msgTypes,
	}.Build()
	File_alameda_api_v1alpha1_datahub_gpu_types_proto = out.File
	file_alameda_api_v1alpha1_datahub_gpu_types_proto_rawDesc = nil
	file_alameda_api_v1alpha1_datahub_gpu_types_proto_goTypes = nil
	file_alameda_api_v1alpha1_datahub_gpu_types_proto_depIdxs = nil
}
