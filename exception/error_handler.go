package exception

import (
	"github.com/go-playground/validator"
	"github.com/jajangratis/money-manager-clone-api/helper"
	"github.com/jajangratis/money-manager-clone-api/model/web"
	"net/http"
	"strconv"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, error interface{}) {
	if notFoundError(writer, request, error) {
		return
	}
	if validationErrors(writer, request, error) {
		return
	}
	internalServerError(writer, request, error)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, error interface{}) bool {
	exception, ok := error.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		helper.LogError(`URL : ` + request.URL.Path + `, CODE: ` + strconv.Itoa(http.StatusBadRequest))
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, error interface{}) bool {
	exception, ok := error.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		helper.LogError(`URL : ` + request.URL.Path + `, CODE: ` + strconv.Itoa(http.StatusNotFound))
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Data Not Found",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, error interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	helper.LogError(`URL : ` + request.URL.Path + `, CODE: ` + strconv.Itoa(http.StatusInternalServerError))
	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   error,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
