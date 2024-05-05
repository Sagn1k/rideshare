package handlers

import (
	"rideshare/models"
	"rideshare/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// VehicleHandler handles vehicle-related requests
type VehicleHandler struct {
    vehicleService *services.VehicleService
    logger         *zap.Logger
}

// NewVehicleHandler creates a new instance of VehicleHandler with the provided services
func NewVehicleHandler(vehicleService *services.VehicleService, logger *zap.Logger) *VehicleHandler {
    return &VehicleHandler{
        vehicleService: vehicleService,
        logger:         logger,
    }
}

// AddVehicle adds a new vehicle for the specified user
func (h *VehicleHandler) AddVehicle(c *fiber.Ctx) error {

    var vehicle models.Vehicle
    if err := c.BodyParser(&vehicle); err != nil {
        h.logger.Error("Failed to parse request body", zap.Error(err))
        return err
    }

    // Call the service function to add the vehicle for the user
    if err := h.vehicleService.AddVehicle(vehicle.UserID, vehicle); err != nil {
        h.logger.Error("Failed to add vehicle", zap.Error(err))
        return err
    }

    // Respond with a success status
    return c.SendStatus(fiber.StatusCreated)
}

func (h *VehicleHandler) GetVehicle(c *fiber.Ctx) error {
    userID, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        h.logger.Error("Invalid user ID", zap.Error(err))
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid user ID",
        })
    }

    // Call the service function to get the vehicle for the user
    vehicle, err := h.vehicleService.GetVehicle(userID)
    if err != nil {
        h.logger.Error("Failed to get vehicle", zap.Error(err))
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    // Respond with the vehicle details
    return c.JSON(vehicle)
}