package router

import (
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/Willsem/golang-api-template/internal/logger"
)

func LoggerMiddleware(log logger.Logger) echo.MiddlewareFunc {
	log = log.With(logger.ComponentKey, "http middleware")

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestLog := log.
				With("method", c.Request().Method).
				With("path", c.Path()).
				With("request ID", uuid.New())

			requestLog.Info("request started")

			startTime := time.Now()
			err := next(c)
			duration := time.Since(startTime)

			status := c.Response().Status
			if err != nil {
				requestLog.
					WithError(err).
					With("status", status).
					Warnf("request failed with duration %f seconds", duration.Seconds())
			} else {
				requestLog.
					With("status", status).
					Infof("request finished with duration %f seconds", duration.Seconds())
			}

			return err
		}
	}
}

func MetricsMiddleware() echo.MiddlewareFunc {
	countMetric := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_responces_total",
		Help: "Count of responses to the service",
	}, []string{"method", "path"})

	latencyMetric := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "latency_cache_service",
		Help:    "Duration of CacheService methods",
		Buckets: []float64{0.0001, 0.0005, 0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1.0, 2.0},
	}, []string{"method", "path"})

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			startTime := time.Now()
			err := next(c)
			duration := time.Since(startTime)

			countMetric.WithLabelValues(c.Request().Method, c.Path()).Inc()
			latencyMetric.WithLabelValues(c.Request().Method, c.Path()).Observe(duration.Seconds())

			return err
		}
	}
}
