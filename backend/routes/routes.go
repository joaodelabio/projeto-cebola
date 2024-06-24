package routes

import (
	"projeto-semestral/ps-backend-joao-gabriel-ramon/controllers"

	"github.com/gorilla/mux"
)

func AuthRoutes(router *mux.Router) {
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
}
