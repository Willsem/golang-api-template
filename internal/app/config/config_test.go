package config_test

import (
	"testing"
	"time"

	"github.com/Willsem/golang-api-template/internal/app"
	"github.com/Willsem/golang-api-template/internal/app/config"
	"github.com/Willsem/golang-api-template/internal/app/startup"
	"github.com/Willsem/golang-api-template/internal/http/server"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap/zapcore"
)

type ConfigTestSuite struct {
	suite.Suite
}

func (s *ConfigTestSuite) TestNewConfigSuccess() {
	var testConfigEnv = map[string]string{
		"LOG_LEVEL": "debug",
		"LOG_ENV":   "local",

		"APP_SHUTDOWN_TIMEOUT": "2m",

		"SERVER_LISTEN_PORT": "3000",
	}

	var testConfig = config.Config{
		Log: startup.LogConfig{
			Level: zapcore.DebugLevel,
			Env:   "local",
		},
		App: app.Config{
			ShutdownTimeout: 2 * time.Minute,
		},
		Server: server.Config{
			ListenPort: 3000,
		},
	}

	for key, value := range testConfigEnv {
		s.T().Setenv(key, value)
	}

	actual, err := config.New()

	s.Require().NoError(err)
	s.Require().Equal(testConfig, *actual)
}

func (s *ConfigTestSuite) TestNewConfigError() {
	_, err := config.New()
	s.Require().Error(err)
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
