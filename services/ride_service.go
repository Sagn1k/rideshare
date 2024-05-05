package services

import (
	"errors"
	"rideshare/models"

	"go.uber.org/zap"
)

type RideService struct {
	logger *zap.Logger
	rides  map[int]models.Ride
}

// NewRideService creates a new instance of RideService with the provided logger
func NewRideService(logger *zap.Logger) *RideService {
	return &RideService{logger: logger,
		rides: make(map[int]models.Ride)}
}

func (s *RideService) OfferRide(rideDetail models.Ride) error {
    
    s.logger.Info("Offering ride", zap.Int("driverID", rideDetail.DriverID))
    
    rideID := len(s.rides) + 1
    rideDetail.ID = rideID
    
    // Store the ride in the map
    s.rides[rideID] = rideDetail
    
    return nil
}

func (s *RideService) SelectRide(selectionDetail models.RideSelection) (*models.Ride, error) {

    s.logger.Info("Selecting ride", zap.String("source", selectionDetail.Source), zap.String("destination", selectionDetail.Destination))
    
    for _, ride := range s.rides {
        if ride.Source == selectionDetail.Source && ride.Destination == selectionDetail.Destination && ride.SeatsAvailable >= selectionDetail.Seats {
            return &ride, nil
        }
    }
    
    return nil, errors.New("no matching ride found")
}

func (s *RideService) EndRide(rideDetail models.RideManagement) error {
    
    s.logger.Info("Ending ride", zap.Int("rideID", rideDetail.RideID))

	ride := s.rides[rideDetail.RideID]

	ride.Active = false
    
    s.rides[rideDetail.RideID] = ride
    
    return nil
}

func (s *RideService) GetActiveRides() ([]models.Ride, error) {
    
    var activeRides []models.Ride

    for _, ride := range s.rides {
        if ride.Active {
            activeRides = append(activeRides, ride)
        }
    }
    return activeRides, nil
}
