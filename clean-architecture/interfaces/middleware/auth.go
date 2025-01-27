package middleware

import("net/http"
"os"
"log"
"github.com/joho/godotenv"
"strings"
"github.com/golang-jwt/jwt/v5"
"context")

func JWTMiddleware(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		authHeader := r.Header.Get("Authorization")

		if authHeader == ""{
			http.Error(w,"Missing Token",http.StatusUnauthorized)
			return
		}

		//Accessing the secret key
		if err:= godotenv.Load();err!=nil{
			log.Fatalf("Error in loading environmental vairable")
		}

		jwt_key := os.Getenv("JWT_SECRET_KEY")

		tokenString := strings.TrimPrefix(authHeader,"Bearer ")
		token,err := jwt.Parse(tokenString,func(token *jwt.Token)(interface{},error){
			return []byte(jwt_key),nil
		})

		if err!=nil || !token.Valid{
			http.Error(w,"Invalid Token",http.StatusUnauthorized)
			return
		}

		claims,ok := token.Claims.(jwt.MapClaims)

		if !ok{
			http.Error(w,"Invalid Token Claims",http.StatusUnauthorized)
			return
		}

		ctx:= context.WithValue(r.Context(),"userID",claims["userID"])

		next.ServeHTTP(w,r.WithContext(ctx))
	})
}