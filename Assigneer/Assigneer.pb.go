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

	KeeperID uint32 `protobuf:"varint,1,opt,name=keeperID,proto3" json:"keeperID,omitempty"`
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

func (x *SyncLocationReq) GetKeeperID() uint32 {
	if x != nil {
		return x.KeeperID
	}
	return 0
}

type SyncLocationRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeeperIDs []uint32 `protobuf:"varint,1,rep,packed,name=keeperIDs,proto3" json:"keeperIDs,omitempty"`
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

func (x *SyncLocationRsp) GetKeeperIDs() []uint32 {
	if x != nil {
		return x.KeeperIDs
	}
	return nil
}

type RemoveKeeperReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeeperID uint32 `protobuf:"varint,1,opt,name=keeperID,proto3" json:"keeperID,omitempty"`
}

func (x *RemoveKeeperReq) Reset() {
	*x = RemoveKeeperReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Assigneer_Assigneer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveKeeperReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveKeeperReq) ProtoMessage() {}

func (x *RemoveKeeperReq) ProtoReflect() protoreflect.Message {
	mi := &file_Assigneer_Assigneer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveKeeperReq.ProtoReflect.Descriptor instead.
func (*RemoveKeeperReq) Descriptor() ([]byte, []int) {
	return file_Assigneer_Assigneer_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveKeeperReq) GetKeeperID() uint32 {
	if x != nil {
		return x.KeeperID
	}
	return 0
}

type AddKeeperReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeeperID uint32 `protobuf:"varint,1,opt,name=keeperID,proto3" json:"keeperID,omitempty"`
}

func (x *AddKeeperReq) Reset() {
	*x = AddKeeperReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Assigneer_Assigneer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddKeeperReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddKeeperReq) ProtoMessage() {}

func (x *AddKeeperReq) ProtoReflect() protoreflect.Message {
	mi := &file_Assigneer_Assigneer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddKeeperReq.ProtoReflect.Descriptor instead.
func (*AddKeeperReq) Descriptor() ([]byte, []int) {
	return file_Assigneer_Assigneer_proto_rawDescGZIP(), []int{3}
}

func (x *AddKeeperReq) GetKeeperID() uint32 {
	if x != nil {
		return x.KeeperID
	}
	return 0
}

type AssignAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AckMessage string `protobuf:"bytes,1,opt,name=ackMessage,proto3" json:"ackMessage,omitempty"`
}

func (x *AssignAck) Reset() {
	*x = AssignAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Assigneer_Assigneer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignAck) ProtoMessage() {}

func (x *AssignAck) ProtoReflect() protoreflect.Message {
	mi := &file_Assigneer_Assigneer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignAck.ProtoReflect.Descriptor instead.
func (*AssignAck) Descriptor() ([]byte, []int) {
	return file_Assigneer_Assigneer_proto_rawDescGZIP(), []int{4}
}

func (x *AssignAck) GetAckMessage() string {
	if x != nil {
		return x.AckMessage
	}
	return ""
}

var File_Assigneer_Assigneer_proto protoreflect.FileDescriptor

var file_Assigneer_Assigneer_proto_rawDesc = []byte{
	0x0a, 0x19, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x72, 0x2f, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x41, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x0f, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x49, 0x44, 0x22, 0x2f, 0x0a, 0x0f, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x6b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x49, 0x44, 0x73, 0x22, 0x2d, 0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x49, 0x44, 0x22, 0x2a, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x4b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x49,
	0x44, 0x22, 0x2b, 0x0a, 0x09, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x41, 0x63, 0x6b, 0x12, 0x1e,
	0x0a, 0x0a, 0x61, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd1,
	0x01, 0x0a, 0x09, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x0c,
	0x53, 0x79, 0x6e, 0x63, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x73, 0x70, 0x12, 0x40, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4b, 0x65,
	0x65, 0x70, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x14, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x41, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x41, 0x63, 0x6b, 0x12, 0x3a, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x4b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x72, 0x2e,
	0x41, 0x64, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x41,
	0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_Assigneer_Assigneer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Assigneer_Assigneer_proto_goTypes = []interface{}{
	(*SyncLocationReq)(nil), // 0: Assigneer.SyncLocationReq
	(*SyncLocationRsp)(nil), // 1: Assigneer.SyncLocationRsp
	(*RemoveKeeperReq)(nil), // 2: Assigneer.RemoveKeeperReq
	(*AddKeeperReq)(nil),    // 3: Assigneer.AddKeeperReq
	(*AssignAck)(nil),       // 4: Assigneer.AssignAck
}
var file_Assigneer_Assigneer_proto_depIdxs = []int32{
	0, // 0: Assigneer.Assigneer.SyncLocation:input_type -> Assigneer.SyncLocationReq
	2, // 1: Assigneer.Assigneer.RemoveKeeper:input_type -> Assigneer.RemoveKeeperReq
	3, // 2: Assigneer.Assigneer.AddKeeper:input_type -> Assigneer.AddKeeperReq
	1, // 3: Assigneer.Assigneer.SyncLocation:output_type -> Assigneer.SyncLocationRsp
	4, // 4: Assigneer.Assigneer.RemoveKeeper:output_type -> Assigneer.AssignAck
	4, // 5: Assigneer.Assigneer.AddKeeper:output_type -> Assigneer.AssignAck
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
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
		file_Assigneer_Assigneer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveKeeperReq); i {
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
		file_Assigneer_Assigneer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddKeeperReq); i {
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
		file_Assigneer_Assigneer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignAck); i {
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
			NumMessages:   5,
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
	RemoveKeeper(ctx context.Context, in *RemoveKeeperReq, opts ...grpc.CallOption) (*AssignAck, error)
	AddKeeper(ctx context.Context, in *AddKeeperReq, opts ...grpc.CallOption) (*AssignAck, error)
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

func (c *assigneerClient) RemoveKeeper(ctx context.Context, in *RemoveKeeperReq, opts ...grpc.CallOption) (*AssignAck, error) {
	out := new(AssignAck)
	err := c.cc.Invoke(ctx, "/Assigneer.Assigneer/RemoveKeeper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assigneerClient) AddKeeper(ctx context.Context, in *AddKeeperReq, opts ...grpc.CallOption) (*AssignAck, error) {
	out := new(AssignAck)
	err := c.cc.Invoke(ctx, "/Assigneer.Assigneer/AddKeeper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssigneerServer is the server API for Assigneer service.
type AssigneerServer interface {
	SyncLocation(context.Context, *SyncLocationReq) (*SyncLocationRsp, error)
	RemoveKeeper(context.Context, *RemoveKeeperReq) (*AssignAck, error)
	AddKeeper(context.Context, *AddKeeperReq) (*AssignAck, error)
}

// UnimplementedAssigneerServer can be embedded to have forward compatible implementations.
type UnimplementedAssigneerServer struct {
}

func (*UnimplementedAssigneerServer) SyncLocation(context.Context, *SyncLocationReq) (*SyncLocationRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncLocation not implemented")
}
func (*UnimplementedAssigneerServer) RemoveKeeper(context.Context, *RemoveKeeperReq) (*AssignAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveKeeper not implemented")
}
func (*UnimplementedAssigneerServer) AddKeeper(context.Context, *AddKeeperReq) (*AssignAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddKeeper not implemented")
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

func _Assigneer_RemoveKeeper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveKeeperReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssigneerServer).RemoveKeeper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Assigneer.Assigneer/RemoveKeeper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssigneerServer).RemoveKeeper(ctx, req.(*RemoveKeeperReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assigneer_AddKeeper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddKeeperReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssigneerServer).AddKeeper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Assigneer.Assigneer/AddKeeper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssigneerServer).AddKeeper(ctx, req.(*AddKeeperReq))
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
		{
			MethodName: "RemoveKeeper",
			Handler:    _Assigneer_RemoveKeeper_Handler,
		},
		{
			MethodName: "AddKeeper",
			Handler:    _Assigneer_AddKeeper_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Assigneer/Assigneer.proto",
}