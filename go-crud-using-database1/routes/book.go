package routes

//import the handlers function and chi

import ("go-crud-using-database1/handlers"
"github.com/go-chi/chi/v5"
"go-crud-using-database1/middleware"
)

// Set all the routes in BookRoutes and bind it together
func BookRoutes() chi.Router{
	//set the new router
	r:= chi.NewRouter()

	//Declare all the routes and functionality of the routes will be inside the handlers


	//Group the routes which need jwt authentication

	r.Group(func(r chi.Router){

	r.Use(middleware.JWTAuth)

	//list
	r.Get("/",handlers.ListBooks)
	//create
	r.Post("/",handlers.CreateBook)
	//update
	r.Put("/{id}",handlers.UpdateBook)
	//delete
	r.Delete("/{id}",handlers.DeleteBook)
	//get
	r.Get("/{id}",handlers.GetBook)

	})


	//After setting all the routes return all the routes

	return r
}