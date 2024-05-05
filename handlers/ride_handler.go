package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"rideshare/models"
	"rideshare/services"
)

type RideHandler struct {
	rideService    *services.RideService
	vehicleService *services.VehicleService
	logger         *zap.Logger
}

func NewRideHandler(rideService *services.RideService, vehicleService *services.VehicleService,
	logger *zap.Logger) *RideHandler {
	return &RideHandler{
		rideService:    rideService,
		vehicleService: vehicleService,
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
