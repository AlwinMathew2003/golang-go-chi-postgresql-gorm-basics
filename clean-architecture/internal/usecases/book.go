package usecases

import("clean-architecture/internal/repositories"
"clean-architecture/internal/models"
"errors")
//It will validate the book data and calls the repository to save it.
//Inorder to call the repository we should get the repository type as struct
//Since BookRepository is an interface type it assigns reference to that type rather than copying it
//So since we are giving an interface type we do not need to use a pointer

type BookService struct{
	//By making it as struct we are doing an dependency injection
	//This decouples the business logic from the implementation details of the repository
	//important: Once you are using struct we can add methods to it
	//we can also create multiple instance 
	BookRepo repositories.BookRepository
}

//CreateBook handles creating a new book
//CreateBook() method validates the data and get the new entity for data and send it to the Save() method
//To save it in the database
func(s *BookService)CreateBook(title,author string)(*models.Book,error){

	//Data validation
	if title== "" || author == "" {
		return nil,errors.New("Invalid Book Data")
	}

	//if the received data contains data then create a new data entity.
	book:=&models.Book{
		Title: title,
		Author: author,
	}

	//It is give a pointer of the model to the Save() method.
	return s.BookRepo.Save(book)
}

func(s *BookService)FindByID(id string)(*models.Book,error){
	
	return s.BookRepo.GetByID(id)
}
