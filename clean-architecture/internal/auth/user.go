package auth

import(
	"os"
	"log"
	"github.com/joho/godotenv"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint)(string,error){
if err:=godotenv.Load();err!=nil{
	log.Fatalf("Error in fetching the environmental vairable")
}

//Accessing the secret key
jwt_key := os.Getenv("JWT_SECRET_KEY")

//Defining the structure of claims
claims := jwt.MapClaims{
	"userID":userID,
	"exp":time.Now().Add(24*time.Hour).Unix(),
}

//Generating the token
token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

//Signing the token
return token.SignedString([]byte(jwt_key))

}