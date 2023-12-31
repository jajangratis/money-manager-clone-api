// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/jajangratis/money-manager-clone-api/app"
	"github.com/jajangratis/money-manager-clone-api/controller"
	"github.com/jajangratis/money-manager-clone-api/middleware"
	"github.com/jajangratis/money-manager-clone-api/model/repository"
	"github.com/jajangratis/money-manager-clone-api/service"
	"net/http"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from injector.go:

func InitializedServer() *http.Server {
	mstCategoryRepositoryImpl := repository.NewMstCategoryRepositoryImpl()
	db := app.NewDb()
	validate := validator.New()
	mstCategoryService := service.NewMstCategoryService(mstCategoryRepositoryImpl, db, validate)
	mstCategoryController := controller.NewMstCategoryController(mstCategoryService)
	mstAccountRepositoryImpl := repository.NewMstAccountRepositoryImpl()
	mstAccountTypeService := service.NewMstAccountService(mstAccountRepositoryImpl, db, validate)
	mstAccountTypeController := controller.NewMstAccountTypeController(mstAccountTypeService)
	mstTransactionMethodRepositoryImpl := repository.NewMstTransactionMethodRepositoryImpl()
	mstTransactionMethodService := service.NewTransactionMethodService(mstTransactionMethodRepositoryImpl, db, validate)
	mstTransactionMethodController := controller.NewMstTransactionMethodController(mstTransactionMethodService)
	router := app.NewRouter(mstCategoryController, mstAccountTypeController, mstTransactionMethodController)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := NewServer(authMiddleware)
	return server
}

// injector.go:

var mstCategoryRepositorySet = wire.NewSet(repository.NewMstCategoryRepositoryImpl, wire.Bind(new(repository.MstCategoryRepository), new(*repository.MstCategoryRepositoryImpl)))

var mstAccountRepositorySet = wire.NewSet(repository.NewMstAccountRepositoryImpl, wire.Bind(new(repository.MstAccountTypeRepository), new(*repository.MstAccountRepositoryImpl)))

var mstTransactionMethodSet = wire.NewSet(repository.NewMstTransactionMethodRepositoryImpl, wire.Bind(new(repository.MstTransactionMethodRepository), new(*repository.MstTransactionMethodRepositoryImpl)))
