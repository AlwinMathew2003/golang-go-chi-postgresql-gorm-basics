package controllers

import (
	"blog-server-using-clean-architecture/internal/usecases"
	"encoding/json"
	"log"
	"net/http"
	"blog-server-using-clean-architecture/internal/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{
	AuthService *usecases.UserService
}

func (c *AuthController)Login(w http.ResponseWriter,r *http.Request){
	var credential struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err:=json.NewDecoder(r.Body).Decode(&credential);err!=nil{
		http.Error(w,"Invalid JSON",http.StatusBadRequest)
		return
	}

	user,err:=c.AuthService.FindByEmailService(credential.Email)

	if err!=nil{
		http.Error(w,"Password not found",http.StatusInternalServerError)
		return
	}

	if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(credential.Password));err!=nil{
		http.Error(w,"Invalid Credentail",http.StatusUnauthorized)
		return
	}

	//token generation
	token,err:= auth.GenerateToken(user.ID)

	if err!=nil{
		log.Fatalf("Failed to generate token")
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(map[string]string{"token":token})
}