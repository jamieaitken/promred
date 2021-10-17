package handler

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	operationCount *prometheus.CounterVec
	errorCount     *prometheus.CounterVec
	duration       *prometheus.HistogramVec
)

func init() {
	operationCount = withRate()
	errorCount = withError()
	duration = withDuration()
}

type Handler struct {
	operationCount *prometheus.CounterVec
	errorCount     *prometheus.CounterVec
	duration       *prometheus.HistogramVec
}

func New() Handler {
	return Handler{
		operationCount: operationCount,
		errorCount:     errorCount,
		duration:       duration,
	}
}

func (h Handler) HandleFor(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		timer := prometheus.NewTimer(h.duration.WithLabelValues(r.URL.Path, r.Method))
		defer timer.ObserveDuration()

		rw := newResponseWriter(w)

		next(rw, r)

		statusCodeString := fmt.Sprint(rw.StatusCode)

		h.operationCount.WithLabelValues(r.URL.Path, r.Method, statusCodeString).Inc()

		if rw.StatusCode >= 400 {
			h.errorCount.WithLabelValues(r.URL.Path, r.Method, statusCodeString).Inc()
		}
	}
}
