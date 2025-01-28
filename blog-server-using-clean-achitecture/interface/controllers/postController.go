package controllers

import (
	"github.com/go-chi/chi/v5"
	"blog-server-using-clean-architecture/internal/models"
	"blog-server-using-clean-architecture/internal/usecases"
	"encoding/json"
	"net/http"
)

type PostController struct {
	PostService *usecases.PostService
}

//Get post by id
func (c *PostController)GetPostByID(w http.ResponseWriter,r *http.Request){

	id := chi.URLParam(r,"id")

	post,err := c.PostService.GetPostByIDService(id)//post can be declare like this

	if err!=nil{
		http.Error(w,"No post found",http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(post)
}


func (c *PostController)GetAllPost(w http.ResponseWriter,r *http.Request){

	posts,err:= c.PostService.GetAllPostService()

	if err!=nil{
		http.Error(w,"Something went wrong",http.StatusInternalServerError)
		return
	}

	count := len(posts)

	if count == 0{
		json.NewEncoder(w).Encode(map[string]string{"message":"No posts found"})
	}else{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(posts)
	}

}

func (c *PostController)CreatePost(w http.ResponseWriter,r *http.Request){

	//Access user id from context

	user_id := r.Context().Value("user_id")

	// Attempt to assert userIDValue to float64
	userIDFloat, ok := user_id.(float64)
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusUnauthorized)
		return
	}

	// Convert float64 to uint
	userID := uint(userIDFloat)
		
	//Access the request body
	var post *models.Post

	if err:=json.NewDecoder(r.Body).Decode(&post);err!=nil{
		http.Error(w,"Invalid JSON",http.StatusBadRequest)
		return
	}

	post1,err := c.PostService.CreatePostService(post,userID)

	if err!=nil{
		http.Error(w,"Error creating new post",http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post1)
}

func (c *PostController)UpdatePost(w http.ResponseWriter,r *http.Request){

	var post *models.Post

	id:= chi.URLParam(r,"id")

	//Access the updated post from the url
	if err:= json.NewDecoder(r.Body).Decode(&post);err!=nil{
		http.Error(w,"Invalid JSON",http.StatusBadRequest)
		return
	}

	post1,err := c.PostService.UpdatePostService(post,id)

	if err!=nil{
		http.Error(w,"Error in Updating Post",http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post1)
}

func (c *PostController)DeletePost(w http.ResponseWriter,r *http.Request){
	id:= chi.URLParam(r,"id")

	if err:=c.PostService.DeletePostService(id);err!=nil{
		http.Error(w,"Error in Deleting Post",http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message":"Deleted successfully"})
}