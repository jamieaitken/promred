package doer

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
	testtool "github.com/jamieaitken/promred/testing"
)

func TestDoer_Do_Success(t *testing.T) {
	tests := []struct {
		name                   string
		givenDoer              doerProvider
		givenRequest           *http.Request
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name: "given success, expect operation count to be 1 and error count to be 0",
			givenDoer: mockDoer{
				GivenResponse: &http.Response{
					Status:     http.StatusText(http.StatusOK),
					StatusCode: http.StatusOK,
				},
			},
			givenRequest: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					Scheme: "https://",
					Host:   "example.com",
					Path:   "/test",
				},
			},
			expectedErrorCount:     0,
			expectedOperationCount: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			doer := New(test.givenDoer)

			_, err := doer.Do(test.givenRequest)
			if err != nil {
				t.Fatalf("expected nil, got %v", err)
			}

			actualOperationCount, err := testtool.GetCounterVecValue(*doer.operationCount, test.givenRequest.URL.Path, test.givenRequest.Method)
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*doer.errorCount, test.givenRequest.URL.Path, test.givenRequest.Method)
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

func TestDoer_Do_Fail(t *testing.T) {
	tests := []struct {
		name                   string
		givenDoer              doerProvider
		givenRequest           *http.Request
		expectError            bool
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name: "given doer fail, expect operation count to be 2 and error count to be 1",
			givenDoer: mockDoer{
				GivenError: errors.New("fail"),
			},
			givenRequest: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					Scheme: "https://",
					Host:   "example.com",
					Path:   "/test",
				},
			},
			expectError:            true,
			expectedErrorCount:     1,
			expectedOperationCount: 2,
		},
		{
			name:      "given nil response, expect operation count to be 3 and error count to be 2",
			givenDoer: mockDoer{},
			givenRequest: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					Scheme: "https://",
					Host:   "example.com",
					Path:   "/test",
				},
			},
			expectError:            false,
			expectedErrorCount:     2,
			expectedOperationCount: 3,
		},
		{
			name: "given response of 404, expect operation count to be 4 and error count to be 3",
			givenDoer: mockDoer{
				GivenResponse: &http.Response{
					Status:     http.StatusText(http.StatusNotFound),
					StatusCode: http.StatusNotFound,
				},
			},
			givenRequest: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					Scheme: "https://",
					Host:   "example.com",
					Path:   "/test",
				},
			},
			expectError:            false,
			expectedErrorCount:     3,
			expectedOperationCount: 4,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			doer := New(test.givenDoer)

			_, err := doer.Do(test.givenRequest)
			if err == nil && test.expectError {
				t.Fatalf("got nil")
			}

			actualOperationCount, err := testtool.GetCounterVecValue(*doer.operationCount, test.givenRequest.URL.Path, test.givenRequest.Method)
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatal(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*doer.errorCount, test.givenRequest.URL.Path, test.givenRequest.Method)
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatal(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

type mockDoer struct {
	GivenResponse *http.Response
	GivenError    error
}

func (m mockDoer) Do(_ *http.Request) (*http.Response, error) {
	return m.GivenResponse, m.GivenError
}
