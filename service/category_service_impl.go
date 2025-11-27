package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/rozanlaudzai/go-mysql-restful-api/model/domain"
	"github.com/rozanlaudzai/go-mysql-restful-api/model/web"
	"github.com/rozanlaudzai/go-mysql-restful-api/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) ([]web.CategoryResponse, error) {

	var categoryResponses []web.CategoryResponse

	tx, err := service.DB.Begin()
	if err != nil {
		return categoryResponses, err
	}

	// rollback if an error exists
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	categories, err := service.CategoryRepository.FindAll(ctx, tx)
	if err != nil {
		return categoryResponses, err
	}

	if err = tx.Commit(); err != nil {
		return categoryResponses, err
	}

	categoryResponses = make([]web.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		categoryResponses = append(categoryResponses, web.CategoryResponse{
			Id:   category.Id,
			Name: category.Name,
		})
	}
	return categoryResponses, nil
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) (web.CategoryResponse, error) {

	var response web.CategoryResponse

	// validate request
	err := service.Validate.Struct(request)
	if err != nil {
		return response, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return response, err
	}

	// rollback if an error exists
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	category := domain.Category{
		Name: request.Name,
	}
	category, err = service.CategoryRepository.Create(ctx, tx, category)
	if err != nil {
		return response, err
	}

	if err = tx.Commit(); err != nil {
		return response, err
	}

	response = web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
	return response, nil
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) (web.CategoryResponse, error) {

	var response web.CategoryResponse

	tx, err := service.DB.Begin()
	if err != nil {
		return response, err
	}

	// rollback if an error exists
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		return response, err
	}

	if err = tx.Commit(); err != nil {
		return response, err
	}

	response = web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
	return response, nil
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) (web.CategoryResponse, error) {

	var response web.CategoryResponse

	// validate request
	err := service.Validate.Struct(request)
	if err != nil {
		return response, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return response, err
	}

	// rollback if an error exists
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// check if the category with that id exists or not
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return response, err
	}

	category = domain.Category{
		Id:   request.Id,
		Name: request.Name,
	}

	category, err = service.CategoryRepository.Update(ctx, tx, category)
	if err != nil {
		return response, err
	}

	if err = tx.Commit(); err != nil {
		return response, err
	}

	response = web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
	return response, nil
}

func (service *CategoryServiceImpl) DeleteById(ctx context.Context, categoryId int) error {

	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}

	// rollback if an error exists
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	if err = service.CategoryRepository.DeleteById(ctx, tx, categoryId); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
