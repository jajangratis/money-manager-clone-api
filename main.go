package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jajangratis/money-manager-clone-api/helper"
	"github.com/jajangratis/money-manager-clone-api/middleware"
	"github.com/joho/godotenv"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {

	helper.LogInfo(`SERVER RUNNING ON 8000`)
	return &http.Server{
		Addr:    "localhost:8000",
		Handler: authMiddleware,
	}
}

func main() {
	//db := app.NewDb()
	//validate := validator.New()
	//carrierScoreRepository := repository.NewCarrierScoreRepositoryImpl()
	//carrierScoreService := service.NewCarrierScoreService(carrierScoreRepository, db, validate)
	//carrierScoreController := controller.NewCarrierScoreController(carrierScoreService)
	//
	//router := app.NewRouter(carrierScoreController)
	//server := http.Server{
	//	Addr:    "localhost:8000",
	//	Handler: router,
	//}
	//
	//err := server.ListenAndServe()
	//helper.PanicIfError(err)
	godotenv.Load()
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
