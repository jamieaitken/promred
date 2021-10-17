package testing

import (
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

func GetCounterVecValue(vec prometheus.CounterVec, lvs ...string) (int, error) {
	counter, err := vec.GetMetricWithLabelValues(lvs...)
	if err != nil {
		return 0, err
	}

	return getCounterValue(counter)
}

func getCounterValue(counter prometheus.Counter) (int, error) {
	dtoMetric := dto.Metric{}
	err := counter.Write(&dtoMetric)
	if err != nil {
		return 0, err
	}

	return int(*dtoMetric.Counter.Value), nil
}
