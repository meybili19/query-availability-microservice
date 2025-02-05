package controllers

import (
	"encoding/json"
	"net/http"
	"query-availability-microservice/services"
	"strconv"

	"github.com/gorilla/mux"
)

type CapacityResponse struct {
	ParkingLotID int `json:"id"`
	Capacity     int `json:"capacity"`
}

func GetParkingCapacity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		http.Error(w, "Invalid parking lot ID", http.StatusBadRequest)
		return
	}

	capacity, err := services.FetchParkingCapacity(id)
	if err != nil {
		http.Error(w, "Parking lot not found", http.StatusNotFound)
		return
	}

	response := CapacityResponse{
		ParkingLotID: id,
		Capacity:     capacity,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
