package repositories

//Firstly we need to get the model so we need to import it from the models folder
import("clean-architecture/internal/models")

//Book repository defines the methods to interact with the book data
type BookRepository interface{
	//This repository will be implemented in the database layer
	//In this function it recieves the books data so its type should be *models.Book
	Save(book *models.Book) (*models.Book,error)
	//Using pointers allows to avoid making a copy of book object.
	//If we want to modify the original model such that updating an ID we use pointers.
	//Return a pointer allows the user to access the modified data
	//If we do not use pointers then it will not make any changes in the original data outside this function

	GetByID(ID string) (*models.Book,error)
}