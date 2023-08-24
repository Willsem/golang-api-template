package handlers

import (
	"net/http"

	"github.com/Willsem/golang-api-template/internal/http/router"
	"github.com/labstack/echo/v4"
)

//go:generate mockery --name=ReadyStatus --dir . --output ./mocks
type ReadyStatus interface {
	IsReady() int
}

type ReadinessHandler struct {
	readyStatus ReadyStatus
}

func NewReadinessHandler(readyStatus ReadyStatus) *ReadinessHandler {
	return &ReadinessHandler{
		readyStatus: readyStatus,
	}
}

func (h *ReadinessHandler) Routes() []router.Route {
	return []router.Route{
		{
			Method:  http.MethodGet,
			Path:    "/readiness",
			Handler: h.getReadiness,
		},
	}
}

func (h *ReadinessHandler) getReadiness(c echo.Context) error {
	c.Response().WriteHeader(h.readyStatus.IsReady())
	return nil
}
