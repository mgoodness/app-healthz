package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	// HTTPRequestsTotal counts number of HTTP requests made. Grouped by code,
	// handler, and method
	HTTPRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests made.",
		},
		[]string{"code", "handler", "method"},
	)
)

var registerMetrics sync.Once

// Register registers all metrics to Prometheus
func Register() {
	// Register the metrics.
	registerMetrics.Do(func() {
		prometheus.MustRegister(HTTPRequestsTotal)
	})
}
