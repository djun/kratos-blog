// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: front/service/v1/i_link.proto

package servicev1

import (
	_ "github.com/google/gnostic/openapiv3"
	v1 "github.com/tx7do/kratos-bootstrap/gen/api/go/pagination/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v11 "kratos-cms/gen/api/go/content/service/v1"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_front_service_v1_i_link_proto protoreflect.FileDescriptor

var file_front_service_v1_i_link_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x24, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xbe, 0x05, 0x0a, 0x0b, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x82, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x19,
	0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x35, 0xba, 0x47, 0x1c, 0x5a, 0x1a, 0x0a, 0x18, 0x0a, 0x14, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x83, 0x01, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b,
	0x22, 0x3a, 0xba, 0x47, 0x1c, 0x5a, 0x1a, 0x0a, 0x18, 0x0a, 0x14, 0x4f, 0x41, 0x75, 0x74, 0x68,
	0x32, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12,
	0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x87, 0x01, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x25, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x38, 0xba, 0x47,
	0x1c, 0x5a, 0x1a, 0x0a, 0x18, 0x0a, 0x14, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x8f, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x40, 0xba, 0x47, 0x1c, 0x5a, 0x1a, 0x0a, 0x18, 0x0a,
	0x14, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x04, 0x6c,
	0x69, 0x6e, 0x6b, 0x1a, 0x13, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69,
	0x6e, 0x6b, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x87, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3a, 0xba, 0x47, 0x1c, 0x5a, 0x1a, 0x0a, 0x18, 0x0a,
	0x14, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a, 0x13, 0x2f,
	0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x42, 0xb6, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x49, 0x4c, 0x69,
	0x6e, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2d, 0x63, 0x6d, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f,
	0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x46, 0x53,
	0x58, 0xaa, 0x02, 0x10, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x3a, 0x3a,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_front_service_v1_i_link_proto_goTypes = []interface{}{
	(*v1.PagingRequest)(nil),      // 0: pagination.PagingRequest
	(*v11.GetLinkRequest)(nil),    // 1: content.service.v1.GetLinkRequest
	(*v11.CreateLinkRequest)(nil), // 2: content.service.v1.CreateLinkRequest
	(*v11.UpdateLinkRequest)(nil), // 3: content.service.v1.UpdateLinkRequest
	(*v11.DeleteLinkRequest)(nil), // 4: content.service.v1.DeleteLinkRequest
	(*v11.ListLinkResponse)(nil),  // 5: content.service.v1.ListLinkResponse
	(*v11.Link)(nil),              // 6: content.service.v1.Link
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_front_service_v1_i_link_proto_depIdxs = []int32{
	0, // 0: front.service.v1.LinkService.ListLink:input_type -> pagination.PagingRequest
	1, // 1: front.service.v1.LinkService.GetLink:input_type -> content.service.v1.GetLinkRequest
	2, // 2: front.service.v1.LinkService.CreateLink:input_type -> content.service.v1.CreateLinkRequest
	3, // 3: front.service.v1.LinkService.UpdateLink:input_type -> content.service.v1.UpdateLinkRequest
	4, // 4: front.service.v1.LinkService.DeleteLink:input_type -> content.service.v1.DeleteLinkRequest
	5, // 5: front.service.v1.LinkService.ListLink:output_type -> content.service.v1.ListLinkResponse
	6, // 6: front.service.v1.LinkService.GetLink:output_type -> content.service.v1.Link
	6, // 7: front.service.v1.LinkService.CreateLink:output_type -> content.service.v1.Link
	6, // 8: front.service.v1.LinkService.UpdateLink:output_type -> content.service.v1.Link
	7, // 9: front.service.v1.LinkService.DeleteLink:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_front_service_v1_i_link_proto_init() }
func file_front_service_v1_i_link_proto_init() {
	if File_front_service_v1_i_link_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_front_service_v1_i_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_front_service_v1_i_link_proto_goTypes,
		DependencyIndexes: file_front_service_v1_i_link_proto_depIdxs,
	}.Build()
	File_front_service_v1_i_link_proto = out.File
	file_front_service_v1_i_link_proto_rawDesc = nil
	file_front_service_v1_i_link_proto_goTypes = nil
	file_front_service_v1_i_link_proto_depIdxs = nil
}
