package controllers

import("encoding/json"
"net/http"
"golang.org/x/crypto/bcrypt"
"clean-architecture/internal/models"
"clean-architecture/internal/repositories"
"clean-architecture/internal/auth"
)

type AuthController struct{
	UserRepo repositories.UserRepo
}

func(a *AuthController)Login(w http.ResponseWriter,r *http.Request){
	var creds struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err:=json.NewDecoder(r.Body).Decode(&creds);err!=nil{
		http.Error(w,"Invalid JSON",http.StatusBadRequest)
		return
	}

	user, err:= a.UserRepo.FindByName(creds.Username)

	if err!=nil || bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(creds.Password))!=nil{
		http.Error(w,"Invalid Credential",http.StatusUnauthorized)
		return
	}

	token,err := auth.GenerateToken(user.ID)
	if err!=nil{
		http.Error(w,"Failed to generate token",http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(map[string]string{"token":token})
}
func(a *AuthController)Register(w http.ResponseWriter,r *http.Request){
	
	var user *models.User

	if err:=json.NewDecoder(r.Body).Decode(&user);err!=nil{
		http.Error(w,"Invalid JSON",http.StatusBadRequest)
		return
	}

	hashedPassword, err:= bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)

	if err!= nil{
		http.Error(w,"Failed to hash password",http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)

	if err:= a.UserRepo.Register(user);err!=nil{
		http.Error(w,"Could not create new user",http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message":
	"Successfully Created user"})
}

