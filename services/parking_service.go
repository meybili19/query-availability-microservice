package services

import (
	"errors"
	"query-availability-microservice/config"
	"query-availability-microservice/models"
)

func GetParkingCapacity(id int) (int, error) {
	var parkingLot models.ParkingLot

	// Query para obtener el parqueadero por ID
	err := config.DB.QueryRow(
		"SELECT id, name, address, capacity FROM ParkingLot WHERE id = ?", id,
	).Scan(&parkingLot.ID, &parkingLot.Name, &parkingLot.Address, &parkingLot.Capacity)

	if err != nil {
		return 0, errors.New("Parking lot not found")
	}

	return parkingLot.Capacity, nil
}
