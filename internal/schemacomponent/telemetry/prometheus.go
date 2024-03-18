package telemetry

//
//const Subsystem = "schema"
//
//// eventsSkipped counts skipped events.
//var eventsSkipped = promauto.NewCounterVec(
//	prometheus.CounterOpts{
//		Namespace: globaltelemtry.Namespace,
//		Subsystem: Subsystem,
//		Name:      "events_skipped_total",
//		Help:      "Count of skipped events",
//	},
//	[]string{
//		"gateway", "reason",
//	},
//)
//
//// eventsIgnored counts ignored events.
//var eventsIgnored = promauto.NewCounterVec(
//	prometheus.CounterOpts{
//		Namespace: globaltelemtry.Namespace,
//		Subsystem: Subsystem,
//		Name:      "events_ignored_total",
//		Help:      "Count of ignored events",
//	},
//	[]string{
//		"event_type", "reason",
//	},
//)
//
//// orphanRefunds counts orphan refunds.
//var orphanRefunds = promauto.NewCounter(
//	prometheus.CounterOpts{
//		Namespace: globaltelemtry.Namespace,
//		Subsystem: Subsystem,
//		Name:      "orphan_refunds_total",
//		Help:      "Count of orphan refunds",
//	},
//)
