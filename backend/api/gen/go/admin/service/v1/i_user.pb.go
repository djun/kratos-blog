// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: admin/service/v1/i_user.proto

package servicev1

import (
	_ "github.com/google/gnostic/openapiv3"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v11 "kratos-cms/api/gen/go/user/service/v1"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_admin_service_v1_i_user_proto protoreflect.FileDescriptor

var file_admin_service_v1_i_user_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x24, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8e,
	0x04, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61,
	0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11,
	0x12, 0x0f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x5f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x66, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x1d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x3a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x0f, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x6b, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x1a, 0x14, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x66, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42,
	0xb6, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x49, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x63,
	0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x53, 0x58, 0xaa, 0x02,
	0x10, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x10, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x3a, 0x3a, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_admin_service_v1_i_user_proto_goTypes = []any{
	(*v1.PagingRequest)(nil),      // 0: pagination.PagingRequest
	(*v11.GetUserRequest)(nil),    // 1: user.service.v1.GetUserRequest
	(*v11.CreateUserRequest)(nil), // 2: user.service.v1.CreateUserRequest
	(*v11.UpdateUserRequest)(nil), // 3: user.service.v1.UpdateUserRequest
	(*v11.DeleteUserRequest)(nil), // 4: user.service.v1.DeleteUserRequest
	(*v11.ListUserResponse)(nil),  // 5: user.service.v1.ListUserResponse
	(*v11.User)(nil),              // 6: user.service.v1.User
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_admin_service_v1_i_user_proto_depIdxs = []int32{
	0, // 0: admin.service.v1.UserService.ListUser:input_type -> pagination.PagingRequest
	1, // 1: admin.service.v1.UserService.GetUser:input_type -> user.service.v1.GetUserRequest
	2, // 2: admin.service.v1.UserService.CreateUser:input_type -> user.service.v1.CreateUserRequest
	3, // 3: admin.service.v1.UserService.UpdateUser:input_type -> user.service.v1.UpdateUserRequest
	4, // 4: admin.service.v1.UserService.DeleteUser:input_type -> user.service.v1.DeleteUserRequest
	5, // 5: admin.service.v1.UserService.ListUser:output_type -> user.service.v1.ListUserResponse
	6, // 6: admin.service.v1.UserService.GetUser:output_type -> user.service.v1.User
	6, // 7: admin.service.v1.UserService.CreateUser:output_type -> user.service.v1.User
	6, // 8: admin.service.v1.UserService.UpdateUser:output_type -> user.service.v1.User
	7, // 9: admin.service.v1.UserService.DeleteUser:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_admin_service_v1_i_user_proto_init() }
func file_admin_service_v1_i_user_proto_init() {
	if File_admin_service_v1_i_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_service_v1_i_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_service_v1_i_user_proto_goTypes,
		DependencyIndexes: file_admin_service_v1_i_user_proto_depIdxs,
	}.Build()
	File_admin_service_v1_i_user_proto = out.File
	file_admin_service_v1_i_user_proto_rawDesc = nil
	file_admin_service_v1_i_user_proto_goTypes = nil
	file_admin_service_v1_i_user_proto_depIdxs = nil
}
