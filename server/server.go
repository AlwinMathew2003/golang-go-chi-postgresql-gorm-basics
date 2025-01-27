package main

import ("fmt"
"log"
"net/http"
"github.com/go-chi/chi/v5")

func main(){
	router:= chi.NewRouter()
	router.Get("/",func(w http.ResponseWriter,router *http.Request){
		fmt.Fprintf(w,"Hello,world!")
	})

	log.Println("Server is running on localhost:8080")
	http.ListenAndServe(":8080",router)

}