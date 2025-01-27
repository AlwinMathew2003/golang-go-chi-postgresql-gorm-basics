package main

import("os"
"github.com/joho/godotenv"
"net/http"
"log"
"github.com/go-chi/chi/v5"
"github.com/go-chi/chi/v5/middleware"
)

func main(){

	//Loading the environmental variable
	err:= godotenv.Load()

	if err != nil{
		log.Println("Error in loading the environmental variable")
	}

	//creating new router
	r := chi.NewRouter()

	//logging information of the request
	r.Use(middleware.Logger)

	//Accepting the request
	r.Get("/",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("OK"))
	})

	//Accessing the environmental variables
	port := os.Getenv("PORT")
	if port == ""{
		log.Fatal("Port is not set in environmental variable")
	}

	//Mount the Books Routes
	r.Mount("/books",BooksRoutes())

	//Starting the server if the port is fetched correctly
	log.Printf("Server running at port:%s",port)
	http.ListenAndServe(":"+port,r)
}

func BooksRoutes()chi.Router{
	
	//Create new router
	r:=chi.NewRouter()
	bookHandler := BookHandler{}
	
	//Defining the routes
	r.Get("/",bookHandler.ListBooks)
	r.Post("/",bookHandler.CreateBook)
	r.Get("/{id}",bookHandler.GetBooks)
	r.Put("/{id}",bookHandler.UpdateBook)
	r.Delete("/{id}",bookHandler.DeleteBook)

	//Returning the routes
	return r
}
//Running the code:
// go run main.go
// curl localhost:3000/ (To see the output)
// Inorder to run all the go files: go run *.go (not working)