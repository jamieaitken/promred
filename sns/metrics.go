package sns

import "github.com/prometheus/client_golang/prometheus"

var (
	labels = []string{"invoker", "operation"}
)

func withRate() *prometheus.CounterVec {
	r := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "sns_operation_total",
		Help: "The number of requests",
	}, labels)

	prometheus.MustRegister(r)

	return r
}

func withError() *prometheus.CounterVec {
	r := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "sns_error_total",
		Help: "The number of those requests that have failed",
	}, labels)

	prometheus.MustRegister(r)

	return r
}

func withDuration() *prometheus.HistogramVec {
	d := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "sns_duration_seconds",
		Help: "The amount of time those requests take",
	}, labels)

	prometheus.MustRegister(d)

	return d
}
