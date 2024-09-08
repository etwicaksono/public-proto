// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: example/example.proto

package example

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Multimedia Message
type MultimediaStorage int32

const (
	MultimediaStorage_LOCAL MultimediaStorage = 0
	MultimediaStorage_MINIO MultimediaStorage = 1
)

// Enum value maps for MultimediaStorage.
var (
	MultimediaStorage_name = map[int32]string{
		0: "LOCAL",
		1: "MINIO",
	}
	MultimediaStorage_value = map[string]int32{
		"LOCAL": 0,
		"MINIO": 1,
	}
)

func (x MultimediaStorage) Enum() *MultimediaStorage {
	p := new(MultimediaStorage)
	*p = x
	return p
}

func (x MultimediaStorage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MultimediaStorage) Descriptor() protoreflect.EnumDescriptor {
	return file_example_example_proto_enumTypes[0].Descriptor()
}

func (MultimediaStorage) Type() protoreflect.EnumType {
	return &file_example_example_proto_enumTypes[0]
}

func (x MultimediaStorage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MultimediaStorage.Descriptor instead.
func (MultimediaStorage) EnumDescriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{0}
}

// Text Message
type SendTextMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender   string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Message  string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendTextMessageRequest) Reset() {
	*x = SendTextMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTextMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTextMessageRequest) ProtoMessage() {}

func (x *SendTextMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTextMessageRequest.ProtoReflect.Descriptor instead.
func (*SendTextMessageRequest) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{0}
}

func (x *SendTextMessageRequest) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *SendTextMessageRequest) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *SendTextMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MessageTextItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender   string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver string `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Message  string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MessageTextItem) Reset() {
	*x = MessageTextItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageTextItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageTextItem) ProtoMessage() {}

func (x *MessageTextItem) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageTextItem.ProtoReflect.Descriptor instead.
func (*MessageTextItem) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{1}
}

func (x *MessageTextItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MessageTextItem) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *MessageTextItem) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *MessageTextItem) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetTextMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*MessageTextItem `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetTextMessageResponse) Reset() {
	*x = GetTextMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTextMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTextMessageResponse) ProtoMessage() {}

func (x *GetTextMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTextMessageResponse.ProtoReflect.Descriptor instead.
func (*GetTextMessageResponse) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{2}
}

func (x *GetTextMessageResponse) GetData() []*MessageTextItem {
	if x != nil {
		return x.Data
	}
	return nil
}

type MultimediaFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MultimediaFile) Reset() {
	*x = MultimediaFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultimediaFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultimediaFile) ProtoMessage() {}

func (x *MultimediaFile) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultimediaFile.ProtoReflect.Descriptor instead.
func (*MultimediaFile) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{3}
}

func (x *MultimediaFile) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *MultimediaFile) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SendMultimediaMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender   string            `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver string            `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Message  string            `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Storage  MultimediaStorage `protobuf:"varint,4,opt,name=storage,proto3,enum=example.MultimediaStorage" json:"storage,omitempty"`
	Files    []*MultimediaFile `protobuf:"bytes,5,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *SendMultimediaMessageRequest) Reset() {
	*x = SendMultimediaMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMultimediaMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMultimediaMessageRequest) ProtoMessage() {}

func (x *SendMultimediaMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMultimediaMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMultimediaMessageRequest) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{4}
}

func (x *SendMultimediaMessageRequest) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *SendMultimediaMessageRequest) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *SendMultimediaMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SendMultimediaMessageRequest) GetStorage() MultimediaStorage {
	if x != nil {
		return x.Storage
	}
	return MultimediaStorage_LOCAL
}

func (x *SendMultimediaMessageRequest) GetFiles() []*MultimediaFile {
	if x != nil {
		return x.Files
	}
	return nil
}

type MessageMultimediaItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender   string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver string `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Message  string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	FileUrl  string `protobuf:"bytes,5,opt,name=fileUrl,proto3" json:"fileUrl,omitempty"`
}

func (x *MessageMultimediaItem) Reset() {
	*x = MessageMultimediaItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageMultimediaItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageMultimediaItem) ProtoMessage() {}

func (x *MessageMultimediaItem) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageMultimediaItem.ProtoReflect.Descriptor instead.
func (*MessageMultimediaItem) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{5}
}

func (x *MessageMultimediaItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MessageMultimediaItem) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *MessageMultimediaItem) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *MessageMultimediaItem) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MessageMultimediaItem) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

type GetMultimediaMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*MessageMultimediaItem `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetMultimediaMessageResponse) Reset() {
	*x = GetMultimediaMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMultimediaMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMultimediaMessageResponse) ProtoMessage() {}

func (x *GetMultimediaMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMultimediaMessageResponse.ProtoReflect.Descriptor instead.
func (*GetMultimediaMessageResponse) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{6}
}

func (x *GetMultimediaMessageResponse) GetData() []*MessageMultimediaItem {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_example_example_proto protoreflect.FileDescriptor

var file_example_example_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a,
	0x16, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6f, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x65, 0x78, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x46, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x65, 0x78, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x40,
	0x0a, 0x0e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0xd1, 0x01, 0x0a, 0x1c, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x34, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x07, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x22, 0x8f, 0x01, 0x0a, 0x15, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x66, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66,
	0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x52, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x29, 0x0a, 0x11, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x49,
	0x4e, 0x49, 0x4f, 0x10, 0x01, 0x32, 0xd6, 0x02, 0x0a, 0x0e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64,
	0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x49, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x56, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x55, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x25, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f,
	0x5a, 0x0d, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_example_proto_rawDescOnce sync.Once
	file_example_example_proto_rawDescData = file_example_example_proto_rawDesc
)

func file_example_example_proto_rawDescGZIP() []byte {
	file_example_example_proto_rawDescOnce.Do(func() {
		file_example_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_example_proto_rawDescData)
	})
	return file_example_example_proto_rawDescData
}

var file_example_example_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_example_example_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_example_example_proto_goTypes = []interface{}{
	(MultimediaStorage)(0),               // 0: example.MultimediaStorage
	(*SendTextMessageRequest)(nil),       // 1: example.SendTextMessageRequest
	(*MessageTextItem)(nil),              // 2: example.MessageTextItem
	(*GetTextMessageResponse)(nil),       // 3: example.GetTextMessageResponse
	(*MultimediaFile)(nil),               // 4: example.MultimediaFile
	(*SendMultimediaMessageRequest)(nil), // 5: example.SendMultimediaMessageRequest
	(*MessageMultimediaItem)(nil),        // 6: example.MessageMultimediaItem
	(*GetMultimediaMessageResponse)(nil), // 7: example.GetMultimediaMessageResponse
	(*emptypb.Empty)(nil),                // 8: google.protobuf.Empty
}
var file_example_example_proto_depIdxs = []int32{
	2, // 0: example.GetTextMessageResponse.data:type_name -> example.MessageTextItem
	0, // 1: example.SendMultimediaMessageRequest.storage:type_name -> example.MultimediaStorage
	4, // 2: example.SendMultimediaMessageRequest.files:type_name -> example.MultimediaFile
	6, // 3: example.GetMultimediaMessageResponse.data:type_name -> example.MessageMultimediaItem
	1, // 4: example.ExampleService.SendTextMessage:input_type -> example.SendTextMessageRequest
	8, // 5: example.ExampleService.GetTextMessage:input_type -> google.protobuf.Empty
	5, // 6: example.ExampleService.SendMultimediaMessage:input_type -> example.SendMultimediaMessageRequest
	8, // 7: example.ExampleService.GetMultimediaMessage:input_type -> google.protobuf.Empty
	8, // 8: example.ExampleService.SendTextMessage:output_type -> google.protobuf.Empty
	3, // 9: example.ExampleService.GetTextMessage:output_type -> example.GetTextMessageResponse
	8, // 10: example.ExampleService.SendMultimediaMessage:output_type -> google.protobuf.Empty
	7, // 11: example.ExampleService.GetMultimediaMessage:output_type -> example.GetMultimediaMessageResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_example_example_proto_init() }
func file_example_example_proto_init() {
	if File_example_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTextMessageRequest); i {
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
		file_example_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageTextItem); i {
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
		file_example_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTextMessageResponse); i {
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
		file_example_example_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultimediaFile); i {
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
		file_example_example_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMultimediaMessageRequest); i {
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
		file_example_example_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageMultimediaItem); i {
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
		file_example_example_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMultimediaMessageResponse); i {
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
			RawDescriptor: file_example_example_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_example_proto_goTypes,
		DependencyIndexes: file_example_example_proto_depIdxs,
		EnumInfos:         file_example_example_proto_enumTypes,
		MessageInfos:      file_example_example_proto_msgTypes,
	}.Build()
	File_example_example_proto = out.File
	file_example_example_proto_rawDesc = nil
	file_example_example_proto_goTypes = nil
	file_example_example_proto_depIdxs = nil
}
