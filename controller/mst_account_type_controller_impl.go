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

type MstAccountTypeControllerImpl struct {
	MstAccountTypeService service.MstAccountTypeService
}

func NewMstAccountTypeController(mstAccountService service.MstAccountTypeService) MstAccountTypeController {
	return &MstAccountTypeControllerImpl{
		MstAccountTypeService: mstAccountService,
	}
}

func (controller *MstAccountTypeControllerImpl) FindAll(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	carrierScoreResponse := controller.MstAccountTypeService.FindAll(request.Context())
	helper.LogError(`URL : ` + request.URL.Path + `, CODE: ` + strconv.Itoa(http.StatusOK))
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   carrierScoreResponse,
	}

	helper.WriteToResponseBody(write, webResponse)
}

func (controller *MstAccountTypeControllerImpl) Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	Id := params.ByName("id")
	input := domain.MstAccountType{Id: Id}
	fmt.Println(input)
	controller.MstAccountTypeService.Delete(request.Context(), input)
	helper.LogError(`URL : ` + request.URL.Path + `, CODE: ` + strconv.Itoa(http.StatusOK))
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}

	helper.WriteToResponseBody(write, webResponse)
}

func (controller *MstAccountTypeControllerImpl) Create(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	accountCreate := web.MstAccountCreateRequest{}
	helper.ReadFromRequestBody(request, &accountCreate)
	categoryResponse := controller.MstAccountTypeService.Create(request.Context(), accountCreate)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}
	helper.WriteToResponseBody(write, webResponse)
}
