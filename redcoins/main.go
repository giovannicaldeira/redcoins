package main

import (
	"fmt"
	"net/http"
	"os"

	"redcoins/app"
	"redcoins/controllers"
	_ "redcoins/models"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router.Use(app.JwtAuthentication)

	router.HandleFunc("/", homeLink)
	router.HandleFunc("/redcoins/api/user/new", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/redcoins/api/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/redcoins/api/operation/new", controllers.CreateOperation).Methods("POST")

	fmt.Println("Server up on port: " + port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}
}
