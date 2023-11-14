package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_ServeHTTP(t *testing.T) {
	s := &Server{} // Create an instance of your Server type

	// Create a test request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()

	// Call the ServeHTTP method with the created Request and ResponseRecorder.
	s.ServeHTTP(rr, req)

	// Check the response status code.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body.
	expected := `{"message": "hello world"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
