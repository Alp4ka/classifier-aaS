package telemetry

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var sessionsTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Name:      "sessions_total",
		Help:      "Count of sessions",
	},
	[]string{
		"gateway",
	},
)

var _buckets = []float64{0.1, 1, 5, 10, 15, 30, 60, 120}

var processOverallDurationSeconds = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace: Namespace,
		Name:      "process_overall_duration_seconds",
		Help:      "Process overall duration in seconds",
		Buckets:   _buckets,
	},
	[]string{
		"gateway",
	},
)
