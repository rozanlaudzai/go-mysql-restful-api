package service

import (
	"context"

	"github.com/rozanlaudzai/go-mysql-restful-api/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) (web.CategoryResponse, error)
	Update(ctx context.Context, request web.CategoryUpdateRequest) (web.CategoryResponse, error)
	DeleteById(ctx context.Context, categoryId int) error
	FindById(ctx context.Context, categoryId int) (web.CategoryResponse, error)
	FindAll(ctx context.Context) ([]web.CategoryResponse, error)
}
