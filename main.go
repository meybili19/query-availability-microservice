package main

import (
	"log"
	"net/http"
	"query-availability-microservice/config"
	"query-availability-microservice/controllers"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	router := mux.NewRouter()
	router.HandleFunc("/webhook/availability", controllers.CheckAvailability).Methods("POST")

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
