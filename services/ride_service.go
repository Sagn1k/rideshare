package services

import (
	"errors"
	"rideshare/models"
	"strings"

	"go.uber.org/zap"
)

type RideService struct {
	logger    *zap.Logger
	rides     map[int]models.Ride
}

// NewRideService creates a new instance of RideService with the provided logger
func NewRideService(logger *zap.Logger) *RideService {
	return &RideService{logger: logger,
		rides:     make(map[int]models.Ride)}
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

	if strings.Contains(selectionDetail.SelectionStrategy, "Preferred Vehicle=") {

		parts := strings.Split(selectionDetail.SelectionStrategy, "=")

		if len(parts) != 2 {
			return nil, errors.New("invalid vehicle details")
		}

		vehicleName := strings.TrimSpace(parts[1])

		for _, ride := range s.rides {

			if ride.Source == selectionDetail.Source && ride.Destination == selectionDetail.Destination &&
				ride.SeatsAvailable >= selectionDetail.Seats && ride.Vehicle == vehicleName {

				ride.Status = "BOOKED"
				ride.SeatsAvailable -= selectionDetail.Seats
				ride.Passenger = selectionDetail.UserId

				s.rides[ride.ID] = ride

				return &ride, nil
			}
		}

	} else {
		for _, ride := range s.rides {
			if ride.Source == selectionDetail.Source && ride.Destination == selectionDetail.Destination && ride.SeatsAvailable >= selectionDetail.Seats {

				ride.Status = "BOOKED"
				ride.SeatsAvailable -= selectionDetail.Seats
				ride.Passenger = selectionDetail.UserId

				s.rides[ride.ID] = ride

				return &ride, nil
			}
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

func (s *RideService) GetRidesByUser(userId int) ([]models.Ride, error) {

	var userRides []models.Ride

	for _, ride := range s.rides {
		if ride.DriverID == userId || ride.Passenger == userId{
			userRides = append(userRides, ride)
		}
	}
	return userRides, nil
}
