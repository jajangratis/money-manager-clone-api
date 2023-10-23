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

type MstAccountServiceImpl struct {
	MstAccountTypeRepository repository.MstAccountTypeRepository
	DB                       *sql.DB
	Validate                 *validator.Validate
}

func NewMstAccountService(mstAccountTypeRepository repository.MstAccountTypeRepository, DB *sql.DB, validate *validator.Validate) MstAccountTypeService {
	return &MstAccountServiceImpl{
		MstAccountTypeRepository: mstAccountTypeRepository,
		DB:                       DB,
		Validate:                 validate,
	}
}

func (service *MstAccountServiceImpl) FindAll(ctx context.Context) []web.MstAccountTypeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	mstAccounts := service.MstAccountTypeRepository.FindAll(ctx, tx)
	return helper.ToMstAccountTypeResponses(mstAccounts)
}

func (service *MstAccountServiceImpl) Delete(ctx context.Context, accountType domain.MstAccountType) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	service.MstAccountTypeRepository.Delete(ctx, tx, accountType)
}

func (service *MstAccountServiceImpl) Create(ctx context.Context, request web.MstAccountCreateRequest) web.MstAccountTypeResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	mstAccount := domain.MstAccountType{
		AccountName: request.AccountName,
		AccountId:   request.AccountId,
	}
	service.MstAccountTypeRepository.Save(ctx, tx, mstAccount)
	return helper.ToMstAccountTypeResponse(mstAccount)
}
