package routes

import (
	"clean-architecture/interfaces/controllers"
	"clean-architecture/interfaces/middleware"
	"github.com/go-chi/chi/v5"
)

func BookRoutes(bookController *controllers.BookController)chi.Router{
	r := chi.NewRouter()

	r.Group(func(r chi.Router){
		r.Use(middleware.JWTMiddleware)
		r.Post("/books",bookController.CreateBook)
		r.Get("/books/{id}",bookController.GetBookByID)

	})

	return r
}