package controllers

import (
	"encoding/json"
	"net/http"
	"query-availability-microservice/services"
)

type RequestPayload struct {
	ParkingLotID int `json:"id"`
}

type ResponsePayload struct {
	ParkingLotID int `json:"id"`
	Capacity     int `json:"capacity"`
}

func CheckAvailability(w http.ResponseWriter, r *http.Request) {
	// Leer el cuerpo de la solicitud
	var payload RequestPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil || payload.ParkingLotID <= 0 {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Obtener la capacidad del parqueadero
	capacity, err := services.GetParkingCapacity(payload.ParkingLotID)
	if err != nil {
		http.Error(w, "Parking lot not found", http.StatusNotFound)
		return
	}

	// Responder con la capacidad
	response := ResponsePayload{
		ParkingLotID: payload.ParkingLotID,
		Capacity:     capacity,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
