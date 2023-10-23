package middleware

import (
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	middleware.Handler.ServeHTTP(writer, request)
	//fmt.Println(request.Header.Get("talent-key"))
	//if "5uydbVS97M-awPATBv0tR" == request.Header.Get("talent-key") {
	//	middleware.Handler.ServeHTTP(writer, request)
	//} else {
	//	writer.Header().Set("Content-Type", "application/json")
	//	writer.WriteHeader(http.StatusUnauthorized)
	//
	//	webResponse := web.WebResponse{
	//		Code:   http.StatusUnauthorized,
	//		Status: "Unauthorized",
	//	}
	//
	//	helper.WriteToResponseBody(writer, webResponse)
	//}
}
