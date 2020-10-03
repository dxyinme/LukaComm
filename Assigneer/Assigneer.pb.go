// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: Assigneer/Assigneer.proto

package Assigneer

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

type SyncLocationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UidHash uint32 `protobuf:"varint,1,opt,name=uidHash,proto3" json:"uidHash,omitempty"`
}

func (x *SyncLocationReq) Reset() {
	*x = SyncLocationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Assigneer_Assigneer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncLocationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncLocationReq) ProtoMessage() {}

func (x *SyncLocationReq) ProtoReflect() protoreflect.Message {
	mi := &file_Assigneer_Assigneer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncLocationReq.ProtoReflect.Descriptor instead.
func (*SyncLocationReq) Descriptor() ([]byte, []int) {
	return file_Assigneer_Assigneer_proto_rawDescGZIP(), []int{0}
}

func (x *SyncLocationReq) GetUidHash() uint32 {
	if x != nil {
		return x.UidHash
	}
	return 0
}

type SyncLocationRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the location server name
	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *SyncLocationRsp) Reset() {
	*x = SyncLocationRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Assigneer_Assigneer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncLocationRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncLocationRsp) ProtoMessage() {}

func (x *SyncLocationRsp) ProtoReflect() protoreflect.Message {
	mi := &file_Assigneer_Assigneer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncLocationRsp.ProtoReflect.Descriptor instead.
func (*SyncLocationRsp) Descriptor() ([]byte, []int) {
	return file_Assigneer_Assigneer_proto_rawDescGZIP(), []int{1}
}

func (x *SyncLocationRsp) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

var File_Assigneer_Assigneer_proto protoreflect.FileDescriptor

var file_Assigneer_Assigneer_proto_rawDesc = []byte{
	0x0a, 0x19, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x72, 0x2f, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x41, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x0f, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x69, 0x64,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x75, 0x69, 0x64, 0x48,
	0x61, 0x73, 0x68, 0x22, 0x2d, 0x0a, 0x0f, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x32, 0x53, 0x0a, 0x09, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x72, 0x12,
	0x46, 0x0a, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Assigneer_Assigneer_proto_rawDescOnce sync.Once
	file_Assigneer_Assigneer_proto_rawDescData = file_Assigneer_Assigneer_proto_rawDesc
)

func file_Assigneer_Assigneer_proto_rawDescGZIP() []byte {
	file_Assigneer_Assigneer_proto_rawDescOnce.Do(func() {
		file_Assigneer_Assigneer_proto_rawDescData = protoimpl.X.CompressGZIP(file_Assigneer_Assigneer_proto_rawDescData)
	})
	return file_Assigneer_Assigneer_proto_rawDescData
}

var file_Assigneer_Assigneer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Assigneer_Assigneer_proto_goTypes = []interface{}{
	(*SyncLocationReq)(nil), // 0: Assigneer.SyncLocationReq
	(*SyncLocationRsp)(nil), // 1: Assigneer.SyncLocationRsp
}
var file_Assigneer_Assigneer_proto_depIdxs = []int32{
	0, // 0: Assigneer.Assigneer.SyncLocation:input_type -> Assigneer.SyncLocationReq
	1, // 1: Assigneer.Assigneer.SyncLocation:output_type -> Assigneer.SyncLocationRsp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Assigneer_Assigneer_proto_init() }
func file_Assigneer_Assigneer_proto_init() {
	if File_Assigneer_Assigneer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Assigneer_Assigneer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncLocationReq); i {
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
		file_Assigneer_Assigneer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncLocationRsp); i {
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
			RawDescriptor: file_Assigneer_Assigneer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Assigneer_Assigneer_proto_goTypes,
		DependencyIndexes: file_Assigneer_Assigneer_proto_depIdxs,
		MessageInfos:      file_Assigneer_Assigneer_proto_msgTypes,
	}.Build()
	File_Assigneer_Assigneer_proto = out.File
	file_Assigneer_Assigneer_proto_rawDesc = nil
	file_Assigneer_Assigneer_proto_goTypes = nil
	file_Assigneer_Assigneer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AssigneerClient is the client API for Assigneer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AssigneerClient interface {
	SyncLocation(ctx context.Context, in *SyncLocationReq, opts ...grpc.CallOption) (*SyncLocationRsp, error)
}

type assigneerClient struct {
	cc grpc.ClientConnInterface
}

func NewAssigneerClient(cc grpc.ClientConnInterface) AssigneerClient {
	return &assigneerClient{cc}
}

func (c *assigneerClient) SyncLocation(ctx context.Context, in *SyncLocationReq, opts ...grpc.CallOption) (*SyncLocationRsp, error) {
	out := new(SyncLocationRsp)
	err := c.cc.Invoke(ctx, "/Assigneer.Assigneer/SyncLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssigneerServer is the server API for Assigneer service.
type AssigneerServer interface {
	SyncLocation(context.Context, *SyncLocationReq) (*SyncLocationRsp, error)
}

// UnimplementedAssigneerServer can be embedded to have forward compatible implementations.
type UnimplementedAssigneerServer struct {
}

func (*UnimplementedAssigneerServer) SyncLocation(context.Context, *SyncLocationReq) (*SyncLocationRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncLocation not implemented")
}

func RegisterAssigneerServer(s *grpc.Server, srv AssigneerServer) {
	s.RegisterService(&_Assigneer_serviceDesc, srv)
}

func _Assigneer_SyncLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncLocationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssigneerServer).SyncLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Assigneer.Assigneer/SyncLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssigneerServer).SyncLocation(ctx, req.(*SyncLocationReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Assigneer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Assigneer.Assigneer",
	HandlerType: (*AssigneerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncLocation",
			Handler:    _Assigneer_SyncLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Assigneer/Assigneer.proto",
}
