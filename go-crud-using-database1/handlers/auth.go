package handlers

import (
	"encoding/json"
	"go-crud-using-database1/db"
	"go-crud-using-database1/models"
	"go-crud-using-database1/utils"
	"net/http"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request){
	//Declare a variable to store the request data
	var cred struct{
		Username string `json:"username"`
		Password string	`json:"password"`
	}

	//Accept the data from the request and store it in cred variable
	if err:=json.NewDecoder(r.Body).Decode(&cred); err!=nil{
		http.Error(w,"Invalid JSON",http.StatusBadRequest)
		return
	}

	//Get the data from the database where cred username matches
	var person models.User

	if err:= db.DB.Where("username=?",cred.Username).First(&person).Error;err!=nil{
		if err == gorm.ErrRecordNotFound{
			http.Error(w,"User not found",http.StatusUnauthorized)
			return
		}else{
			http.Error(w,"Something went wrong",http.StatusInternalServerError)
			return
		}
	}

	//Verify the password
	//user.password contains the hashed passowrd and it will be compared the the plain text password
	if err:= utils.VerifyPassword(person.Password,cred.Password);err!=nil{
		http.Error(w,"Invalid Credential",http.StatusUnauthorized)
		return
	}

	//Generate token
	token,err:= utils.GenerateToken(person.ID)

	//Check if there is any error while creating token
	if err!=nil{
		http.Error(w,"Could not generate token",http.StatusInternalServerError)
		return
	}

	//Send the token to the client header
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(map[string]string{"token":token})
}