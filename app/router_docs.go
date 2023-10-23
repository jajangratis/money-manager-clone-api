package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/zc2638/swag"
	"github.com/zc2638/swag/endpoint"
	"github.com/zc2638/swag/option"
	"net/http"
)

func NewDocs(router *httprouter.Router) *swag.API {
	api := swag.New(
		option.Title("Money Manager Clone Doc"),
		option.Security("petstore_auth", "read:pets"),
		option.SecurityScheme("petstore_auth",
			option.OAuth2Security("accessCode", "http://example.com/oauth/authorize", "http://example.com/oauth/token"),
			option.OAuth2Scope("write:pets", "modify pets in your account"),
			option.OAuth2Scope("read:pets", "read your pets"),
		),
	)
	api.AddEndpoint(
		endpoint.New(
			http.MethodGet, "/pet",
			endpoint.Handler(handle),
			endpoint.Summary("Add a new pet to the store"),
			endpoint.Description("Additional information on adding a pet to the store"),
			endpoint.Body(Pet{}, "Pet object that needs to be added to the store", true),
			//endpoint.Response(http.StatusOK, "Successfully added pet", endpoint.Schema(Pet{})),
			endpoint.Security("petstore_auth", "read:pets", "write:pets"),
		),
	)
	api.Walk(func(path string, e *swag.Endpoint) {
		h := e.Handler.(http.Handler)
		path = swag.ColonPath(path)
		router.Handler(e.Method, path, h)
	})

	return api
}
