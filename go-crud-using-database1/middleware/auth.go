package middleware

import(
	"context"
	"strings"
	"net/http"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/golang-jwt/jwt/v5"
)

//http.Handler is the interface(it means it can be of any type)
func JWTAuth(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		
		//Get the data from the header
		authHeader:= r.Header.Get("Authorization")

		//Check whether the auth header is empty
		if authHeader==""{
			http.Error(w,"Missing token",http.StatusUnauthorized)
			return
		}

		//Retreive the token by trimming out the prefix
		tokenStr:= strings.TrimPrefix(authHeader,"Bearer ")


		//Access the secret key from the environmental variabl
		if err:= godotenv.Load();err!=nil{
			log.Fatalln("Error access the environmental variable")
			return
		}
		jwt_key := os.Getenv("JWT_SECRET")

		//Parse the token to get the actual token
		//In this one parameter is token and other is the secret key to check with the token
		token,err := jwt.Parse(tokenStr,func(token *jwt.Token)(interface{},error){
			return []byte(jwt_key),nil
		})

		if err!=nil || !token.Valid{
			http.Error(w,"Invalid Token",http.StatusUnauthorized)
			return
		}

		claims,ok := token.Claims.(jwt.MapClaims)

		if!ok{
			http.Error(w,"Invalid token claims",http.StatusUnauthorized)
			return
		}
		
		ctx:= context.WithValue(r.Context(),"user_id",claims["usre_id"])

		next.ServeHTTP(w,r.WithContext(ctx))
	})

}