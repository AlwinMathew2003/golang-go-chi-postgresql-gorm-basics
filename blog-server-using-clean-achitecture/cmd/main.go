package main

import (
	"net/http"
	// "fmt"
	"blog-server-using-clean-architecture/interface/controllers"
	"blog-server-using-clean-architecture/interface/database"
	"blog-server-using-clean-architecture/interface/routes"
	"blog-server-using-clean-architecture/internal/usecases"
	"log"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)
func main(){
	//Access the environmental variables(loading and then accessing)
	//Database connection
	//Migration
	//Initialize the new router to recieve the routes
	//Implementing the Controllers
	//Implementing the Routes
	//start the server

	if err:= godotenv.Load();err!=nil{
		log.Fatalf("Error in loading the environmental variables")
		return
	}

	//Database Connection and Migration
	db := database.Connect()

	database.Migrate(db)
	
	database.SeedDB(db)



	//Initializing the Router
	r:= chi.NewRouter()

	//Checking the route
	// r.Get("/",func(w http.ResponseWriter, r *http.Request){
	// 	fmt.Fprintf(w,"Hello World")
	// })

	//setting the Post Controller
	postRepo := &database.PostRepositoryDB{DB:db}
	postService := &usecases.PostService{PostRepo:postRepo}
	postController := &controllers.PostController{PostService: postService}

	//setting the user Controller
	userRepo := &database.UserRepositoryDB{DB:db}
	userService :=&usecases.UserService{UserRepo: userRepo}
	userController :=&controllers.UserController{UserService: userService}
	AuthController :=&controllers.AuthController{AuthService: userService}

	//Setting the routes
	r.Mount("/posts",routes.PostRoutes(postController))
	r.Mount("/user",routes.UserRoutes(AuthController,userController))
	//Accessing the port
	port := os.Getenv("PORT")

	http.ListenAndServe(":"+port,r)
}