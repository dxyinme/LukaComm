// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        (unknown)
// source: Auth/Auth.proto

package Auth

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

type UserPassword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserPassword) Reset() {
	*x = UserPassword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Auth_Auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPassword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPassword) ProtoMessage() {}

func (x *UserPassword) ProtoReflect() protoreflect.Message {
	mi := &file_Auth_Auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPassword.ProtoReflect.Descriptor instead.
func (*UserPassword) Descriptor() ([]byte, []int) {
	return file_Auth_Auth_proto_rawDescGZIP(), []int{0}
}

func (x *UserPassword) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *UserPassword) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserGroupInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupNames []string `protobuf:"bytes,1,rep,name=groupNames,proto3" json:"groupNames,omitempty"`
	Name       string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UserGroupInfo) Reset() {
	*x = UserGroupInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Auth_Auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGroupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroupInfo) ProtoMessage() {}

func (x *UserGroupInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Auth_Auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroupInfo.ProtoReflect.Descriptor instead.
func (*UserGroupInfo) Descriptor() ([]byte, []int) {
	return file_Auth_Auth_proto_rawDescGZIP(), []int{1}
}

func (x *UserGroupInfo) GetGroupNames() []string {
	if x != nil {
		return x.GroupNames
	}
	return nil
}

func (x *UserGroupInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserPassword *UserPassword `protobuf:"bytes,1,opt,name=userPassword,proto3" json:"userPassword,omitempty"`
	Name         string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RsaPubKey    []byte        `protobuf:"bytes,3,opt,name=rsaPubKey,proto3" json:"rsaPubKey,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Auth_Auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Auth_Auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_Auth_Auth_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfo) GetUserPassword() *UserPassword {
	if x != nil {
		return x.UserPassword
	}
	return nil
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetRsaPubKey() []byte {
	if x != nil {
		return x.RsaPubKey
	}
	return nil
}

type LoginRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorMsg string `protobuf:"bytes,1,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
}

func (x *LoginRsp) Reset() {
	*x = LoginRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Auth_Auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRsp) ProtoMessage() {}

func (x *LoginRsp) ProtoReflect() protoreflect.Message {
	mi := &file_Auth_Auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRsp.ProtoReflect.Descriptor instead.
func (*LoginRsp) Descriptor() ([]byte, []int) {
	return file_Auth_Auth_proto_rawDescGZIP(), []int{3}
}

func (x *LoginRsp) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

type SignUpRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorMsg string `protobuf:"bytes,1,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
}

func (x *SignUpRsp) Reset() {
	*x = SignUpRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Auth_Auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRsp) ProtoMessage() {}

func (x *SignUpRsp) ProtoReflect() protoreflect.Message {
	mi := &file_Auth_Auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRsp.ProtoReflect.Descriptor instead.
func (*SignUpRsp) Descriptor() ([]byte, []int) {
	return file_Auth_Auth_proto_rawDescGZIP(), []int{4}
}

func (x *SignUpRsp) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

type ChangeInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorMsg string `protobuf:"bytes,1,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
}

func (x *ChangeInfoRsp) Reset() {
	*x = ChangeInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Auth_Auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeInfoRsp) ProtoMessage() {}

func (x *ChangeInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_Auth_Auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeInfoRsp.ProtoReflect.Descriptor instead.
func (*ChangeInfoRsp) Descriptor() ([]byte, []int) {
	return file_Auth_Auth_proto_rawDescGZIP(), []int{5}
}

func (x *ChangeInfoRsp) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

type GetAuthPubKeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=Uid,json=uid,proto3" json:"Uid,omitempty"`
}

func (x *GetAuthPubKeyReq) Reset() {
	*x = GetAuthPubKeyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Auth_Auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthPubKeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthPubKeyReq) ProtoMessage() {}

func (x *GetAuthPubKeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_Auth_Auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthPubKeyReq.ProtoReflect.Descriptor instead.
func (*GetAuthPubKeyReq) Descriptor() ([]byte, []int) {
	return file_Auth_Auth_proto_rawDescGZIP(), []int{6}
}

func (x *GetAuthPubKeyReq) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type GetAuthPubKeyRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorMsg      string `protobuf:"bytes,1,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
	AuthRsaPubKey []byte `protobuf:"bytes,2,opt,name=authRsaPubKey,proto3" json:"authRsaPubKey,omitempty"`
}

func (x *GetAuthPubKeyRsp) Reset() {
	*x = GetAuthPubKeyRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Auth_Auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthPubKeyRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthPubKeyRsp) ProtoMessage() {}

func (x *GetAuthPubKeyRsp) ProtoReflect() protoreflect.Message {
	mi := &file_Auth_Auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthPubKeyRsp.ProtoReflect.Descriptor instead.
func (*GetAuthPubKeyRsp) Descriptor() ([]byte, []int) {
	return file_Auth_Auth_proto_rawDescGZIP(), []int{7}
}

func (x *GetAuthPubKeyRsp) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

func (x *GetAuthPubKeyRsp) GetAuthRsaPubKey() []byte {
	if x != nil {
		return x.AuthRsaPubKey
	}
	return nil
}

var File_Auth_Auth_proto protoreflect.FileDescriptor

var file_Auth_Auth_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x41, 0x75, 0x74, 0x68, 0x2f, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x41, 0x75, 0x74, 0x68, 0x22, 0x3c, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x43, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x74, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x73, 0x61, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x72, 0x73, 0x61, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79,
	0x22, 0x26, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x22, 0x27, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e,
	0x55, 0x70, 0x52, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73,
	0x67, 0x22, 0x2b, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x22, 0x24,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x22, 0x54, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x73, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x52, 0x73, 0x61, 0x50,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x61, 0x75, 0x74,
	0x68, 0x52, 0x73, 0x61, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x32, 0xd2, 0x01, 0x0a, 0x04, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x2b, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x1a, 0x0e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70,
	0x12, 0x29, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x0e, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0f, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x0a, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x13, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x3f,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12,
	0x16, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75,
	0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x73, 0x70, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Auth_Auth_proto_rawDescOnce sync.Once
	file_Auth_Auth_proto_rawDescData = file_Auth_Auth_proto_rawDesc
)

func file_Auth_Auth_proto_rawDescGZIP() []byte {
	file_Auth_Auth_proto_rawDescOnce.Do(func() {
		file_Auth_Auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_Auth_Auth_proto_rawDescData)
	})
	return file_Auth_Auth_proto_rawDescData
}

var file_Auth_Auth_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_Auth_Auth_proto_goTypes = []interface{}{
	(*UserPassword)(nil),     // 0: Auth.UserPassword
	(*UserGroupInfo)(nil),    // 1: Auth.UserGroupInfo
	(*UserInfo)(nil),         // 2: Auth.UserInfo
	(*LoginRsp)(nil),         // 3: Auth.LoginRsp
	(*SignUpRsp)(nil),        // 4: Auth.SignUpRsp
	(*ChangeInfoRsp)(nil),    // 5: Auth.ChangeInfoRsp
	(*GetAuthPubKeyReq)(nil), // 6: Auth.GetAuthPubKeyReq
	(*GetAuthPubKeyRsp)(nil), // 7: Auth.GetAuthPubKeyRsp
}
var file_Auth_Auth_proto_depIdxs = []int32{
	0, // 0: Auth.UserInfo.userPassword:type_name -> Auth.UserPassword
	0, // 1: Auth.Auth.Login:input_type -> Auth.UserPassword
	2, // 2: Auth.Auth.SignUp:input_type -> Auth.UserInfo
	2, // 3: Auth.Auth.ChangeInfo:input_type -> Auth.UserInfo
	6, // 4: Auth.Auth.GetAuthPubKey:input_type -> Auth.GetAuthPubKeyReq
	3, // 5: Auth.Auth.Login:output_type -> Auth.LoginRsp
	4, // 6: Auth.Auth.SignUp:output_type -> Auth.SignUpRsp
	5, // 7: Auth.Auth.ChangeInfo:output_type -> Auth.ChangeInfoRsp
	7, // 8: Auth.Auth.GetAuthPubKey:output_type -> Auth.GetAuthPubKeyRsp
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Auth_Auth_proto_init() }
func file_Auth_Auth_proto_init() {
	if File_Auth_Auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Auth_Auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPassword); i {
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
		file_Auth_Auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGroupInfo); i {
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
		file_Auth_Auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_Auth_Auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRsp); i {
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
		file_Auth_Auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpRsp); i {
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
		file_Auth_Auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeInfoRsp); i {
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
		file_Auth_Auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthPubKeyReq); i {
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
		file_Auth_Auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthPubKeyRsp); i {
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
			RawDescriptor: file_Auth_Auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Auth_Auth_proto_goTypes,
		DependencyIndexes: file_Auth_Auth_proto_depIdxs,
		MessageInfos:      file_Auth_Auth_proto_msgTypes,
	}.Build()
	File_Auth_Auth_proto = out.File
	file_Auth_Auth_proto_rawDesc = nil
	file_Auth_Auth_proto_goTypes = nil
	file_Auth_Auth_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	Login(ctx context.Context, in *UserPassword, opts ...grpc.CallOption) (*LoginRsp, error)
	SignUp(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*SignUpRsp, error)
	ChangeInfo(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*ChangeInfoRsp, error)
	GetAuthPubKey(ctx context.Context, in *GetAuthPubKeyReq, opts ...grpc.CallOption) (*GetAuthPubKeyRsp, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Login(ctx context.Context, in *UserPassword, opts ...grpc.CallOption) (*LoginRsp, error) {
	out := new(LoginRsp)
	err := c.cc.Invoke(ctx, "/Auth.Auth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SignUp(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*SignUpRsp, error) {
	out := new(SignUpRsp)
	err := c.cc.Invoke(ctx, "/Auth.Auth/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ChangeInfo(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*ChangeInfoRsp, error) {
	out := new(ChangeInfoRsp)
	err := c.cc.Invoke(ctx, "/Auth.Auth/ChangeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetAuthPubKey(ctx context.Context, in *GetAuthPubKeyReq, opts ...grpc.CallOption) (*GetAuthPubKeyRsp, error) {
	out := new(GetAuthPubKeyRsp)
	err := c.cc.Invoke(ctx, "/Auth.Auth/GetAuthPubKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	Login(context.Context, *UserPassword) (*LoginRsp, error)
	SignUp(context.Context, *UserInfo) (*SignUpRsp, error)
	ChangeInfo(context.Context, *UserInfo) (*ChangeInfoRsp, error)
	GetAuthPubKey(context.Context, *GetAuthPubKeyReq) (*GetAuthPubKeyRsp, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) Login(context.Context, *UserPassword) (*LoginRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthServer) SignUp(context.Context, *UserInfo) (*SignUpRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (*UnimplementedAuthServer) ChangeInfo(context.Context, *UserInfo) (*ChangeInfoRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeInfo not implemented")
}
func (*UnimplementedAuthServer) GetAuthPubKey(context.Context, *GetAuthPubKeyReq) (*GetAuthPubKeyRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthPubKey not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPassword)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth.Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*UserPassword))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth.Auth/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SignUp(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ChangeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ChangeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth.Auth/ChangeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ChangeInfo(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetAuthPubKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthPubKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetAuthPubKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth.Auth/GetAuthPubKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetAuthPubKey(ctx, req.(*GetAuthPubKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "SignUp",
			Handler:    _Auth_SignUp_Handler,
		},
		{
			MethodName: "ChangeInfo",
			Handler:    _Auth_ChangeInfo_Handler,
		},
		{
			MethodName: "GetAuthPubKey",
			Handler:    _Auth_GetAuthPubKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Auth/Auth.proto",
}
