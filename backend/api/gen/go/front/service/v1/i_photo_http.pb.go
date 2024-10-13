// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             (unknown)
// source: front/service/v1/i_photo.proto

package servicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v11 "kratos-cms/api/gen/go/content/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationPhotoServiceCreatePhoto = "/front.service.v1.PhotoService/CreatePhoto"
const OperationPhotoServiceDeletePhoto = "/front.service.v1.PhotoService/DeletePhoto"
const OperationPhotoServiceGetPhoto = "/front.service.v1.PhotoService/GetPhoto"
const OperationPhotoServiceListPhoto = "/front.service.v1.PhotoService/ListPhoto"
const OperationPhotoServiceUpdatePhoto = "/front.service.v1.PhotoService/UpdatePhoto"

type PhotoServiceHTTPServer interface {
	// CreatePhoto 创建照片
	CreatePhoto(context.Context, *v11.CreatePhotoRequest) (*v11.Photo, error)
	// DeletePhoto 删除照片
	DeletePhoto(context.Context, *v11.DeletePhotoRequest) (*emptypb.Empty, error)
	// GetPhoto 获取照片数据
	GetPhoto(context.Context, *v11.GetPhotoRequest) (*v11.Photo, error)
	// ListPhoto 获取照片列表
	ListPhoto(context.Context, *v1.PagingRequest) (*v11.ListPhotoResponse, error)
	// UpdatePhoto 更新照片
	UpdatePhoto(context.Context, *v11.UpdatePhotoRequest) (*v11.Photo, error)
}

func RegisterPhotoServiceHTTPServer(s *http.Server, srv PhotoServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/blog/v1/photos", _PhotoService_ListPhoto0_HTTP_Handler(srv))
	r.GET("/blog/v1/photos/{id}", _PhotoService_GetPhoto0_HTTP_Handler(srv))
	r.POST("/blog/v1/photos", _PhotoService_CreatePhoto0_HTTP_Handler(srv))
	r.PUT("/blog/v1/photos/{id}", _PhotoService_UpdatePhoto0_HTTP_Handler(srv))
	r.DELETE("/blog/v1/photos/{id}", _PhotoService_DeletePhoto0_HTTP_Handler(srv))
}

func _PhotoService_ListPhoto0_HTTP_Handler(srv PhotoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPhotoServiceListPhoto)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPhoto(ctx, req.(*v1.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.ListPhotoResponse)
		return ctx.Result(200, reply)
	}
}

func _PhotoService_GetPhoto0_HTTP_Handler(srv PhotoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.GetPhotoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPhotoServiceGetPhoto)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPhoto(ctx, req.(*v11.GetPhotoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.Photo)
		return ctx.Result(200, reply)
	}
}

func _PhotoService_CreatePhoto0_HTTP_Handler(srv PhotoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.CreatePhotoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPhotoServiceCreatePhoto)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePhoto(ctx, req.(*v11.CreatePhotoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.Photo)
		return ctx.Result(200, reply)
	}
}

func _PhotoService_UpdatePhoto0_HTTP_Handler(srv PhotoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.UpdatePhotoRequest
		if err := ctx.Bind(&in.Photo); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPhotoServiceUpdatePhoto)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePhoto(ctx, req.(*v11.UpdatePhotoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.Photo)
		return ctx.Result(200, reply)
	}
}

func _PhotoService_DeletePhoto0_HTTP_Handler(srv PhotoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.DeletePhotoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPhotoServiceDeletePhoto)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeletePhoto(ctx, req.(*v11.DeletePhotoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type PhotoServiceHTTPClient interface {
	CreatePhoto(ctx context.Context, req *v11.CreatePhotoRequest, opts ...http.CallOption) (rsp *v11.Photo, err error)
	DeletePhoto(ctx context.Context, req *v11.DeletePhotoRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetPhoto(ctx context.Context, req *v11.GetPhotoRequest, opts ...http.CallOption) (rsp *v11.Photo, err error)
	ListPhoto(ctx context.Context, req *v1.PagingRequest, opts ...http.CallOption) (rsp *v11.ListPhotoResponse, err error)
	UpdatePhoto(ctx context.Context, req *v11.UpdatePhotoRequest, opts ...http.CallOption) (rsp *v11.Photo, err error)
}

type PhotoServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewPhotoServiceHTTPClient(client *http.Client) PhotoServiceHTTPClient {
	return &PhotoServiceHTTPClientImpl{client}
}

func (c *PhotoServiceHTTPClientImpl) CreatePhoto(ctx context.Context, in *v11.CreatePhotoRequest, opts ...http.CallOption) (*v11.Photo, error) {
	var out v11.Photo
	pattern := "/blog/v1/photos"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPhotoServiceCreatePhoto))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PhotoServiceHTTPClientImpl) DeletePhoto(ctx context.Context, in *v11.DeletePhotoRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/blog/v1/photos/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPhotoServiceDeletePhoto))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PhotoServiceHTTPClientImpl) GetPhoto(ctx context.Context, in *v11.GetPhotoRequest, opts ...http.CallOption) (*v11.Photo, error) {
	var out v11.Photo
	pattern := "/blog/v1/photos/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPhotoServiceGetPhoto))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PhotoServiceHTTPClientImpl) ListPhoto(ctx context.Context, in *v1.PagingRequest, opts ...http.CallOption) (*v11.ListPhotoResponse, error) {
	var out v11.ListPhotoResponse
	pattern := "/blog/v1/photos"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPhotoServiceListPhoto))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PhotoServiceHTTPClientImpl) UpdatePhoto(ctx context.Context, in *v11.UpdatePhotoRequest, opts ...http.CallOption) (*v11.Photo, error) {
	var out v11.Photo
	pattern := "/blog/v1/photos/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPhotoServiceUpdatePhoto))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Photo, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
