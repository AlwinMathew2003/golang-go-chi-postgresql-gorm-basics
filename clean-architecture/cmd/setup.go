// package main

// import (
// 	"clean-architecture/interfaces/controllers"
// 	"clean-architecture/interfaces/database"
// 	"clean-architecture/internal/usecases"

// 	"gorm.io/gorm"
// )

// func InitializeComponents(db *gorm.DB)(*controllers.BookController,error){
	
// 	//Initialize the Book repository
// 	bookRepo := &database.BookRepository{DB:db}

// 	//Initialize the Book service
// 	bookService := &usecases.BookService{BookRepo:bookRepo}

// 	//Initialize the Book Controller
// 	bookController := &controllers.BookController{BookService:bookService}

// 	return bookController,nil
// }