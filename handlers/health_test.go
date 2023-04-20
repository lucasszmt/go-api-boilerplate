package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	tests := []struct {
		name         string
		expectedCode int
		expectedBody string
	}{
		{"valid request", 200, `{"status":"ok"}`},
	}
	for _, tt := range tests {
		req, err := http.NewRequest(http.MethodGet, "/health", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Health)
		handler.ServeHTTP(rr, req)

		if rr.Code != tt.expectedCode {
			t.Errorf("Health() = %v, want %v", rr.Code, tt.expectedCode)
		}

		if rr.Body.String() != tt.expectedBody {
			t.Log(rr.Body.String())
			t.Errorf("Health() = %v, want %v", rr.Body.String(), tt.expectedBody)
		}
	}
}
