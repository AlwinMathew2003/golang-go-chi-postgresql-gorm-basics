package controllers

import (
	"blog-server-using-clean-architecture/internal/models"
	"blog-server-using-clean-architecture/internal/usecases"
	"net/http"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
)
type UserController struct{
	UserService *usecases.UserService
}

func (c *UserController)Register(w http.ResponseWriter,r *http.Request){
	var user *models.User

	if err:= json.NewDecoder(r.Body).Decode(&user);err!=nil{
		http.Error(w,"Invalid JSON",http.StatusBadRequest)
		return
	}

	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)

	if err!=nil{
		http.Error(w,"Could not hash the password",http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword) //It is important to convert hashed password to string

	if err:=c.UserService.RegisterService(user);err!=nil{
		http.Error(w,"User is not created",http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}