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

type MstCategoryControllerImpl struct {
	MstCategoryService service.MstCategoryService
}

func NewMstCategoryController(mstCategoryService service.MstCategoryService) MstCategoryController {
	return &MstCategoryControllerImpl{
		MstCategoryService: mstCategoryService,
	}
}

func (controller MstCategoryControllerImpl) FindAll(write http.ResponseWriter, request *http.Request, params httprouter.Params) {

	mstCategoryWebRequest := web.CarrierScoreRequest{}
	helper.ReadFromRequestBody(request, &mstCategoryWebRequest)

	//parsedBody, err := helper.ParsedRequestBody(request)
	//helper.PanicIfError(err)
	//stringMap, err := helper.MapStringToStrings(parsedBody)
	//helper.PanicIfError(err)
	//ByJobId := stringMap["by_job_id"]
	//ByNik := stringMap["by_nik"]
	//
	//LimitQuery := stringMap["limit"]
	//OffsetQuery := stringMap["offset"]
	//var nik int = helper.StringToInt(ByNik)
	//var jobId int = helper.StringToInt(ByJobId)
	//var limit int = helper.StringToInt(LimitQuery)
	//var offset int = helper.StringToInt(OffsetQuery)

	carrierScoreResponse := controller.MstCategoryService.FindAll(request.Context())
	helper.LogError(`URL : ` + request.URL.Path + `, CODE: ` + strconv.Itoa(http.StatusOK))
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   carrierScoreResponse,
	}

	helper.WriteToResponseBody(write, webResponse)
}

func (controller MstCategoryControllerImpl) Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	Id := params.ByName("id")
	input := domain.MstCategory{Id: Id}
	fmt.Println(input)
	controller.MstCategoryService.Delete(request.Context(), input)
	helper.LogError(`URL : ` + request.URL.Path + `, CODE: ` + strconv.Itoa(http.StatusOK))
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}

	helper.WriteToResponseBody(write, webResponse)
}

func (controller MstCategoryControllerImpl) Create(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	mstCategoryCreate := web.MstCategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &mstCategoryCreate)
	categoryResponse := controller.MstCategoryService.Create(request.Context(), mstCategoryCreate)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}
	helper.WriteToResponseBody(write, webResponse)
}
