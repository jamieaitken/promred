package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	testtool "github.com/jamieaitken/promred/testing"
)

func TestHandler_HandleFor(t *testing.T) {
	tests := []struct {
		name                   string
		givenHandler           mockHandler
		expectedStatus         int
		expectedOperationCount int
		expectedErrorCount     int
	}{
		{
			name: "given valid request, expect operation count to be 1 and error count to be 0",
			givenHandler: mockHandler{
				GivenStatusCode: http.StatusOK,
				GivenPath:       "/v1/code",
				GivenMethod:     http.MethodGet,
			},
			expectedStatus:         http.StatusOK,
			expectedOperationCount: 1,
			expectedErrorCount:     0,
		},
		{
			name: "given valid response with 409, expect operation count to be 2 and error count to be 1",
			givenHandler: mockHandler{
				GivenStatusCode: http.StatusConflict,
				GivenPath:       "/v1/code",
				GivenMethod:     http.MethodGet,
			},
			expectedStatus:         http.StatusConflict,
			expectedOperationCount: 1,
			expectedErrorCount:     1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := New()

			router := new(http.ServeMux)

			router.HandleFunc("/v1/code", h.HandleFor(test.givenHandler.Get))

			rr := httptest.NewRecorder()
			req := httptest.NewRequest(test.givenHandler.GivenMethod, "/v1/code", nil)
			router.ServeHTTP(rr, req)

			if !cmp.Equal(rr.Code, test.expectedStatus) {
				t.Fatalf(cmp.Diff(rr.Code, test.expectedStatus))
			}

			statusCode := fmt.Sprint(rr.Code)

			actualOperationCount, err := testtool.GetCounterVecValue(*h.operationCount, test.givenHandler.GivenPath,
				test.givenHandler.GivenMethod, statusCode)
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualOperationCount, test.expectedOperationCount) {
				t.Fatalf(cmp.Diff(actualOperationCount, test.expectedOperationCount))
			}

			actualErrorCount, err := testtool.GetCounterVecValue(*h.errorCount, test.givenHandler.GivenPath,
				test.givenHandler.GivenMethod, statusCode)
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(actualErrorCount, test.expectedErrorCount) {
				t.Fatalf(cmp.Diff(actualErrorCount, test.expectedErrorCount))
			}
		})
	}
}

type mockHandler struct {
	GivenStatusCode int
	GivenPath       string
	GivenMethod     string
}

func (m mockHandler) Get(w http.ResponseWriter, r *http.Request) {
	r.Method = m.GivenMethod
	r.URL.Path = m.GivenPath

	w.WriteHeader(m.GivenStatusCode)
}
