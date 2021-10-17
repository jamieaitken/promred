package sns

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/google/go-cmp/cmp"
	testtool "github.com/jamieaitken/promred/testing"
)

func TestSNS_Publish_Success(t *testing.T) {
	tests := []struct {
		name                   string
		givenSNS               snsProvider
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name:                   "given success, expect operation count to be 1 and error count to be 0",
			givenSNS:               mockSNS{},
			expectedOperationCount: 1,
			expectedErrorCount:     0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			snsNew := New(test.givenSNS)

			_, err := snsNew.Publish(context.Background(), nil, nil, "test")
			if err != nil {
				t.Fatalf("expected nil, got %v", err)
			}

			actualOperationCount, err := testtool.GetCounterVecValue(*snsNew.operationCount, "test", "Publish")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*snsNew.errorCount, "test", "Publish")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

func TestSNS_Publish_Fail(t *testing.T) {
	tests := []struct {
		name                   string
		givenSNS               mockSNS
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name: "given error, expect operation count to be 2 and error count to be 1",
			givenSNS: mockSNS{
				GivenError: errors.New("fail"),
			},
			expectedOperationCount: 2,
			expectedErrorCount:     1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			snsNew := New(test.givenSNS)

			_, err := snsNew.Publish(context.Background(), nil, nil, "test")
			if err == nil {
				t.Fatalf("expected %v, got nil", test.givenSNS.GivenError)
			}

			actualOperationCount, err := testtool.GetCounterVecValue(*snsNew.operationCount, "test", "Publish")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*snsNew.errorCount, "test", "Publish")
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

type mockSNS struct {
	GivenOutput *sns.PublishOutput
	GivenError  error
}

func (m mockSNS) Publish(_ context.Context, _ *sns.PublishInput, _ ...func(*sns.Options)) (*sns.PublishOutput, error) {
	return m.GivenOutput, m.GivenError
}
