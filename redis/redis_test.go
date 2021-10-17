package redis

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/go-cmp/cmp"
	testtool "github.com/jamieaitken/promred/testing"
)

func TestRedis_Get(t *testing.T) {
	failCmd := redis.NewStringCmd(context.Background())
	failCmd.SetErr(errors.New("fail"))

	tests := []struct {
		name                   string
		givenRedis             redisProvider
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name: "given success, expect operation count to be 1 and error count to be 0",
			givenRedis: mockRedis{
				GivenGetCmd: redis.NewStringCmd(context.Background()),
			},
			expectedOperationCount: 1,
			expectedErrorCount:     0,
		},
		{
			name: "given failure, expect operation count to be 2 and error count to be 1",
			givenRedis: mockRedis{
				GivenGetCmd: failCmd,
			},
			expectedOperationCount: 2,
			expectedErrorCount:     1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := New(test.givenRedis)

			_ = r.Get(context.Background(), "", "test")

			actualOperationCount, err := testtool.GetCounterVecValue(*r.operationCount, "test", "Get")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*r.errorCount, "test", "Get")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

func TestRedis_HGet(t *testing.T) {
	failCmd := redis.NewStringCmd(context.Background())
	failCmd.SetErr(errors.New("fail"))

	tests := []struct {
		name                   string
		givenRedis             redisProvider
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name: "given success, expect operation count to be 1 and error count to be 0",
			givenRedis: mockRedis{
				GivenHGetCmd: redis.NewStringCmd(context.Background()),
			},
			expectedOperationCount: 1,
			expectedErrorCount:     0,
		},
		{
			name: "given failure, expect operation count to be 2 and error count to be 1",
			givenRedis: mockRedis{
				GivenHGetCmd: failCmd,
			},
			expectedOperationCount: 2,
			expectedErrorCount:     1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := New(test.givenRedis)

			_ = r.HGet(context.Background(), "", "", "test")

			actualOperationCount, err := testtool.GetCounterVecValue(*r.operationCount, "test", "HGet")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*r.errorCount, "test", "HGet")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

func TestRedis_MGet(t *testing.T) {
	failCmd := redis.NewSliceCmd(context.Background())
	failCmd.SetErr(errors.New("fail"))

	tests := []struct {
		name                   string
		givenRedis             redisProvider
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name: "given success, expect operation count to be 1 and error count to be 0",
			givenRedis: mockRedis{
				GivenMGetCmd: redis.NewSliceCmd(context.Background()),
			},
			expectedOperationCount: 1,
			expectedErrorCount:     0,
		},
		{
			name: "given failure, expect operation count to be 2 and error count to be 1",
			givenRedis: mockRedis{
				GivenMGetCmd: failCmd,
			},
			expectedOperationCount: 2,
			expectedErrorCount:     1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := New(test.givenRedis)

			_ = r.MGet(context.Background(), nil, "test")

			actualOperationCount, err := testtool.GetCounterVecValue(*r.operationCount, "test", "MGet")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*r.errorCount, "test", "MGet")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

func TestRedis_MSet(t *testing.T) {
	failCmd := redis.NewStatusCmd(context.Background())
	failCmd.SetErr(errors.New("fail"))

	tests := []struct {
		name                   string
		givenRedis             redisProvider
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name: "given success, expect operation count to be 1 and error count to be 0",
			givenRedis: mockRedis{
				GivenMSetCmd: redis.NewStatusCmd(context.Background()),
			},
			expectedOperationCount: 1,
			expectedErrorCount:     0,
		},
		{
			name: "given failure, expect operation count to be 2 and error count to be 1",
			givenRedis: mockRedis{
				GivenMSetCmd: failCmd,
			},
			expectedOperationCount: 2,
			expectedErrorCount:     1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := New(test.givenRedis)

			_ = r.MSet(context.Background(), nil, "test")

			actualOperationCount, err := testtool.GetCounterVecValue(*r.operationCount, "test", "MSet")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*r.errorCount, "test", "MSet")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

func TestRedis_Set(t *testing.T) {
	failCmd := redis.NewStatusCmd(context.Background())
	failCmd.SetErr(errors.New("fail"))

	tests := []struct {
		name                   string
		givenRedis             redisProvider
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name: "given success, expect operation count to be 1 and error count to be 0",
			givenRedis: mockRedis{
				GivenSetCmd: redis.NewStatusCmd(context.Background()),
			},
			expectedOperationCount: 1,
			expectedErrorCount:     0,
		},
		{
			name: "given failure, expect operation count to be 2 and error count to be 1",
			givenRedis: mockRedis{
				GivenSetCmd: failCmd,
			},
			expectedOperationCount: 2,
			expectedErrorCount:     1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := New(test.givenRedis)

			_ = r.Set(context.Background(), "", "", time.Hour*1, "test")

			actualOperationCount, err := testtool.GetCounterVecValue(*r.operationCount, "test", "Set")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*r.errorCount, "test", "Set")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

func TestRedis_SetEX(t *testing.T) {
	failCmd := redis.NewStatusCmd(context.Background())
	failCmd.SetErr(errors.New("fail"))

	tests := []struct {
		name                   string
		givenRedis             redisProvider
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name: "given success, expect operation count to be 1 and error count to be 0",
			givenRedis: mockRedis{
				GivenSetEXCmd: redis.NewStatusCmd(context.Background()),
			},
			expectedOperationCount: 1,
			expectedErrorCount:     0,
		},
		{
			name: "given failure, expect operation count to be 2 and error count to be 1",
			givenRedis: mockRedis{
				GivenSetEXCmd: failCmd,
			},
			expectedOperationCount: 2,
			expectedErrorCount:     1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := New(test.givenRedis)

			_ = r.SetEX(context.Background(), "", "", time.Hour*1, "test")

			actualOperationCount, err := testtool.GetCounterVecValue(*r.operationCount, "test", "SetEX")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*r.errorCount, "test", "SetEX")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

func TestRedis_Ping(t *testing.T) {
	failCmd := redis.NewStatusCmd(context.Background())
	failCmd.SetErr(errors.New("fail"))

	tests := []struct {
		name                   string
		givenRedis             redisProvider
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name: "given success, expect operation count to be 1 and error count to be 0",
			givenRedis: mockRedis{
				GivenPingCmd: redis.NewStatusCmd(context.Background()),
			},
			expectedOperationCount: 1,
			expectedErrorCount:     0,
		},
		{
			name: "given failure, expect operation count to be 2 and error count to be 1",
			givenRedis: mockRedis{
				GivenPingCmd: failCmd,
			},
			expectedOperationCount: 2,
			expectedErrorCount:     1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := New(test.givenRedis)

			_ = r.Ping(context.Background(), "test")

			actualOperationCount, err := testtool.GetCounterVecValue(*r.operationCount, "test", "Ping")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*r.errorCount, "test", "Ping")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

type mockRedis struct {
	GivenGetCmd   *redis.StringCmd
	GivenHGetCmd  *redis.StringCmd
	GivenMGetCmd  *redis.SliceCmd
	GivenMSetCmd  *redis.StatusCmd
	GivenSetCmd   *redis.StatusCmd
	GivenSetEXCmd *redis.StatusCmd
	GivenPingCmd  *redis.StatusCmd
}

func (m mockRedis) Get(_ context.Context, _ string) *redis.StringCmd {
	return m.GivenGetCmd
}

func (m mockRedis) HGet(_ context.Context, _, _ string) *redis.StringCmd {
	return m.GivenHGetCmd
}

func (m mockRedis) MGet(_ context.Context, _ ...string) *redis.SliceCmd {
	return m.GivenMGetCmd
}

func (m mockRedis) MSet(_ context.Context, _ ...interface{}) *redis.StatusCmd {
	return m.GivenMSetCmd
}

func (m mockRedis) Set(_ context.Context, _ string, _ interface{}, _ time.Duration) *redis.StatusCmd {
	return m.GivenSetCmd
}

func (m mockRedis) SetEX(_ context.Context, _ string, _ interface{}, _ time.Duration) *redis.StatusCmd {
	return m.GivenSetEXCmd
}

func (m mockRedis) Ping(_ context.Context) *redis.StatusCmd {
	return m.GivenPingCmd
}
