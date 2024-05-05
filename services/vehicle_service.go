package services

import "rideshare/models"

var vehicles map[int][]models.Vehicle
var vehicleIDCounter int

func addVehicle(userID int, vehicleDetail models.Vehicle) error {
	vehicleDetail.ID = vehicleIDCounter
	vehicles[userID] = append(vehicles[userID], vehicleDetail)

	vehicleIDCounter++

	return nil
}
