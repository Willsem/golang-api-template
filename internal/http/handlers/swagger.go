package handlers

import (
	"net/http"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/Willsem/golang-api-template/internal/http/router"
	// Import swagger file.
	_ "github.com/Willsem/golang-api-template/api"
)

type SwaggerHandler struct{}

func NewSwaggerHandler() SwaggerHandler {
	return SwaggerHandler{}
}

func (SwaggerHandler) Routes() []router.Route {
	return []router.Route{
		{
			Method:  http.MethodGet,
			Path:    "/docs/*",
			Handler: echoSwagger.WrapHandler,
		},
	}
}
