package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthCheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	reqHandler := http.HandlerFunc(healthCheckHandler)

	reqHandler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Expected status code not returned : got %v, expected %v", status, http.StatusOK)
	}

	expectedStatus := `{"alive": true}`
	if recorder.Body.String() != expectedStatus {
		t.Errorf("Expected response body not found : got %v, expected %v", recorder.Body.String(), expectedStatus)
	}
}
