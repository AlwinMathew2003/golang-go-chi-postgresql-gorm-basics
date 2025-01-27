package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

// Book represents the book model
type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

// In-memory storage for books
var books = []Book{
    {ID: "1", Title: "The Go Programming Language", Author: "Alan A. A. Donovan"},
    {ID: "2", Title: "Clean Code", Author: "Robert C. Martin"},
}

func main() {
    r := chi.NewRouter()

    // Add middlewares
    r.Use(middleware.Logger)

    // CRUD routes
    r.Get("/books", listBooks)         // List all books
    r.Get("/books/{id}", getBookByID)  // Get book by ID
    r.Post("/books", createBook)       // Create a new book
    r.Put("/books/{id}", updateBook)   // Update a book
    r.Delete("/books/{id}", deleteBook) // Delete a book

    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", r)
}

// Handler to list all books
func listBooks(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(books)
}

// Handler to get a book by ID
func getBookByID(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    for _, book := range books {
        if book.ID == id {
            json.NewEncoder(w).Encode(book)
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}

// Handler to create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
    var book Book
    json.NewDecoder(r.Body).Decode(&book)
    books = append(books, book)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(book)
}

// Handler to update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    var updatedBook Book
    json.NewDecoder(r.Body).Decode(&updatedBook)

    for i, book := range books {
        if book.ID == id {
            books[i] = updatedBook
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode(updatedBook)
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}

// Handler to delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    for i, book := range books {
        if book.ID == id {
            books = append(books[:i], books[i+1:]...)
            w.WriteHeader(http.StatusOK)
            w.Write([]byte("Book deleted"))
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}
