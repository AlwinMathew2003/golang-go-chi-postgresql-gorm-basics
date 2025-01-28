package routes

import("github.com/go-chi/chi/v5"
"blog-server-using-clean-architecture/interface/controllers")

func UserRoutes(AuthController *controllers.AuthController,UserController *controllers.UserController)chi.Router{

	r:= chi.NewRouter()

	r.Post("/login",AuthController.Login)
	r.Post("/Register",UserController.Register)

	return r
}