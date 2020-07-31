/// This file has messages and services (methods) related to the AI engine. This file will be deprated in v0.2.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.0.0
// source: alameda_api/v1alpha1/ai_service/ai_service.proto

package ai_service

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
//  Recommendation policy. A policy may be either stable or compact.
type RecommendationPolicy int32

const (
	RecommendationPolicy_STABLE  RecommendationPolicy = 0
	RecommendationPolicy_COMPACT RecommendationPolicy = 1
)

// Enum value maps for RecommendationPolicy.
var (
	RecommendationPolicy_name = map[int32]string{
		0: "STABLE",
		1: "COMPACT",
	}
	RecommendationPolicy_value = map[string]int32{
		"STABLE":  0,
		"COMPACT": 1,
	}
)

func (x RecommendationPolicy) Enum() *RecommendationPolicy {
	p := new(RecommendationPolicy)
	*p = x
	return p
}

func (x RecommendationPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecommendationPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_alameda_api_v1alpha1_ai_service_ai_service_proto_enumTypes[0].Descriptor()
}

func (RecommendationPolicy) Type() protoreflect.EnumType {
	return &file_alameda_api_v1alpha1_ai_service_ai_service_proto_enumTypes[0]
}

func (x RecommendationPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecommendationPolicy.Descriptor instead.
func (RecommendationPolicy) EnumDescriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_ai_service_ai_service_proto_rawDescGZIP(), []int{0}
}

/// Types of an object
type Object_Type int32

const (
	Object_POD  Object_Type = 0
	Object_NODE Object_Type = 1
)

// Enum value maps for Object_Type.
var (
	Object_Type_name = map[int32]string{
		0: "POD",
		1: "NODE",
	}
	Object_Type_value = map[string]int32{
		"POD":  0,
		"NODE": 1,
	}
)

func (x Object_Type) Enum() *Object_Type {
	p := new(Object_Type)
	*p = x
	return p
}

func (x Object_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Object_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_alameda_api_v1alpha1_ai_service_ai_service_proto_enumTypes[1].Descriptor()
}

func (Object_Type) Type() protoreflect.EnumType {
	return &file_alameda_api_v1alpha1_ai_service_ai_service_proto_enumTypes[1]
}

func (x Object_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Object_Type.Descriptor instead.
func (Object_Type) EnumDescriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_ai_service_ai_service_proto_rawDescGZIP(), []int{1, 0}
}

//*
// Represents a Kubernetes pod.
//
type Pod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName     string `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ResourceLink string `protobuf:"bytes,2,opt,name=resourceLink,proto3" json:"resourceLink,omitempty"`
}

func (x *Pod) Reset() {
	*x = Pod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_ai_service_ai_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pod) ProtoMessage() {}

func (x *Pod) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_ai_service_ai_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pod.ProtoReflect.Descriptor instead.
func (*Pod) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_ai_service_ai_service_proto_rawDescGZIP(), []int{0}
}

func (x *Pod) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *Pod) GetResourceLink() string {
	if x != nil {
		return x.ResourceLink
	}
	return ""
}

//*
// Represents a Kubernetes object.
type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      Object_Type          `protobuf:"varint,1,opt,name=type,proto3,enum=containers_ai.alameda.v1alpha1.ai_service.Object_Type" json:"type,omitempty"`
	Policy    RecommendationPolicy `protobuf:"varint,2,opt,name=policy,proto3,enum=containers_ai.alameda.v1alpha1.ai_service.RecommendationPolicy" json:"policy,omitempty"`
	Uid       string               `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Namespace string               `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string               `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Pod       *Pod                 `protobuf:"bytes,6,opt,name=pod,proto3" json:"pod,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_ai_service_ai_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_ai_service_ai_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_ai_service_ai_service_proto_rawDescGZIP(), []int{1}
}

func (x *Object) GetType() Object_Type {
	if x != nil {
		return x.Type
	}
	return Object_POD
}

func (x *Object) GetPolicy() RecommendationPolicy {
	if x != nil {
		return x.Policy
	}
	return RecommendationPolicy_STABLE
}

func (x *Object) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Object) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Object) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Object) GetPod() *Pod {
	if x != nil {
		return x.Pod
	}
	return nil
}

//*
// Represents a request for creating a list of prediction objects
type PredictionObjectListCreationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Objects []*Object `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *PredictionObjectListCreationRequest) Reset() {
	*x = PredictionObjectListCreationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_ai_service_ai_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictionObjectListCreationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictionObjectListCreationRequest) ProtoMessage() {}

func (x *PredictionObjectListCreationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_ai_service_ai_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictionObjectListCreationRequest.ProtoReflect.Descriptor instead.
func (*PredictionObjectListCreationRequest) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_ai_service_ai_service_proto_rawDescGZIP(), []int{2}
}

func (x *PredictionObjectListCreationRequest) GetObjects() []*Object {
	if x != nil {
		return x.Objects
	}
	return nil
}

//*
// Represents a request for removing a list of prediction objects
type PredictionObjectListDeletionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Objects []*Object `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *PredictionObjectListDeletionRequest) Reset() {
	*x = PredictionObjectListDeletionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_ai_service_ai_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictionObjectListDeletionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictionObjectListDeletionRequest) ProtoMessage() {}

func (x *PredictionObjectListDeletionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_ai_service_ai_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictionObjectListDeletionRequest.ProtoReflect.Descriptor instead.
func (*PredictionObjectListDeletionRequest) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_ai_service_ai_service_proto_rawDescGZIP(), []int{3}
}

func (x *PredictionObjectListDeletionRequest) GetObjects() []*Object {
	if x != nil {
		return x.Objects
	}
	return nil
}

//*
// Represents a reponse of a request
type RequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RequestResponse) Reset() {
	*x = RequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_ai_service_ai_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestResponse) ProtoMessage() {}

func (x *RequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_ai_service_ai_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestResponse.ProtoReflect.Descriptor instead.
func (*RequestResponse) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_ai_service_ai_service_proto_rawDescGZIP(), []int{4}
}

func (x *RequestResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_alameda_api_v1alpha1_ai_service_ai_service_proto protoreflect.FileDescriptor

var file_alameda_api_v1alpha1_ai_service_ai_service_proto_rawDesc = []byte{
	0x0a, 0x30, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x61, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x29, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x5f, 0x61,
	0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x61, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x03, 0x50, 0x6f,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x22, 0xce, 0x02, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x4a, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x5f, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d,
	0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x69, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x57, 0x0a, 0x06, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x5f, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64,
	0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x69, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x03, 0x70, 0x6f, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x5f, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x50, 0x6f, 0x64, 0x52, 0x03, 0x70, 0x6f, 0x64, 0x22, 0x19, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x07, 0x0a, 0x03, 0x50, 0x4f, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x44,
	0x45, 0x10, 0x01, 0x22, 0x72, 0x0a, 0x23, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x07, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x5f, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d,
	0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x69, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x72, 0x0a, 0x23, 0x50, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b,
	0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x5f, 0x61, 0x69, 0x2e,
	0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x61, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x2b, 0x0a, 0x0f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x2f, 0x0a, 0x14, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x43, 0x4f, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x10, 0x01, 0x32, 0xc2, 0x02, 0x0a, 0x10, 0x41, 0x6c,
	0x61, 0x6d, 0x65, 0x64, 0x61, 0x41, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa7,
	0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x4e, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x5f, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65,
	0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x69, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x5f, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65,
	0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x69, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x83, 0x01, 0x0a, 0x17, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x12, 0x4e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x73, 0x5f, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x61, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x3e,
	0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2d, 0x61, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x61, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alameda_api_v1alpha1_ai_service_ai_service_proto_rawDescOnce sync.Once
	file_alameda_api_v1alpha1_ai_service_ai_service_proto_rawDescData = file_alameda_api_v1alpha1_ai_service_ai_service_proto_rawDesc
)

func file_alameda_api_v1alpha1_ai_service_ai_service_proto_rawDescGZIP() []byte {
	file_alameda_api_v1alpha1_ai_service_ai_service_proto_rawDescOnce.Do(func() {
		file_alameda_api_v1alpha1_ai_service_ai_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_alameda_api_v1alpha1_ai_service_ai_service_proto_rawDescData)
	})
	return file_alameda_api_v1alpha1_ai_service_ai_service_proto_rawDescData
}

var file_alameda_api_v1alpha1_ai_service_ai_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_alameda_api_v1alpha1_ai_service_ai_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_alameda_api_v1alpha1_ai_service_ai_service_proto_goTypes = []interface{}{
	(RecommendationPolicy)(0), // 0: containers_ai.alameda.v1alpha1.ai_service.RecommendationPolicy
	(Object_Type)(0),          // 1: containers_ai.alameda.v1alpha1.ai_service.Object.Type
	(*Pod)(nil),               // 2: containers_ai.alameda.v1alpha1.ai_service.Pod
	(*Object)(nil),            // 3: containers_ai.alameda.v1alpha1.ai_service.Object
	(*PredictionObjectListCreationRequest)(nil), // 4: containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListCreationRequest
	(*PredictionObjectListDeletionRequest)(nil), // 5: containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListDeletionRequest
	(*RequestResponse)(nil),                     // 6: containers_ai.alameda.v1alpha1.ai_service.RequestResponse
	(*empty.Empty)(nil),                         // 7: google.protobuf.Empty
}
var file_alameda_api_v1alpha1_ai_service_ai_service_proto_depIdxs = []int32{
	1, // 0: containers_ai.alameda.v1alpha1.ai_service.Object.type:type_name -> containers_ai.alameda.v1alpha1.ai_service.Object.Type
	0, // 1: containers_ai.alameda.v1alpha1.ai_service.Object.policy:type_name -> containers_ai.alameda.v1alpha1.ai_service.RecommendationPolicy
	2, // 2: containers_ai.alameda.v1alpha1.ai_service.Object.pod:type_name -> containers_ai.alameda.v1alpha1.ai_service.Pod
	3, // 3: containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListCreationRequest.objects:type_name -> containers_ai.alameda.v1alpha1.ai_service.Object
	3, // 4: containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListDeletionRequest.objects:type_name -> containers_ai.alameda.v1alpha1.ai_service.Object
	4, // 5: containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService.CreatePredictionObjects:input_type -> containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListCreationRequest
	5, // 6: containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService.DeletePredictionObjects:input_type -> containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListDeletionRequest
	6, // 7: containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService.CreatePredictionObjects:output_type -> containers_ai.alameda.v1alpha1.ai_service.RequestResponse
	7, // 8: containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService.DeletePredictionObjects:output_type -> google.protobuf.Empty
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_alameda_api_v1alpha1_ai_service_ai_service_proto_init() }
func file_alameda_api_v1alpha1_ai_service_ai_service_proto_init() {
	if File_alameda_api_v1alpha1_ai_service_ai_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alameda_api_v1alpha1_ai_service_ai_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pod); i {
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
		file_alameda_api_v1alpha1_ai_service_ai_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_alameda_api_v1alpha1_ai_service_ai_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictionObjectListCreationRequest); i {
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
		file_alameda_api_v1alpha1_ai_service_ai_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictionObjectListDeletionRequest); i {
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
		file_alameda_api_v1alpha1_ai_service_ai_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestResponse); i {
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
			RawDescriptor: file_alameda_api_v1alpha1_ai_service_ai_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_alameda_api_v1alpha1_ai_service_ai_service_proto_goTypes,
		DependencyIndexes: file_alameda_api_v1alpha1_ai_service_ai_service_proto_depIdxs,
		EnumInfos:         file_alameda_api_v1alpha1_ai_service_ai_service_proto_enumTypes,
		MessageInfos:      file_alameda_api_v1alpha1_ai_service_ai_service_proto_msgTypes,
	}.Build()
	File_alameda_api_v1alpha1_ai_service_ai_service_proto = out.File
	file_alameda_api_v1alpha1_ai_service_ai_service_proto_rawDesc = nil
	file_alameda_api_v1alpha1_ai_service_ai_service_proto_goTypes = nil
	file_alameda_api_v1alpha1_ai_service_ai_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AlamedaAIServiceClient is the client API for AlamedaAIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlamedaAIServiceClient interface {
	/// Used to create prediction objects
	CreatePredictionObjects(ctx context.Context, in *PredictionObjectListCreationRequest, opts ...grpc.CallOption) (*RequestResponse, error)
	/// Used to remove prediction objects
	DeletePredictionObjects(ctx context.Context, in *PredictionObjectListDeletionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type alamedaAIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlamedaAIServiceClient(cc grpc.ClientConnInterface) AlamedaAIServiceClient {
	return &alamedaAIServiceClient{cc}
}

func (c *alamedaAIServiceClient) CreatePredictionObjects(ctx context.Context, in *PredictionObjectListCreationRequest, opts ...grpc.CallOption) (*RequestResponse, error) {
	out := new(RequestResponse)
	err := c.cc.Invoke(ctx, "/containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService/CreatePredictionObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alamedaAIServiceClient) DeletePredictionObjects(ctx context.Context, in *PredictionObjectListDeletionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService/DeletePredictionObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlamedaAIServiceServer is the server API for AlamedaAIService service.
type AlamedaAIServiceServer interface {
	/// Used to create prediction objects
	CreatePredictionObjects(context.Context, *PredictionObjectListCreationRequest) (*RequestResponse, error)
	/// Used to remove prediction objects
	DeletePredictionObjects(context.Context, *PredictionObjectListDeletionRequest) (*empty.Empty, error)
}

// UnimplementedAlamedaAIServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAlamedaAIServiceServer struct {
}

func (*UnimplementedAlamedaAIServiceServer) CreatePredictionObjects(context.Context, *PredictionObjectListCreationRequest) (*RequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePredictionObjects not implemented")
}
func (*UnimplementedAlamedaAIServiceServer) DeletePredictionObjects(context.Context, *PredictionObjectListDeletionRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePredictionObjects not implemented")
}

func RegisterAlamedaAIServiceServer(s *grpc.Server, srv AlamedaAIServiceServer) {
	s.RegisterService(&_AlamedaAIService_serviceDesc, srv)
}

func _AlamedaAIService_CreatePredictionObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictionObjectListCreationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlamedaAIServiceServer).CreatePredictionObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService/CreatePredictionObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlamedaAIServiceServer).CreatePredictionObjects(ctx, req.(*PredictionObjectListCreationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlamedaAIService_DeletePredictionObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictionObjectListDeletionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlamedaAIServiceServer).DeletePredictionObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService/DeletePredictionObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlamedaAIServiceServer).DeletePredictionObjects(ctx, req.(*PredictionObjectListDeletionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlamedaAIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService",
	HandlerType: (*AlamedaAIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePredictionObjects",
			Handler:    _AlamedaAIService_CreatePredictionObjects_Handler,
		},
		{
			MethodName: "DeletePredictionObjects",
			Handler:    _AlamedaAIService_DeletePredictionObjects_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alameda_api/v1alpha1/ai_service/ai_service.proto",
}
