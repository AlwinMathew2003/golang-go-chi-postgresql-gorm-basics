package routes

import (
    "github.com/go-chi/chi/v5"
    "go-crud-using-database1/handlers"
)

func AuthRoutes() chi.Router {
    r := chi.NewRouter()

    // Define the login route
    r.Post("/login", handlers.Login)

    return r
}
