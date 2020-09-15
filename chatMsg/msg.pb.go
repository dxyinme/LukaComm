// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: chatMsg/msg.proto

package chatMsg

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

type MsgContentType int32

const (
	MsgContentType_MsgContentTypeErr MsgContentType = 0
	MsgContentType_Text              MsgContentType = 1
	MsgContentType_Img               MsgContentType = 2
	MsgContentType_Video             MsgContentType = 3
)

// Enum value maps for MsgContentType.
var (
	MsgContentType_name = map[int32]string{
		0: "MsgContentTypeErr",
		1: "Text",
		2: "Img",
		3: "Video",
	}
	MsgContentType_value = map[string]int32{
		"MsgContentTypeErr": 0,
		"Text":              1,
		"Img":               2,
		"Video":             3,
	}
)

func (x MsgContentType) Enum() *MsgContentType {
	p := new(MsgContentType)
	*p = x
	return p
}

func (x MsgContentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgContentType) Descriptor() protoreflect.EnumDescriptor {
	return file_chatMsg_msg_proto_enumTypes[0].Descriptor()
}

func (MsgContentType) Type() protoreflect.EnumType {
	return &file_chatMsg_msg_proto_enumTypes[0]
}

func (x MsgContentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgContentType.Descriptor instead.
func (MsgContentType) EnumDescriptor() ([]byte, []int) {
	return file_chatMsg_msg_proto_rawDescGZIP(), []int{0}
}

type MsgType int32

const (
	MsgType_MsgTypeErr MsgType = 0
	MsgType_Single     MsgType = 1
	MsgType_Group      MsgType = 2
)

// Enum value maps for MsgType.
var (
	MsgType_name = map[int32]string{
		0: "MsgTypeErr",
		1: "Single",
		2: "Group",
	}
	MsgType_value = map[string]int32{
		"MsgTypeErr": 0,
		"Single":     1,
		"Group":      2,
	}
)

func (x MsgType) Enum() *MsgType {
	p := new(MsgType)
	*p = x
	return p
}

func (x MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_chatMsg_msg_proto_enumTypes[1].Descriptor()
}

func (MsgType) Type() protoreflect.EnumType {
	return &file_chatMsg_msg_proto_enumTypes[1]
}

func (x MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgType.Descriptor instead.
func (MsgType) EnumDescriptor() ([]byte, []int) {
	return file_chatMsg_msg_proto_rawDescGZIP(), []int{1}
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From           string         `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Target         string         `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Content        []byte         `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	MsgType        MsgType        `protobuf:"varint,4,opt,name=msgType,proto3,enum=chatMsg.MsgType" json:"msgType,omitempty"`
	MsgContentType MsgContentType `protobuf:"varint,5,opt,name=msgContentType,proto3,enum=chatMsg.MsgContentType" json:"msgContentType,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatMsg_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_chatMsg_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_chatMsg_msg_proto_rawDescGZIP(), []int{0}
}

func (x *Msg) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Msg) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *Msg) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Msg) GetMsgType() MsgType {
	if x != nil {
		return x.MsgType
	}
	return MsgType_MsgTypeErr
}

func (x *Msg) GetMsgContentType() MsgContentType {
	if x != nil {
		return x.MsgContentType
	}
	return MsgContentType_MsgContentTypeErr
}

var File_chatMsg_msg_proto protoreflect.FileDescriptor

var file_chatMsg_msg_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x22, 0xb8, 0x01, 0x0a,
	0x03, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x6d, 0x73,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x4d, 0x73, 0x67, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6d,
	0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x6d, 0x73, 0x67, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x6d, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2a, 0x45, 0x0a, 0x0e, 0x4d, 0x73, 0x67, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x73, 0x67,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x72, 0x72, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x6d,
	0x67, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x10, 0x03, 0x2a, 0x30,
	0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x73, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x72, 0x72, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x10, 0x02,
	0x42, 0x09, 0x5a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_chatMsg_msg_proto_rawDescOnce sync.Once
	file_chatMsg_msg_proto_rawDescData = file_chatMsg_msg_proto_rawDesc
)

func file_chatMsg_msg_proto_rawDescGZIP() []byte {
	file_chatMsg_msg_proto_rawDescOnce.Do(func() {
		file_chatMsg_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_chatMsg_msg_proto_rawDescData)
	})
	return file_chatMsg_msg_proto_rawDescData
}

var file_chatMsg_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_chatMsg_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chatMsg_msg_proto_goTypes = []interface{}{
	(MsgContentType)(0), // 0: chatMsg.MsgContentType
	(MsgType)(0),        // 1: chatMsg.MsgType
	(*Msg)(nil),         // 2: chatMsg.Msg
}
var file_chatMsg_msg_proto_depIdxs = []int32{
	1, // 0: chatMsg.Msg.msgType:type_name -> chatMsg.MsgType
	0, // 1: chatMsg.Msg.msgContentType:type_name -> chatMsg.MsgContentType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_chatMsg_msg_proto_init() }
func file_chatMsg_msg_proto_init() {
	if File_chatMsg_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chatMsg_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
			RawDescriptor: file_chatMsg_msg_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chatMsg_msg_proto_goTypes,
		DependencyIndexes: file_chatMsg_msg_proto_depIdxs,
		EnumInfos:         file_chatMsg_msg_proto_enumTypes,
		MessageInfos:      file_chatMsg_msg_proto_msgTypes,
	}.Build()
	File_chatMsg_msg_proto = out.File
	file_chatMsg_msg_proto_rawDesc = nil
	file_chatMsg_msg_proto_goTypes = nil
	file_chatMsg_msg_proto_depIdxs = nil
}
