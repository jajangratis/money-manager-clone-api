package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"github.com/jajangratis/money-manager-clone-api/helper"
	"github.com/jajangratis/money-manager-clone-api/model/domain"
	"github.com/jajangratis/money-manager-clone-api/model/repository"
	"github.com/jajangratis/money-manager-clone-api/model/web"
)

type MstCategoryImpl struct {
	MstCategoryRepository repository.MstCategoryRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewMstCategoryService(mstCategoryRepository repository.MstCategoryRepository, DB *sql.DB, validate *validator.Validate) MstCategoryService {
	return &MstCategoryImpl{
		MstCategoryRepository: mstCategoryRepository,
		DB:                    DB,
		Validate:              validate,
	}
}

func (service *MstCategoryImpl) FindAll(ctx context.Context) []web.MstCategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	mstCategories := service.MstCategoryRepository.FindAll(ctx, tx)
	return helper.ToMstCategoryResponses(mstCategories)
}

func (service *MstCategoryImpl) Delete(ctx context.Context, mstCategory domain.MstCategory) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	service.MstCategoryRepository.Delete(ctx, tx, mstCategory)
}

func (service *MstCategoryImpl) Create(ctx context.Context, request web.MstCategoryCreateRequest) web.MstCategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	mstCategory := domain.MstCategory{
		CategoryId:   request.CategoryId,
		CategoryName: request.CategoryName,
	}
	service.MstCategoryRepository.Save(ctx, tx, mstCategory)
	return helper.ToMstCategoryResponse(mstCategory)
}
