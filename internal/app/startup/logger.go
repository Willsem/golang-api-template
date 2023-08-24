package startup

import (
	"go.uber.org/zap/zapcore"

	"github.com/Willsem/golang-api-template/internal/logger"
	"github.com/Willsem/golang-api-template/internal/logger/zap"
)

type LogConfig struct {
	Level zapcore.Level `envconfig:"LEVEL" default:"debug"`
	Env   string        `envconfig:"ENV" required:"true"`
}

func NewLogger(name string, config LogConfig) logger.Logger {
	return zap.NewLogger(
		zap.Name(name),
		zap.LogLevel(config.Level),
		zap.Env(config.Env),
	)
}

func NewFallbackLogger(name string) logger.Logger {
	return zap.NewLogger(zap.Name(name))
}
