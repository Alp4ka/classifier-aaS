package telemetry

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type telemetry struct {
	sessionsCounter     *prometheus.CounterVec
	processDurationHist *prometheus.HistogramVec
}

var _ Telemetry = &telemetry{}

func New() Telemetry {
	return &telemetry{
		sessionsCounter:     sessionsTotal,
		processDurationHist: processOverallDurationSeconds,
	}
}

func (t *telemetry) ObserveProcessDuration(gateway string, duration time.Duration) {
	t.processDurationHist.WithLabelValues(gateway).Observe(duration.Seconds())
}

func (t *telemetry) IncrementSessionCount(gateway string) {
	t.sessionsCounter.WithLabelValues(gateway).Inc()
}
