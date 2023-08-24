package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/Willsem/golang-api-template/internal/http/router"
)

type MetricsHandler struct {
	promHandler http.Handler
}

func NewMetricsHandler() *MetricsHandler {
	return &MetricsHandler{
		promHandler: promhttp.Handler(),
	}
}

func (h *MetricsHandler) Routes() []router.Route {
	return []router.Route{
		{
			Method:  http.MethodGet,
			Path:    "/metrics",
			Handler: h.getMetrics,
		},
	}
}

func (h *MetricsHandler) getMetrics(c echo.Context) error {
	h.promHandler.ServeHTTP(c.Response(), c.Request())
	return nil
}
