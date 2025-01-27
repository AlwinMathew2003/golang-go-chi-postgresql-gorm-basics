//Token generation we need header,payload,signature in jwt token
package utils

import(
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"os"
	"log"
)

func GenerateToken(userID uint)(string,error){ // it will return the token as string 

	//set the payload
	claims := jwt.MapClaims{ //It is used for the payload
		"user_id": userID,
		"exp":time.Now().Add(24*time.Hour).Unix(),
	}

	//attatch the payload and generate the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	//access the secret key from environmental variable
	if err:=godotenv.Load();err!=nil{
		log.Println("Error in accessing environmental variable")
	}
	jwt_key := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(jwt_key))
}