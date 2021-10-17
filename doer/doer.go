package doer

import (
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

type doerProvider interface {
	Do(req *http.Request) (*http.Response, error)
}

type Doer struct {
	doerProvider   doerProvider
	operationCount *prometheus.CounterVec
	errorCount     *prometheus.CounterVec
	duration       *prometheus.HistogramVec
}

func New(doer doerProvider) Doer {
	return Doer{
		doerProvider:   doer,
		operationCount: operationCount,
		errorCount:     errorCount,
		duration:       duration,
	}
}

func (d Doer) Do(req *http.Request) (*http.Response, error) {
	timer := prometheus.NewTimer(d.duration.WithLabelValues(req.URL.Path, req.Method))
	defer timer.ObserveDuration()

	d.operationCount.WithLabelValues(req.URL.Path, req.Method).Inc()

	res, err := d.doerProvider.Do(req)
	if err != nil {
		d.errorCount.WithLabelValues(req.URL.Path, req.Method).Inc()

		return res, err
	}

	if res == nil {
		d.errorCount.WithLabelValues(req.URL.Path, req.Method).Inc()

		return res, err
	}

	if res.StatusCode >= 400 {
		d.errorCount.WithLabelValues(req.URL.Path, req.Method).Inc()
	}

	return res, err
}
