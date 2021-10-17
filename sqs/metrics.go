package sqs

import "github.com/prometheus/client_golang/prometheus"

var (
	labels = []string{"invoker", "operation"}
)

func withRate() *prometheus.CounterVec {
	r := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "sqs_operation_count",
		Help: "The number of requests",
	}, labels)

	return r
}

func withError() *prometheus.CounterVec {
	r := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "sqs_error_count",
		Help: "The number of those requests that have failed",
	}, labels)

	return r
}

func withDuration() *prometheus.HistogramVec {
	d := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "sqs_duration_total",
		Help: "The amount of time those requests take",
	}, labels)

	prometheus.MustRegister(d)

	return d
}
