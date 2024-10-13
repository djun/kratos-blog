// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             (unknown)
// source: front/service/v1/i_category.proto

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

const OperationCategoryServiceCreateCategory = "/front.service.v1.CategoryService/CreateCategory"
const OperationCategoryServiceDeleteCategory = "/front.service.v1.CategoryService/DeleteCategory"
const OperationCategoryServiceGetCategory = "/front.service.v1.CategoryService/GetCategory"
const OperationCategoryServiceListCategory = "/front.service.v1.CategoryService/ListCategory"
const OperationCategoryServiceUpdateCategory = "/front.service.v1.CategoryService/UpdateCategory"

type CategoryServiceHTTPServer interface {
	// CreateCategory 创建类别
	CreateCategory(context.Context, *v11.CreateCategoryRequest) (*v11.Category, error)
	// DeleteCategory 删除类别
	DeleteCategory(context.Context, *v11.DeleteCategoryRequest) (*emptypb.Empty, error)
	// GetCategory 获取类别数据
	GetCategory(context.Context, *v11.GetCategoryRequest) (*v11.Category, error)
	// ListCategory 获取类别列表
	ListCategory(context.Context, *v1.PagingRequest) (*v11.ListCategoryResponse, error)
	// UpdateCategory 更新类别
	UpdateCategory(context.Context, *v11.UpdateCategoryRequest) (*v11.Category, error)
}

func RegisterCategoryServiceHTTPServer(s *http.Server, srv CategoryServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/blog/v1/categories", _CategoryService_ListCategory0_HTTP_Handler(srv))
	r.GET("/blog/v1/categories/{id}", _CategoryService_GetCategory0_HTTP_Handler(srv))
	r.POST("/blog/v1/categories", _CategoryService_CreateCategory0_HTTP_Handler(srv))
	r.PUT("/blog/v1/categories/{id}", _CategoryService_UpdateCategory0_HTTP_Handler(srv))
	r.DELETE("/blog/v1/categories/{id}", _CategoryService_DeleteCategory0_HTTP_Handler(srv))
}

func _CategoryService_ListCategory0_HTTP_Handler(srv CategoryServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCategoryServiceListCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCategory(ctx, req.(*v1.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.ListCategoryResponse)
		return ctx.Result(200, reply)
	}
}

func _CategoryService_GetCategory0_HTTP_Handler(srv CategoryServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.GetCategoryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCategoryServiceGetCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCategory(ctx, req.(*v11.GetCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.Category)
		return ctx.Result(200, reply)
	}
}

func _CategoryService_CreateCategory0_HTTP_Handler(srv CategoryServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.CreateCategoryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCategoryServiceCreateCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCategory(ctx, req.(*v11.CreateCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.Category)
		return ctx.Result(200, reply)
	}
}

func _CategoryService_UpdateCategory0_HTTP_Handler(srv CategoryServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.UpdateCategoryRequest
		if err := ctx.Bind(&in.Category); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCategoryServiceUpdateCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCategory(ctx, req.(*v11.UpdateCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.Category)
		return ctx.Result(200, reply)
	}
}

func _CategoryService_DeleteCategory0_HTTP_Handler(srv CategoryServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.DeleteCategoryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCategoryServiceDeleteCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteCategory(ctx, req.(*v11.DeleteCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type CategoryServiceHTTPClient interface {
	CreateCategory(ctx context.Context, req *v11.CreateCategoryRequest, opts ...http.CallOption) (rsp *v11.Category, err error)
	DeleteCategory(ctx context.Context, req *v11.DeleteCategoryRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetCategory(ctx context.Context, req *v11.GetCategoryRequest, opts ...http.CallOption) (rsp *v11.Category, err error)
	ListCategory(ctx context.Context, req *v1.PagingRequest, opts ...http.CallOption) (rsp *v11.ListCategoryResponse, err error)
	UpdateCategory(ctx context.Context, req *v11.UpdateCategoryRequest, opts ...http.CallOption) (rsp *v11.Category, err error)
}

type CategoryServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewCategoryServiceHTTPClient(client *http.Client) CategoryServiceHTTPClient {
	return &CategoryServiceHTTPClientImpl{client}
}

func (c *CategoryServiceHTTPClientImpl) CreateCategory(ctx context.Context, in *v11.CreateCategoryRequest, opts ...http.CallOption) (*v11.Category, error) {
	var out v11.Category
	pattern := "/blog/v1/categories"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCategoryServiceCreateCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CategoryServiceHTTPClientImpl) DeleteCategory(ctx context.Context, in *v11.DeleteCategoryRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/blog/v1/categories/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCategoryServiceDeleteCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CategoryServiceHTTPClientImpl) GetCategory(ctx context.Context, in *v11.GetCategoryRequest, opts ...http.CallOption) (*v11.Category, error) {
	var out v11.Category
	pattern := "/blog/v1/categories/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCategoryServiceGetCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CategoryServiceHTTPClientImpl) ListCategory(ctx context.Context, in *v1.PagingRequest, opts ...http.CallOption) (*v11.ListCategoryResponse, error) {
	var out v11.ListCategoryResponse
	pattern := "/blog/v1/categories"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCategoryServiceListCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CategoryServiceHTTPClientImpl) UpdateCategory(ctx context.Context, in *v11.UpdateCategoryRequest, opts ...http.CallOption) (*v11.Category, error) {
	var out v11.Category
	pattern := "/blog/v1/categories/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCategoryServiceUpdateCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Category, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
