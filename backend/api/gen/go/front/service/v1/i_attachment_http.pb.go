// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.3
// - protoc             (unknown)
// source: front/service/v1/i_attachment.proto

package servicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v11 "kratos-cms/api/gen/go/file/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAttachmentServiceCreateAttachment = "/front.service.v1.AttachmentService/CreateAttachment"
const OperationAttachmentServiceDeleteAttachment = "/front.service.v1.AttachmentService/DeleteAttachment"
const OperationAttachmentServiceGetAttachment = "/front.service.v1.AttachmentService/GetAttachment"
const OperationAttachmentServiceListAttachment = "/front.service.v1.AttachmentService/ListAttachment"
const OperationAttachmentServiceUpdateAttachment = "/front.service.v1.AttachmentService/UpdateAttachment"

type AttachmentServiceHTTPServer interface {
	// CreateAttachment 创建附件
	CreateAttachment(context.Context, *v11.CreateAttachmentRequest) (*v11.Attachment, error)
	// DeleteAttachment 删除附件
	DeleteAttachment(context.Context, *v11.DeleteAttachmentRequest) (*emptypb.Empty, error)
	// GetAttachment 获取附件数据
	GetAttachment(context.Context, *v11.GetAttachmentRequest) (*v11.Attachment, error)
	// ListAttachment 获取附件列表
	ListAttachment(context.Context, *v1.PagingRequest) (*v11.ListAttachmentResponse, error)
	// UpdateAttachment 更新附件
	UpdateAttachment(context.Context, *v11.UpdateAttachmentRequest) (*v11.Attachment, error)
}

func RegisterAttachmentServiceHTTPServer(s *http.Server, srv AttachmentServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/blog/v1/attachments", _AttachmentService_ListAttachment0_HTTP_Handler(srv))
	r.GET("/blog/v1/attachments/{id}", _AttachmentService_GetAttachment0_HTTP_Handler(srv))
	r.POST("/blog/v1/attachments", _AttachmentService_CreateAttachment0_HTTP_Handler(srv))
	r.PUT("/blog/v1/attachments/{id}", _AttachmentService_UpdateAttachment0_HTTP_Handler(srv))
	r.DELETE("/blog/v1/attachments/{id}", _AttachmentService_DeleteAttachment0_HTTP_Handler(srv))
}

func _AttachmentService_ListAttachment0_HTTP_Handler(srv AttachmentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAttachmentServiceListAttachment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListAttachment(ctx, req.(*v1.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.ListAttachmentResponse)
		return ctx.Result(200, reply)
	}
}

func _AttachmentService_GetAttachment0_HTTP_Handler(srv AttachmentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.GetAttachmentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAttachmentServiceGetAttachment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAttachment(ctx, req.(*v11.GetAttachmentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.Attachment)
		return ctx.Result(200, reply)
	}
}

func _AttachmentService_CreateAttachment0_HTTP_Handler(srv AttachmentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.CreateAttachmentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAttachmentServiceCreateAttachment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAttachment(ctx, req.(*v11.CreateAttachmentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.Attachment)
		return ctx.Result(200, reply)
	}
}

func _AttachmentService_UpdateAttachment0_HTTP_Handler(srv AttachmentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.UpdateAttachmentRequest
		if err := ctx.Bind(&in.Attachment); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAttachmentServiceUpdateAttachment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateAttachment(ctx, req.(*v11.UpdateAttachmentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.Attachment)
		return ctx.Result(200, reply)
	}
}

func _AttachmentService_DeleteAttachment0_HTTP_Handler(srv AttachmentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.DeleteAttachmentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAttachmentServiceDeleteAttachment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteAttachment(ctx, req.(*v11.DeleteAttachmentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type AttachmentServiceHTTPClient interface {
	CreateAttachment(ctx context.Context, req *v11.CreateAttachmentRequest, opts ...http.CallOption) (rsp *v11.Attachment, err error)
	DeleteAttachment(ctx context.Context, req *v11.DeleteAttachmentRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetAttachment(ctx context.Context, req *v11.GetAttachmentRequest, opts ...http.CallOption) (rsp *v11.Attachment, err error)
	ListAttachment(ctx context.Context, req *v1.PagingRequest, opts ...http.CallOption) (rsp *v11.ListAttachmentResponse, err error)
	UpdateAttachment(ctx context.Context, req *v11.UpdateAttachmentRequest, opts ...http.CallOption) (rsp *v11.Attachment, err error)
}

type AttachmentServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewAttachmentServiceHTTPClient(client *http.Client) AttachmentServiceHTTPClient {
	return &AttachmentServiceHTTPClientImpl{client}
}

func (c *AttachmentServiceHTTPClientImpl) CreateAttachment(ctx context.Context, in *v11.CreateAttachmentRequest, opts ...http.CallOption) (*v11.Attachment, error) {
	var out v11.Attachment
	pattern := "/blog/v1/attachments"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAttachmentServiceCreateAttachment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AttachmentServiceHTTPClientImpl) DeleteAttachment(ctx context.Context, in *v11.DeleteAttachmentRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/blog/v1/attachments/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAttachmentServiceDeleteAttachment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AttachmentServiceHTTPClientImpl) GetAttachment(ctx context.Context, in *v11.GetAttachmentRequest, opts ...http.CallOption) (*v11.Attachment, error) {
	var out v11.Attachment
	pattern := "/blog/v1/attachments/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAttachmentServiceGetAttachment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AttachmentServiceHTTPClientImpl) ListAttachment(ctx context.Context, in *v1.PagingRequest, opts ...http.CallOption) (*v11.ListAttachmentResponse, error) {
	var out v11.ListAttachmentResponse
	pattern := "/blog/v1/attachments"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAttachmentServiceListAttachment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AttachmentServiceHTTPClientImpl) UpdateAttachment(ctx context.Context, in *v11.UpdateAttachmentRequest, opts ...http.CallOption) (*v11.Attachment, error) {
	var out v11.Attachment
	pattern := "/blog/v1/attachments/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAttachmentServiceUpdateAttachment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Attachment, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
