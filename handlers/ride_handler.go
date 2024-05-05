package handlers

import (
	"rideshare/models"
	"rideshare/services"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type RideHandler struct {
	rideService    *services.RideService
	vehicleService *services.VehicleService
	userService    *services.UserService
	logger         *zap.Logger
}

func NewRideHandler(rideService *services.RideService, vehicleService *services.VehicleService,
	userService *services.UserService, logger *zap.Logger) *RideHandler {
	return &RideHandler{
		rideService:    rideService,
		vehicleService: vehicleService,
		userService:    userService,
		logger:         logger,
	}
}

// OfferRide allows users to offer a shared ride on a route with specific details
func (h *RideHandler) OfferRide(c *fiber.Ctx) error {

	var request models.RideOfferRequest
	if err := c.BodyParser(&request); err != nil {
		h.logger.Error("Failed to parse request body", zap.Error(err))
		return err
	}

	if err := h.vehicleService.CheckVehicleOwnership(request.DriverID, request.Vehicle); err != nil {
		h.logger.Error("Vehicle does not belong to user", zap.Error(err))
		return fiber.NewError(fiber.StatusBadRequest, "Vehicle does not belong to user")
	}

	rideDetail := models.Ride{
		DriverID:       request.DriverID,
		Source:         request.Source,
		Destination:    request.Destination,
		SeatsAvailable: request.SeatsAvailable,
		Vehicle:        strings.Split(request.Vehicle, ",")[0],
		Active:         true,
		Status:         "OFFERED",
	}

	if err := h.rideService.OfferRide(rideDetail); err != nil {
		h.logger.Error("Failed to offer ride", zap.Error(err))
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

// SelectRide allows users to select a ride from multiple offered rides
func (h *RideHandler) SelectRide(c *fiber.Ctx) error {
	var selectionDetail models.RideSelection
	if err := c.BodyParser(&selectionDetail); err != nil {
		h.logger.Error("Failed to parse request body", zap.Error(err))
		return err
	}

	ride, err := h.rideService.SelectRide(selectionDetail)
	if err != nil {
		h.logger.Error("Failed to select ride", zap.Error(err))
		return err
	}

	return c.JSON(ride)
}

// EndRide ends a ride based on ride details
func (h *RideHandler) EndRide(c *fiber.Ctx) error {
	var rideDetail models.RideManagement
	if err := c.BodyParser(&rideDetail); err != nil {
		h.logger.Error("Failed to parse request body", zap.Error(err))
		return err
	}

	if err := h.rideService.EndRide(rideDetail); err != nil {
		h.logger.Error("Failed to end ride", zap.Error(err))
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *RideHandler) GetAllRides(c *fiber.Ctx) error {

	rides, err := h.rideService.GetActiveRides()
	if err != nil {
		h.logger.Error("Failed to get active rides", zap.Error(err))
		return err
	}

	return c.JSON(rides)
}

func (h *RideHandler) PrintRideStatsByUser(c *fiber.Ctx) error {

	rideStats := make(map[string]*models.RideStats)

	users, err := h.userService.GetAllUsers()

	if err != nil {
		h.logger.Error("Failed to get all users", zap.Error(err))
		return err
	}

	for _, user := range users {

	    // Retrieve rides for the user
	    rides, err := h.rideService.GetRidesByUser(user.ID)
	    if err != nil {
	        
			h.logger.Error("Unable to fetch stats for User", zap.Error(err))
	        continue
	    }
	
	    
	    totalTaken := len(rides)
	    totalOffered := 0
	    for _, ride := range rides {
	        if ride.DriverID == user.ID {
	            totalOffered++
	        }
	    }
	
	    rideStats[user.Name] = &models.RideStats{
	        Taken:  totalTaken,
	        Offered: totalOffered,
	    }
	}

	return c.JSON(rideStats)
}
