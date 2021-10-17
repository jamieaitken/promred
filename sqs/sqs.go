package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
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

type sqsProvider interface {
	ReceiveMessage(ctx context.Context, params *sqs.ReceiveMessageInput, optFns ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error)
	SendMessage(ctx context.Context, params *sqs.SendMessageInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
}

type SQS struct {
	provider       sqsProvider
	operationCount *prometheus.CounterVec
	errorCount     *prometheus.CounterVec
	duration       *prometheus.HistogramVec
}

func New(provider sqsProvider) SQS {
	return SQS{
		provider:       provider,
		operationCount: operationCount,
		errorCount:     errorCount,
		duration:       duration,
	}
}

func (s SQS) ReceiveMessage(ctx context.Context, params *sqs.ReceiveMessageInput, optFns []func(*sqs.Options),
	invoker string) (*sqs.ReceiveMessageOutput, error) {
	timer := prometheus.NewTimer(s.duration.WithLabelValues(invoker, "ReceiveMessage"))
	defer timer.ObserveDuration()

	s.operationCount.WithLabelValues(invoker, "ReceiveMessage").Inc()

	out, err := s.provider.ReceiveMessage(ctx, params, optFns...)
	if err != nil {
		s.errorCount.WithLabelValues(invoker, "ReceiveMessage").Inc()
	}

	return out, err
}

func (s SQS) SendMessage(ctx context.Context, params *sqs.SendMessageInput, optFns []func(*sqs.Options),
	invoker string) (*sqs.SendMessageOutput, error) {
	timer := prometheus.NewTimer(s.duration.WithLabelValues(invoker, "SendMessage"))
	defer timer.ObserveDuration()

	s.operationCount.WithLabelValues(invoker, "SendMessage").Inc()

	out, err := s.provider.SendMessage(ctx, params, optFns...)
	if err != nil {
		s.errorCount.WithLabelValues(invoker, "SendMessage").Inc()
	}

	return out, err
}
