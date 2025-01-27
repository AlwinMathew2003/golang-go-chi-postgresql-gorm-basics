package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock storage for testing
type fakeStorage struct{}

func (s fakeStorage) Get(_ string) *Book {
	return &Book{
		ID:               "1",
		Title:            "7 Habits of Highly Effective People",
		Author:           "Stephen Covey",
		PublishedDate:    "15/08/1989",
		OriginalLanguage: "English",
	}
}

func (s fakeStorage) Delete(_ string) *Book {
	return nil
}

func (s fakeStorage) List() []*Book {
	return []*Book{
		{
			ID:               "1",
			Title:            "7 Habits of Highly Effective People",
			Author:           "Stephen Covey",
			PublishedDate:    "15/08/1989",
			OriginalLanguage: "English",
		},
	}
}

func (s fakeStorage) Create(_ Book) {}

func (s fakeStorage) Update(id string, book Book) *Book {
	return &Book{
		ID:               id,
		Title:            book.Title,
		Author:           book.Author,
		PublishedDate:    book.PublishedDate,
		OriginalLanguage: book.OriginalLanguage,
	}
}

// TestGetBooksHandler tests the GetBooks handler
func TestGetBooksHandler(t *testing.T) {
	// Create a fake GET request
	req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
	w := httptest.NewRecorder()

	// Create a BookHandler with fakeStorage
	bookHandler := BookHandler{
		storage: fakeStorage{},
	}

	// Call the handler
	bookHandler.GetBooks(w, req)

	// Get the response
	res := w.Result()
	defer res.Body.Close()

	// Read and parse the response body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	book := Book{}
	json.Unmarshal(data, &book)

	// Assert the response
	if book.Title != "7 Habits of Highly Effective People" {
		t.Errorf("expected '7 Habits of Highly Effective People', got %v", book.Title)
	}
}
