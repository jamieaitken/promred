package sqs

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/google/go-cmp/cmp"
	testtool "github.com/jamieaitken/promred/testing"
)

func TestSQS_ReceiveMessageWithContext_Success(t *testing.T) {
	tests := []struct {
		name                   string
		givenSQS               sqsProvider
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name:                   "given success, expect operation count to be 1 and error count to be 0",
			givenSQS:               mockSQS{},
			expectedOperationCount: 1,
			expectedErrorCount:     0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := New(test.givenSQS)

			_, err := s.ReceiveMessage(context.Background(), nil, nil, "test")
			if err != nil {
				t.Fatalf("expected nil, got %v", err)
			}

			actualOperationCount, err := testtool.GetCounterVecValue(*s.operationCount, "test", "ReceiveMessage")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*s.errorCount, "test", "ReceiveMessage")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

func TestSQS_ReceiveMessageWithContext_Fail(t *testing.T) {
	tests := []struct {
		name                   string
		givenSQS               mockSQS
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name: "given failure, expect operation count to be 2 and error count to be 1",
			givenSQS: mockSQS{
				GivenReceiveError: errors.New("fail"),
			},
			expectedOperationCount: 2,
			expectedErrorCount:     1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := New(test.givenSQS)

			_, err := s.ReceiveMessage(context.Background(), nil, nil, "test")
			if err == nil {
				t.Fatalf("expected %v, got nil", test.givenSQS.GivenReceiveError)
			}

			actualOperationCount, err := testtool.GetCounterVecValue(*s.operationCount, "test", "ReceiveMessage")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*s.errorCount, "test", "ReceiveMessage")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

func TestSQS_SendMessageWithContext_Success(t *testing.T) {
	tests := []struct {
		name                   string
		givenSQS               sqsProvider
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name:                   "given success, expect operation count to be 1 and error count to be 0",
			givenSQS:               mockSQS{},
			expectedOperationCount: 1,
			expectedErrorCount:     0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := New(test.givenSQS)

			_, err := s.SendMessage(context.Background(), nil, nil, "test")
			if err != nil {
				t.Fatalf("expected nil, got %v", err)
			}

			actualOperationCount, err := testtool.GetCounterVecValue(*s.operationCount, "test", "SendMessage")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*s.errorCount, "test", "SendMessage")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

func TestSQS_SendMessageWithContext_Fail(t *testing.T) {
	tests := []struct {
		name                   string
		givenSQS               mockSQS
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name: "given failure, expect operation count to be 2 and error count to be 1",
			givenSQS: mockSQS{
				GivenSendError: errors.New("fail"),
			},
			expectedOperationCount: 2,
			expectedErrorCount:     1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := New(test.givenSQS)

			_, err := s.SendMessage(context.Background(), nil, nil, "test")
			if err == nil {
				t.Fatalf("expected %v, got nil", test.givenSQS.GivenSendError)
			}

			actualOperationCount, err := testtool.GetCounterVecValue(*s.operationCount, "test", "SendMessage")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*s.errorCount, "test", "SendMessage")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

type mockSQS struct {
	GivenReceiveMessage *sqs.ReceiveMessageOutput
	GivenReceiveError   error
	GivenSendMessage    *sqs.SendMessageOutput
	GivenSendError      error
}

func (m mockSQS) ReceiveMessage(_ context.Context, _ *sqs.ReceiveMessageInput, _ ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error) {
	return m.GivenReceiveMessage, m.GivenReceiveError
}

func (m mockSQS) SendMessage(_ context.Context, _ *sqs.SendMessageInput, _ ...func(*sqs.Options)) (*sqs.SendMessageOutput, error) {
	return m.GivenSendMessage, m.GivenSendError
}
