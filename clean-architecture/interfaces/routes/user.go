package routes

import(
"github.com/go-chi/chi/v5"
"clean-architecture/interfaces/controllers"
)


func UserRoutes(userController *controllers.AuthController)chi.Router{
	
	r:= chi.NewRouter()

	r.Post("/Login",userController.Login)
	r.Post("/Register",userController.Register)

	return r

}
