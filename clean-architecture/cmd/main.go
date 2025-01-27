package main

import (
	"clean-architecture/interfaces/controllers"
	"clean-architecture/interfaces/database"
	"clean-architecture/interfaces/routes"
	"clean-architecture/internal/usecases"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)



func main(){

	//setting the database connection and migration
	db := database.Connect()
	
	database.Migrate()

	//Setting the router
	r:=chi.NewRouter()

	//Initialize the Book repository
	//Lets think of a library management where BookRepo is the librarian and we give the key to the database
	//in this 
	bookRepo := &database.BookRepository{DB:db}

	//Initialize the Book service
	//It acts as basically the manager he donot do the work rather he instructs the librarian.
	bookService := &usecases.BookService{BookRepo:bookRepo}

	//Initialize the Book Controller
	//It is like the front desk where all people come and ask for help, They will guide you to manager.
	bookController := &controllers.BookController{BookService:bookService}


	//Setting up of userController
	userRepo := &database.UserRepo{DB:db}
	userController := &controllers.AuthController{UserRepo:userRepo} 

	r.Mount("/user",routes.UserRoutes(userController))

	//Accessing the routes
	r.Mount("/",routes.BookRoutes(bookController))

	if err:=godotenv.Load();err!=nil{
		log.Fatalf("Error in fetching environmental variables")
	}

	port:=os.Getenv("PORT")

	http.ListenAndServe(":"+port,r)
	
}