package main

import (
	"log"
	"net/http"
	"query-availability-microservice/config"
	"query-availability-microservice/controllers"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	router := mux.NewRouter()
	router.HandleFunc("/parkinglot/capacity/{id}", controllers.GetParkingCapacity).Methods("GET")

	log.Println("Server running on port 8080")
	http.ListenAndServe(":6005", router)
}
