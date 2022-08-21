// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: CacheMsg.proto

package cacheserver_grpc_go

import (
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

// 请求
type GenSnowflakeIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatacenterId int64 `protobuf:"varint,1,opt,name=DatacenterId,proto3" json:"DatacenterId"`
	WorkerId     int64 `protobuf:"varint,2,opt,name=WorkerId,proto3" json:"WorkerId"`
}

func (x *GenSnowflakeIdRequest) Reset() {
	*x = GenSnowflakeIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CacheMsg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenSnowflakeIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenSnowflakeIdRequest) ProtoMessage() {}

func (x *GenSnowflakeIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_CacheMsg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenSnowflakeIdRequest.ProtoReflect.Descriptor instead.
func (*GenSnowflakeIdRequest) Descriptor() ([]byte, []int) {
	return file_CacheMsg_proto_rawDescGZIP(), []int{0}
}

func (x *GenSnowflakeIdRequest) GetDatacenterId() int64 {
	if x != nil {
		return x.DatacenterId
	}
	return 0
}

func (x *GenSnowflakeIdRequest) GetWorkerId() int64 {
	if x != nil {
		return x.WorkerId
	}
	return 0
}

type SetStringRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key            string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key"`
	Value          string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value"`
	ExpMillisecond int64  `protobuf:"varint,3,opt,name=ExpMillisecond,proto3" json:"ExpMillisecond"`
}

func (x *SetStringRequest) Reset() {
	*x = SetStringRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CacheMsg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetStringRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStringRequest) ProtoMessage() {}

func (x *SetStringRequest) ProtoReflect() protoreflect.Message {
	mi := &file_CacheMsg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStringRequest.ProtoReflect.Descriptor instead.
func (*SetStringRequest) Descriptor() ([]byte, []int) {
	return file_CacheMsg_proto_rawDescGZIP(), []int{1}
}

func (x *SetStringRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetStringRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *SetStringRequest) GetExpMillisecond() int64 {
	if x != nil {
		return x.ExpMillisecond
	}
	return 0
}

type SetIntRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key            string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key"`
	Value          int64  `protobuf:"varint,2,opt,name=Value,proto3" json:"Value"`
	ExpMillisecond int64  `protobuf:"varint,3,opt,name=ExpMillisecond,proto3" json:"ExpMillisecond"`
}

func (x *SetIntRequest) Reset() {
	*x = SetIntRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CacheMsg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetIntRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetIntRequest) ProtoMessage() {}

func (x *SetIntRequest) ProtoReflect() protoreflect.Message {
	mi := &file_CacheMsg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetIntRequest.ProtoReflect.Descriptor instead.
func (*SetIntRequest) Descriptor() ([]byte, []int) {
	return file_CacheMsg_proto_rawDescGZIP(), []int{2}
}

func (x *SetIntRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetIntRequest) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *SetIntRequest) GetExpMillisecond() int64 {
	if x != nil {
		return x.ExpMillisecond
	}
	return 0
}

type SetExpireRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key            string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key"`
	ExpMillisecond int64  `protobuf:"varint,2,opt,name=ExpMillisecond,proto3" json:"ExpMillisecond"`
}

func (x *SetExpireRequest) Reset() {
	*x = SetExpireRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CacheMsg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetExpireRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetExpireRequest) ProtoMessage() {}

func (x *SetExpireRequest) ProtoReflect() protoreflect.Message {
	mi := &file_CacheMsg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetExpireRequest.ProtoReflect.Descriptor instead.
func (*SetExpireRequest) Descriptor() ([]byte, []int) {
	return file_CacheMsg_proto_rawDescGZIP(), []int{3}
}

func (x *SetExpireRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetExpireRequest) GetExpMillisecond() int64 {
	if x != nil {
		return x.ExpMillisecond
	}
	return 0
}

type OneKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key"`
}

func (x *OneKeyRequest) Reset() {
	*x = OneKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CacheMsg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneKeyRequest) ProtoMessage() {}

func (x *OneKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_CacheMsg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneKeyRequest.ProtoReflect.Descriptor instead.
func (*OneKeyRequest) Descriptor() ([]byte, []int) {
	return file_CacheMsg_proto_rawDescGZIP(), []int{4}
}

func (x *OneKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type MultipleKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []string `protobuf:"bytes,1,rep,name=Keys,proto3" json:"Keys"`
}

func (x *MultipleKeyRequest) Reset() {
	*x = MultipleKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CacheMsg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipleKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleKeyRequest) ProtoMessage() {}

func (x *MultipleKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_CacheMsg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleKeyRequest.ProtoReflect.Descriptor instead.
func (*MultipleKeyRequest) Descriptor() ([]byte, []int) {
	return file_CacheMsg_proto_rawDescGZIP(), []int{5}
}

func (x *MultipleKeyRequest) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type PfAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string   `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key"`
	Elements []string `protobuf:"bytes,2,rep,name=Elements,proto3" json:"Elements"`
}

func (x *PfAddRequest) Reset() {
	*x = PfAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CacheMsg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PfAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PfAddRequest) ProtoMessage() {}

func (x *PfAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_CacheMsg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PfAddRequest.ProtoReflect.Descriptor instead.
func (*PfAddRequest) Descriptor() ([]byte, []int) {
	return file_CacheMsg_proto_rawDescGZIP(), []int{6}
}

func (x *PfAddRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PfAddRequest) GetElements() []string {
	if x != nil {
		return x.Elements
	}
	return nil
}

type ResultDataWithInt64 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message"`
	Data    int64  `protobuf:"varint,3,opt,name=Data,proto3" json:"Data"`
}

func (x *ResultDataWithInt64) Reset() {
	*x = ResultDataWithInt64{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CacheMsg_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultDataWithInt64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultDataWithInt64) ProtoMessage() {}

func (x *ResultDataWithInt64) ProtoReflect() protoreflect.Message {
	mi := &file_CacheMsg_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultDataWithInt64.ProtoReflect.Descriptor instead.
func (*ResultDataWithInt64) Descriptor() ([]byte, []int) {
	return file_CacheMsg_proto_rawDescGZIP(), []int{7}
}

func (x *ResultDataWithInt64) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResultDataWithInt64) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResultDataWithInt64) GetData() int64 {
	if x != nil {
		return x.Data
	}
	return 0
}

type ResultDataWithInt64Array struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code"`
	Message string  `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message"`
	Data    []int64 `protobuf:"varint,3,rep,packed,name=Data,proto3" json:"Data"`
}

func (x *ResultDataWithInt64Array) Reset() {
	*x = ResultDataWithInt64Array{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CacheMsg_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultDataWithInt64Array) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultDataWithInt64Array) ProtoMessage() {}

func (x *ResultDataWithInt64Array) ProtoReflect() protoreflect.Message {
	mi := &file_CacheMsg_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultDataWithInt64Array.ProtoReflect.Descriptor instead.
func (*ResultDataWithInt64Array) Descriptor() ([]byte, []int) {
	return file_CacheMsg_proto_rawDescGZIP(), []int{8}
}

func (x *ResultDataWithInt64Array) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResultDataWithInt64Array) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResultDataWithInt64Array) GetData() []int64 {
	if x != nil {
		return x.Data
	}
	return nil
}

type ResultDataWithString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message"`
	Data    string `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data"`
}

func (x *ResultDataWithString) Reset() {
	*x = ResultDataWithString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CacheMsg_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultDataWithString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultDataWithString) ProtoMessage() {}

func (x *ResultDataWithString) ProtoReflect() protoreflect.Message {
	mi := &file_CacheMsg_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultDataWithString.ProtoReflect.Descriptor instead.
func (*ResultDataWithString) Descriptor() ([]byte, []int) {
	return file_CacheMsg_proto_rawDescGZIP(), []int{9}
}

func (x *ResultDataWithString) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResultDataWithString) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResultDataWithString) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type ResultDataWithBool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message"`
	Data    bool   `protobuf:"varint,3,opt,name=Data,proto3" json:"Data"`
}

func (x *ResultDataWithBool) Reset() {
	*x = ResultDataWithBool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CacheMsg_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultDataWithBool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultDataWithBool) ProtoMessage() {}

func (x *ResultDataWithBool) ProtoReflect() protoreflect.Message {
	mi := &file_CacheMsg_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultDataWithBool.ProtoReflect.Descriptor instead.
func (*ResultDataWithBool) Descriptor() ([]byte, []int) {
	return file_CacheMsg_proto_rawDescGZIP(), []int{10}
}

func (x *ResultDataWithBool) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResultDataWithBool) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResultDataWithBool) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

var File_CacheMsg_proto protoreflect.FileDescriptor

var file_CacheMsg_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x22, 0x57, 0x0a, 0x15, 0x47, 0x65, 0x6e, 0x53, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61,
	0x6b, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x44,
	0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x22, 0x62, 0x0a, 0x10, 0x53,
	0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x4d, 0x69,
	0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0e, 0x45, 0x78, 0x70, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x22,
	0x5f, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x4d,
	0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x45, 0x78, 0x70, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x22, 0x4c, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x4d, 0x69, 0x6c,
	0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x45, 0x78, 0x70, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x22, 0x21,
	0x0a, 0x0d, 0x4f, 0x6e, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65,
	0x79, 0x22, 0x28, 0x0a, 0x12, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4b, 0x65, 0x79, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x4b, 0x65, 0x79, 0x73, 0x22, 0x3c, 0x0a, 0x0c, 0x50,
	0x66, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x57, 0x0a, 0x13, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x44, 0x61, 0x74, 0x61, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x5c, 0x0a, 0x18, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x58, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x61, 0x74, 0x61, 0x57, 0x69,
	0x74, 0x68, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x56, 0x0a, 0x12, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x44, 0x61, 0x74, 0x61, 0x57, 0x69, 0x74, 0x68, 0x42, 0x6f, 0x6f, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_CacheMsg_proto_rawDescOnce sync.Once
	file_CacheMsg_proto_rawDescData = file_CacheMsg_proto_rawDesc
)

func file_CacheMsg_proto_rawDescGZIP() []byte {
	file_CacheMsg_proto_rawDescOnce.Do(func() {
		file_CacheMsg_proto_rawDescData = protoimpl.X.CompressGZIP(file_CacheMsg_proto_rawDescData)
	})
	return file_CacheMsg_proto_rawDescData
}

var file_CacheMsg_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_CacheMsg_proto_goTypes = []interface{}{
	(*GenSnowflakeIdRequest)(nil),    // 0: cacheserver_grpc.GenSnowflakeIdRequest
	(*SetStringRequest)(nil),         // 1: cacheserver_grpc.SetStringRequest
	(*SetIntRequest)(nil),            // 2: cacheserver_grpc.SetIntRequest
	(*SetExpireRequest)(nil),         // 3: cacheserver_grpc.SetExpireRequest
	(*OneKeyRequest)(nil),            // 4: cacheserver_grpc.OneKeyRequest
	(*MultipleKeyRequest)(nil),       // 5: cacheserver_grpc.MultipleKeyRequest
	(*PfAddRequest)(nil),             // 6: cacheserver_grpc.PfAddRequest
	(*ResultDataWithInt64)(nil),      // 7: cacheserver_grpc.ResultDataWithInt64
	(*ResultDataWithInt64Array)(nil), // 8: cacheserver_grpc.ResultDataWithInt64Array
	(*ResultDataWithString)(nil),     // 9: cacheserver_grpc.ResultDataWithString
	(*ResultDataWithBool)(nil),       // 10: cacheserver_grpc.ResultDataWithBool
}
var file_CacheMsg_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CacheMsg_proto_init() }
func file_CacheMsg_proto_init() {
	if File_CacheMsg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CacheMsg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenSnowflakeIdRequest); i {
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
		file_CacheMsg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetStringRequest); i {
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
		file_CacheMsg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetIntRequest); i {
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
		file_CacheMsg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetExpireRequest); i {
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
		file_CacheMsg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneKeyRequest); i {
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
		file_CacheMsg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultipleKeyRequest); i {
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
		file_CacheMsg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PfAddRequest); i {
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
		file_CacheMsg_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultDataWithInt64); i {
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
		file_CacheMsg_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultDataWithInt64Array); i {
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
		file_CacheMsg_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultDataWithString); i {
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
		file_CacheMsg_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultDataWithBool); i {
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
			RawDescriptor: file_CacheMsg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CacheMsg_proto_goTypes,
		DependencyIndexes: file_CacheMsg_proto_depIdxs,
		MessageInfos:      file_CacheMsg_proto_msgTypes,
	}.Build()
	File_CacheMsg_proto = out.File
	file_CacheMsg_proto_rawDesc = nil
	file_CacheMsg_proto_goTypes = nil
	file_CacheMsg_proto_depIdxs = nil
}
