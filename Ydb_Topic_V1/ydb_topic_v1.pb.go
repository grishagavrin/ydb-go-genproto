// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: ydb_topic_v1.proto

package Ydb_Topic_V1

import (
	Ydb_Topic "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Topic"
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

var File_ydb_topic_v1_proto protoreflect.FileDescriptor

var file_ydb_topic_v1_proto_rawDesc = []byte{
	0x0a, 0x12, 0x79, 0x64, 0x62, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x2e,
	0x56, 0x31, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbc, 0x05, 0x0a, 0x0c, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x0b, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x59, 0x64, 0x62,
	0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x1a, 0x28, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x28, 0x01,
	0x30, 0x01, 0x12, 0x62, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x61, 0x64,
	0x12, 0x27, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x46,
	0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x27, 0x2e, 0x59, 0x64, 0x62, 0x2e,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x61, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x28, 0x01, 0x30, 0x01, 0x12, 0x4f, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1d, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x10, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x12, 0x22, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x1c, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x41,
	0x6c, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x46, 0x0a, 0x09, 0x44, 0x72, 0x6f, 0x70, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1b,
	0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x59, 0x64,
	0x62, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4c, 0x0a, 0x11, 0x74, 0x65, 0x63,
	0x68, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x59, 0x64, 0x62, 0x5f, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x5f, 0x56, 0x31, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_ydb_topic_v1_proto_goTypes = []interface{}{
	(*Ydb_Topic.StreamWriteMessage_FromClient)(nil), // 0: Ydb.Topic.StreamWriteMessage.FromClient
	(*Ydb_Topic.StreamReadMessage_FromClient)(nil),  // 1: Ydb.Topic.StreamReadMessage.FromClient
	(*Ydb_Topic.CommitOffsetRequest)(nil),           // 2: Ydb.Topic.CommitOffsetRequest
	(*Ydb_Topic.CreateTopicRequest)(nil),            // 3: Ydb.Topic.CreateTopicRequest
	(*Ydb_Topic.DescribeTopicRequest)(nil),          // 4: Ydb.Topic.DescribeTopicRequest
	(*Ydb_Topic.DescribeConsumerRequest)(nil),       // 5: Ydb.Topic.DescribeConsumerRequest
	(*Ydb_Topic.AlterTopicRequest)(nil),             // 6: Ydb.Topic.AlterTopicRequest
	(*Ydb_Topic.DropTopicRequest)(nil),              // 7: Ydb.Topic.DropTopicRequest
	(*Ydb_Topic.StreamWriteMessage_FromServer)(nil), // 8: Ydb.Topic.StreamWriteMessage.FromServer
	(*Ydb_Topic.StreamReadMessage_FromServer)(nil),  // 9: Ydb.Topic.StreamReadMessage.FromServer
	(*Ydb_Topic.CommitOffsetResponse)(nil),          // 10: Ydb.Topic.CommitOffsetResponse
	(*Ydb_Topic.CreateTopicResponse)(nil),           // 11: Ydb.Topic.CreateTopicResponse
	(*Ydb_Topic.DescribeTopicResponse)(nil),         // 12: Ydb.Topic.DescribeTopicResponse
	(*Ydb_Topic.DescribeConsumerResponse)(nil),      // 13: Ydb.Topic.DescribeConsumerResponse
	(*Ydb_Topic.AlterTopicResponse)(nil),            // 14: Ydb.Topic.AlterTopicResponse
	(*Ydb_Topic.DropTopicResponse)(nil),             // 15: Ydb.Topic.DropTopicResponse
}
var file_ydb_topic_v1_proto_depIdxs = []int32{
	0,  // 0: Ydb.Topic.V1.TopicService.StreamWrite:input_type -> Ydb.Topic.StreamWriteMessage.FromClient
	1,  // 1: Ydb.Topic.V1.TopicService.StreamRead:input_type -> Ydb.Topic.StreamReadMessage.FromClient
	2,  // 2: Ydb.Topic.V1.TopicService.CommitOffset:input_type -> Ydb.Topic.CommitOffsetRequest
	3,  // 3: Ydb.Topic.V1.TopicService.CreateTopic:input_type -> Ydb.Topic.CreateTopicRequest
	4,  // 4: Ydb.Topic.V1.TopicService.DescribeTopic:input_type -> Ydb.Topic.DescribeTopicRequest
	5,  // 5: Ydb.Topic.V1.TopicService.DescribeConsumer:input_type -> Ydb.Topic.DescribeConsumerRequest
	6,  // 6: Ydb.Topic.V1.TopicService.AlterTopic:input_type -> Ydb.Topic.AlterTopicRequest
	7,  // 7: Ydb.Topic.V1.TopicService.DropTopic:input_type -> Ydb.Topic.DropTopicRequest
	8,  // 8: Ydb.Topic.V1.TopicService.StreamWrite:output_type -> Ydb.Topic.StreamWriteMessage.FromServer
	9,  // 9: Ydb.Topic.V1.TopicService.StreamRead:output_type -> Ydb.Topic.StreamReadMessage.FromServer
	10, // 10: Ydb.Topic.V1.TopicService.CommitOffset:output_type -> Ydb.Topic.CommitOffsetResponse
	11, // 11: Ydb.Topic.V1.TopicService.CreateTopic:output_type -> Ydb.Topic.CreateTopicResponse
	12, // 12: Ydb.Topic.V1.TopicService.DescribeTopic:output_type -> Ydb.Topic.DescribeTopicResponse
	13, // 13: Ydb.Topic.V1.TopicService.DescribeConsumer:output_type -> Ydb.Topic.DescribeConsumerResponse
	14, // 14: Ydb.Topic.V1.TopicService.AlterTopic:output_type -> Ydb.Topic.AlterTopicResponse
	15, // 15: Ydb.Topic.V1.TopicService.DropTopic:output_type -> Ydb.Topic.DropTopicResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_ydb_topic_v1_proto_init() }
func file_ydb_topic_v1_proto_init() {
	if File_ydb_topic_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ydb_topic_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ydb_topic_v1_proto_goTypes,
		DependencyIndexes: file_ydb_topic_v1_proto_depIdxs,
	}.Build()
	File_ydb_topic_v1_proto = out.File
	file_ydb_topic_v1_proto_rawDesc = nil
	file_ydb_topic_v1_proto_goTypes = nil
	file_ydb_topic_v1_proto_depIdxs = nil
}
