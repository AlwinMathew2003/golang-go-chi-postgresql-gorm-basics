package middleware

import("net/http"
"os"
"log"
"github.com/joho/godotenv"
"strings"
"github.com/golang-jwt/jwt/v5"
"context"
)

func JWTAuth(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		authHeader:=r.Header.Get("Authorization")

		if authHeader == ""{
			http.Error(w,"Token not found",http.StatusUnauthorized)
			return
		}

		if err:=godotenv.Load();err!=nil{
			log.Fatalf("Error in loading the environmental variable")
		}

		jwt_key:= os.Getenv("JWT_SECRET_KEY")

		tokenString:= strings.TrimPrefix(authHeader,"Bearer ")

		token,err:=jwt.Parse(tokenString,func(token *jwt.Token)(interface{},error){
			return []byte(jwt_key),nil
		})

		if err!=nil || !token.Valid{
			http.Error(w,"Invalid Token",http.StatusUnauthorized)
			return
		}

		claims,ok:= token.Claims.(jwt.MapClaims)


		if !ok{
			http.Error(w,"Invalid token claims",http.StatusUnauthorized)
			return
		}

	

		ctx:=context.WithValue(r.Context(),"user_id",claims["user_id"])

		next.ServeHTTP(w,r.WithContext(ctx))
	})
}