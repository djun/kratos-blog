// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             (unknown)
// source: admin/service/v1/i_link.proto

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

const OperationLinkServiceCreateLink = "/admin.service.v1.LinkService/CreateLink"
const OperationLinkServiceDeleteLink = "/admin.service.v1.LinkService/DeleteLink"
const OperationLinkServiceGetLink = "/admin.service.v1.LinkService/GetLink"
const OperationLinkServiceListLink = "/admin.service.v1.LinkService/ListLink"
const OperationLinkServiceUpdateLink = "/admin.service.v1.LinkService/UpdateLink"

type LinkServiceHTTPServer interface {
	// CreateLink 创建链接
	CreateLink(context.Context, *v11.CreateLinkRequest) (*v11.Link, error)
	// DeleteLink 删除链接
	DeleteLink(context.Context, *v11.DeleteLinkRequest) (*emptypb.Empty, error)
	// GetLink 获取链接数据
	GetLink(context.Context, *v11.GetLinkRequest) (*v11.Link, error)
	// ListLink 获取链接列表
	ListLink(context.Context, *v1.PagingRequest) (*v11.ListLinkResponse, error)
	// UpdateLink 更新链接
	UpdateLink(context.Context, *v11.UpdateLinkRequest) (*v11.Link, error)
}

func RegisterLinkServiceHTTPServer(s *http.Server, srv LinkServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/links", _LinkService_ListLink0_HTTP_Handler(srv))
	r.GET("/admin/v1/links/{id}", _LinkService_GetLink0_HTTP_Handler(srv))
	r.POST("/admin/v1/links", _LinkService_CreateLink0_HTTP_Handler(srv))
	r.PUT("/admin/v1/links/{id}", _LinkService_UpdateLink0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/links/{id}", _LinkService_DeleteLink0_HTTP_Handler(srv))
}

func _LinkService_ListLink0_HTTP_Handler(srv LinkServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLinkServiceListLink)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListLink(ctx, req.(*v1.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.ListLinkResponse)
		return ctx.Result(200, reply)
	}
}

func _LinkService_GetLink0_HTTP_Handler(srv LinkServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.GetLinkRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLinkServiceGetLink)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetLink(ctx, req.(*v11.GetLinkRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.Link)
		return ctx.Result(200, reply)
	}
}

func _LinkService_CreateLink0_HTTP_Handler(srv LinkServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.CreateLinkRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLinkServiceCreateLink)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateLink(ctx, req.(*v11.CreateLinkRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.Link)
		return ctx.Result(200, reply)
	}
}

func _LinkService_UpdateLink0_HTTP_Handler(srv LinkServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.UpdateLinkRequest
		if err := ctx.Bind(&in.Link); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLinkServiceUpdateLink)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateLink(ctx, req.(*v11.UpdateLinkRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.Link)
		return ctx.Result(200, reply)
	}
}

func _LinkService_DeleteLink0_HTTP_Handler(srv LinkServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.DeleteLinkRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLinkServiceDeleteLink)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteLink(ctx, req.(*v11.DeleteLinkRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type LinkServiceHTTPClient interface {
	CreateLink(ctx context.Context, req *v11.CreateLinkRequest, opts ...http.CallOption) (rsp *v11.Link, err error)
	DeleteLink(ctx context.Context, req *v11.DeleteLinkRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetLink(ctx context.Context, req *v11.GetLinkRequest, opts ...http.CallOption) (rsp *v11.Link, err error)
	ListLink(ctx context.Context, req *v1.PagingRequest, opts ...http.CallOption) (rsp *v11.ListLinkResponse, err error)
	UpdateLink(ctx context.Context, req *v11.UpdateLinkRequest, opts ...http.CallOption) (rsp *v11.Link, err error)
}

type LinkServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewLinkServiceHTTPClient(client *http.Client) LinkServiceHTTPClient {
	return &LinkServiceHTTPClientImpl{client}
}

func (c *LinkServiceHTTPClientImpl) CreateLink(ctx context.Context, in *v11.CreateLinkRequest, opts ...http.CallOption) (*v11.Link, error) {
	var out v11.Link
	pattern := "/admin/v1/links"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLinkServiceCreateLink))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LinkServiceHTTPClientImpl) DeleteLink(ctx context.Context, in *v11.DeleteLinkRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/links/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLinkServiceDeleteLink))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LinkServiceHTTPClientImpl) GetLink(ctx context.Context, in *v11.GetLinkRequest, opts ...http.CallOption) (*v11.Link, error) {
	var out v11.Link
	pattern := "/admin/v1/links/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLinkServiceGetLink))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LinkServiceHTTPClientImpl) ListLink(ctx context.Context, in *v1.PagingRequest, opts ...http.CallOption) (*v11.ListLinkResponse, error) {
	var out v11.ListLinkResponse
	pattern := "/admin/v1/links"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLinkServiceListLink))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LinkServiceHTTPClientImpl) UpdateLink(ctx context.Context, in *v11.UpdateLinkRequest, opts ...http.CallOption) (*v11.Link, error) {
	var out v11.Link
	pattern := "/admin/v1/links/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLinkServiceUpdateLink))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Link, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
