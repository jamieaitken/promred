package handler

import "github.com/prometheus/client_golang/prometheus"

var (
	requestLabels  = []string{"path", "http_method"}
	responseLabels = []string{"path", "http_method", "status_code"}
)

func withRate() *prometheus.CounterVec {
	r := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "handler_operation_total",
		Help: "The number of requests",
	}, responseLabels)

	prometheus.MustRegister(r)

	return r
}

func withError() *prometheus.CounterVec {
	r := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "handler_error_total",
		Help: "The number of those requests that have failed",
	}, responseLabels)

	prometheus.MustRegister(r)

	return r
}

func withDuration() *prometheus.HistogramVec {
	d := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "handler_duration_seconds",
		Help: "The amount of time those requests take",
	}, requestLabels)

	prometheus.MustRegister(d)

	return d
}
