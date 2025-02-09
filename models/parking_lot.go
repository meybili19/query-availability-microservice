package models

type ParkingLot struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	TotalSpace int    `json:"total_space"`
	Capacity   int    `json:"capacity"`
}
