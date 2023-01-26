// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: admin/service/v1/i_system.proto

package v1

import (
	_ "github.com/google/gnostic/openapiv3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pagination "kratos-blog/gen/api/go/pagination"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type System struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Theme        *int32  `protobuf:"varint,2,opt,name=theme,proto3,oneof" json:"theme,omitempty"`
	Title        *string `protobuf:"bytes,3,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Keywords     *string `protobuf:"bytes,4,opt,name=keywords,proto3,oneof" json:"keywords,omitempty"`
	Description  *string `protobuf:"bytes,5,opt,name=description,proto3,oneof" json:"description,omitempty"`
	RecordNumber *string `protobuf:"bytes,6,opt,name=recordNumber,proto3,oneof" json:"recordNumber,omitempty"`
}

func (x *System) Reset() {
	*x = System{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_service_v1_i_system_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *System) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*System) ProtoMessage() {}

func (x *System) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_v1_i_system_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use System.ProtoReflect.Descriptor instead.
func (*System) Descriptor() ([]byte, []int) {
	return file_admin_service_v1_i_system_proto_rawDescGZIP(), []int{0}
}

func (x *System) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *System) GetTheme() int32 {
	if x != nil && x.Theme != nil {
		return *x.Theme
	}
	return 0
}

func (x *System) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *System) GetKeywords() string {
	if x != nil && x.Keywords != nil {
		return *x.Keywords
	}
	return ""
}

func (x *System) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *System) GetRecordNumber() string {
	if x != nil && x.RecordNumber != nil {
		return *x.RecordNumber
	}
	return ""
}

// 回应 - 系统设置列表
type ListSystemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*System `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int32     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListSystemResponse) Reset() {
	*x = ListSystemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_service_v1_i_system_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSystemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSystemResponse) ProtoMessage() {}

func (x *ListSystemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_v1_i_system_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSystemResponse.ProtoReflect.Descriptor instead.
func (*ListSystemResponse) Descriptor() ([]byte, []int) {
	return file_admin_service_v1_i_system_proto_rawDescGZIP(), []int{1}
}

func (x *ListSystemResponse) GetItems() []*System {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListSystemResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

// 请求 - 系统设置数据
type GetSystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSystemRequest) Reset() {
	*x = GetSystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_service_v1_i_system_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSystemRequest) ProtoMessage() {}

func (x *GetSystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_v1_i_system_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSystemRequest.ProtoReflect.Descriptor instead.
func (*GetSystemRequest) Descriptor() ([]byte, []int) {
	return file_admin_service_v1_i_system_proto_rawDescGZIP(), []int{2}
}

func (x *GetSystemRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 请求 - 创建系统设置
type CreateSystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	System     *System `protobuf:"bytes,1,opt,name=system,proto3" json:"system,omitempty"`
	OperatorId *uint64 `protobuf:"varint,2,opt,name=operatorId,proto3,oneof" json:"operatorId,omitempty"`
}

func (x *CreateSystemRequest) Reset() {
	*x = CreateSystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_service_v1_i_system_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSystemRequest) ProtoMessage() {}

func (x *CreateSystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_v1_i_system_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSystemRequest.ProtoReflect.Descriptor instead.
func (*CreateSystemRequest) Descriptor() ([]byte, []int) {
	return file_admin_service_v1_i_system_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSystemRequest) GetSystem() *System {
	if x != nil {
		return x.System
	}
	return nil
}

func (x *CreateSystemRequest) GetOperatorId() uint64 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

// 请求 - 更新系统设置
type UpdateSystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	System     *System `protobuf:"bytes,2,opt,name=system,proto3" json:"system,omitempty"`
	OperatorId *uint64 `protobuf:"varint,3,opt,name=operatorId,proto3,oneof" json:"operatorId,omitempty"`
}

func (x *UpdateSystemRequest) Reset() {
	*x = UpdateSystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_service_v1_i_system_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSystemRequest) ProtoMessage() {}

func (x *UpdateSystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_v1_i_system_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSystemRequest.ProtoReflect.Descriptor instead.
func (*UpdateSystemRequest) Descriptor() ([]byte, []int) {
	return file_admin_service_v1_i_system_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSystemRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateSystemRequest) GetSystem() *System {
	if x != nil {
		return x.System
	}
	return nil
}

func (x *UpdateSystemRequest) GetOperatorId() uint64 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

// 请求 - 删除系统设置
type DeleteSystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OperatorId *uint64 `protobuf:"varint,2,opt,name=operatorId,proto3,oneof" json:"operatorId,omitempty"`
}

func (x *DeleteSystemRequest) Reset() {
	*x = DeleteSystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_service_v1_i_system_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSystemRequest) ProtoMessage() {}

func (x *DeleteSystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_v1_i_system_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSystemRequest.ProtoReflect.Descriptor instead.
func (*DeleteSystemRequest) Descriptor() ([]byte, []int) {
	return file_admin_service_v1_i_system_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteSystemRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteSystemRequest) GetOperatorId() uint64 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

var File_admin_service_v1_i_system_proto protoreflect.FileDescriptor

var file_admin_service_v1_i_system_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x27, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x74, 0x68,
	0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x81, 0x02, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x74,
	0x68, 0x65, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x74, 0x68,
	0x65, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x04, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88,
	0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x5a, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x7b, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x23, 0x0a,
	0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49,
	0x64, 0x22, 0x8b, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x23, 0x0a, 0x0a, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x48,
	0x00, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22,
	0x59, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x32, 0xb1, 0x04, 0x0a, 0x0d, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x0a,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x12, 0x67, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x12, 0x22, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22,
	0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6b, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x25, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x62, 0x6c, 0x6f, 0x67,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x75, 0x0a, 0x0c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x25, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x24, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1e, 0x3a, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x14, 0x2f, 0x62, 0x6c,
	0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x6b, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x12, 0x25, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x64,
	0x5a, 0x2a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0xba, 0x47, 0x35, 0x12,
	0x33, 0x0a, 0x18, 0xe7, 0xb3, 0xbb, 0xe7, 0xbb, 0x9f, 0xe9, 0x85, 0x8d, 0xe7, 0xbd, 0xae, 0xe7,
	0xae, 0xa1, 0xe7, 0x90, 0x86, 0xe6, 0x8e, 0xa5, 0xe5, 0x8f, 0xa3, 0x12, 0x12, 0xe7, 0xb3, 0xbb,
	0xe7, 0xbb, 0x9f, 0xe9, 0x85, 0x8d, 0xe7, 0xbd, 0xae, 0xe7, 0xae, 0xa1, 0xe7, 0x90, 0x86, 0x32,
	0x03, 0x31, 0x2e, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_service_v1_i_system_proto_rawDescOnce sync.Once
	file_admin_service_v1_i_system_proto_rawDescData = file_admin_service_v1_i_system_proto_rawDesc
)

func file_admin_service_v1_i_system_proto_rawDescGZIP() []byte {
	file_admin_service_v1_i_system_proto_rawDescOnce.Do(func() {
		file_admin_service_v1_i_system_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_service_v1_i_system_proto_rawDescData)
	})
	return file_admin_service_v1_i_system_proto_rawDescData
}

var file_admin_service_v1_i_system_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_admin_service_v1_i_system_proto_goTypes = []interface{}{
	(*System)(nil),                   // 0: admin.service.v1.System
	(*ListSystemResponse)(nil),       // 1: admin.service.v1.ListSystemResponse
	(*GetSystemRequest)(nil),         // 2: admin.service.v1.GetSystemRequest
	(*CreateSystemRequest)(nil),      // 3: admin.service.v1.CreateSystemRequest
	(*UpdateSystemRequest)(nil),      // 4: admin.service.v1.UpdateSystemRequest
	(*DeleteSystemRequest)(nil),      // 5: admin.service.v1.DeleteSystemRequest
	(*pagination.PagingRequest)(nil), // 6: pagination.PagingRequest
	(*emptypb.Empty)(nil),            // 7: google.protobuf.Empty
}
var file_admin_service_v1_i_system_proto_depIdxs = []int32{
	0, // 0: admin.service.v1.ListSystemResponse.items:type_name -> admin.service.v1.System
	0, // 1: admin.service.v1.CreateSystemRequest.system:type_name -> admin.service.v1.System
	0, // 2: admin.service.v1.UpdateSystemRequest.system:type_name -> admin.service.v1.System
	6, // 3: admin.service.v1.SystemService.ListSystem:input_type -> pagination.PagingRequest
	2, // 4: admin.service.v1.SystemService.GetSystem:input_type -> admin.service.v1.GetSystemRequest
	3, // 5: admin.service.v1.SystemService.CreateSystem:input_type -> admin.service.v1.CreateSystemRequest
	4, // 6: admin.service.v1.SystemService.UpdateSystem:input_type -> admin.service.v1.UpdateSystemRequest
	5, // 7: admin.service.v1.SystemService.DeleteSystem:input_type -> admin.service.v1.DeleteSystemRequest
	1, // 8: admin.service.v1.SystemService.ListSystem:output_type -> admin.service.v1.ListSystemResponse
	0, // 9: admin.service.v1.SystemService.GetSystem:output_type -> admin.service.v1.System
	0, // 10: admin.service.v1.SystemService.CreateSystem:output_type -> admin.service.v1.System
	0, // 11: admin.service.v1.SystemService.UpdateSystem:output_type -> admin.service.v1.System
	7, // 12: admin.service.v1.SystemService.DeleteSystem:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_admin_service_v1_i_system_proto_init() }
func file_admin_service_v1_i_system_proto_init() {
	if File_admin_service_v1_i_system_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_service_v1_i_system_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*System); i {
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
		file_admin_service_v1_i_system_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSystemResponse); i {
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
		file_admin_service_v1_i_system_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSystemRequest); i {
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
		file_admin_service_v1_i_system_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSystemRequest); i {
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
		file_admin_service_v1_i_system_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSystemRequest); i {
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
		file_admin_service_v1_i_system_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSystemRequest); i {
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
	file_admin_service_v1_i_system_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_admin_service_v1_i_system_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_admin_service_v1_i_system_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_admin_service_v1_i_system_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_service_v1_i_system_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_service_v1_i_system_proto_goTypes,
		DependencyIndexes: file_admin_service_v1_i_system_proto_depIdxs,
		MessageInfos:      file_admin_service_v1_i_system_proto_msgTypes,
	}.Build()
	File_admin_service_v1_i_system_proto = out.File
	file_admin_service_v1_i_system_proto_rawDesc = nil
	file_admin_service_v1_i_system_proto_goTypes = nil
	file_admin_service_v1_i_system_proto_depIdxs = nil
}
