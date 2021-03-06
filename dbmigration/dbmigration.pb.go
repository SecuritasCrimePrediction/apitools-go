// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: dbmigration/dbmigration.proto

package dbmigration

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type MigrationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version       uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	LatestVersion int32  `protobuf:"varint,2,opt,name=latest_version,json=latestVersion,proto3" json:"latest_version,omitempty"`
	UpToDate      bool   `protobuf:"varint,3,opt,name=up_to_date,json=upToDate,proto3" json:"up_to_date,omitempty"`
	Dirty         bool   `protobuf:"varint,4,opt,name=dirty,proto3" json:"dirty,omitempty"`
}

func (x *MigrationStatus) Reset() {
	*x = MigrationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbmigration_dbmigration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigrationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrationStatus) ProtoMessage() {}

func (x *MigrationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_dbmigration_dbmigration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrationStatus.ProtoReflect.Descriptor instead.
func (*MigrationStatus) Descriptor() ([]byte, []int) {
	return file_dbmigration_dbmigration_proto_rawDescGZIP(), []int{0}
}

func (x *MigrationStatus) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *MigrationStatus) GetLatestVersion() int32 {
	if x != nil {
		return x.LatestVersion
	}
	return 0
}

func (x *MigrationStatus) GetUpToDate() bool {
	if x != nil {
		return x.UpToDate
	}
	return false
}

func (x *MigrationStatus) GetDirty() bool {
	if x != nil {
		return x.Dirty
	}
	return false
}

type DatabaseMigrateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DatabaseMigrateRequest) Reset() {
	*x = DatabaseMigrateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbmigration_dbmigration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseMigrateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseMigrateRequest) ProtoMessage() {}

func (x *DatabaseMigrateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dbmigration_dbmigration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseMigrateRequest.ProtoReflect.Descriptor instead.
func (*DatabaseMigrateRequest) Descriptor() ([]byte, []int) {
	return file_dbmigration_dbmigration_proto_rawDescGZIP(), []int{1}
}

type DatabaseMigrateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *MigrationStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DatabaseMigrateResponse) Reset() {
	*x = DatabaseMigrateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbmigration_dbmigration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseMigrateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseMigrateResponse) ProtoMessage() {}

func (x *DatabaseMigrateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dbmigration_dbmigration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseMigrateResponse.ProtoReflect.Descriptor instead.
func (*DatabaseMigrateResponse) Descriptor() ([]byte, []int) {
	return file_dbmigration_dbmigration_proto_rawDescGZIP(), []int{2}
}

func (x *DatabaseMigrateResponse) GetStatus() *MigrationStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type DatabaseStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DatabaseStatusRequest) Reset() {
	*x = DatabaseStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbmigration_dbmigration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseStatusRequest) ProtoMessage() {}

func (x *DatabaseStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dbmigration_dbmigration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseStatusRequest.ProtoReflect.Descriptor instead.
func (*DatabaseStatusRequest) Descriptor() ([]byte, []int) {
	return file_dbmigration_dbmigration_proto_rawDescGZIP(), []int{3}
}

type DatabaseStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *MigrationStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DatabaseStatusResponse) Reset() {
	*x = DatabaseStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbmigration_dbmigration_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseStatusResponse) ProtoMessage() {}

func (x *DatabaseStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dbmigration_dbmigration_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseStatusResponse.ProtoReflect.Descriptor instead.
func (*DatabaseStatusResponse) Descriptor() ([]byte, []int) {
	return file_dbmigration_dbmigration_proto_rawDescGZIP(), []int{4}
}

func (x *DatabaseStatusResponse) GetStatus() *MigrationStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type DatabaseRollbackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DatabaseRollbackRequest) Reset() {
	*x = DatabaseRollbackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbmigration_dbmigration_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseRollbackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseRollbackRequest) ProtoMessage() {}

func (x *DatabaseRollbackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dbmigration_dbmigration_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseRollbackRequest.ProtoReflect.Descriptor instead.
func (*DatabaseRollbackRequest) Descriptor() ([]byte, []int) {
	return file_dbmigration_dbmigration_proto_rawDescGZIP(), []int{5}
}

type DatabaseRollbackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *MigrationStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DatabaseRollbackResponse) Reset() {
	*x = DatabaseRollbackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbmigration_dbmigration_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseRollbackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseRollbackResponse) ProtoMessage() {}

func (x *DatabaseRollbackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dbmigration_dbmigration_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseRollbackResponse.ProtoReflect.Descriptor instead.
func (*DatabaseRollbackResponse) Descriptor() ([]byte, []int) {
	return file_dbmigration_dbmigration_proto_rawDescGZIP(), []int{6}
}

func (x *DatabaseRollbackResponse) GetStatus() *MigrationStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type DatabaseForceVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *DatabaseForceVersionRequest) Reset() {
	*x = DatabaseForceVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbmigration_dbmigration_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseForceVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseForceVersionRequest) ProtoMessage() {}

func (x *DatabaseForceVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dbmigration_dbmigration_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseForceVersionRequest.ProtoReflect.Descriptor instead.
func (*DatabaseForceVersionRequest) Descriptor() ([]byte, []int) {
	return file_dbmigration_dbmigration_proto_rawDescGZIP(), []int{7}
}

func (x *DatabaseForceVersionRequest) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type DatabaseForceVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *MigrationStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DatabaseForceVersionResponse) Reset() {
	*x = DatabaseForceVersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbmigration_dbmigration_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseForceVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseForceVersionResponse) ProtoMessage() {}

func (x *DatabaseForceVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dbmigration_dbmigration_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseForceVersionResponse.ProtoReflect.Descriptor instead.
func (*DatabaseForceVersionResponse) Descriptor() ([]byte, []int) {
	return file_dbmigration_dbmigration_proto_rawDescGZIP(), []int{8}
}

func (x *DatabaseForceVersionResponse) GetStatus() *MigrationStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_dbmigration_dbmigration_proto protoreflect.FileDescriptor

var file_dbmigration_dbmigration_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x64, 0x62, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x62,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x73, 0x69, 0x73, 0x2e, 0x72, 0x70, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a,
	0x0f, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x0a, 0x75, 0x70, 0x5f, 0x74, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x75, 0x70, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x64, 0x69, 0x72, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x64, 0x69, 0x72, 0x74, 0x79, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x51, 0x0a, 0x17, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x69, 0x73,
	0x2e, 0x72, 0x70, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x50, 0x0a, 0x16, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x69, 0x73, 0x2e, 0x72, 0x70, 0x2e, 0x64,
	0x65, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x19, 0x0a,
	0x17, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x52, 0x0a, 0x18, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x69, 0x73, 0x2e, 0x72, 0x70, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x37, 0x0a, 0x1b,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x56, 0x0a, 0x1c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x69, 0x73, 0x2e, 0x72, 0x70, 0x2e, 0x64,
	0x65, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xab, 0x04,
	0x0a, 0x10, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x7c, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x73, 0x69, 0x73, 0x2e, 0x72, 0x70, 0x2e, 0x64,
	0x65, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73,
	0x69, 0x73, 0x2e, 0x72, 0x70, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x64,
	0x65, 0x76, 0x2f, 0x64, 0x62, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a,
	0x12, 0x78, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x24, 0x2e, 0x73, 0x69, 0x73, 0x2e, 0x72, 0x70, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x69, 0x73, 0x2e, 0x72,
	0x70, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x64, 0x65, 0x76, 0x2f, 0x64, 0x62,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x80, 0x01, 0x0a, 0x10, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12,
	0x26, 0x2e, 0x73, 0x69, 0x73, 0x2e, 0x72, 0x70, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x69, 0x73, 0x2e, 0x72, 0x70,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x64, 0x65, 0x76, 0x2f, 0x64,
	0x62, 0x2f, 0x72, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x3a, 0x01, 0x2a, 0x12, 0x9b, 0x01,
	0x0a, 0x14, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x2e, 0x73, 0x69, 0x73, 0x2e, 0x72, 0x70, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x46,
	0x6f, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x69, 0x73, 0x2e, 0x72, 0x70, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x63, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x22, 0x1f, 0x2f, 0x64, 0x65, 0x76, 0x2f, 0x64, 0x62,
	0x2f, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x7b,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x3a, 0x01, 0x2a, 0x42, 0x49, 0x5a, 0x47, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x61, 0x73, 0x43, 0x72, 0x69, 0x6d, 0x65, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x64,
	0x62, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x64, 0x62, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dbmigration_dbmigration_proto_rawDescOnce sync.Once
	file_dbmigration_dbmigration_proto_rawDescData = file_dbmigration_dbmigration_proto_rawDesc
)

func file_dbmigration_dbmigration_proto_rawDescGZIP() []byte {
	file_dbmigration_dbmigration_proto_rawDescOnce.Do(func() {
		file_dbmigration_dbmigration_proto_rawDescData = protoimpl.X.CompressGZIP(file_dbmigration_dbmigration_proto_rawDescData)
	})
	return file_dbmigration_dbmigration_proto_rawDescData
}

var file_dbmigration_dbmigration_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_dbmigration_dbmigration_proto_goTypes = []interface{}{
	(*MigrationStatus)(nil),              // 0: sis.rp.dev.v1.MigrationStatus
	(*DatabaseMigrateRequest)(nil),       // 1: sis.rp.dev.v1.DatabaseMigrateRequest
	(*DatabaseMigrateResponse)(nil),      // 2: sis.rp.dev.v1.DatabaseMigrateResponse
	(*DatabaseStatusRequest)(nil),        // 3: sis.rp.dev.v1.DatabaseStatusRequest
	(*DatabaseStatusResponse)(nil),       // 4: sis.rp.dev.v1.DatabaseStatusResponse
	(*DatabaseRollbackRequest)(nil),      // 5: sis.rp.dev.v1.DatabaseRollbackRequest
	(*DatabaseRollbackResponse)(nil),     // 6: sis.rp.dev.v1.DatabaseRollbackResponse
	(*DatabaseForceVersionRequest)(nil),  // 7: sis.rp.dev.v1.DatabaseForceVersionRequest
	(*DatabaseForceVersionResponse)(nil), // 8: sis.rp.dev.v1.DatabaseForceVersionResponse
}
var file_dbmigration_dbmigration_proto_depIdxs = []int32{
	0, // 0: sis.rp.dev.v1.DatabaseMigrateResponse.status:type_name -> sis.rp.dev.v1.MigrationStatus
	0, // 1: sis.rp.dev.v1.DatabaseStatusResponse.status:type_name -> sis.rp.dev.v1.MigrationStatus
	0, // 2: sis.rp.dev.v1.DatabaseRollbackResponse.status:type_name -> sis.rp.dev.v1.MigrationStatus
	0, // 3: sis.rp.dev.v1.DatabaseForceVersionResponse.status:type_name -> sis.rp.dev.v1.MigrationStatus
	1, // 4: sis.rp.dev.v1.MigrationService.DatabaseMigrate:input_type -> sis.rp.dev.v1.DatabaseMigrateRequest
	3, // 5: sis.rp.dev.v1.MigrationService.DatabaseStatus:input_type -> sis.rp.dev.v1.DatabaseStatusRequest
	5, // 6: sis.rp.dev.v1.MigrationService.DatabaseRollback:input_type -> sis.rp.dev.v1.DatabaseRollbackRequest
	7, // 7: sis.rp.dev.v1.MigrationService.DatabaseForceVersion:input_type -> sis.rp.dev.v1.DatabaseForceVersionRequest
	2, // 8: sis.rp.dev.v1.MigrationService.DatabaseMigrate:output_type -> sis.rp.dev.v1.DatabaseMigrateResponse
	4, // 9: sis.rp.dev.v1.MigrationService.DatabaseStatus:output_type -> sis.rp.dev.v1.DatabaseStatusResponse
	6, // 10: sis.rp.dev.v1.MigrationService.DatabaseRollback:output_type -> sis.rp.dev.v1.DatabaseRollbackResponse
	8, // 11: sis.rp.dev.v1.MigrationService.DatabaseForceVersion:output_type -> sis.rp.dev.v1.DatabaseForceVersionResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_dbmigration_dbmigration_proto_init() }
func file_dbmigration_dbmigration_proto_init() {
	if File_dbmigration_dbmigration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dbmigration_dbmigration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigrationStatus); i {
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
		file_dbmigration_dbmigration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseMigrateRequest); i {
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
		file_dbmigration_dbmigration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseMigrateResponse); i {
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
		file_dbmigration_dbmigration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseStatusRequest); i {
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
		file_dbmigration_dbmigration_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseStatusResponse); i {
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
		file_dbmigration_dbmigration_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseRollbackRequest); i {
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
		file_dbmigration_dbmigration_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseRollbackResponse); i {
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
		file_dbmigration_dbmigration_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseForceVersionRequest); i {
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
		file_dbmigration_dbmigration_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseForceVersionResponse); i {
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
			RawDescriptor: file_dbmigration_dbmigration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dbmigration_dbmigration_proto_goTypes,
		DependencyIndexes: file_dbmigration_dbmigration_proto_depIdxs,
		MessageInfos:      file_dbmigration_dbmigration_proto_msgTypes,
	}.Build()
	File_dbmigration_dbmigration_proto = out.File
	file_dbmigration_dbmigration_proto_rawDesc = nil
	file_dbmigration_dbmigration_proto_goTypes = nil
	file_dbmigration_dbmigration_proto_depIdxs = nil
}
