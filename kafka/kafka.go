package kafka

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/segmentio/kafka-go"
)

var (
	operationCount *prometheus.CounterVec
	errorCount     *prometheus.CounterVec
	duration       *prometheus.HistogramVec
)

type readerProvider interface {
	ReadMessage(ctx context.Context) (kafka.Message, error)
	Close() error
}

type writerProvider interface {
	WriteMessages(ctx context.Context, msgs ...kafka.Message) error
	Close() error
}

type heartbeatProvider interface {
	Heartbeat(ctx context.Context, req *kafka.HeartbeatRequest) (*kafka.HeartbeatResponse, error)
}

func init() {
	operationCount = withRate()
	errorCount = withError()
	duration = withDuration()
}

type Reader struct {
	provider       readerProvider
	operationCount *prometheus.CounterVec
	errorCount     *prometheus.CounterVec
	duration       *prometheus.HistogramVec
}

func NewReader(reader readerProvider) Reader {
	return Reader{
		provider:       reader,
		operationCount: operationCount,
		errorCount:     errorCount,
		duration:       duration,
	}
}

func (r Reader) ReadMessage(ctx context.Context, invoker string) (kafka.Message, error) {
	timer := prometheus.NewTimer(r.duration.WithLabelValues(invoker, "ReadMessage"))
	defer timer.ObserveDuration()

	r.operationCount.WithLabelValues(invoker, "ReadMessage").Inc()

	msg, err := r.provider.ReadMessage(ctx)
	if err != nil {
		r.errorCount.WithLabelValues(invoker, "ReadMessage").Inc()
	}

	return msg, err
}

func (r Reader) Close(invoker string) error {
	timer := prometheus.NewTimer(r.duration.WithLabelValues(invoker, "ReaderClose"))
	defer timer.ObserveDuration()

	r.operationCount.WithLabelValues(invoker, "ReaderClose").Inc()

	err := r.provider.Close()
	if err != nil {
		r.errorCount.WithLabelValues(invoker, "ReaderClose").Inc()
	}

	return err
}

type Writer struct {
	provider       writerProvider
	operationCount *prometheus.CounterVec
	errorCount     *prometheus.CounterVec
	duration       *prometheus.HistogramVec
}

func NewWriter(writer writerProvider) Writer {
	return Writer{
		provider:       writer,
		operationCount: operationCount,
		errorCount:     errorCount,
		duration:       duration,
	}
}

func (w Writer) WriteMessages(ctx context.Context, msgs []kafka.Message, invoker string) error {
	timer := prometheus.NewTimer(w.duration.WithLabelValues(invoker, "WriteMessages"))
	defer timer.ObserveDuration()

	w.operationCount.WithLabelValues(invoker, "WriteMessages").Inc()

	err := w.provider.WriteMessages(ctx, msgs...)
	if err != nil {
		w.errorCount.WithLabelValues(invoker, "WriteMessages").Inc()
	}

	return err
}

func (w Writer) Close(invoker string) error {
	timer := prometheus.NewTimer(w.duration.WithLabelValues(invoker, "WriterClose"))
	defer timer.ObserveDuration()

	w.operationCount.WithLabelValues(invoker, "WriterClose").Inc()

	err := w.provider.Close()
	if err != nil {
		w.errorCount.WithLabelValues(invoker, "WriterClose").Inc()
	}

	return err
}

type Heartbeater struct {
	provider       heartbeatProvider
	operationCount *prometheus.CounterVec
	errorCount     *prometheus.CounterVec
	duration       *prometheus.HistogramVec
}

func NewHeartbeater(client heartbeatProvider) Heartbeater {
	return Heartbeater{
		provider:       client,
		operationCount: operationCount,
		errorCount:     errorCount,
		duration:       duration,
	}
}

func (h Heartbeater) Heartbeat(ctx context.Context, req *kafka.HeartbeatRequest, invoker string) (*kafka.HeartbeatResponse, error) {
	timer := prometheus.NewTimer(h.duration.WithLabelValues(invoker, "Heartbeat"))
	defer timer.ObserveDuration()

	h.operationCount.WithLabelValues(invoker, "Heartbeat").Inc()

	res, err := h.provider.Heartbeat(ctx, req)
	if err != nil {
		h.errorCount.WithLabelValues(invoker, "Heartbeat").Inc()
	}

	if res.Error != nil {
		h.errorCount.WithLabelValues(invoker, "Heartbeat").Inc()
	}

	return res, err
}
