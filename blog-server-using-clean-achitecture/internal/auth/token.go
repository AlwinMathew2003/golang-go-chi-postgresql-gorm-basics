package auth

import (
	"log"
	"os"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"time"
)
func GenerateToken(userID uint)(string,error){

	//Access the secret key
	if err:=godotenv.Load();err!=nil{
		log.Fatalf("Error in accessing the environmental variable")
	}

	jwt_key:=os.Getenv("JWT_SECRET_KEY")

	log.Printf("%T:%v",userID,userID)
	claims:= jwt.MapClaims{
		"user_id":userID,
		"exp":time.Now().Add(24*time.Hour).Unix(),
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	return token.SignedString([]byte(jwt_key))
}