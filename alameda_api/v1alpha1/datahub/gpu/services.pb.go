// This file has messages related to gpu info, metrics and predictions

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.0.0
// source: alameda_api/v1alpha1/datahub/gpu/services.proto

package gpu

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
// Represents a request for listing graphics processing units that need to be predicted.
type ListGpusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryCondition *common.QueryCondition `protobuf:"bytes,1,opt,name=query_condition,json=queryCondition,proto3" json:"query_condition,omitempty"`
	Host           string                 `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	MinorNumber    string                 `protobuf:"bytes,3,opt,name=minor_number,json=minorNumber,proto3" json:"minor_number,omitempty"`
}

func (x *ListGpusRequest) Reset() {
	*x = ListGpusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGpusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGpusRequest) ProtoMessage() {}

func (x *ListGpusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGpusRequest.ProtoReflect.Descriptor instead.
func (*ListGpusRequest) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_gpu_services_proto_rawDescGZIP(), []int{0}
}

func (x *ListGpusRequest) GetQueryCondition() *common.QueryCondition {
	if x != nil {
		return x.QueryCondition
	}
	return nil
}

func (x *ListGpusRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ListGpusRequest) GetMinorNumber() string {
	if x != nil {
		return x.MinorNumber
	}
	return ""
}

//*
// Represents a response for a listing graphics processing units request.
type ListGpusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Gpus   []*Gpu         `protobuf:"bytes,2,rep,name=gpus,proto3" json:"gpus,omitempty"`
}

func (x *ListGpusResponse) Reset() {
	*x = ListGpusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGpusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGpusResponse) ProtoMessage() {}

func (x *ListGpusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGpusResponse.ProtoReflect.Descriptor instead.
func (*ListGpusResponse) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_gpu_services_proto_rawDescGZIP(), []int{1}
}

func (x *ListGpusResponse) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ListGpusResponse) GetGpus() []*Gpu {
	if x != nil {
		return x.Gpus
	}
	return nil
}

//*
// Represents a request for listing metric data of graphics processing units.
type ListGpuMetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryCondition *common.QueryCondition `protobuf:"bytes,1,opt,name=query_condition,json=queryCondition,proto3" json:"query_condition,omitempty"`
	MetricTypes    []common.MetricType    `protobuf:"varint,2,rep,packed,name=metric_types,json=metricTypes,proto3,enum=containersai.alameda.v1alpha1.datahub.common.MetricType" json:"metric_types,omitempty"`
	Host           string                 `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	MinorNumber    string                 `protobuf:"bytes,4,opt,name=minor_number,json=minorNumber,proto3" json:"minor_number,omitempty"`
}

func (x *ListGpuMetricsRequest) Reset() {
	*x = ListGpuMetricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGpuMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGpuMetricsRequest) ProtoMessage() {}

func (x *ListGpuMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGpuMetricsRequest.ProtoReflect.Descriptor instead.
func (*ListGpuMetricsRequest) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_gpu_services_proto_rawDescGZIP(), []int{2}
}

func (x *ListGpuMetricsRequest) GetQueryCondition() *common.QueryCondition {
	if x != nil {
		return x.QueryCondition
	}
	return nil
}

func (x *ListGpuMetricsRequest) GetMetricTypes() []common.MetricType {
	if x != nil {
		return x.MetricTypes
	}
	return nil
}

func (x *ListGpuMetricsRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ListGpuMetricsRequest) GetMinorNumber() string {
	if x != nil {
		return x.MinorNumber
	}
	return ""
}

//*
// Represents a response for a listing graphics processing units metric data request.
type ListGpuMetricsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	GpuMetrics []*GpuMetric   `protobuf:"bytes,2,rep,name=gpu_metrics,json=gpuMetrics,proto3" json:"gpu_metrics,omitempty"`
}

func (x *ListGpuMetricsResponse) Reset() {
	*x = ListGpuMetricsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGpuMetricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGpuMetricsResponse) ProtoMessage() {}

func (x *ListGpuMetricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGpuMetricsResponse.ProtoReflect.Descriptor instead.
func (*ListGpuMetricsResponse) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_gpu_services_proto_rawDescGZIP(), []int{3}
}

func (x *ListGpuMetricsResponse) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ListGpuMetricsResponse) GetGpuMetrics() []*GpuMetric {
	if x != nil {
		return x.GpuMetrics
	}
	return nil
}

//*
// Represents a request for creating predictions of graphics processing units' metric data.
type CreateGpuPredictionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GpuPredictions []*GpuPrediction `protobuf:"bytes,1,rep,name=gpu_predictions,json=gpuPredictions,proto3" json:"gpu_predictions,omitempty"`
}

func (x *CreateGpuPredictionsRequest) Reset() {
	*x = CreateGpuPredictionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGpuPredictionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGpuPredictionsRequest) ProtoMessage() {}

func (x *CreateGpuPredictionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGpuPredictionsRequest.ProtoReflect.Descriptor instead.
func (*CreateGpuPredictionsRequest) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_gpu_services_proto_rawDescGZIP(), []int{4}
}

func (x *CreateGpuPredictionsRequest) GetGpuPredictions() []*GpuPrediction {
	if x != nil {
		return x.GpuPredictions
	}
	return nil
}

//*
// Represents a list of predicted metric data of graphics processing units.
type ListGpuPredictionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryCondition *common.QueryCondition `protobuf:"bytes,1,opt,name=query_condition,json=queryCondition,proto3" json:"query_condition,omitempty"`
	Host           string                 `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	MinorNumber    string                 `protobuf:"bytes,3,opt,name=minor_number,json=minorNumber,proto3" json:"minor_number,omitempty"`
	Granularity    int64                  `protobuf:"varint,4,opt,name=granularity,proto3" json:"granularity,omitempty"`
	ModelId        string                 `protobuf:"bytes,5,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	PredictionId   string                 `protobuf:"bytes,6,opt,name=prediction_id,json=predictionId,proto3" json:"prediction_id,omitempty"`
}

func (x *ListGpuPredictionsRequest) Reset() {
	*x = ListGpuPredictionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGpuPredictionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGpuPredictionsRequest) ProtoMessage() {}

func (x *ListGpuPredictionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGpuPredictionsRequest.ProtoReflect.Descriptor instead.
func (*ListGpuPredictionsRequest) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_gpu_services_proto_rawDescGZIP(), []int{5}
}

func (x *ListGpuPredictionsRequest) GetQueryCondition() *common.QueryCondition {
	if x != nil {
		return x.QueryCondition
	}
	return nil
}

func (x *ListGpuPredictionsRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ListGpuPredictionsRequest) GetMinorNumber() string {
	if x != nil {
		return x.MinorNumber
	}
	return ""
}

func (x *ListGpuPredictionsRequest) GetGranularity() int64 {
	if x != nil {
		return x.Granularity
	}
	return 0
}

func (x *ListGpuPredictionsRequest) GetModelId() string {
	if x != nil {
		return x.ModelId
	}
	return ""
}

func (x *ListGpuPredictionsRequest) GetPredictionId() string {
	if x != nil {
		return x.PredictionId
	}
	return ""
}

//*
// Represents a response for a listing predictions of graphics processing units request.
type ListGpuPredictionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status         *status.Status   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	GpuPredictions []*GpuPrediction `protobuf:"bytes,2,rep,name=gpu_predictions,json=gpuPredictions,proto3" json:"gpu_predictions,omitempty"`
}

func (x *ListGpuPredictionsResponse) Reset() {
	*x = ListGpuPredictionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGpuPredictionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGpuPredictionsResponse) ProtoMessage() {}

func (x *ListGpuPredictionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGpuPredictionsResponse.ProtoReflect.Descriptor instead.
func (*ListGpuPredictionsResponse) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_gpu_services_proto_rawDescGZIP(), []int{6}
}

func (x *ListGpuPredictionsResponse) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ListGpuPredictionsResponse) GetGpuPredictions() []*GpuPrediction {
	if x != nil {
		return x.GpuPredictions
	}
	return nil
}

var File_alameda_api_v1alpha1_datahub_gpu_services_proto protoreflect.FileDescriptor

var file_alameda_api_v1alpha1_datahub_gpu_services_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x67,
	0x70, 0x75, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x29, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e,
	0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x70, 0x75, 0x1a, 0x31, 0x61, 0x6c,
	0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x31, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2a, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62,
	0x2f, 0x67, 0x70, 0x75, 0x2f, 0x67, 0x70, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74,
	0x47, 0x70, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x65, 0x0a, 0x0f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x69,
	0x6e, 0x6f, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x82, 0x01, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x47, 0x70, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x42, 0x0a, 0x04, 0x67, 0x70,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62,
	0x2e, 0x67, 0x70, 0x75, 0x2e, 0x47, 0x70, 0x75, 0x52, 0x04, 0x67, 0x70, 0x75, 0x73, 0x22, 0x92,
	0x02, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x70, 0x75, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x65, 0x0a, 0x0f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69,
	0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x5b, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0x9b, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x70, 0x75, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x55, 0x0a, 0x0b, 0x67, 0x70,
	0x75, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x34, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61,
	0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x70, 0x75, 0x2e, 0x47, 0x70, 0x75, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x0a, 0x67, 0x70, 0x75, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x22, 0x80, 0x01, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x70, 0x75, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x61, 0x0a, 0x0f, 0x67, 0x70, 0x75, 0x5f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64,
	0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68,
	0x75, 0x62, 0x2e, 0x67, 0x70, 0x75, 0x2e, 0x47, 0x70, 0x75, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x67, 0x70, 0x75, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x9b, 0x02, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x70, 0x75,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x65, 0x0a, 0x0f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65,
	0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x20, 0x0a, 0x0b, 0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0xab, 0x01, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x70, 0x75, 0x50, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x61, 0x0a,
	0x0f, 0x67, 0x70, 0x75, 0x5f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x67,
	0x70, 0x75, 0x2e, 0x47, 0x70, 0x75, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0e, 0x67, 0x70, 0x75, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2d, 0x61, 0x69, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x67, 0x70,
	0x75, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alameda_api_v1alpha1_datahub_gpu_services_proto_rawDescOnce sync.Once
	file_alameda_api_v1alpha1_datahub_gpu_services_proto_rawDescData = file_alameda_api_v1alpha1_datahub_gpu_services_proto_rawDesc
)

func file_alameda_api_v1alpha1_datahub_gpu_services_proto_rawDescGZIP() []byte {
	file_alameda_api_v1alpha1_datahub_gpu_services_proto_rawDescOnce.Do(func() {
		file_alameda_api_v1alpha1_datahub_gpu_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_alameda_api_v1alpha1_datahub_gpu_services_proto_rawDescData)
	})
	return file_alameda_api_v1alpha1_datahub_gpu_services_proto_rawDescData
}

var file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_alameda_api_v1alpha1_datahub_gpu_services_proto_goTypes = []interface{}{
	(*ListGpusRequest)(nil),             // 0: containersai.alameda.v1alpha1.datahub.gpu.ListGpusRequest
	(*ListGpusResponse)(nil),            // 1: containersai.alameda.v1alpha1.datahub.gpu.ListGpusResponse
	(*ListGpuMetricsRequest)(nil),       // 2: containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsRequest
	(*ListGpuMetricsResponse)(nil),      // 3: containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsResponse
	(*CreateGpuPredictionsRequest)(nil), // 4: containersai.alameda.v1alpha1.datahub.gpu.CreateGpuPredictionsRequest
	(*ListGpuPredictionsRequest)(nil),   // 5: containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsRequest
	(*ListGpuPredictionsResponse)(nil),  // 6: containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsResponse
	(*common.QueryCondition)(nil),       // 7: containersai.alameda.v1alpha1.datahub.common.QueryCondition
	(*status.Status)(nil),               // 8: google.rpc.Status
	(*Gpu)(nil),                         // 9: containersai.alameda.v1alpha1.datahub.gpu.Gpu
	(common.MetricType)(0),              // 10: containersai.alameda.v1alpha1.datahub.common.MetricType
	(*GpuMetric)(nil),                   // 11: containersai.alameda.v1alpha1.datahub.gpu.GpuMetric
	(*GpuPrediction)(nil),               // 12: containersai.alameda.v1alpha1.datahub.gpu.GpuPrediction
}
var file_alameda_api_v1alpha1_datahub_gpu_services_proto_depIdxs = []int32{
	7,  // 0: containersai.alameda.v1alpha1.datahub.gpu.ListGpusRequest.query_condition:type_name -> containersai.alameda.v1alpha1.datahub.common.QueryCondition
	8,  // 1: containersai.alameda.v1alpha1.datahub.gpu.ListGpusResponse.status:type_name -> google.rpc.Status
	9,  // 2: containersai.alameda.v1alpha1.datahub.gpu.ListGpusResponse.gpus:type_name -> containersai.alameda.v1alpha1.datahub.gpu.Gpu
	7,  // 3: containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsRequest.query_condition:type_name -> containersai.alameda.v1alpha1.datahub.common.QueryCondition
	10, // 4: containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsRequest.metric_types:type_name -> containersai.alameda.v1alpha1.datahub.common.MetricType
	8,  // 5: containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsResponse.status:type_name -> google.rpc.Status
	11, // 6: containersai.alameda.v1alpha1.datahub.gpu.ListGpuMetricsResponse.gpu_metrics:type_name -> containersai.alameda.v1alpha1.datahub.gpu.GpuMetric
	12, // 7: containersai.alameda.v1alpha1.datahub.gpu.CreateGpuPredictionsRequest.gpu_predictions:type_name -> containersai.alameda.v1alpha1.datahub.gpu.GpuPrediction
	7,  // 8: containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsRequest.query_condition:type_name -> containersai.alameda.v1alpha1.datahub.common.QueryCondition
	8,  // 9: containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsResponse.status:type_name -> google.rpc.Status
	12, // 10: containersai.alameda.v1alpha1.datahub.gpu.ListGpuPredictionsResponse.gpu_predictions:type_name -> containersai.alameda.v1alpha1.datahub.gpu.GpuPrediction
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_alameda_api_v1alpha1_datahub_gpu_services_proto_init() }
func file_alameda_api_v1alpha1_datahub_gpu_services_proto_init() {
	if File_alameda_api_v1alpha1_datahub_gpu_services_proto != nil {
		return
	}
	file_alameda_api_v1alpha1_datahub_gpu_gpu_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGpusRequest); i {
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
		file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGpusResponse); i {
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
		file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGpuMetricsRequest); i {
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
		file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGpuMetricsResponse); i {
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
		file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGpuPredictionsRequest); i {
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
		file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGpuPredictionsRequest); i {
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
		file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGpuPredictionsResponse); i {
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
			RawDescriptor: file_alameda_api_v1alpha1_datahub_gpu_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_alameda_api_v1alpha1_datahub_gpu_services_proto_goTypes,
		DependencyIndexes: file_alameda_api_v1alpha1_datahub_gpu_services_proto_depIdxs,
		MessageInfos:      file_alameda_api_v1alpha1_datahub_gpu_services_proto_msgTypes,
	}.Build()
	File_alameda_api_v1alpha1_datahub_gpu_services_proto = out.File
	file_alameda_api_v1alpha1_datahub_gpu_services_proto_rawDesc = nil
	file_alameda_api_v1alpha1_datahub_gpu_services_proto_goTypes = nil
	file_alameda_api_v1alpha1_datahub_gpu_services_proto_depIdxs = nil
}
