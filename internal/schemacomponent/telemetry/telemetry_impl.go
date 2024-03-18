package telemetry

//type telemetry struct {
//	eventsSkippedCounter *prometheus.CounterVec
//	eventsIgnoredCounter *prometheus.CounterVec
//	orphanRefundsCounter prometheus.Counter
//}
//
//var _ Telemetry = &telemetry{}
//
//func New() Telemetry {
//	return &telemetry{
//		eventsSkippedCounter: eventsSkipped,
//		eventsIgnoredCounter: eventsIgnored,
//		orphanRefundsCounter: orphanRefunds,
//	}
//}
//
//func (t *telemetry) SkipEvent(eventType EventType, reason string) {
//	t.eventsSkippedCounter.WithLabelValues(string(eventType), reason).Add(1)
//}
//
//func (t *telemetry) IgnoreEvent(eventType EventType, reason string) {
//	t.eventsIgnoredCounter.WithLabelValues(string(eventType), reason).Add(1)
//}
//
//func (t *telemetry) IncrementOrphanRefund() {
//	t.orphanRefundsCounter.Inc()
//}
