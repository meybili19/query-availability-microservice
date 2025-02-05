package services

import (
	"errors"
	"query-availability-microservice/config"
)

func FetchParkingCapacity(id int) (int, error) {
	var capacity int

	err := config.DB.QueryRow(
		"SELECT capacity FROM ParkingLot WHERE id = ?", id,
	).Scan(&capacity)

	if err != nil {
		return 0, errors.New("Parking lot not found")
	}

	return capacity, nil
}
