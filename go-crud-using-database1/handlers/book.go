package handlers

import("net/http"
"encoding/json" //Inorder to work with json data
"github.com/go-chi/chi/v5" //inorder to get the data from the url parameters
"go-crud-using-database1/db"
"go-crud-using-database1/models"
)

// var books = []models.Books{
// 	{ID: "1", Name: "The Go Programming Language", Author: "Alan A. A. Donovan"},
//     {ID: "2", Name: "Clean Code", Author: "Robert C. Martin"},
// }

//Listing all the books
func ListBooks(w http.ResponseWriter, r *http.Request){

	//Setting the variable to store the details fetched from the database
	var books []models.Books

	//We donot need to mention the table name based on the struct type it will 
	//Automatically map to the required table and retreive the data
	//As we are giving &books it will check the slice type and figure out the table
	if err:=db.DB.Find(&books).Error; err!=nil{ //Retrieve all the data in the books table in the database
		http.Error(w,"Could not fetch books",http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books) //Encoder() will convert it into json and it is send to response using
	// the NewEncoder() objec created
}

//Getting a Book by ID
func GetBook(w http.ResponseWriter,r *http.Request){
	//Access the id from the url
	id:= chi.URLParam(r,"id") //first we should mention the request, then the parameter

	//Setting the variable to the details from the database
	var book models.Books

	if err:=db.DB.First(&book,id).Error;err!=nil{
		http.Error(w,"Book not found",http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book) 
	//pointer or values can be passed for large struct pointer is prefered
}


//Deleting a Book by ID
func DeleteBook(w http.ResponseWriter,r *http.Request){
	id:= chi.URLParam(r,"id")

	//Deleting the data from the database
	if err:=db.DB.Delete(&models.Books{},id).Error; err!=nil{
		http.Error(w,"Could not delete book",http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Book deleted successfully"))

}

//Updating a Book by ID
func UpdateBook(w http.ResponseWriter,r *http.Request){

	//Access the id from the url
	id:= chi.URLParam(r,"id") //We should specify the request router and id

	//Setting a variable to store the request body
	var Updatedbook models.Books

	//Retreiving the data from the request body
	if err:=json.NewDecoder(r.Body).Decode(&Updatedbook);err!=nil{
		http.Error(w,"Invalid JSON",http.StatusInternalServerError)
		return
	}

	//Setting the variable to store the data retreived from the database
	var book models.Books

	//Accessing the data from the database
	if err:=db.DB.First(&book,id).Error;err!=nil{
		http.Error(w,"Book not found",http.StatusNotFound)
		return
	}

	//Updating the present data
	book.Name = Updatedbook.Name
	book.Author = Updatedbook.Author

	//Saving to the database
	if err:=db.DB.Save(&book).Error;err!=nil{
		http.Error(w,"Error in saving data",http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Updated successfully"))
	json.NewEncoder(w).Encode(book)
}


//Creating a Book
func CreateBook(w http.ResponseWriter,r *http.Request){

	//Setting a variable to store data
	var book models.Books

	//Getting the details from the request body
	if err:=json.NewDecoder(r.Body).Decode(&book);err!=nil{
		http.Error(w,"Invalid JSON",http.StatusBadRequest)
		return
	}

	//Creating the data in the database
	if err:=db.DB.Create(&book).Error;err!=nil{
		http.Error(w,"Something went wrong",http.StatusInternalServerError)
		return
	}

	//Showing the result
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}
