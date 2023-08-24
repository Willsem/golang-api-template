package config

import (
	"github.com/kelseyhightower/envconfig"

	"github.com/Willsem/golang-api-template/internal/app"
	"github.com/Willsem/golang-api-template/internal/app/startup"
	"github.com/Willsem/golang-api-template/internal/http/server"
)

type Config struct {
	Log    startup.LogConfig `envconfig:"LOG"`
	App    app.Config        `envconfig:"APP"`
	Server server.Config     `envconfig:"SERVER"`
}

func New() (*Config, error) {
	var config Config

	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
