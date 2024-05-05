package services

import (
	"errors"
	"rideshare/models"

	"go.uber.org/zap"
)

var vehicleIDCounter int

type VehicleService struct {
	logger   *zap.Logger
	vehicles map[int]models.Vehicle
}

// NewVehicleService creates a new instance of VehicleService with the provided logger
func NewVehicleService(logger *zap.Logger) *VehicleService {
	return &VehicleService{logger: logger,
		vehicles: make(map[int]models.Vehicle)}
}

// AddVehicle adds a new vehicle for the specified user
func (s *VehicleService) AddVehicle(userID int, vehicle models.Vehicle) error {
	vehicle.ID = vehicleIDCounter

	if _, exists := s.vehicles[userID]; exists {
        return errors.New("user already has a vehicle")
    }

	s.vehicles[userID] = vehicle 

	vehicleIDCounter++
	s.logger.Info("Vehicle added successfully", zap.Int("userID", userID))
	return nil
}

func (s *VehicleService) GetVehicle(userID int) (*models.Vehicle, error) {
    
    vehicle, exists := s.vehicles[userID]
    if !exists {
        return nil, errors.New("user does not have a vehicle")
    }

    return &vehicle, nil
}
