// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: CacheService.proto

package cacheserver_grpc_go

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_CacheService_proto protoreflect.FileDescriptor

var file_CacheService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x0e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4d, 0x73, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x73,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbd, 0x06, 0x0a, 0x0c, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x22, 0x2e, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65,
	0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x61, 0x74, 0x61, 0x57, 0x69, 0x74, 0x68,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x56, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x57, 0x69, 0x74, 0x68, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x5c,
	0x0a, 0x10, 0x53, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x4e, 0x78, 0x12, 0x22, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x57, 0x69, 0x74, 0x68, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x56, 0x0a, 0x0d,
	0x53, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4e, 0x78, 0x12, 0x1f, 0x2e,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x61, 0x74, 0x61, 0x57, 0x69, 0x74, 0x68,
	0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x59, 0x0a, 0x0f, 0x4c, 0x6f, 0x61, 0x64, 0x4b, 0x65, 0x79, 0x49,
	0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x6e, 0x65, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x44, 0x61, 0x74, 0x61, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12,
	0x5d, 0x0a, 0x12, 0x4c, 0x6f, 0x61, 0x64, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x6e, 0x65, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x57, 0x69, 0x74, 0x68, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x55,
	0x0a, 0x0b, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x2e,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x4f, 0x6e, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x61, 0x74, 0x61, 0x57, 0x69, 0x74, 0x68,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x55, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x12, 0x22, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x57, 0x69, 0x74, 0x68, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x59, 0x0a, 0x0a,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x24, 0x2e, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x61, 0x74, 0x61, 0x57, 0x69,
	0x74, 0x68, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_CacheService_proto_goTypes = []interface{}{
	(*SetStringRequest)(nil),     // 0: cacheserver_grpc.SetStringRequest
	(*SetIntRequest)(nil),        // 1: cacheserver_grpc.SetIntRequest
	(*OneKeyRequest)(nil),        // 2: cacheserver_grpc.OneKeyRequest
	(*SetExpireRequest)(nil),     // 3: cacheserver_grpc.SetExpireRequest
	(*MultipleKeyRequest)(nil),   // 4: cacheserver_grpc.MultipleKeyRequest
	(*ResultDataWithString)(nil), // 5: cacheserver_grpc.ResultDataWithString
	(*ResultDataWithBool)(nil),   // 6: cacheserver_grpc.ResultDataWithBool
	(*ResultDataWithInt64)(nil),  // 7: cacheserver_grpc.ResultDataWithInt64
}
var file_CacheService_proto_depIdxs = []int32{
	0, // 0: cacheserver_grpc.CacheService.SetStringValue:input_type -> cacheserver_grpc.SetStringRequest
	1, // 1: cacheserver_grpc.CacheService.SetIntValue:input_type -> cacheserver_grpc.SetIntRequest
	0, // 2: cacheserver_grpc.CacheService.SetStringValueNx:input_type -> cacheserver_grpc.SetStringRequest
	1, // 3: cacheserver_grpc.CacheService.SetIntValueNx:input_type -> cacheserver_grpc.SetIntRequest
	2, // 4: cacheserver_grpc.CacheService.LoadKeyIntValue:input_type -> cacheserver_grpc.OneKeyRequest
	2, // 5: cacheserver_grpc.CacheService.LoadKeyStringValue:input_type -> cacheserver_grpc.OneKeyRequest
	2, // 6: cacheserver_grpc.CacheService.IncreaseKey:input_type -> cacheserver_grpc.OneKeyRequest
	3, // 7: cacheserver_grpc.CacheService.SetExpire:input_type -> cacheserver_grpc.SetExpireRequest
	4, // 8: cacheserver_grpc.CacheService.RemoveKeys:input_type -> cacheserver_grpc.MultipleKeyRequest
	5, // 9: cacheserver_grpc.CacheService.SetStringValue:output_type -> cacheserver_grpc.ResultDataWithString
	5, // 10: cacheserver_grpc.CacheService.SetIntValue:output_type -> cacheserver_grpc.ResultDataWithString
	6, // 11: cacheserver_grpc.CacheService.SetStringValueNx:output_type -> cacheserver_grpc.ResultDataWithBool
	6, // 12: cacheserver_grpc.CacheService.SetIntValueNx:output_type -> cacheserver_grpc.ResultDataWithBool
	7, // 13: cacheserver_grpc.CacheService.LoadKeyIntValue:output_type -> cacheserver_grpc.ResultDataWithInt64
	5, // 14: cacheserver_grpc.CacheService.LoadKeyStringValue:output_type -> cacheserver_grpc.ResultDataWithString
	7, // 15: cacheserver_grpc.CacheService.IncreaseKey:output_type -> cacheserver_grpc.ResultDataWithInt64
	6, // 16: cacheserver_grpc.CacheService.SetExpire:output_type -> cacheserver_grpc.ResultDataWithBool
	7, // 17: cacheserver_grpc.CacheService.RemoveKeys:output_type -> cacheserver_grpc.ResultDataWithInt64
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CacheService_proto_init() }
func file_CacheService_proto_init() {
	if File_CacheService_proto != nil {
		return
	}
	file_CacheMsg_proto_init()
	file_CommonMsg_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_CacheService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_CacheService_proto_goTypes,
		DependencyIndexes: file_CacheService_proto_depIdxs,
	}.Build()
	File_CacheService_proto = out.File
	file_CacheService_proto_rawDesc = nil
	file_CacheService_proto_goTypes = nil
	file_CacheService_proto_depIdxs = nil
}
