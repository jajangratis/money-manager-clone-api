package controller

import (
	"fmt"
	"github.com/jajangratis/money-manager-clone-api/helper"
	"github.com/jajangratis/money-manager-clone-api/model/domain"
	"github.com/jajangratis/money-manager-clone-api/model/web"
	"github.com/jajangratis/money-manager-clone-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type MstTransactionMethodControllerImpl struct {
	MstTransactionMethodService service.MstTransactionMethodService
}

func NewMstTransactionMethodController(mstTransactionService service.MstTransactionMethodService) MstTransactionMethodController {
	return &MstTransactionMethodControllerImpl{
		MstTransactionMethodService: mstTransactionService,
	}
}

func (controller MstTransactionMethodControllerImpl) FindAll(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	data := controller.MstTransactionMethodService.FindAll(request.Context())
	helper.LogError(`URL : ` + request.URL.Path + `, CODE: ` + strconv.Itoa(http.StatusOK))
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   data,
	}

	helper.WriteToResponseBody(write, webResponse)
}

func (controller MstTransactionMethodControllerImpl) Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	Id := params.ByName("id")
	input := domain.MstTransactionMethod{Id: Id}
	fmt.Println(input)
	controller.MstTransactionMethodService.Delete(request.Context(), input)
	helper.LogError(`URL : ` + request.URL.Path + `, CODE: ` + strconv.Itoa(http.StatusOK))
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}

	helper.WriteToResponseBody(write, webResponse)
}

func (controller MstTransactionMethodControllerImpl) Create(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	mstTransactionMethod := web.MstTransactionMethodRequest{}
	helper.ReadFromRequestBody(request, &mstTransactionMethod)
	response := controller.MstTransactionMethodService.Create(request.Context(), mstTransactionMethod)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   response,
	}
	helper.WriteToResponseBody(write, webResponse)
}
