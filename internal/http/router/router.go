package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Willsem/golang-api-template/internal/logger"
)

type Route struct {
	Method      string
	Path        string
	Handler     echo.HandlerFunc
	Middlewares []echo.MiddlewareFunc
}

//go:generate mockery --name=Handler --dir . --output ./mocks
type Handler interface {
	Routes() []Route
}

type HTTPRouter struct {
	router *echo.Echo
	logger logger.Logger
}

func NewHTTPRouter(log logger.Logger, handlers ...Handler) *HTTPRouter {
	r := &HTTPRouter{
		router: echo.New(),
		logger: log.With(logger.ComponentKey, "http router"),
	}

	r.router.HideBanner = true
	r.router.HidePort = true

	r.router.Use(
		middleware.Recover(),
		MetricsMiddleware(),
		LoggerMiddleware(log),
	)

	for _, handler := range handlers {
		for _, route := range handler.Routes() {
			r.register(route)
		}
	}

	return r
}

func (r *HTTPRouter) register(route Route) {
	switch route.Method {
	case http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions:
		r.router.Add(route.Method, route.Path, route.Handler, route.Middlewares...)
	default:
		r.logger.
			With("route", route).
			Warn("unknown method for the route")
	}
}

func (r *HTTPRouter) GetInstance() *echo.Echo {
	return r.router
}
