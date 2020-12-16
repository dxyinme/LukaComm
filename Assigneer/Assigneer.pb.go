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
	Hosts     []string `protobuf:"bytes,2,rep,name=hosts,proto3" json:"hosts,omitempty"`
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

func (x *SyncLocationRsp) GetHosts() []string {
	if x != nil {
		return x.Hosts
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
	Host     string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Pid      string `protobuf:"bytes,3,opt,name=pid,proto3" json:"pid,omitempty"`
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

func (x *AddKeeperReq) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *AddKeeperReq) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
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

type SwitchKeeperReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *SwitchKeeperReq) Reset() {
	*x = SwitchKeeperReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Assigneer_Assigneer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwitchKeeperReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwitchKeeperReq) ProtoMessage() {}

func (x *SwitchKeeperReq) ProtoReflect() protoreflect.Message {
	mi := &file_Assigneer_Assigneer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwitchKeeperReq.ProtoReflect.Descriptor instead.
func (*SwitchKeeperReq) Descriptor() ([]byte, []int) {
	return file_Assigneer_Assigneer_proto_rawDescGZIP(), []int{5}
}

func (x *SwitchKeeperReq) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type SwitchKeeperRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeeperID uint32 `protobuf:"varint,1,opt,name=keeperID,proto3" json:"keeperID,omitempty"`
	Host     string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *SwitchKeeperRsp) Reset() {
	*x = SwitchKeeperRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Assigneer_Assigneer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwitchKeeperRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwitchKeeperRsp) ProtoMessage() {}

func (x *SwitchKeeperRsp) ProtoReflect() protoreflect.Message {
	mi := &file_Assigneer_Assigneer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwitchKeeperRsp.ProtoReflect.Descriptor instead.
func (*SwitchKeeperRsp) Descriptor() ([]byte, []int) {
	return file_Assigneer_Assigneer_proto_rawDescGZIP(), []int{6}
}

func (x *SwitchKeeperRsp) GetKeeperID() uint32 {
	if x != nil {
		return x.KeeperID
	}
	return 0
}

func (x *SwitchKeeperRsp) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type ClusterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperatorID uint32 `protobuf:"varint,1,opt,name=operatorID,proto3" json:"operatorID,omitempty"`
	ReqInfo    []byte `protobuf:"bytes,2,opt,name=reqInfo,proto3" json:"reqInfo,omitempty"`
}

func (x *ClusterReq) Reset() {
	*x = ClusterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Assigneer_Assigneer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterReq) ProtoMessage() {}

func (x *ClusterReq) ProtoReflect() protoreflect.Message {
	mi := &file_Assigneer_Assigneer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterReq.ProtoReflect.Descriptor instead.
func (*ClusterReq) Descriptor() ([]byte, []int) {
	return file_Assigneer_Assigneer_proto_rawDescGZIP(), []int{7}
}

func (x *ClusterReq) GetOperatorID() uint32 {
	if x != nil {
		return x.OperatorID
	}
	return 0
}

func (x *ClusterReq) GetReqInfo() []byte {
	if x != nil {
		return x.ReqInfo
	}
	return nil
}

type ClusterRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperatorID uint32 `protobuf:"varint,1,opt,name=operatorID,proto3" json:"operatorID,omitempty"`
	RspInfo    []byte `protobuf:"bytes,2,opt,name=rspInfo,proto3" json:"rspInfo,omitempty"`
}

func (x *ClusterRsp) Reset() {
	*x = ClusterRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Assigneer_Assigneer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterRsp) ProtoMessage() {}

func (x *ClusterRsp) ProtoReflect() protoreflect.Message {
	mi := &file_Assigneer_Assigneer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterRsp.ProtoReflect.Descriptor instead.
func (*ClusterRsp) Descriptor() ([]byte, []int) {
	return file_Assigneer_Assigneer_proto_rawDescGZIP(), []int{8}
}

func (x *ClusterRsp) GetOperatorID() uint32 {
	if x != nil {
		return x.OperatorID
	}
	return 0
}

func (x *ClusterRsp) GetRspInfo() []byte {
	if x != nil {
		return x.RspInfo
	}
	return nil
}

type RegisterNodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip  string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Pwd string `protobuf:"bytes,2,opt,name=pwd,proto3" json:"pwd,omitempty"`
}

func (x *RegisterNodeReq) Reset() {
	*x = RegisterNodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Assigneer_Assigneer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterNodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNodeReq) ProtoMessage() {}

func (x *RegisterNodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_Assigneer_Assigneer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterNodeReq.ProtoReflect.Descriptor instead.
func (*RegisterNodeReq) Descriptor() ([]byte, []int) {
	return file_Assigneer_Assigneer_proto_rawDescGZIP(), []int{9}
}

func (x *RegisterNodeReq) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *RegisterNodeReq) GetPwd() string {
	if x != nil {
		return x.Pwd
	}
	return ""
}

type RegisterNodeRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegisterInfo string `protobuf:"bytes,1,opt,name=registerInfo,proto3" json:"registerInfo,omitempty"`
}

func (x *RegisterNodeRsp) Reset() {
	*x = RegisterNodeRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Assigneer_Assigneer_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterNodeRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNodeRsp) ProtoMessage() {}

func (x *RegisterNodeRsp) ProtoReflect() protoreflect.Message {
	mi := &file_Assigneer_Assigneer_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterNodeRsp.ProtoReflect.Descriptor instead.
func (*RegisterNodeRsp) Descriptor() ([]byte, []int) {
	return file_Assigneer_Assigneer_proto_rawDescGZIP(), []int{10}
}

func (x *RegisterNodeRsp) GetRegisterInfo() string {
	if x != nil {
		return x.RegisterInfo
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
	0x70, 0x65, 0x72, 0x49, 0x44, 0x22, 0x45, 0x0a, 0x0f, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x6b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x49, 0x44, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x2d, 0x0a, 0x0f,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12,
	0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x49, 0x44, 0x22, 0x50, 0x0a, 0x0c, 0x41,
	0x64, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x2b, 0x0a,
	0x09, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x41, 0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63,
	0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x23, 0x0a, 0x0f, 0x53, 0x77,
	0x69, 0x74, 0x63, 0x68, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22,
	0x41, 0x0a, 0x0f, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x52,
	0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x22, 0x46, 0x0a, 0x0a, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x44,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x72, 0x65, 0x71, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x46, 0x0a, 0x0a, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x73, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x73, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x33, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x77, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x70, 0x77, 0x64, 0x22, 0x35, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x9f,
	0x03, 0x0a, 0x09, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x0c,
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
	0x63, 0x6b, 0x12, 0x46, 0x0a, 0x0c, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x4b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x12, 0x1a, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x53,
	0x77, 0x69, 0x74, 0x63, 0x68, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1a,
	0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x53, 0x77, 0x69, 0x74, 0x63,
	0x68, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x0c, 0x4d, 0x61,
	0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x2e, 0x41, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x15, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x46, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x73, 0x70,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_Assigneer_Assigneer_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_Assigneer_Assigneer_proto_goTypes = []interface{}{
	(*SyncLocationReq)(nil), // 0: Assigneer.SyncLocationReq
	(*SyncLocationRsp)(nil), // 1: Assigneer.SyncLocationRsp
	(*RemoveKeeperReq)(nil), // 2: Assigneer.RemoveKeeperReq
	(*AddKeeperReq)(nil),    // 3: Assigneer.AddKeeperReq
	(*AssignAck)(nil),       // 4: Assigneer.AssignAck
	(*SwitchKeeperReq)(nil), // 5: Assigneer.SwitchKeeperReq
	(*SwitchKeeperRsp)(nil), // 6: Assigneer.SwitchKeeperRsp
	(*ClusterReq)(nil),      // 7: Assigneer.ClusterReq
	(*ClusterRsp)(nil),      // 8: Assigneer.ClusterRsp
	(*RegisterNodeReq)(nil), // 9: Assigneer.RegisterNodeReq
	(*RegisterNodeRsp)(nil), // 10: Assigneer.RegisterNodeRsp
}
var file_Assigneer_Assigneer_proto_depIdxs = []int32{
	0,  // 0: Assigneer.Assigneer.SyncLocation:input_type -> Assigneer.SyncLocationReq
	2,  // 1: Assigneer.Assigneer.RemoveKeeper:input_type -> Assigneer.RemoveKeeperReq
	3,  // 2: Assigneer.Assigneer.AddKeeper:input_type -> Assigneer.AddKeeperReq
	5,  // 3: Assigneer.Assigneer.SwitchKeeper:input_type -> Assigneer.SwitchKeeperReq
	7,  // 4: Assigneer.Assigneer.MaintainInfo:input_type -> Assigneer.ClusterReq
	9,  // 5: Assigneer.Assigneer.RegisterNode:input_type -> Assigneer.RegisterNodeReq
	1,  // 6: Assigneer.Assigneer.SyncLocation:output_type -> Assigneer.SyncLocationRsp
	4,  // 7: Assigneer.Assigneer.RemoveKeeper:output_type -> Assigneer.AssignAck
	4,  // 8: Assigneer.Assigneer.AddKeeper:output_type -> Assigneer.AssignAck
	6,  // 9: Assigneer.Assigneer.SwitchKeeper:output_type -> Assigneer.SwitchKeeperRsp
	8,  // 10: Assigneer.Assigneer.MaintainInfo:output_type -> Assigneer.ClusterRsp
	10, // 11: Assigneer.Assigneer.RegisterNode:output_type -> Assigneer.RegisterNodeRsp
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
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
		file_Assigneer_Assigneer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwitchKeeperReq); i {
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
		file_Assigneer_Assigneer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwitchKeeperRsp); i {
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
		file_Assigneer_Assigneer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterReq); i {
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
		file_Assigneer_Assigneer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterRsp); i {
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
		file_Assigneer_Assigneer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterNodeReq); i {
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
		file_Assigneer_Assigneer_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterNodeRsp); i {
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
			NumMessages:   11,
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
	// sync the keeperIDs and Hosts
	// prepare for keeper
	SyncLocation(ctx context.Context, in *SyncLocationReq, opts ...grpc.CallOption) (*SyncLocationRsp, error)
	RemoveKeeper(ctx context.Context, in *RemoveKeeperReq, opts ...grpc.CallOption) (*AssignAck, error)
	AddKeeper(ctx context.Context, in *AddKeeperReq, opts ...grpc.CallOption) (*AssignAck, error)
	// prepare for client
	SwitchKeeper(ctx context.Context, in *SwitchKeeperReq, opts ...grpc.CallOption) (*SwitchKeeperRsp, error)
	// maintainInfo of each keeper.
	MaintainInfo(ctx context.Context, in *ClusterReq, opts ...grpc.CallOption) (*ClusterRsp, error)
	// registerNode to assigneerServer
	RegisterNode(ctx context.Context, in *RegisterNodeReq, opts ...grpc.CallOption) (*RegisterNodeRsp, error)
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

func (c *assigneerClient) SwitchKeeper(ctx context.Context, in *SwitchKeeperReq, opts ...grpc.CallOption) (*SwitchKeeperRsp, error) {
	out := new(SwitchKeeperRsp)
	err := c.cc.Invoke(ctx, "/Assigneer.Assigneer/SwitchKeeper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assigneerClient) MaintainInfo(ctx context.Context, in *ClusterReq, opts ...grpc.CallOption) (*ClusterRsp, error) {
	out := new(ClusterRsp)
	err := c.cc.Invoke(ctx, "/Assigneer.Assigneer/MaintainInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assigneerClient) RegisterNode(ctx context.Context, in *RegisterNodeReq, opts ...grpc.CallOption) (*RegisterNodeRsp, error) {
	out := new(RegisterNodeRsp)
	err := c.cc.Invoke(ctx, "/Assigneer.Assigneer/RegisterNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssigneerServer is the server API for Assigneer service.
type AssigneerServer interface {
	// sync the keeperIDs and Hosts
	// prepare for keeper
	SyncLocation(context.Context, *SyncLocationReq) (*SyncLocationRsp, error)
	RemoveKeeper(context.Context, *RemoveKeeperReq) (*AssignAck, error)
	AddKeeper(context.Context, *AddKeeperReq) (*AssignAck, error)
	// prepare for client
	SwitchKeeper(context.Context, *SwitchKeeperReq) (*SwitchKeeperRsp, error)
	// maintainInfo of each keeper.
	MaintainInfo(context.Context, *ClusterReq) (*ClusterRsp, error)
	// registerNode to assigneerServer
	RegisterNode(context.Context, *RegisterNodeReq) (*RegisterNodeRsp, error)
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
func (*UnimplementedAssigneerServer) SwitchKeeper(context.Context, *SwitchKeeperReq) (*SwitchKeeperRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwitchKeeper not implemented")
}
func (*UnimplementedAssigneerServer) MaintainInfo(context.Context, *ClusterReq) (*ClusterRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MaintainInfo not implemented")
}
func (*UnimplementedAssigneerServer) RegisterNode(context.Context, *RegisterNodeReq) (*RegisterNodeRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNode not implemented")
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

func _Assigneer_SwitchKeeper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SwitchKeeperReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssigneerServer).SwitchKeeper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Assigneer.Assigneer/SwitchKeeper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssigneerServer).SwitchKeeper(ctx, req.(*SwitchKeeperReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assigneer_MaintainInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssigneerServer).MaintainInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Assigneer.Assigneer/MaintainInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssigneerServer).MaintainInfo(ctx, req.(*ClusterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assigneer_RegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssigneerServer).RegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Assigneer.Assigneer/RegisterNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssigneerServer).RegisterNode(ctx, req.(*RegisterNodeReq))
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
		{
			MethodName: "SwitchKeeper",
			Handler:    _Assigneer_SwitchKeeper_Handler,
		},
		{
			MethodName: "MaintainInfo",
			Handler:    _Assigneer_MaintainInfo_Handler,
		},
		{
			MethodName: "RegisterNode",
			Handler:    _Assigneer_RegisterNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Assigneer/Assigneer.proto",
}
