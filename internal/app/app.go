package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/Willsem/golang-api-template/internal/logger"
)

//go:generate mockery --name=Component --dir . --output ./mocks
type Component interface {
	Start() error
	Stop(ctx context.Context) error
}

//go:generate mockery --name=ReadyStatus --dir . --output ./mocks
type ReadyStatus interface {
	SetReady()
}

type Config struct {
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"1m"`
}

type App struct {
	config      Config
	readyStatus ReadyStatus
	logger      logger.Logger
	components  []Component
}

func New(config Config, readyStatus ReadyStatus, logger logger.Logger, components ...Component) *App {
	return &App{
		config:      config,
		readyStatus: readyStatus,
		logger:      logger,
		components:  components,
	}
}

func (a *App) Run(ctx context.Context) error {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	g := &errgroup.Group{}

	for _, c := range a.components {
		component := c
		g.Go(component.Start)
	}

	if err := g.Wait(); err != nil {
		return err
	}

	a.logger.Info("application has been started")
	a.readyStatus.SetReady()

	select {
	case sig := <-sigs:
		a.logger.Infof("shutdown the app by the signal: %s", sig)
	case <-ctx.Done():
		a.logger.Info("shutdown the app by the end of the main context")
	}

	ctxStop, cancel := context.WithTimeout(ctx, a.config.ShutdownTimeout)
	defer cancel()

	for _, c := range a.components {
		component := c
		g.Go(func() error {
			return component.Stop(ctxStop)
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	a.logger.Info("application has been closed succesfully")

	return nil
}
