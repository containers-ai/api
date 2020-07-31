/// This file has messages related to nodes

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.0.0
// source: datahub/resource/pod/assign/v1alpha2/assign.proto

package v1alpha2

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

//*
// Represents the priority of a node
type NodePriority struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []string `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *NodePriority) Reset() {
	*x = NodePriority{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_resource_pod_assign_v1alpha2_assign_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodePriority) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodePriority) ProtoMessage() {}

func (x *NodePriority) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_resource_pod_assign_v1alpha2_assign_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodePriority.ProtoReflect.Descriptor instead.
func (*NodePriority) Descriptor() ([]byte, []int) {
	return file_datahub_resource_pod_assign_v1alpha2_assign_proto_rawDescGZIP(), []int{0}
}

func (x *NodePriority) GetNodes() []string {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type Selector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Selector map[string]string `protobuf:"bytes,1,rep,name=selector,proto3" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Selector) Reset() {
	*x = Selector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_resource_pod_assign_v1alpha2_assign_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Selector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Selector) ProtoMessage() {}

func (x *Selector) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_resource_pod_assign_v1alpha2_assign_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Selector.ProtoReflect.Descriptor instead.
func (*Selector) Descriptor() ([]byte, []int) {
	return file_datahub_resource_pod_assign_v1alpha2_assign_proto_rawDescGZIP(), []int{1}
}

func (x *Selector) GetSelector() map[string]string {
	if x != nil {
		return x.Selector
	}
	return nil
}

var File_datahub_resource_pod_assign_v1alpha2_assign_proto protoreflect.FileDescriptor

var file_datahub_resource_pod_assign_v1alpha2_assign_proto_rawDesc = []byte{
	0x0a, 0x31, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x70, 0x6f, 0x64, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x31, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61,
	0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x22, 0x24, 0x0a, 0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0xae, 0x01, 0x0a,
	0x08, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x65, 0x0a, 0x08, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68,
	0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x64, 0x2e,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x1a, 0x3b, 0x0a, 0x0d, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x43, 0x5a,
	0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2d, 0x61, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70,
	0x6f, 0x64, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datahub_resource_pod_assign_v1alpha2_assign_proto_rawDescOnce sync.Once
	file_datahub_resource_pod_assign_v1alpha2_assign_proto_rawDescData = file_datahub_resource_pod_assign_v1alpha2_assign_proto_rawDesc
)

func file_datahub_resource_pod_assign_v1alpha2_assign_proto_rawDescGZIP() []byte {
	file_datahub_resource_pod_assign_v1alpha2_assign_proto_rawDescOnce.Do(func() {
		file_datahub_resource_pod_assign_v1alpha2_assign_proto_rawDescData = protoimpl.X.CompressGZIP(file_datahub_resource_pod_assign_v1alpha2_assign_proto_rawDescData)
	})
	return file_datahub_resource_pod_assign_v1alpha2_assign_proto_rawDescData
}

var file_datahub_resource_pod_assign_v1alpha2_assign_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_datahub_resource_pod_assign_v1alpha2_assign_proto_goTypes = []interface{}{
	(*NodePriority)(nil), // 0: containersai.datahub.resource.pod.assign.v1alpha2.NodePriority
	(*Selector)(nil),     // 1: containersai.datahub.resource.pod.assign.v1alpha2.Selector
	nil,                  // 2: containersai.datahub.resource.pod.assign.v1alpha2.Selector.SelectorEntry
}
var file_datahub_resource_pod_assign_v1alpha2_assign_proto_depIdxs = []int32{
	2, // 0: containersai.datahub.resource.pod.assign.v1alpha2.Selector.selector:type_name -> containersai.datahub.resource.pod.assign.v1alpha2.Selector.SelectorEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_datahub_resource_pod_assign_v1alpha2_assign_proto_init() }
func file_datahub_resource_pod_assign_v1alpha2_assign_proto_init() {
	if File_datahub_resource_pod_assign_v1alpha2_assign_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datahub_resource_pod_assign_v1alpha2_assign_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodePriority); i {
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
		file_datahub_resource_pod_assign_v1alpha2_assign_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Selector); i {
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
			RawDescriptor: file_datahub_resource_pod_assign_v1alpha2_assign_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datahub_resource_pod_assign_v1alpha2_assign_proto_goTypes,
		DependencyIndexes: file_datahub_resource_pod_assign_v1alpha2_assign_proto_depIdxs,
		MessageInfos:      file_datahub_resource_pod_assign_v1alpha2_assign_proto_msgTypes,
	}.Build()
	File_datahub_resource_pod_assign_v1alpha2_assign_proto = out.File
	file_datahub_resource_pod_assign_v1alpha2_assign_proto_rawDesc = nil
	file_datahub_resource_pod_assign_v1alpha2_assign_proto_goTypes = nil
	file_datahub_resource_pod_assign_v1alpha2_assign_proto_depIdxs = nil
}
