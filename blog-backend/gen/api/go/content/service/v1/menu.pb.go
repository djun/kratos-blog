// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: content/service/v1/menu.proto

package servicev1

import (
	v1 "github.com/tx7do/kratos-bootstrap/gen/api/go/pagination/v1"
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

// 菜单
type Menu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Url        *string `protobuf:"bytes,3,opt,name=url,proto3,oneof" json:"url,omitempty"`
	Priority   *int32  `protobuf:"varint,4,opt,name=priority,proto3,oneof" json:"priority,omitempty"`
	Target     *string `protobuf:"bytes,5,opt,name=target,proto3,oneof" json:"target,omitempty"`
	Icon       *string `protobuf:"bytes,6,opt,name=icon,proto3,oneof" json:"icon,omitempty"`
	ParentId   *uint32 `protobuf:"varint,7,opt,name=parentId,proto3,oneof" json:"parentId,omitempty"`
	Team       *string `protobuf:"bytes,8,opt,name=team,proto3,oneof" json:"team,omitempty"`
	CreateTime *string `protobuf:"bytes,9,opt,name=createTime,proto3,oneof" json:"createTime,omitempty"`
	UpdateTime *string `protobuf:"bytes,10,opt,name=updateTime,proto3,oneof" json:"updateTime,omitempty"`
}

func (x *Menu) Reset() {
	*x = Menu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_service_v1_menu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Menu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Menu) ProtoMessage() {}

func (x *Menu) ProtoReflect() protoreflect.Message {
	mi := &file_content_service_v1_menu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Menu.ProtoReflect.Descriptor instead.
func (*Menu) Descriptor() ([]byte, []int) {
	return file_content_service_v1_menu_proto_rawDescGZIP(), []int{0}
}

func (x *Menu) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Menu) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Menu) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *Menu) GetPriority() int32 {
	if x != nil && x.Priority != nil {
		return *x.Priority
	}
	return 0
}

func (x *Menu) GetTarget() string {
	if x != nil && x.Target != nil {
		return *x.Target
	}
	return ""
}

func (x *Menu) GetIcon() string {
	if x != nil && x.Icon != nil {
		return *x.Icon
	}
	return ""
}

func (x *Menu) GetParentId() uint32 {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return 0
}

func (x *Menu) GetTeam() string {
	if x != nil && x.Team != nil {
		return *x.Team
	}
	return ""
}

func (x *Menu) GetCreateTime() string {
	if x != nil && x.CreateTime != nil {
		return *x.CreateTime
	}
	return ""
}

func (x *Menu) GetUpdateTime() string {
	if x != nil && x.UpdateTime != nil {
		return *x.UpdateTime
	}
	return ""
}

// 回应 - 菜单列表
type ListMenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Menu `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int32   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListMenuResponse) Reset() {
	*x = ListMenuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_service_v1_menu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMenuResponse) ProtoMessage() {}

func (x *ListMenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_content_service_v1_menu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMenuResponse.ProtoReflect.Descriptor instead.
func (*ListMenuResponse) Descriptor() ([]byte, []int) {
	return file_content_service_v1_menu_proto_rawDescGZIP(), []int{1}
}

func (x *ListMenuResponse) GetItems() []*Menu {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListMenuResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

// 请求 - 菜单数据
type GetMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMenuRequest) Reset() {
	*x = GetMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_service_v1_menu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMenuRequest) ProtoMessage() {}

func (x *GetMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_content_service_v1_menu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMenuRequest.ProtoReflect.Descriptor instead.
func (*GetMenuRequest) Descriptor() ([]byte, []int) {
	return file_content_service_v1_menu_proto_rawDescGZIP(), []int{2}
}

func (x *GetMenuRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 请求 - 创建菜单
type CreateMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Menu       *Menu   `protobuf:"bytes,1,opt,name=menu,proto3" json:"menu,omitempty"`
	OperatorId *uint32 `protobuf:"varint,2,opt,name=operatorId,proto3,oneof" json:"operatorId,omitempty"`
}

func (x *CreateMenuRequest) Reset() {
	*x = CreateMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_service_v1_menu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMenuRequest) ProtoMessage() {}

func (x *CreateMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_content_service_v1_menu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMenuRequest.ProtoReflect.Descriptor instead.
func (*CreateMenuRequest) Descriptor() ([]byte, []int) {
	return file_content_service_v1_menu_proto_rawDescGZIP(), []int{3}
}

func (x *CreateMenuRequest) GetMenu() *Menu {
	if x != nil {
		return x.Menu
	}
	return nil
}

func (x *CreateMenuRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

// 请求 - 更新菜单
type UpdateMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Menu       *Menu   `protobuf:"bytes,2,opt,name=menu,proto3" json:"menu,omitempty"`
	OperatorId *uint32 `protobuf:"varint,3,opt,name=operatorId,proto3,oneof" json:"operatorId,omitempty"`
}

func (x *UpdateMenuRequest) Reset() {
	*x = UpdateMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_service_v1_menu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMenuRequest) ProtoMessage() {}

func (x *UpdateMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_content_service_v1_menu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMenuRequest.ProtoReflect.Descriptor instead.
func (*UpdateMenuRequest) Descriptor() ([]byte, []int) {
	return file_content_service_v1_menu_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateMenuRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateMenuRequest) GetMenu() *Menu {
	if x != nil {
		return x.Menu
	}
	return nil
}

func (x *UpdateMenuRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

// 请求 - 删除菜单
type DeleteMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OperatorId *uint32 `protobuf:"varint,2,opt,name=operatorId,proto3,oneof" json:"operatorId,omitempty"`
}

func (x *DeleteMenuRequest) Reset() {
	*x = DeleteMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_service_v1_menu_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMenuRequest) ProtoMessage() {}

func (x *DeleteMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_content_service_v1_menu_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMenuRequest.ProtoReflect.Descriptor instead.
func (*DeleteMenuRequest) Descriptor() ([]byte, []int) {
	return file_content_service_v1_menu_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteMenuRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteMenuRequest) GetOperatorId() uint32 {
	if x != nil && x.OperatorId != nil {
		return *x.OperatorId
	}
	return 0
}

var File_content_service_v1_menu_proto protoreflect.FileDescriptor

var file_content_service_v1_menu_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x87, 0x03, 0x0a, 0x04, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x05, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x06, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x23, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x06, 0x0a,
	0x04, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x58, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x75, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x6d,
	0x65, 0x6e, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x6e, 0x75, 0x52, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52,
	0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x85, 0x01,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04, 0x6d, 0x65, 0x6e,
	0x75, 0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d,
	0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00,
	0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x32, 0x98,
	0x03, 0x0a, 0x0b, 0x4d, 0x65, 0x6e, 0x75, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d,
	0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0xc1, 0x01, 0x0a, 0x16, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x4d, 0x65, 0x6e, 0x75, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x32, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x63, 0x6d, 0x73, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x12, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_content_service_v1_menu_proto_rawDescOnce sync.Once
	file_content_service_v1_menu_proto_rawDescData = file_content_service_v1_menu_proto_rawDesc
)

func file_content_service_v1_menu_proto_rawDescGZIP() []byte {
	file_content_service_v1_menu_proto_rawDescOnce.Do(func() {
		file_content_service_v1_menu_proto_rawDescData = protoimpl.X.CompressGZIP(file_content_service_v1_menu_proto_rawDescData)
	})
	return file_content_service_v1_menu_proto_rawDescData
}

var file_content_service_v1_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_content_service_v1_menu_proto_goTypes = []interface{}{
	(*Menu)(nil),              // 0: content.service.v1.Menu
	(*ListMenuResponse)(nil),  // 1: content.service.v1.ListMenuResponse
	(*GetMenuRequest)(nil),    // 2: content.service.v1.GetMenuRequest
	(*CreateMenuRequest)(nil), // 3: content.service.v1.CreateMenuRequest
	(*UpdateMenuRequest)(nil), // 4: content.service.v1.UpdateMenuRequest
	(*DeleteMenuRequest)(nil), // 5: content.service.v1.DeleteMenuRequest
	(*v1.PagingRequest)(nil),  // 6: pagination.PagingRequest
	(*emptypb.Empty)(nil),     // 7: google.protobuf.Empty
}
var file_content_service_v1_menu_proto_depIdxs = []int32{
	0, // 0: content.service.v1.ListMenuResponse.items:type_name -> content.service.v1.Menu
	0, // 1: content.service.v1.CreateMenuRequest.menu:type_name -> content.service.v1.Menu
	0, // 2: content.service.v1.UpdateMenuRequest.menu:type_name -> content.service.v1.Menu
	6, // 3: content.service.v1.MenuService.ListMenu:input_type -> pagination.PagingRequest
	2, // 4: content.service.v1.MenuService.GetMenu:input_type -> content.service.v1.GetMenuRequest
	3, // 5: content.service.v1.MenuService.CreateMenu:input_type -> content.service.v1.CreateMenuRequest
	4, // 6: content.service.v1.MenuService.UpdateMenu:input_type -> content.service.v1.UpdateMenuRequest
	5, // 7: content.service.v1.MenuService.DeleteMenu:input_type -> content.service.v1.DeleteMenuRequest
	1, // 8: content.service.v1.MenuService.ListMenu:output_type -> content.service.v1.ListMenuResponse
	0, // 9: content.service.v1.MenuService.GetMenu:output_type -> content.service.v1.Menu
	0, // 10: content.service.v1.MenuService.CreateMenu:output_type -> content.service.v1.Menu
	0, // 11: content.service.v1.MenuService.UpdateMenu:output_type -> content.service.v1.Menu
	7, // 12: content.service.v1.MenuService.DeleteMenu:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_content_service_v1_menu_proto_init() }
func file_content_service_v1_menu_proto_init() {
	if File_content_service_v1_menu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_content_service_v1_menu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Menu); i {
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
		file_content_service_v1_menu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMenuResponse); i {
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
		file_content_service_v1_menu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMenuRequest); i {
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
		file_content_service_v1_menu_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMenuRequest); i {
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
		file_content_service_v1_menu_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMenuRequest); i {
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
		file_content_service_v1_menu_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMenuRequest); i {
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
	file_content_service_v1_menu_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_content_service_v1_menu_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_content_service_v1_menu_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_content_service_v1_menu_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_content_service_v1_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_content_service_v1_menu_proto_goTypes,
		DependencyIndexes: file_content_service_v1_menu_proto_depIdxs,
		MessageInfos:      file_content_service_v1_menu_proto_msgTypes,
	}.Build()
	File_content_service_v1_menu_proto = out.File
	file_content_service_v1_menu_proto_rawDesc = nil
	file_content_service_v1_menu_proto_goTypes = nil
	file_content_service_v1_menu_proto_depIdxs = nil
}
