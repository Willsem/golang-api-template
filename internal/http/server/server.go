package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Willsem/golang-api-template/internal/logger"
)

//go:generate mockery --name=Router --dir . --output ./mocks
type Router interface {
	GetInstance() *echo.Echo
}

type Config struct {
	ListenPort int `envconfig:"LISTEN_PORT" default:"3000"`
}

type HTTPServer struct {
	router *echo.Echo
	config Config
	logger logger.Logger
}

func NewHTTPServer(httpRouter Router, config Config, log logger.Logger) *HTTPServer {
	return &HTTPServer{
		router: httpRouter.GetInstance(),
		config: config,
		logger: log.With(logger.ComponentKey, "http server"),
	}
}

func (s *HTTPServer) Start() error {
	go func() {
		s.logger.Infof("server is starting on port %d", s.config.ListenPort)
		if err := s.router.Start(fmt.Sprintf(":%d", s.config.ListenPort)); !errors.Is(err, http.ErrServerClosed) {
			s.logger.WithError(err).Fatalf("failed to start the server at port %d", s.config.ListenPort)
		}
		s.logger.Info("server has been stopped")
	}()
	return nil
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	return s.router.Shutdown(ctx)
}
