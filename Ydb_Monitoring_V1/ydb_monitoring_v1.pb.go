// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: ydb_monitoring_v1.proto

package Ydb_Monitoring_V1

import (
	Ydb_Monitoring "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Monitoring"
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

var File_ydb_monitoring_v1_proto protoreflect.FileDescriptor

var file_ydb_monitoring_v1_proto_rawDesc = []byte{
	0x0a, 0x17, 0x79, 0x64, 0x62, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x59, 0x64, 0x62, 0x2e, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x1a, 0x1b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb7, 0x01, 0x0a, 0x11, 0x4d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x50, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x20, 0x2e, 0x59,
	0x64, 0x62, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65,
	0x6c, 0x66, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x53, 0x65, 0x6c, 0x66, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x50, 0x0a, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x20,
	0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x59, 0x0a, 0x1c, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x79, 0x64, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x79, 0x64, 0x62,
	0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x59, 0x64, 0x62,
	0x5f, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_ydb_monitoring_v1_proto_goTypes = []interface{}{
	(*Ydb_Monitoring.SelfCheckRequest)(nil),  // 0: Ydb.Monitoring.SelfCheckRequest
	(*Ydb_Monitoring.NodeCheckRequest)(nil),  // 1: Ydb.Monitoring.NodeCheckRequest
	(*Ydb_Monitoring.SelfCheckResponse)(nil), // 2: Ydb.Monitoring.SelfCheckResponse
	(*Ydb_Monitoring.NodeCheckResponse)(nil), // 3: Ydb.Monitoring.NodeCheckResponse
}
var file_ydb_monitoring_v1_proto_depIdxs = []int32{
	0, // 0: Ydb.Monitoring.V1.MonitoringService.SelfCheck:input_type -> Ydb.Monitoring.SelfCheckRequest
	1, // 1: Ydb.Monitoring.V1.MonitoringService.NodeCheck:input_type -> Ydb.Monitoring.NodeCheckRequest
	2, // 2: Ydb.Monitoring.V1.MonitoringService.SelfCheck:output_type -> Ydb.Monitoring.SelfCheckResponse
	3, // 3: Ydb.Monitoring.V1.MonitoringService.NodeCheck:output_type -> Ydb.Monitoring.NodeCheckResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ydb_monitoring_v1_proto_init() }
func file_ydb_monitoring_v1_proto_init() {
	if File_ydb_monitoring_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ydb_monitoring_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ydb_monitoring_v1_proto_goTypes,
		DependencyIndexes: file_ydb_monitoring_v1_proto_depIdxs,
	}.Build()
	File_ydb_monitoring_v1_proto = out.File
	file_ydb_monitoring_v1_proto_rawDesc = nil
	file_ydb_monitoring_v1_proto_goTypes = nil
	file_ydb_monitoring_v1_proto_depIdxs = nil
}
