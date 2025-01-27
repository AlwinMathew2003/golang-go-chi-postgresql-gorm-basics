package main

import("log" //for writing the log details
"os"
"net/http"
"github.com/joho/godotenv"
"go-crud-using-database1/db"
"go-crud-using-database1/routes"
"github.com/go-chi/chi/v5"
"github.com/go-chi/chi/v5/middleware"
)

func main(){

	//Connect to the Database
	db.Connect()

	//Migrage the Database schema
	db.Migrate()

	
	//Seed the dummy user data
	db.SeedDB()


	//Create a router
	r:= chi.NewRouter()

	//Setting middlewares
	r.Use(middleware.Logger)//To show logger informations
	r.Use(middleware.Recoverer)//To prevent unexpected crashes

	//Mounting the routes
	r.Mount("/books",routes.BookRoutes())
	r.Mount("/auth", routes.AuthRoutes())   // Authentication-related routes

	//Load the .env file and access the port
	if err:=godotenv.Load();err!=nil{
		log.Println("Error in accessing the environmental variable")
	}

	port:= os.Getenv("PORT")

	//start the server
	log.Printf("Server running at port:%s",port)
	http.ListenAndServe(":"+port,r)
}