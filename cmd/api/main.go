package main

import (
	"context"

	"github.com/Willsem/golang-api-template/internal/app"
	"github.com/Willsem/golang-api-template/internal/app/build"
	"github.com/Willsem/golang-api-template/internal/app/config"
	"github.com/Willsem/golang-api-template/internal/app/startup"
	"github.com/Willsem/golang-api-template/internal/health"
	"github.com/Willsem/golang-api-template/internal/http/handlers"
	"github.com/Willsem/golang-api-template/internal/http/router"
	"github.com/Willsem/golang-api-template/internal/http/server"
	"github.com/Willsem/golang-api-template/internal/logger"
)

const appName = "golang-api-template"

// @title       Golang API Template
// @version     1.0
// @description golang api template

func main() {
	cfg, err := config.New()
	if err != nil {
		startup.NewFallbackLogger(appName).WithError(err).Fatal("failed to parse configuration")
	}

	logger := startup.NewLogger(appName, cfg.Log)

	if err := run(cfg, logger); err != nil {
		logger.WithError(err).Fatal("error during the running app")
	}
}

func run(cfg *config.Config, logger logger.Logger) error {
	logger.Infof(
		"%s has version %s built from %s on %s by %s",
		appName, build.Version, build.VersionCommit, build.BuildDate, build.GoVersion,
	)

	logger.With("config", cfg).Info("application is starting with config")

	probe := health.NewProbe()
	readyStatus := health.NewReadyStatus()

	router := router.NewHTTPRouter(
		logger,
		handlers.NewMetricsHandler(),
		handlers.NewSwaggerHandler(),
		handlers.NewLivenessHandler(probe),
		handlers.NewReadinessHandler(readyStatus),
	)

	return app.New(
		cfg.App, readyStatus, logger,
		server.NewHTTPServer(router, cfg.Server, logger),
	).Run(context.Background())
}
