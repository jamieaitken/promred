package doer

import "github.com/prometheus/client_golang/prometheus"

var (
	labels = []string{"path", "http_method"}
)

func withRate() *prometheus.CounterVec {
	r := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "doer_operation_total",
		Help: "The number of requests",
	}, labels)

	return r
}

func withError() *prometheus.CounterVec {
	r := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "doer_error_total",
		Help: "The number of those requests that have failed",
	}, labels)

	return r
}

func withDuration() *prometheus.HistogramVec {
	d := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "doer_duration_seconds",
		Help: "The amount of time those requests take",
	}, labels)

	prometheus.MustRegister(d)

	return d
}
