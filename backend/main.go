package main

import (
	"net/http"
	"projeto-semestral/ps-backend-joao-gabriel-ramon/config"
	"projeto-semestral/ps-backend-joao-gabriel-ramon/controllers"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	router := mux.NewRouter()

	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	http.ListenAndServe(":8080", router)
}
