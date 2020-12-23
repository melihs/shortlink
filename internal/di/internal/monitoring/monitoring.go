package monitoring

import (
	"net/http"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/heptiolabs/healthcheck"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func New(sentryHandler *sentryhttp.Handler) *http.ServeMux {
	// Create a new Prometheus registry
	registry := prometheus.NewRegistry()

	// Create a metrics-exposing Handler for the Prometheus registry
	// The healthcheck related metrics will be prefixed with the provided namespace
	health := healthcheck.NewMetricsHandler(registry, "common")

	// Our app is not happy if we've got more than 100 goroutines running.
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))

	// Create an "common" listener
	commonMux := http.NewServeMux()

	// Expose prometheus metrics on /metrics
	commonMux.Handle("/metrics", sentryHandler.Handle(promhttp.Handler()))

	// Expose a liveness check on /live
	commonMux.HandleFunc("/live", sentryHandler.HandleFunc(health.LiveEndpoint))

	// Expose a readiness check on /ready
	commonMux.HandleFunc("/ready", sentryHandler.HandleFunc(health.ReadyEndpoint))

	return commonMux
}
