package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckStatusHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckStatusHandler)

	handler.ServeHTTP(rec, req)

	expectedStatusCode := http.StatusOK
	if statusCode := rec.Code; statusCode != expectedStatusCode {
		t.Errorf("handler returned wrong status code: got %v want %v",
			statusCode, expectedStatusCode)
	}

	expectedContentType := "application/json; charset=UTF-8"
	if contentType := rec.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong Content-Type: got %v want %v",
			contentType, expectedContentType)
	}

	var body HealthCheckStatusResponse
	json.Unmarshal(rec.Body.Bytes(), &body)
	expectedBody := HealthCheckStatusResponse{"OK"}
	if body != expectedBody {
		t.Errorf("handler returned unexpected body: got `%v` want `%v`",
			body, expectedBody)
	}
}
