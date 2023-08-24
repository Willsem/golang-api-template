package health_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/Willsem/golang-api-template/internal/health"
)

type ReadyStatusTestSuite struct {
	suite.Suite

	ready *health.ReadyStatus
}

func (s *ReadyStatusTestSuite) SetupTest() {
	s.ready = health.NewReadyStatus()
}

func (s *ReadyStatusTestSuite) TestIsReady() {
	s.Require().Equal(http.StatusServiceUnavailable, s.ready.IsReady())
}

func (s *ReadyStatusTestSuite) TestSetReady() {
	s.ready.SetReady()
	s.Require().Equal(http.StatusOK, s.ready.IsReady())
}

func TestReadyStatusTestSute(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(ReadyStatusTestSuite))
}
