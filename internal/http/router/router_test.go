package router_test

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"

	"github.com/Willsem/golang-api-template/internal/http/router"
	"github.com/Willsem/golang-api-template/internal/http/router/mocks"
	"github.com/Willsem/golang-api-template/internal/logger"
	"github.com/Willsem/golang-api-template/internal/testdata"
)

type RouterTestSuite struct {
	suite.Suite

	logger   logger.Logger
	handlers []*mocks.Handler
}

func (s *RouterTestSuite) SetupTest() {
	s.logger = testdata.NewLogger()

	s.handlers = []*mocks.Handler{
		mocks.NewHandler(s.T()),
		mocks.NewHandler(s.T()),
		mocks.NewHandler(s.T()),
	}
}

func (s *RouterTestSuite) TearDownTest() {
	s.handlers[0].AssertExpectations(s.T())
	s.handlers[1].AssertExpectations(s.T())
	s.handlers[2].AssertExpectations(s.T())
}

func (s *RouterTestSuite) equalRoutes(expected []*echo.Route, actual []*echo.Route) {
	s.Require().Equal(len(expected), len(actual))
	for i := range expected {
		wasChecked := false

		for j := range actual {
			if expected[i].Method == actual[j].Method {
				s.Require().Equal(*expected[i], *actual[j])
				wasChecked = true
				break
			}
		}

		s.Require().True(wasChecked)
	}
}

func (s *RouterTestSuite) TestRoutes() {
	defaultHandler := func(c echo.Context) error {
		return nil
	}

	routeName :=
		"github.com/Willsem/golang-api-template/internal/http/router_test.(*RouterTestSuite).TestRoutes.func1"
	expectedRoutes := []*echo.Route{
		{
			Method: http.MethodGet,
			Path:   "/handler1/get",
			Name:   routeName,
		},
		{
			Method: http.MethodPut,
			Path:   "/handler1/put",
			Name:   routeName,
		},
		{
			Method: http.MethodPost,
			Path:   "/handler2/post",
			Name:   routeName,
		},
		{
			Method: http.MethodDelete,
			Path:   "/handler2/delete",
			Name:   routeName,
		},
		{
			Method: http.MethodOptions,
			Path:   "/handler3/options",
			Name:   routeName,
		},
	}

	s.handlers[0].On("Routes").Return([]router.Route{
		{
			Method:  http.MethodGet,
			Path:    "/handler1/get",
			Handler: defaultHandler,
		},
		{
			Method:  "unkown",
			Path:    "/handler1/unknown",
			Handler: defaultHandler,
		},
		{
			Method:  http.MethodPut,
			Path:    "/handler1/put",
			Handler: defaultHandler,
		},
	}).Once()
	s.handlers[1].On("Routes").Return([]router.Route{
		{
			Method:  http.MethodPost,
			Path:    "/handler2/post",
			Handler: defaultHandler,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/handler2/delete",
			Handler: defaultHandler,
		},
	}).Once()
	s.handlers[2].On("Routes").Return([]router.Route{
		{
			Method:  http.MethodOptions,
			Path:    "/handler3/options",
			Handler: defaultHandler,
		},
		{
			Method:  "unkown",
			Path:    "/handler3/unknown",
			Handler: defaultHandler,
		},
	}).Once()

	router := router.NewHTTPRouter(s.logger,
		s.handlers[0],
		s.handlers[1],
		s.handlers[2],
	)

	actualRoutes := router.GetInstance().Routes()

	s.equalRoutes(expectedRoutes, actualRoutes)
}

func TestRouterTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(RouterTestSuite))
}
