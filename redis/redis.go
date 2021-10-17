package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus"
)

type redisProvider interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	HGet(ctx context.Context, key, field string) *redis.StringCmd
	MGet(ctx context.Context, keys ...string) *redis.SliceCmd
	MSet(ctx context.Context, values ...interface{}) *redis.StatusCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Ping(ctx context.Context) *redis.StatusCmd
}

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

type Redis struct {
	provider       redisProvider
	operationCount *prometheus.CounterVec
	errorCount     *prometheus.CounterVec
	duration       *prometheus.HistogramVec
}

func New(client redisProvider) Redis {
	return Redis{
		provider:       client,
		operationCount: operationCount,
		errorCount:     errorCount,
		duration:       duration,
	}
}

func (r Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration, invoker string) *redis.StatusCmd {
	timer := prometheus.NewTimer(r.duration.WithLabelValues(invoker, "Set"))
	defer timer.ObserveDuration()

	r.operationCount.WithLabelValues(invoker, "Set").Inc()

	cmd := r.provider.Set(ctx, key, value, expiration)
	if cmd.Err() != nil {
		r.errorCount.WithLabelValues(invoker, "Set").Inc()
	}

	return cmd
}

func (r Redis) Get(ctx context.Context, key, invoker string) *redis.StringCmd {
	timer := prometheus.NewTimer(r.duration.WithLabelValues(invoker, "Get"))
	defer timer.ObserveDuration()

	r.operationCount.WithLabelValues(invoker, "Get").Inc()

	cmd := r.provider.Get(ctx, key)
	if cmd.Err() != nil {
		r.errorCount.WithLabelValues(invoker, "Get").Inc()
	}

	return cmd
}

func (r Redis) HGet(ctx context.Context, key, field, invoker string) *redis.StringCmd {
	timer := prometheus.NewTimer(r.duration.WithLabelValues(invoker, "HGet"))
	defer timer.ObserveDuration()

	r.operationCount.WithLabelValues(invoker, "HGet").Inc()

	cmd := r.provider.HGet(ctx, key, field)
	if cmd.Err() != nil {
		r.errorCount.WithLabelValues(invoker, "HGet").Inc()
	}

	return cmd
}

func (r Redis) MGet(ctx context.Context, keys []string, invoker string) *redis.SliceCmd {
	timer := prometheus.NewTimer(r.duration.WithLabelValues(invoker, "MGet"))
	defer timer.ObserveDuration()

	r.operationCount.WithLabelValues(invoker, "MGet").Inc()

	cmd := r.provider.MGet(ctx, keys...)
	if cmd.Err() != nil {
		r.errorCount.WithLabelValues(invoker, "MGet").Inc()
	}

	return cmd
}

func (r Redis) MSet(ctx context.Context, values []interface{}, invoker string) *redis.StatusCmd {
	timer := prometheus.NewTimer(r.duration.WithLabelValues(invoker, "MSet"))
	defer timer.ObserveDuration()

	r.operationCount.WithLabelValues(invoker, "MSet").Inc()

	cmd := r.provider.MSet(ctx, values...)
	if cmd.Err() != nil {
		r.errorCount.WithLabelValues(invoker, "MSet").Inc()
	}

	return cmd
}

func (r Redis) SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration, invoker string) *redis.StatusCmd {
	timer := prometheus.NewTimer(r.duration.WithLabelValues(invoker, "SetEX"))
	defer timer.ObserveDuration()

	r.operationCount.WithLabelValues(invoker, "SetEX").Inc()

	cmd := r.provider.SetEX(ctx, key, value, expiration)
	if cmd.Err() != nil {
		r.errorCount.WithLabelValues(invoker, "SetEX").Inc()
	}

	return cmd
}

func (r Redis) Ping(ctx context.Context, invoker string) *redis.StatusCmd {
	timer := prometheus.NewTimer(r.duration.WithLabelValues(invoker, "Ping"))
	defer timer.ObserveDuration()

	r.operationCount.WithLabelValues(invoker, "Ping").Inc()

	cmd := r.provider.Ping(ctx)
	if cmd.Err() != nil {
		r.errorCount.WithLabelValues(invoker, "Ping").Inc()
	}

	return cmd
}
