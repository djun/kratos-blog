// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.3
// - protoc             (unknown)
// source: admin/service/v1/i_tag.proto

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

const OperationTagServiceCreateTag = "/admin.service.v1.TagService/CreateTag"
const OperationTagServiceDeleteTag = "/admin.service.v1.TagService/DeleteTag"
const OperationTagServiceGetTag = "/admin.service.v1.TagService/GetTag"
const OperationTagServiceListTag = "/admin.service.v1.TagService/ListTag"
const OperationTagServiceUpdateTag = "/admin.service.v1.TagService/UpdateTag"

type TagServiceHTTPServer interface {
	// CreateTag 创建标签
	CreateTag(context.Context, *v11.CreateTagRequest) (*v11.Tag, error)
	// DeleteTag 删除标签
	DeleteTag(context.Context, *v11.DeleteTagRequest) (*emptypb.Empty, error)
	// GetTag 获取标签数据
	GetTag(context.Context, *v11.GetTagRequest) (*v11.Tag, error)
	// ListTag 获取标签列表
	ListTag(context.Context, *v1.PagingRequest) (*v11.ListTagResponse, error)
	// UpdateTag 更新标签
	UpdateTag(context.Context, *v11.UpdateTagRequest) (*v11.Tag, error)
}

func RegisterTagServiceHTTPServer(s *http.Server, srv TagServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/tags", _TagService_ListTag0_HTTP_Handler(srv))
	r.GET("/admin/v1/tags/{id}", _TagService_GetTag0_HTTP_Handler(srv))
	r.POST("/admin/v1/tags", _TagService_CreateTag0_HTTP_Handler(srv))
	r.PUT("/admin/v1/tags/{id}", _TagService_UpdateTag0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/tags/{id}", _TagService_DeleteTag0_HTTP_Handler(srv))
}

func _TagService_ListTag0_HTTP_Handler(srv TagServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagServiceListTag)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTag(ctx, req.(*v1.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.ListTagResponse)
		return ctx.Result(200, reply)
	}
}

func _TagService_GetTag0_HTTP_Handler(srv TagServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.GetTagRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagServiceGetTag)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTag(ctx, req.(*v11.GetTagRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.Tag)
		return ctx.Result(200, reply)
	}
}

func _TagService_CreateTag0_HTTP_Handler(srv TagServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.CreateTagRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagServiceCreateTag)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateTag(ctx, req.(*v11.CreateTagRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.Tag)
		return ctx.Result(200, reply)
	}
}

func _TagService_UpdateTag0_HTTP_Handler(srv TagServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.UpdateTagRequest
		if err := ctx.Bind(&in.Tag); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagServiceUpdateTag)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTag(ctx, req.(*v11.UpdateTagRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.Tag)
		return ctx.Result(200, reply)
	}
}

func _TagService_DeleteTag0_HTTP_Handler(srv TagServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.DeleteTagRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagServiceDeleteTag)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteTag(ctx, req.(*v11.DeleteTagRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type TagServiceHTTPClient interface {
	CreateTag(ctx context.Context, req *v11.CreateTagRequest, opts ...http.CallOption) (rsp *v11.Tag, err error)
	DeleteTag(ctx context.Context, req *v11.DeleteTagRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetTag(ctx context.Context, req *v11.GetTagRequest, opts ...http.CallOption) (rsp *v11.Tag, err error)
	ListTag(ctx context.Context, req *v1.PagingRequest, opts ...http.CallOption) (rsp *v11.ListTagResponse, err error)
	UpdateTag(ctx context.Context, req *v11.UpdateTagRequest, opts ...http.CallOption) (rsp *v11.Tag, err error)
}

type TagServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewTagServiceHTTPClient(client *http.Client) TagServiceHTTPClient {
	return &TagServiceHTTPClientImpl{client}
}

func (c *TagServiceHTTPClientImpl) CreateTag(ctx context.Context, in *v11.CreateTagRequest, opts ...http.CallOption) (*v11.Tag, error) {
	var out v11.Tag
	pattern := "/admin/v1/tags"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTagServiceCreateTag))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TagServiceHTTPClientImpl) DeleteTag(ctx context.Context, in *v11.DeleteTagRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/tags/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTagServiceDeleteTag))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TagServiceHTTPClientImpl) GetTag(ctx context.Context, in *v11.GetTagRequest, opts ...http.CallOption) (*v11.Tag, error) {
	var out v11.Tag
	pattern := "/admin/v1/tags/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTagServiceGetTag))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TagServiceHTTPClientImpl) ListTag(ctx context.Context, in *v1.PagingRequest, opts ...http.CallOption) (*v11.ListTagResponse, error) {
	var out v11.ListTagResponse
	pattern := "/admin/v1/tags"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTagServiceListTag))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TagServiceHTTPClientImpl) UpdateTag(ctx context.Context, in *v11.UpdateTagRequest, opts ...http.CallOption) (*v11.Tag, error) {
	var out v11.Tag
	pattern := "/admin/v1/tags/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTagServiceUpdateTag))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Tag, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
