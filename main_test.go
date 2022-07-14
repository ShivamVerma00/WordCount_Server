package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWordCount(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(WordCount)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong Status Code: got %v want %v", status, http.StatusOK)
	}
}
