package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type MstCategoryController interface {
	FindAll(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	Create(write http.ResponseWriter, request *http.Request, params httprouter.Params)
}
