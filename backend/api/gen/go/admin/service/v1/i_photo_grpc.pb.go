// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: admin/service/v1/i_photo.proto

package servicev1

import (
	context "context"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v11 "kratos-cms/api/gen/go/content/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PhotoService_ListPhoto_FullMethodName   = "/admin.service.v1.PhotoService/ListPhoto"
	PhotoService_GetPhoto_FullMethodName    = "/admin.service.v1.PhotoService/GetPhoto"
	PhotoService_CreatePhoto_FullMethodName = "/admin.service.v1.PhotoService/CreatePhoto"
	PhotoService_UpdatePhoto_FullMethodName = "/admin.service.v1.PhotoService/UpdatePhoto"
	PhotoService_DeletePhoto_FullMethodName = "/admin.service.v1.PhotoService/DeletePhoto"
)

// PhotoServiceClient is the client API for PhotoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 照片服务
type PhotoServiceClient interface {
	// 获取照片列表
	ListPhoto(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*v11.ListPhotoResponse, error)
	// 获取照片数据
	GetPhoto(ctx context.Context, in *v11.GetPhotoRequest, opts ...grpc.CallOption) (*v11.Photo, error)
	// 创建照片
	CreatePhoto(ctx context.Context, in *v11.CreatePhotoRequest, opts ...grpc.CallOption) (*v11.Photo, error)
	// 更新照片
	UpdatePhoto(ctx context.Context, in *v11.UpdatePhotoRequest, opts ...grpc.CallOption) (*v11.Photo, error)
	// 删除照片
	DeletePhoto(ctx context.Context, in *v11.DeletePhotoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type photoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPhotoServiceClient(cc grpc.ClientConnInterface) PhotoServiceClient {
	return &photoServiceClient{cc}
}

func (c *photoServiceClient) ListPhoto(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*v11.ListPhotoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v11.ListPhotoResponse)
	err := c.cc.Invoke(ctx, PhotoService_ListPhoto_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoServiceClient) GetPhoto(ctx context.Context, in *v11.GetPhotoRequest, opts ...grpc.CallOption) (*v11.Photo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v11.Photo)
	err := c.cc.Invoke(ctx, PhotoService_GetPhoto_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoServiceClient) CreatePhoto(ctx context.Context, in *v11.CreatePhotoRequest, opts ...grpc.CallOption) (*v11.Photo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v11.Photo)
	err := c.cc.Invoke(ctx, PhotoService_CreatePhoto_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoServiceClient) UpdatePhoto(ctx context.Context, in *v11.UpdatePhotoRequest, opts ...grpc.CallOption) (*v11.Photo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v11.Photo)
	err := c.cc.Invoke(ctx, PhotoService_UpdatePhoto_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoServiceClient) DeletePhoto(ctx context.Context, in *v11.DeletePhotoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PhotoService_DeletePhoto_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhotoServiceServer is the server API for PhotoService service.
// All implementations must embed UnimplementedPhotoServiceServer
// for forward compatibility.
//
// 照片服务
type PhotoServiceServer interface {
	// 获取照片列表
	ListPhoto(context.Context, *v1.PagingRequest) (*v11.ListPhotoResponse, error)
	// 获取照片数据
	GetPhoto(context.Context, *v11.GetPhotoRequest) (*v11.Photo, error)
	// 创建照片
	CreatePhoto(context.Context, *v11.CreatePhotoRequest) (*v11.Photo, error)
	// 更新照片
	UpdatePhoto(context.Context, *v11.UpdatePhotoRequest) (*v11.Photo, error)
	// 删除照片
	DeletePhoto(context.Context, *v11.DeletePhotoRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPhotoServiceServer()
}

// UnimplementedPhotoServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPhotoServiceServer struct{}

func (UnimplementedPhotoServiceServer) ListPhoto(context.Context, *v1.PagingRequest) (*v11.ListPhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPhoto not implemented")
}
func (UnimplementedPhotoServiceServer) GetPhoto(context.Context, *v11.GetPhotoRequest) (*v11.Photo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPhoto not implemented")
}
func (UnimplementedPhotoServiceServer) CreatePhoto(context.Context, *v11.CreatePhotoRequest) (*v11.Photo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePhoto not implemented")
}
func (UnimplementedPhotoServiceServer) UpdatePhoto(context.Context, *v11.UpdatePhotoRequest) (*v11.Photo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePhoto not implemented")
}
func (UnimplementedPhotoServiceServer) DeletePhoto(context.Context, *v11.DeletePhotoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePhoto not implemented")
}
func (UnimplementedPhotoServiceServer) mustEmbedUnimplementedPhotoServiceServer() {}
func (UnimplementedPhotoServiceServer) testEmbeddedByValue()                      {}

// UnsafePhotoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PhotoServiceServer will
// result in compilation errors.
type UnsafePhotoServiceServer interface {
	mustEmbedUnimplementedPhotoServiceServer()
}

func RegisterPhotoServiceServer(s grpc.ServiceRegistrar, srv PhotoServiceServer) {
	// If the following call pancis, it indicates UnimplementedPhotoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PhotoService_ServiceDesc, srv)
}

func _PhotoService_ListPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoServiceServer).ListPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoService_ListPhoto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoServiceServer).ListPhoto(ctx, req.(*v1.PagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoService_GetPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.GetPhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoServiceServer).GetPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoService_GetPhoto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoServiceServer).GetPhoto(ctx, req.(*v11.GetPhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoService_CreatePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.CreatePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoServiceServer).CreatePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoService_CreatePhoto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoServiceServer).CreatePhoto(ctx, req.(*v11.CreatePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoService_UpdatePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.UpdatePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoServiceServer).UpdatePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoService_UpdatePhoto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoServiceServer).UpdatePhoto(ctx, req.(*v11.UpdatePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoService_DeletePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.DeletePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoServiceServer).DeletePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoService_DeletePhoto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoServiceServer).DeletePhoto(ctx, req.(*v11.DeletePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PhotoService_ServiceDesc is the grpc.ServiceDesc for PhotoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PhotoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.service.v1.PhotoService",
	HandlerType: (*PhotoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPhoto",
			Handler:    _PhotoService_ListPhoto_Handler,
		},
		{
			MethodName: "GetPhoto",
			Handler:    _PhotoService_GetPhoto_Handler,
		},
		{
			MethodName: "CreatePhoto",
			Handler:    _PhotoService_CreatePhoto_Handler,
		},
		{
			MethodName: "UpdatePhoto",
			Handler:    _PhotoService_UpdatePhoto_Handler,
		},
		{
			MethodName: "DeletePhoto",
			Handler:    _PhotoService_DeletePhoto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/service/v1/i_photo.proto",
}
