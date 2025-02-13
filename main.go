package main

import (
	"log"
	"net/http"
	"query-availability-microservice/config"
	"query-availability-microservice/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	router := mux.NewRouter()

	router.HandleFunc("/parkinglot/capacity/{id}", controllers.GetParkingCapacity).Methods("GET")

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),                                       // Permitir cualquier origen
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // Métodos permitidos
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),           // Headers permitidos
	)

	log.Println("🚀 -Server running on 0.0.0.0:7004")
	http.ListenAndServe("0.0.0.0:7004", corsHandler(router))
}
