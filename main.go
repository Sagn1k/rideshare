package main

import (
	"rideshare/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/user", handlers.AddUser)
	app.Listen(":3000")
}
