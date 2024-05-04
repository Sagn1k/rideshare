package handlers

import "github.com/gofiber/fiber/v2"

func AddUser(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusCreated)
}