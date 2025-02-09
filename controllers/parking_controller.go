package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"query-availability-microservice/services"
	"strconv"

	"github.com/gorilla/mux"
)

// CapacityResponse represents the API response structure
type CapacityResponse struct {
	ParkingLotID int `json:"id"`
	TotalSpace   int `json:"total_space"`
	Capacity     int `json:"capacity"`
}

// GetParkingCapacity handles requests to retrieve parking lot capacity
func GetParkingCapacity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		http.Error(w, `{"error": "invalid parking lot ID"}`, http.StatusBadRequest)
		return
	}

	capacity, totalSpace, err := services.FetchParkingCapacity(id)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err), http.StatusNotFound)
		return
	}

	response := CapacityResponse{
		ParkingLotID: id,
		TotalSpace:   totalSpace,
		Capacity:     capacity,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
