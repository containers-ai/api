// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ai-service/v1alpha1/ai_service.proto

package ai_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Object_Type int32

const (
	Object_POD Object_Type = 0
)

var Object_Type_name = map[int32]string{
	0: "POD",
}
var Object_Type_value = map[string]int32{
	"POD": 0,
}

func (x Object_Type) String() string {
	return proto.EnumName(Object_Type_name, int32(x))
}
func (Object_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ai_service_b78ec710570d78c2, []int{0, 0}
}

type Object struct {
	Type                 Object_Type `protobuf:"varint,1,opt,name=type,proto3,enum=Object_Type" json:"type,omitempty"`
	Uid                  string      `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Namespace            string      `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string      `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_ai_service_b78ec710570d78c2, []int{0}
}
func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (dst *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(dst, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetType() Object_Type {
	if m != nil {
		return m.Type
	}
	return Object_POD
}

func (m *Object) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Object) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Object) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PredictionObjectListCreationRequest struct {
	Objects              []*Object `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PredictionObjectListCreationRequest) Reset()         { *m = PredictionObjectListCreationRequest{} }
func (m *PredictionObjectListCreationRequest) String() string { return proto.CompactTextString(m) }
func (*PredictionObjectListCreationRequest) ProtoMessage()    {}
func (*PredictionObjectListCreationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ai_service_b78ec710570d78c2, []int{1}
}
func (m *PredictionObjectListCreationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictionObjectListCreationRequest.Unmarshal(m, b)
}
func (m *PredictionObjectListCreationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictionObjectListCreationRequest.Marshal(b, m, deterministic)
}
func (dst *PredictionObjectListCreationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictionObjectListCreationRequest.Merge(dst, src)
}
func (m *PredictionObjectListCreationRequest) XXX_Size() int {
	return xxx_messageInfo_PredictionObjectListCreationRequest.Size(m)
}
func (m *PredictionObjectListCreationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictionObjectListCreationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PredictionObjectListCreationRequest proto.InternalMessageInfo

func (m *PredictionObjectListCreationRequest) GetObjects() []*Object {
	if m != nil {
		return m.Objects
	}
	return nil
}

type PredictionObjectListDeletionRequest struct {
	Objects              []*Object `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PredictionObjectListDeletionRequest) Reset()         { *m = PredictionObjectListDeletionRequest{} }
func (m *PredictionObjectListDeletionRequest) String() string { return proto.CompactTextString(m) }
func (*PredictionObjectListDeletionRequest) ProtoMessage()    {}
func (*PredictionObjectListDeletionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ai_service_b78ec710570d78c2, []int{2}
}
func (m *PredictionObjectListDeletionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictionObjectListDeletionRequest.Unmarshal(m, b)
}
func (m *PredictionObjectListDeletionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictionObjectListDeletionRequest.Marshal(b, m, deterministic)
}
func (dst *PredictionObjectListDeletionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictionObjectListDeletionRequest.Merge(dst, src)
}
func (m *PredictionObjectListDeletionRequest) XXX_Size() int {
	return xxx_messageInfo_PredictionObjectListDeletionRequest.Size(m)
}
func (m *PredictionObjectListDeletionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictionObjectListDeletionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PredictionObjectListDeletionRequest proto.InternalMessageInfo

func (m *PredictionObjectListDeletionRequest) GetObjects() []*Object {
	if m != nil {
		return m.Objects
	}
	return nil
}

type RequestResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestResponse) Reset()         { *m = RequestResponse{} }
func (m *RequestResponse) String() string { return proto.CompactTextString(m) }
func (*RequestResponse) ProtoMessage()    {}
func (*RequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ai_service_b78ec710570d78c2, []int{3}
}
func (m *RequestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestResponse.Unmarshal(m, b)
}
func (m *RequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestResponse.Marshal(b, m, deterministic)
}
func (dst *RequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestResponse.Merge(dst, src)
}
func (m *RequestResponse) XXX_Size() int {
	return xxx_messageInfo_RequestResponse.Size(m)
}
func (m *RequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RequestResponse proto.InternalMessageInfo

func (m *RequestResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Object)(nil), "Object")
	proto.RegisterType((*PredictionObjectListCreationRequest)(nil), "PredictionObjectListCreationRequest")
	proto.RegisterType((*PredictionObjectListDeletionRequest)(nil), "PredictionObjectListDeletionRequest")
	proto.RegisterType((*RequestResponse)(nil), "RequestResponse")
	proto.RegisterEnum("Object_Type", Object_Type_name, Object_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AlamendaAIServiceClient is the client API for AlamendaAIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlamendaAIServiceClient interface {
	CreatePredictionObjects(ctx context.Context, in *PredictionObjectListCreationRequest, opts ...grpc.CallOption) (*RequestResponse, error)
	DeletePredictionObjects(ctx context.Context, in *PredictionObjectListDeletionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type alamendaAIServiceClient struct {
	cc *grpc.ClientConn
}

func NewAlamendaAIServiceClient(cc *grpc.ClientConn) AlamendaAIServiceClient {
	return &alamendaAIServiceClient{cc}
}

func (c *alamendaAIServiceClient) CreatePredictionObjects(ctx context.Context, in *PredictionObjectListCreationRequest, opts ...grpc.CallOption) (*RequestResponse, error) {
	out := new(RequestResponse)
	err := c.cc.Invoke(ctx, "/AlamendaAIService/CreatePredictionObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alamendaAIServiceClient) DeletePredictionObjects(ctx context.Context, in *PredictionObjectListDeletionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/AlamendaAIService/DeletePredictionObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlamendaAIServiceServer is the server API for AlamendaAIService service.
type AlamendaAIServiceServer interface {
	CreatePredictionObjects(context.Context, *PredictionObjectListCreationRequest) (*RequestResponse, error)
	DeletePredictionObjects(context.Context, *PredictionObjectListDeletionRequest) (*empty.Empty, error)
}

func RegisterAlamendaAIServiceServer(s *grpc.Server, srv AlamendaAIServiceServer) {
	s.RegisterService(&_AlamendaAIService_serviceDesc, srv)
}

func _AlamendaAIService_CreatePredictionObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictionObjectListCreationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlamendaAIServiceServer).CreatePredictionObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AlamendaAIService/CreatePredictionObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlamendaAIServiceServer).CreatePredictionObjects(ctx, req.(*PredictionObjectListCreationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlamendaAIService_DeletePredictionObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictionObjectListDeletionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlamendaAIServiceServer).DeletePredictionObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AlamendaAIService/DeletePredictionObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlamendaAIServiceServer).DeletePredictionObjects(ctx, req.(*PredictionObjectListDeletionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlamendaAIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AlamendaAIService",
	HandlerType: (*AlamendaAIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePredictionObjects",
			Handler:    _AlamendaAIService_CreatePredictionObjects_Handler,
		},
		{
			MethodName: "DeletePredictionObjects",
			Handler:    _AlamendaAIService_DeletePredictionObjects_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ai-service/v1alpha1/ai_service.proto",
}

func init() {
	proto.RegisterFile("ai-service/v1alpha1/ai_service.proto", fileDescriptor_ai_service_b78ec710570d78c2)
}

var fileDescriptor_ai_service_b78ec710570d78c2 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0x9b, 0x7f, 0x42, 0x43, 0xe7, 0x2f, 0xb6, 0xee, 0xc1, 0x86, 0xea, 0x21, 0xc6, 0x1e,
	0x0a, 0xe2, 0x86, 0xd6, 0x27, 0x28, 0x56, 0x50, 0x10, 0x5a, 0x52, 0x2f, 0x9e, 0x64, 0x9b, 0x8e,
	0x75, 0x25, 0xc9, 0xae, 0xd9, 0x6d, 0xa1, 0x27, 0x1f, 0xce, 0x17, 0x93, 0xec, 0xb6, 0x08, 0x01,
	0xa1, 0xde, 0x66, 0x7e, 0xdf, 0xf2, 0xed, 0x7c, 0xc3, 0x40, 0x9f, 0xf1, 0x6b, 0x85, 0xe5, 0x86,
	0xa7, 0x18, 0x6f, 0x86, 0x2c, 0x93, 0x6f, 0x6c, 0x18, 0x33, 0xfe, 0xb2, 0x63, 0x54, 0x96, 0x42,
	0x8b, 0xde, 0xd9, 0x4a, 0x88, 0x55, 0x86, 0xb1, 0xe9, 0x16, 0xeb, 0xd7, 0x18, 0x73, 0xa9, 0xb7,
	0x56, 0x8c, 0x3e, 0xa1, 0x39, 0x5d, 0xbc, 0x63, 0xaa, 0x49, 0x08, 0x9e, 0xde, 0x4a, 0x0c, 0x9c,
	0xd0, 0x19, 0x1c, 0x8f, 0x8e, 0xa8, 0xc5, 0xf4, 0x69, 0x2b, 0x31, 0x31, 0x0a, 0xe9, 0x80, 0xbb,
	0xe6, 0xcb, 0xe0, 0x5f, 0xe8, 0x0c, 0x5a, 0x49, 0x55, 0x92, 0x73, 0x68, 0x15, 0x2c, 0x47, 0x25,
	0x59, 0x8a, 0x81, 0x6b, 0xf8, 0x0f, 0x20, 0x04, 0xbc, 0xaa, 0x09, 0x3c, 0x23, 0x98, 0x3a, 0x6a,
	0x83, 0x57, 0x39, 0x12, 0x1f, 0xdc, 0xd9, 0x74, 0xd2, 0x69, 0x44, 0xf7, 0x70, 0x39, 0x2b, 0x71,
	0xc9, 0x53, 0xcd, 0x45, 0x61, 0xff, 0x7c, 0xe4, 0x4a, 0xdf, 0x96, 0xc8, 0x2a, 0x92, 0xe0, 0xc7,
	0x1a, 0x95, 0x26, 0x17, 0xe0, 0x0b, 0x23, 0xaa, 0xc0, 0x09, 0xdd, 0xc1, 0xff, 0x91, 0xbf, 0x1b,
	0x30, 0xd9, 0xf3, 0xdf, 0x9c, 0x26, 0x98, 0xe1, 0x1f, 0x9d, 0xae, 0xa0, 0xbd, 0x7b, 0x9d, 0xa0,
	0x92, 0xa2, 0x50, 0x48, 0x02, 0xf0, 0x73, 0x54, 0x8a, 0xad, 0xec, 0x82, 0x5a, 0xc9, 0xbe, 0x1d,
	0x7d, 0x39, 0x70, 0x32, 0xce, 0x58, 0x8e, 0xc5, 0x92, 0x8d, 0x1f, 0xe6, 0x76, 0xf5, 0x64, 0x0e,
	0x5d, 0x13, 0x01, 0xeb, 0x23, 0x29, 0xd2, 0xa7, 0x07, 0x04, 0xee, 0x75, 0x68, 0x6d, 0x84, 0xa8,
	0x41, 0x9e, 0xa1, 0x6b, 0xd2, 0x1c, 0x6c, 0x5a, 0xcb, 0xde, 0x3b, 0xa5, 0xf6, 0x16, 0xe8, 0xfe,
	0x16, 0xe8, 0x5d, 0x75, 0x0b, 0x51, 0x63, 0xd1, 0x34, 0xe4, 0xe6, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0xbf, 0xf6, 0xf4, 0xbd, 0x53, 0x02, 0x00, 0x00,
}
