//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/jajangratis/money-manager-clone-api/app"
	"github.com/jajangratis/money-manager-clone-api/controller"
	"github.com/jajangratis/money-manager-clone-api/middleware"
	"github.com/jajangratis/money-manager-clone-api/model/repository"
	"github.com/jajangratis/money-manager-clone-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var mstCategoryRepositorySet = wire.NewSet(
	repository.NewMstCategoryRepositoryImpl,
	wire.Bind(new(repository.MstCategoryRepository), new(*repository.MstCategoryRepositoryImpl)),
)

var mstAccountRepositorySet = wire.NewSet(
	repository.NewMstAccountRepositoryImpl,
	wire.Bind(new(repository.MstAccountTypeRepository), new(*repository.MstAccountRepositoryImpl)),
)

var mstTransactionMethodSet = wire.NewSet(
	repository.NewMstTransactionMethodRepositoryImpl,
	wire.Bind(new(repository.MstTransactionMethodRepository), new(*repository.MstTransactionMethodRepositoryImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDb,
		validator.New,
		app.NewRouter,
		middleware.NewAuthMiddleware,
		NewServer,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		//Package
		mstCategoryRepositorySet,
		service.NewMstCategoryService,
		controller.NewMstCategoryController,
		//MST ACCOUNT TYPE
		mstAccountRepositorySet,
		service.NewMstAccountService,
		controller.NewMstAccountTypeController,
		//MST TRANSACTION METHOD
		mstTransactionMethodSet,
		service.NewTransactionMethodService,
		controller.NewMstTransactionMethodController,
	)
	return nil
}
