package routes

import("github.com/go-chi/chi/v5"
"blog-server-using-clean-architecture/interface/controllers"
"blog-server-using-clean-architecture/interface/middleware")

func PostRoutes(postController *controllers.PostController)chi.Router{
	r:= chi.NewRouter()

	r.Get("/",postController.GetAllPost)
	
	r.Group(func(r chi.Router){
		r.Use(middleware.JWTAuth)
		r.Get("/{id}",postController.GetPostByID)
		r.Post("/create",postController.CreatePost)
		r.Put("/update/{id}",postController.UpdatePost)
		r.Delete("/delete/{id}",postController.DeletePost)
	})


	return r
}