// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: TestPb/grpcTest.proto

package TestPb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type TestMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TestMsg) Reset() {
	*x = TestMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TestPb_grpcTest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMsg) ProtoMessage() {}

func (x *TestMsg) ProtoReflect() protoreflect.Message {
	mi := &file_TestPb_grpcTest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMsg.ProtoReflect.Descriptor instead.
func (*TestMsg) Descriptor() ([]byte, []int) {
	return file_TestPb_grpcTest_proto_rawDescGZIP(), []int{0}
}

func (x *TestMsg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TestMsgReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ans string `protobuf:"bytes,1,opt,name=ans,proto3" json:"ans,omitempty"`
}

func (x *TestMsgReply) Reset() {
	*x = TestMsgReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TestPb_grpcTest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMsgReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMsgReply) ProtoMessage() {}

func (x *TestMsgReply) ProtoReflect() protoreflect.Message {
	mi := &file_TestPb_grpcTest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMsgReply.ProtoReflect.Descriptor instead.
func (*TestMsgReply) Descriptor() ([]byte, []int) {
	return file_TestPb_grpcTest_proto_rawDescGZIP(), []int{1}
}

func (x *TestMsgReply) GetAns() string {
	if x != nil {
		return x.Ans
	}
	return ""
}

var File_TestPb_grpcTest_proto protoreflect.FileDescriptor

var file_TestPb_grpcTest_proto_rawDesc = []byte{
	0x0a, 0x15, 0x54, 0x65, 0x73, 0x74, 0x50, 0x62, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x54, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x54, 0x65, 0x73, 0x74, 0x50, 0x62, 0x22,
	0x1d, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20,
	0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x6e, 0x73,
	0x32, 0x3f, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x31,
	0x0a, 0x08, 0x67, 0x72, 0x70, 0x63, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x0f, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x50, 0x62, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x14, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x50, 0x62, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x42, 0x08, 0x5a, 0x06, 0x54, 0x65, 0x73, 0x74, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_TestPb_grpcTest_proto_rawDescOnce sync.Once
	file_TestPb_grpcTest_proto_rawDescData = file_TestPb_grpcTest_proto_rawDesc
)

func file_TestPb_grpcTest_proto_rawDescGZIP() []byte {
	file_TestPb_grpcTest_proto_rawDescOnce.Do(func() {
		file_TestPb_grpcTest_proto_rawDescData = protoimpl.X.CompressGZIP(file_TestPb_grpcTest_proto_rawDescData)
	})
	return file_TestPb_grpcTest_proto_rawDescData
}

var file_TestPb_grpcTest_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_TestPb_grpcTest_proto_goTypes = []interface{}{
	(*TestMsg)(nil),      // 0: TestPb.TestMsg
	(*TestMsgReply)(nil), // 1: TestPb.TestMsgReply
}
var file_TestPb_grpcTest_proto_depIdxs = []int32{
	0, // 0: TestPb.TestServer.grpcCall:input_type -> TestPb.TestMsg
	1, // 1: TestPb.TestServer.grpcCall:output_type -> TestPb.TestMsgReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TestPb_grpcTest_proto_init() }
func file_TestPb_grpcTest_proto_init() {
	if File_TestPb_grpcTest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TestPb_grpcTest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMsg); i {
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
		file_TestPb_grpcTest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMsgReply); i {
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
			RawDescriptor: file_TestPb_grpcTest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_TestPb_grpcTest_proto_goTypes,
		DependencyIndexes: file_TestPb_grpcTest_proto_depIdxs,
		MessageInfos:      file_TestPb_grpcTest_proto_msgTypes,
	}.Build()
	File_TestPb_grpcTest_proto = out.File
	file_TestPb_grpcTest_proto_rawDesc = nil
	file_TestPb_grpcTest_proto_goTypes = nil
	file_TestPb_grpcTest_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TestServerClient is the client API for TestServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServerClient interface {
	GrpcCall(ctx context.Context, in *TestMsg, opts ...grpc.CallOption) (*TestMsgReply, error)
}

type testServerClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServerClient(cc grpc.ClientConnInterface) TestServerClient {
	return &testServerClient{cc}
}

func (c *testServerClient) GrpcCall(ctx context.Context, in *TestMsg, opts ...grpc.CallOption) (*TestMsgReply, error) {
	out := new(TestMsgReply)
	err := c.cc.Invoke(ctx, "/TestPb.TestServer/grpcCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServerServer is the server API for TestServer service.
type TestServerServer interface {
	GrpcCall(context.Context, *TestMsg) (*TestMsgReply, error)
}

// UnimplementedTestServerServer can be embedded to have forward compatible implementations.
type UnimplementedTestServerServer struct {
}

func (*UnimplementedTestServerServer) GrpcCall(context.Context, *TestMsg) (*TestMsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrpcCall not implemented")
}

func RegisterTestServerServer(s *grpc.Server, srv TestServerServer) {
	s.RegisterService(&_TestServer_serviceDesc, srv)
}

func _TestServer_GrpcCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServerServer).GrpcCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TestPb.TestServer/GrpcCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServerServer).GrpcCall(ctx, req.(*TestMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TestPb.TestServer",
	HandlerType: (*TestServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "grpcCall",
			Handler:    _TestServer_GrpcCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "TestPb/grpcTest.proto",
}