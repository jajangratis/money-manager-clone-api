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

type MstTransactionMethodServiceImpl struct {
	MstTransactionMethodRepository repository.MstTransactionMethodRepository
	DB                             *sql.DB
	Validate                       *validator.Validate
}

func NewTransactionMethodService(mstTransactionMethodRepository repository.MstTransactionMethodRepository, DB *sql.DB, validate *validator.Validate) MstTransactionMethodService {
	return &MstTransactionMethodServiceImpl{
		MstTransactionMethodRepository: mstTransactionMethodRepository,
		DB:                             DB,
		Validate:                       validate,
	}
}

func (service *MstTransactionMethodServiceImpl) FindAll(ctx context.Context) []web.MstTransactionMethodResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	transactionMethodList := service.MstTransactionMethodRepository.FindAll(ctx, tx)
	return helper.ToMstTransactionMethodResponses(transactionMethodList)
}

func (service *MstTransactionMethodServiceImpl) Delete(ctx context.Context, mstTransactionMethod domain.MstTransactionMethod) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	service.MstTransactionMethodRepository.Delete(ctx, tx, mstTransactionMethod)

}

func (service *MstTransactionMethodServiceImpl) Create(ctx context.Context, request web.MstTransactionMethodRequest) web.MstTransactionMethodResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	data := domain.MstTransactionMethod{
		MethodName: request.MethodName,
		MethodId:   request.MethodId,
	}

	transactionMethod := service.MstTransactionMethodRepository.Save(ctx, tx, data)
	return helper.ToMstTransactionMethodResponse(transactionMethod)
}
