package handlers

import (
	"rideshare/models"
	"rideshare/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserHandler struct {
	userService *services.UserService
	logger      *zap.Logger
}

func NewUserHandler(userService *services.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger,
	}
}

func (h *UserHandler) AddUser(c *fiber.Ctx) error {
	var userDetail models.User
	if err := c.BodyParser(&userDetail); err != nil {
		return err
	}

	if err := h.userService.AddUser(userDetail); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusCreated)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {

	userIDStr := c.Params("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	user, err := h.userService.GetUser(userID)
	if err != nil {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(user)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {

	users, err := h.userService.GetAllUsers()
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get users",
		})
	}

	return c.JSON(users)
}
