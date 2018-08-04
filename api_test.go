package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func executeRequest(s *server, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.router.ServeHTTP(rr, req)

	return rr
}

func TestApiRunning(t *testing.T) {
	s := newServer()

	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(s, req)

	if response.Code != 200 {
		t.Errorf("Expected response code 200, got %d", response.Code)
	}
}
