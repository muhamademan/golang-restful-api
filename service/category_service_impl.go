package service

import (
	"context"
	"database/sql"
	"emansulaeman/belajar-golang-restful-api/exception"
	"emansulaeman/belajar-golang-restful-api/helper"
	"emansulaeman/belajar-golang-restful-api/model/domain"
	"emansulaeman/belajar-golang-restful-api/model/web"
	"emansulaeman/belajar-golang-restful-api/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)

}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// cari dulu id nya menggunakan FindById
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	// helper.PanicIfError(err) // jika id nya tidak ada makan akan panic error
	if err != nil {
		panic(exception.NewNotFoundError(err.Error())) // memanggil fungsi NewNotFoundError jika data not found
	}

	// jika id ketemu / ada maka :
	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// untuk mencari id jika ada
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	// helper.PanicIfError(err) // jika id tidak ada makan akan panic error
	if err != nil {
		panic(exception.NewNotFoundError(err.Error())) // memanggil fungsi NewNotFoundError jika data not found
	}

	// jika Id ada makan akan mengembalikan value
	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// untuk mencari id jika ada
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	// helper.PanicIfError(err) // jika id tidak ada makan akan panic error
	if err != nil {
		panic(exception.NewNotFoundError(err.Error())) // memanggil fungsi NewNotFoundError jika data not found
	}

	// jika Id ada makan akan mengembalikan value sesuai Id
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
