package main

import (
	"io"             // For reading HTTP response body
	"net/http"              // For HTTP server and client
	"net/http/httptest"     // For creating test servers and recording HTTP requests
	"testing"               // For writing test cases
	"github.com/go-chi/chi/v5" // For creating HTTP routes and handlers
)

func runTestServer() *httptest.Server {
    return httptest.NewServer(setupServer())
}

func TestIntegrationGetBooksHandler(t *testing.T) {
	// Create a test server
	fakeStorage := fakeStorage{}
	bookHandler := BookHandler{storage: fakeStorage}

	r := chi.NewRouter()
	r.Get("/books", bookHandler.ListBooks)

	ts := httptest.NewServer(r)
	defer ts.Close()

	// Make a GET request to the test server
	res, err := http.Get(ts.URL + "/books")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer res.Body.Close()

	// Verify the response
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %v", res.StatusCode)
	}

	body, _ := io.ReadAll(res.Body)
	if len(body) == 0 {
		t.Errorf("Expected a non-empty response body")
	}
}
