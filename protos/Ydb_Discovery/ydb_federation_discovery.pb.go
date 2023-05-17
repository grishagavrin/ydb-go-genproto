// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: protos/ydb_federation_discovery.proto

package Ydb_Discovery

import (
	Ydb_Operations "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Operations"
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

type DatabaseInfo_Status int32

const (
	DatabaseInfo_STATUS_UNSPECIFIED DatabaseInfo_Status = 0
	DatabaseInfo_AVAILABLE          DatabaseInfo_Status = 1
	DatabaseInfo_READ_ONLY          DatabaseInfo_Status = 2
	DatabaseInfo_UNAVAILABLE        DatabaseInfo_Status = 3
)

// Enum value maps for DatabaseInfo_Status.
var (
	DatabaseInfo_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "AVAILABLE",
		2: "READ_ONLY",
		3: "UNAVAILABLE",
	}
	DatabaseInfo_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"AVAILABLE":          1,
		"READ_ONLY":          2,
		"UNAVAILABLE":        3,
	}
)

func (x DatabaseInfo_Status) Enum() *DatabaseInfo_Status {
	p := new(DatabaseInfo_Status)
	*p = x
	return p
}

func (x DatabaseInfo_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DatabaseInfo_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_ydb_federation_discovery_proto_enumTypes[0].Descriptor()
}

func (DatabaseInfo_Status) Type() protoreflect.EnumType {
	return &file_protos_ydb_federation_discovery_proto_enumTypes[0]
}

func (x DatabaseInfo_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DatabaseInfo_Status.Descriptor instead.
func (DatabaseInfo_Status) EnumDescriptor() ([]byte, []int) {
	return file_protos_ydb_federation_discovery_proto_rawDescGZIP(), []int{0, 0}
}

type DatabaseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Path     string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Id       string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Endpoint string `protobuf:"bytes,4,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// for single datacenter databases
	Location string              `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	Status   DatabaseInfo_Status `protobuf:"varint,6,opt,name=status,proto3,enum=Ydb.FederationDiscovery.DatabaseInfo_Status" json:"status,omitempty"`
	// to determine this database priority on the client side
	Weight int64 `protobuf:"varint,7,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *DatabaseInfo) Reset() {
	*x = DatabaseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_federation_discovery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseInfo) ProtoMessage() {}

func (x *DatabaseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_federation_discovery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseInfo.ProtoReflect.Descriptor instead.
func (*DatabaseInfo) Descriptor() ([]byte, []int) {
	return file_protos_ydb_federation_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *DatabaseInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DatabaseInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *DatabaseInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DatabaseInfo) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *DatabaseInfo) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *DatabaseInfo) GetStatus() DatabaseInfo_Status {
	if x != nil {
		return x.Status
	}
	return DatabaseInfo_STATUS_UNSPECIFIED
}

func (x *DatabaseInfo) GetWeight() int64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type ListFederationDatabasesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListFederationDatabasesRequest) Reset() {
	*x = ListFederationDatabasesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_federation_discovery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFederationDatabasesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFederationDatabasesRequest) ProtoMessage() {}

func (x *ListFederationDatabasesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_federation_discovery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFederationDatabasesRequest.ProtoReflect.Descriptor instead.
func (*ListFederationDatabasesRequest) Descriptor() ([]byte, []int) {
	return file_protos_ydb_federation_discovery_proto_rawDescGZIP(), []int{1}
}

type ListFederationDatabasesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Operation contains the result of the request. Check the ydb_operation.proto.
	Operation *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *ListFederationDatabasesResponse) Reset() {
	*x = ListFederationDatabasesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_federation_discovery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFederationDatabasesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFederationDatabasesResponse) ProtoMessage() {}

func (x *ListFederationDatabasesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_federation_discovery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFederationDatabasesResponse.ProtoReflect.Descriptor instead.
func (*ListFederationDatabasesResponse) Descriptor() ([]byte, []int) {
	return file_protos_ydb_federation_discovery_proto_rawDescGZIP(), []int{2}
}

func (x *ListFederationDatabasesResponse) GetOperation() *Ydb_Operations.Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

type ListFederationDatabasesResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ControlPlaneEndpoint string          `protobuf:"bytes,1,opt,name=control_plane_endpoint,json=controlPlaneEndpoint,proto3" json:"control_plane_endpoint,omitempty"`
	FederationDatabases  []*DatabaseInfo `protobuf:"bytes,2,rep,name=federation_databases,json=federationDatabases,proto3" json:"federation_databases,omitempty"`
	SelfLocation         string          `protobuf:"bytes,3,opt,name=self_location,json=selfLocation,proto3" json:"self_location,omitempty"`
}

func (x *ListFederationDatabasesResult) Reset() {
	*x = ListFederationDatabasesResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_federation_discovery_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFederationDatabasesResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFederationDatabasesResult) ProtoMessage() {}

func (x *ListFederationDatabasesResult) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_federation_discovery_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFederationDatabasesResult.ProtoReflect.Descriptor instead.
func (*ListFederationDatabasesResult) Descriptor() ([]byte, []int) {
	return file_protos_ydb_federation_discovery_proto_rawDescGZIP(), []int{3}
}

func (x *ListFederationDatabasesResult) GetControlPlaneEndpoint() string {
	if x != nil {
		return x.ControlPlaneEndpoint
	}
	return ""
}

func (x *ListFederationDatabasesResult) GetFederationDatabases() []*DatabaseInfo {
	if x != nil {
		return x.FederationDatabases
	}
	return nil
}

func (x *ListFederationDatabasesResult) GetSelfLocation() string {
	if x != nil {
		return x.SelfLocation
	}
	return ""
}

var File_protos_ydb_federation_discovery_proto protoreflect.FileDescriptor

var file_protos_ydb_federation_discovery_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x66, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x59, 0x64, 0x62, 0x2e, 0x46, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x02, 0x0a,
	0x0c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x4f, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x52, 0x45, 0x41, 0x44, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x03, 0x22, 0x20, 0x0a, 0x1e,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5a,
	0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x37, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd4, 0x01, 0x0a, 0x1d, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x34, 0x0a, 0x16,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x58, 0x0a, 0x14, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x13, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x65, 0x6c, 0x66, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x6c, 0x66, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x67, 0x0a, 0x12, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x42, 0x10, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x59, 0x64, 0x62, 0x5f, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protos_ydb_federation_discovery_proto_rawDescOnce sync.Once
	file_protos_ydb_federation_discovery_proto_rawDescData = file_protos_ydb_federation_discovery_proto_rawDesc
)

func file_protos_ydb_federation_discovery_proto_rawDescGZIP() []byte {
	file_protos_ydb_federation_discovery_proto_rawDescOnce.Do(func() {
		file_protos_ydb_federation_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_ydb_federation_discovery_proto_rawDescData)
	})
	return file_protos_ydb_federation_discovery_proto_rawDescData
}

var file_protos_ydb_federation_discovery_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_ydb_federation_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_ydb_federation_discovery_proto_goTypes = []interface{}{
	(DatabaseInfo_Status)(0),                // 0: Ydb.FederationDiscovery.DatabaseInfo.Status
	(*DatabaseInfo)(nil),                    // 1: Ydb.FederationDiscovery.DatabaseInfo
	(*ListFederationDatabasesRequest)(nil),  // 2: Ydb.FederationDiscovery.ListFederationDatabasesRequest
	(*ListFederationDatabasesResponse)(nil), // 3: Ydb.FederationDiscovery.ListFederationDatabasesResponse
	(*ListFederationDatabasesResult)(nil),   // 4: Ydb.FederationDiscovery.ListFederationDatabasesResult
	(*Ydb_Operations.Operation)(nil),        // 5: Ydb.Operations.Operation
}
var file_protos_ydb_federation_discovery_proto_depIdxs = []int32{
	0, // 0: Ydb.FederationDiscovery.DatabaseInfo.status:type_name -> Ydb.FederationDiscovery.DatabaseInfo.Status
	5, // 1: Ydb.FederationDiscovery.ListFederationDatabasesResponse.operation:type_name -> Ydb.Operations.Operation
	1, // 2: Ydb.FederationDiscovery.ListFederationDatabasesResult.federation_databases:type_name -> Ydb.FederationDiscovery.DatabaseInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_ydb_federation_discovery_proto_init() }
func file_protos_ydb_federation_discovery_proto_init() {
	if File_protos_ydb_federation_discovery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_ydb_federation_discovery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseInfo); i {
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
		file_protos_ydb_federation_discovery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFederationDatabasesRequest); i {
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
		file_protos_ydb_federation_discovery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFederationDatabasesResponse); i {
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
		file_protos_ydb_federation_discovery_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFederationDatabasesResult); i {
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
			RawDescriptor: file_protos_ydb_federation_discovery_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_ydb_federation_discovery_proto_goTypes,
		DependencyIndexes: file_protos_ydb_federation_discovery_proto_depIdxs,
		EnumInfos:         file_protos_ydb_federation_discovery_proto_enumTypes,
		MessageInfos:      file_protos_ydb_federation_discovery_proto_msgTypes,
	}.Build()
	File_protos_ydb_federation_discovery_proto = out.File
	file_protos_ydb_federation_discovery_proto_rawDesc = nil
	file_protos_ydb_federation_discovery_proto_goTypes = nil
	file_protos_ydb_federation_discovery_proto_depIdxs = nil
}
