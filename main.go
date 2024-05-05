package main

import (
	"log"
	"os"
	"os/signal"
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
	vehicleService := services.NewVehicleService(zapLogger)
	rideService := services.NewRideService(zapLogger)

	// Create instances of handlers
	userHandler := handlers.NewUserHandler(userService, zapLogger)
	vehicleHandler := handlers.NewVehicleHandler(vehicleService, zapLogger)
	rideHandler := handlers.NewRideHandler(rideService, vehicleService, zapLogger)

	// Initialize Fiber app
	app := fiber.New()

	app.Use(fiberLogger)

	app.Post("/user", userHandler.AddUser)
	app.Get("/user/:id", userHandler.GetUser)
	app.Get("/users", userHandler.GetAllUsers)

	app.Post("/vehicle", vehicleHandler.AddVehicle)
	app.Get("/user/:id/vehicle", vehicleHandler.GetVehicle)

	app.Post("/ride", rideHandler.OfferRide)
	app.Get("/rides", rideHandler.GetAllRides)
	app.Post("/ride/select", rideHandler.SelectRide)
	app.Post("/ride/end", rideHandler.EndRide)

	// Start Fiber app in a separate goroutine
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("failed to start Fiber app: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	// Gracefully shutdown the server
	if err := app.Shutdown(); err != nil {
		log.Fatalf("graceful shutdown failed: %v", err)
	}
	log.Println("server gracefully shutdown")
}
