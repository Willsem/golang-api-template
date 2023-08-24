package handlers

import (
	"net/http"

	"github.com/Willsem/golang-api-template/internal/http/router"
	"github.com/labstack/echo/v4"
)

type HealthProbe interface {
	GetStatus() int
}

type LivenessHandler struct {
	probe HealthProbe
}

func NewLivenessHandler(probe HealthProbe) *LivenessHandler {
	return &LivenessHandler{
		probe: probe,
	}
}

func (h *LivenessHandler) Routes() []router.Route {
	return []router.Route{
		{
			Method:  http.MethodGet,
			Path:    "/liveness",
			Handler: h.getLiveness,
		},
	}
}

func (h *LivenessHandler) getLiveness(c echo.Context) error {
	c.Response().WriteHeader(h.probe.GetStatus())
	return nil
}
