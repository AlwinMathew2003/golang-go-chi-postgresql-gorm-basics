package handlers

import("net/http"
"encoding/json" //Inorder to work with json data
"github.com/go-chi/chi/v5" //inorder to get the data from the url parameters
"go-crud/models"
)

var books = []models.Books{
	{ID: "1", Name: "The Go Programming Language", Author: "Alan A. A. Donovan"},
    {ID: "2", Name: "Clean Code", Author: "Robert C. Martin"},
}

//Listing all the books
func ListBooks(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(books) //Encoder() will convert it into json and it is send to response using
	// the NewEncoder() objec created
}

//Getting a Book by ID
func GetBook(w http.ResponseWriter,r *http.Request){
	//Access the id from the url
	id:= chi.URLParam(r,"id") //first we should mention the request, then the parameter

	//Search for the url if present
	for _,book:=range books{ //we are retreiving the elements from books so don't take care of model

		//Return the result
		if id == book.ID{
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w,"Book not found",http.StatusNotFound)
}


//Deleting a Book by ID
func DeleteBook(w http.ResponseWriter,r *http.Request){
	id:= chi.URLParam(r,"id")

	for i,book:=range books{
		if id == book.ID{
			books = append(books[:i],books[i+1:]...)//... is used to unpack the elements
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Book is Deleted")) //we should convert into byte array and send to the client

			return
		}
	}
	http.Error(w,"Book not found",http.StatusNotFound)
}

//Updating a Book by ID
func UpdateBook(w http.ResponseWriter,r *http.Request){
	var updatedBook models.Books //models.Books is the type
	
	//we should only mention the address in Decode()
	// access the data from r.Body
	if err:=json.NewDecoder(r.Body).Decode(&updatedBook);err!=nil{
		http.Error(w,"Invalid Json",http.StatusBadRequest)
		return
	}
	
	id := chi.URLParam(r,"id")
	for i,book:=range books{
		if id == book.ID{
			books[i] = updatedBook
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	http.Error(w,"Book not found",http.StatusNotFound)
}


//Creating a Book
func CreateBook(w http.ResponseWriter,r *http.Request){
	var book models.Books
	
	//The NewDecoder() will access the data from the body and Decode it and send to the address mentioned
	if err:=json.NewDecoder(r.Body).Decode(&book);err!=nil{
		http.Error(w,"Invalid Json",http.StatusBadRequest)
		return
	}
	books = append(books,book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}
