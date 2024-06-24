package main

import (
	"log"
	"myapp/config"
	"myapp/handlers"
	"net/http"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/api/register", handlers.Register)
	http.HandleFunc("/api/login", handlers.Login)

	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
