package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-cms/app/core/service/internal/data"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	v1 "kratos-cms/api/gen/go/content/service/v1"
)

type CategoryService struct {
	v1.UnimplementedCategoryServiceServer

	uc  *data.CategoryRepo
	log *log.Helper
}

func NewCategoryService(logger log.Logger, uc *data.CategoryRepo) *CategoryService {
	l := log.NewHelper(log.With(logger, "module", "category/service/core-service"))
	return &CategoryService{
		log: l,
		uc:  uc,
	}
}

// ListCategory 获取类别列表
func (s *CategoryService) ListCategory(ctx context.Context, req *pagination.PagingRequest) (*v1.ListCategoryResponse, error) {
	return s.uc.List(ctx, req)
}

// GetCategory 获取类别数据
func (s *CategoryService) GetCategory(ctx context.Context, req *v1.GetCategoryRequest) (*v1.Category, error) {
	return s.uc.Get(ctx, req)
}

// CreateCategory 创建类别
func (s *CategoryService) CreateCategory(ctx context.Context, req *v1.CreateCategoryRequest) (*v1.Category, error) {
	return s.uc.Create(ctx, req)
}

// UpdateCategory 更新类别
func (s *CategoryService) UpdateCategory(ctx context.Context, req *v1.UpdateCategoryRequest) (*v1.Category, error) {
	return s.uc.Update(ctx, req)
}

// DeleteCategory 删除类别
func (s *CategoryService) DeleteCategory(ctx context.Context, req *v1.DeleteCategoryRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
