package health_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/Willsem/golang-api-template/internal/health"
)

type ProbeTestSuite struct {
	suite.Suite

	probe *health.Probe
}

func (s *ProbeTestSuite) SetupTest() {
	s.probe = health.NewProbe()
}

func (s *ProbeTestSuite) TestGetStatus() {
	s.Require().Equal(http.StatusOK, s.probe.GetStatus())
}

func (s *ProbeTestSuite) TestSetStatus() {
	s.probe.SetStatus(health.ProbeStatusFailed)
	s.Require().Equal(http.StatusServiceUnavailable, s.probe.GetStatus())
	s.probe.SetStatus(health.ProbeStatusOK)
	s.Require().Equal(http.StatusOK, s.probe.GetStatus())
}

func TestProbeTestSute(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(ProbeTestSuite))
}
