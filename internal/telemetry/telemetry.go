package telemetry

import "time"

var _globalTel = New()

func T() Telemetry {
	return _globalTel
}

type Telemetry interface {
	IncrementSessionCount(gateway string)
	ObserveProcessDuration(gateway string, duration time.Duration)
}
