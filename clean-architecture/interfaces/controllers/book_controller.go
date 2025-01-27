package controllers

import (
	"encoding/json"
	"clean-architecture/internal/usecases"
	"clean-architecture/internal/models"
	"net/http"
	"github.com/go-chi/chi/v5"
)

type BookController struct{
	BookService *usecases.BookService
}

//It handles the http request and receives the data and store it in books varibale(temporary)
//Then it passes the data to the usecases->BookServices from where it validates the data
//After validation it updates the globally stored data of the book by creating a new instance
//It then instructs the repository to save the data from the BookService
//Inside the repository the save method is defined in interface and this methods logic is done in the database

func(c *BookController)CreateBook(w http.ResponseWriter,r *http.Request){
	var book models.Book

	if err:=json.NewDecoder(r.Body).Decode(&book); err!=nil{
		http.Error(w,"Invalid Input",http.StatusBadRequest)
		return
	}

	createdBook,err := c.BookService.CreateBook(book.Title,book.Author)

	if err!=nil{
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdBook)
}

func (c *BookController)GetBookByID(w http.ResponseWriter,r *http.Request){
	

	id:= chi.URLParam(r,"id")

	//The book type need not be declared as it automatically infers the type that is returned
	//by the function
	book,err := c.BookService.FindByID(id)

	if err!=nil{
		http.Error(w,"Book not found",http.StatusNotFound)
		return
	}

	if err:=json.NewEncoder(w).Encode(book);err!=nil{
		http.Error(w,"Something went wrong",http.StatusInternalServerError)
		return
	}
}