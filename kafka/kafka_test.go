package kafka

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	testtool "github.com/jamieaitken/promred/testing"
	"github.com/segmentio/kafka-go"
)

func TestReader_ReadMessage(t *testing.T) {
	tests := []struct {
		name                   string
		givenReader            readerProvider
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name:                   "given successful reader response, expect operation count to be 1 and error count to be 0",
			givenReader:            mockReader{},
			expectedErrorCount:     0,
			expectedOperationCount: 1,
		},
		{
			name: "given failed reader response, expect operation count to be 2 and error count to be 1",
			givenReader: mockReader{
				GivenReadMessageError: errors.New("fail"),
			},
			expectedErrorCount:     1,
			expectedOperationCount: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := NewReader(test.givenReader)

			_, _ = r.ReadMessage(context.Background(), "test")

			actualOperationCount, err := testtool.GetCounterVecValue(*r.operationCount, "test", "ReadMessage")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*r.errorCount, "test", "ReadMessage")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

func TestReader_Close(t *testing.T) {
	tests := []struct {
		name                   string
		givenReader            readerProvider
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name:                   "given successful close, expect operation count to be 1 and error count to be 0",
			givenReader:            mockReader{},
			expectedErrorCount:     0,
			expectedOperationCount: 1,
		},
		{
			name: "given failed close, expect operation count to be 2 and error count to be 1",
			givenReader: mockReader{
				GivenCloseError: errors.New("fail"),
			},
			expectedErrorCount:     1,
			expectedOperationCount: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := NewReader(test.givenReader)

			_ = r.Close("test")

			actualOperationCount, err := testtool.GetCounterVecValue(*r.operationCount, "test", "ReaderClose")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*r.errorCount, "test", "ReaderClose")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

func TestWriter_WriteMessages(t *testing.T) {
	tests := []struct {
		name                   string
		givenWriter            writerProvider
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name:                   "given successful write, expect operation count to be 1 and error count to be 0",
			givenWriter:            mockWriter{},
			expectedErrorCount:     0,
			expectedOperationCount: 1,
		},
		{
			name: "given failed write, expect operation count to be 2 and error count to be 1",
			givenWriter: mockWriter{
				GivenWriteMessagesError: errors.New("fail"),
			},
			expectedErrorCount:     1,
			expectedOperationCount: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := NewWriter(test.givenWriter)

			_ = r.WriteMessages(context.Background(), nil, "test")

			actualOperationCount, err := testtool.GetCounterVecValue(*r.operationCount, "test", "WriteMessages")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*r.errorCount, "test", "WriteMessages")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

func TestWriter_Close(t *testing.T) {
	tests := []struct {
		name                   string
		givenWriter            writerProvider
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name:                   "given successful close, expect operation count to be 1 and error count to be 0",
			givenWriter:            mockWriter{},
			expectedErrorCount:     0,
			expectedOperationCount: 1,
		},
		{
			name: "given failed close, expect operation count to be 2 and error count to be 1",
			givenWriter: mockWriter{
				GivenCloseError: errors.New("fail"),
			},
			expectedErrorCount:     1,
			expectedOperationCount: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := NewWriter(test.givenWriter)

			_ = r.Close("test")

			actualOperationCount, err := testtool.GetCounterVecValue(*r.operationCount, "test", "WriterClose")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*r.errorCount, "test", "WriterClose")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

func TestHeartbeater_Heartbeat(t *testing.T) {
	tests := []struct {
		name                   string
		givenHeartbeater       heartbeatProvider
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name: "given successful heartbeat, expect operation count to be 1 and error count to be 0",
			givenHeartbeater: mockHeartbeater{
				GivenHeartbeatRes: &kafka.HeartbeatResponse{},
			},
			expectedErrorCount:     0,
			expectedOperationCount: 1,
		},
		{
			name: "given failed heartbeat, expect operation count to be 2 and error count to be 1",
			givenHeartbeater: mockHeartbeater{
				GivenHeartbeatRes:   &kafka.HeartbeatResponse{},
				GivenHeartbeatError: errors.New("fail"),
			},
			expectedErrorCount:     1,
			expectedOperationCount: 2,
		},
		{
			name: "given failed error in heartbeat response, expect operation count to be 3 and error count to be 2",
			givenHeartbeater: mockHeartbeater{
				GivenHeartbeatRes: &kafka.HeartbeatResponse{
					Error: errors.New("fail"),
				},
			},
			expectedErrorCount:     2,
			expectedOperationCount: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := NewHeartbeater(test.givenHeartbeater)

			_, _ = r.Heartbeat(context.Background(), &kafka.HeartbeatRequest{}, "test")

			actualOperationCount, err := testtool.GetCounterVecValue(*r.operationCount, "test", "Heartbeat")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*r.errorCount, "test", "Heartbeat")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

type mockReader struct {
	GivenReadMessageMsg   kafka.Message
	GivenReadMessageError error
	GivenCloseError       error
}

func (m mockReader) ReadMessage(_ context.Context) (kafka.Message, error) {
	return m.GivenReadMessageMsg, m.GivenReadMessageError
}

func (m mockReader) Close() error {
	return m.GivenCloseError
}

type mockWriter struct {
	GivenWriteMessagesError error
	GivenCloseError         error
}

func (m mockWriter) WriteMessages(_ context.Context, _ ...kafka.Message) error {
	return m.GivenWriteMessagesError
}

func (m mockWriter) Close() error {
	return m.GivenCloseError
}

type mockHeartbeater struct {
	GivenHeartbeatRes   *kafka.HeartbeatResponse
	GivenHeartbeatError error
}

func (m mockHeartbeater) Heartbeat(_ context.Context, _ *kafka.HeartbeatRequest) (*kafka.HeartbeatResponse, error) {
	return m.GivenHeartbeatRes, m.GivenHeartbeatError
}
