package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sns"
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

type snsProvider interface {
	Publish(ctx context.Context, params *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
}

type SNS struct {
	snsProvider    snsProvider
	operationCount *prometheus.CounterVec
	errorCount     *prometheus.CounterVec
	duration       *prometheus.HistogramVec
}

func New(provider snsProvider) SNS {
	return SNS{
		snsProvider:    provider,
		operationCount: operationCount,
		errorCount:     errorCount,
		duration:       duration,
	}
}

func (s SNS) Publish(ctx context.Context, params *sns.PublishInput, optFns []func(*sns.Options),
	invoker string) (*sns.PublishOutput, error) {
	timer := prometheus.NewTimer(s.duration.WithLabelValues(invoker, "Publish"))
	defer timer.ObserveDuration()

	s.operationCount.WithLabelValues(invoker, "Publish").Inc()

	out, err := s.snsProvider.Publish(ctx, params, optFns...)
	if err != nil {
		s.errorCount.WithLabelValues(invoker, "Publish").Inc()
	}

	return out, err
}
