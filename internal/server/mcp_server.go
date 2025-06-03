package server

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Define a counter for the number of requests
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "mcp_requests_total",
			Help: "Total number of requests received",
		},
		[]string{"tool"},
	)

	// Define a gauge for the current number of active requests
	activeRequests = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "mcp_active_requests",
			Help: "Current number of active requests",
		},
		[]string{"tool"},
	)
)

// InitMetrics initializes the Prometheus metrics
func InitMetrics() {
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(activeRequests)

	http.Handle("/metrics", promhttp.Handler())
}

// IncrementRequestCounter increments the request counter for a specific tool
func IncrementRequestCounter(tool string) {
	requestCounter.WithLabelValues(tool).Inc()
}

// SetActiveRequests sets the number of active requests for a specific tool
func SetActiveRequests(tool string, count float64) {
	activeRequests.WithLabelValues(tool).Set(count)
}