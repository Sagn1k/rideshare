package handlers

import (
	"rideshare/models"
	"rideshare/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type VehicleHandler struct {
    vehicleService *services.VehicleService
    logger         *zap.Logger
}

func NewVehicleHandler(vehicleService *services.VehicleService, logger *zap.Logger) *VehicleHandler {
    return &VehicleHandler{
        vehicleService: vehicleService,
        logger:         logger,
    }
}

func (h *VehicleHandler) AddVehicle(c *fiber.Ctx) error {

    var vehicle models.Vehicle
    if err := c.BodyParser(&vehicle); err != nil {
        h.logger.Error("Failed to parse request body", zap.Error(err))
        return err
    }

    if err := h.vehicleService.AddVehicle(vehicle.UserID, vehicle); err != nil {
        h.logger.Error("Failed to add vehicle", zap.Error(err))
        return err
    }

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

    vehicle, err := h.vehicleService.GetVehicle(userID)
    if err != nil {
        h.logger.Error("Failed to get vehicle", zap.Error(err))
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(vehicle)
}