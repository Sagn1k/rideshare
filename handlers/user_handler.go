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

// NewUserHandler creates a new instance of UserHandler with the provided services
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
    
    // Call the service function to add the user
    if err := h.userService.AddUser(userDetail); err != nil {
        return err
    }
	return c.SendStatus(fiber.StatusCreated)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
    // Get the user ID from the request parameters
    userIDStr := c.Params("id")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        // If the ID is not a valid integer, return a bad request error
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid user ID",
        })
    }
    
    // Call the service function to get the user by ID
    user, err := h.userService.GetUser(userID)
    if err != nil {
        // If there was an error retrieving the user, return a not found error
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "User not found",
        })
    }
    
    // Respond with the user details
    return c.JSON(user)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
    // Call the service function to get all users
    users, err := h.userService.GetAllUsers()
    if err != nil {
        // If there was an error retrieving users, return an internal server error
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to get users",
        })
    }
    
    // Respond with the list of users
    return c.JSON(users)
}