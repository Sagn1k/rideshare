package main

import (
	"log"
	"os"
	"rideshare/handlers"
	"rideshare/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:         "json",
		OutputPaths:      []string{"stdout", "app.log"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    zap.NewProductionEncoderConfig(),
	}

	// Create a Zap logger
	zapLogger, err := cfg.Build()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer zapLogger.Sync()

	// Create a Fiber logger middleware using Zap logger
	fiberLogger := logger.New(logger.Config{
		Format: "${time} ${status} - ${latency} ${method} ${path}\n",
		Output: zapcore.AddSync(os.Stdout), // Write logs to stdout
	})

	// Create instances of services
    userService := services.NewUserService(zapLogger)

    // Create instances of handlers
    userHandler := handlers.NewUserHandler(userService, zapLogger)

	// Initialize Fiber app
	app := fiber.New()

	// Register the logger middleware
	app.Use(fiberLogger)

	app.Post("/user", userHandler.AddUser)
	app.Get("/user/:id", userHandler.GetUser)
	app.Get("/users", userHandler.GetAllUsers)

	err = app.Listen(":3000")
	if err != nil {
		log.Fatalf("failed to start Fiber app: %v", err)
	}
}
