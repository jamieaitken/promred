package redis

import "github.com/prometheus/client_golang/prometheus"

var labels = []string{"invoker", "operation"}

func withRate() *prometheus.CounterVec {
	r := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "redis_operation_count",
		Help: "The number of operations",
	}, labels)

	return r
}

func withError() *prometheus.CounterVec {
	r := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "redis_error_count",
		Help: "The number of those operations that have failed",
	}, labels)

	return r
}

func withDuration() *prometheus.HistogramVec {
	d := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "redis_duration_total",
		Help: "The amount of time those operations take",
	}, labels)

	prometheus.MustRegister(d)

	return d
}
