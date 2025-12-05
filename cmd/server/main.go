package main

import (
	"car-rental-backend/config"
	"car-rental-backend/pkg/middleware"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
	cfg := config.LoadConfig()
	app, err := InitializeApp(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	fiberApp := fiber.New(fiber.Config{
		AppName:      "Car Rental Backend",
		ErrorHandler: handleAppError,
	})

	fiberApp.Use(recover.New())
	fiberApp.Use(logger.New())

	fiberApp.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	api := fiberApp.Group("/api/v1")
	app.AuthHandler.RegisterRoutes(api)

	api.Use(middleware.JWTMiddleware(app.JWTManager))
	app.CarHandler.RegisterRoutes(api)
	app.RenterHandler.RegisterRoutes(api)
	app.RentalTransactionHandler.RegisterRoutes(api)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Gracefully shutting down...")
		_ = fiberApp.Shutdown()
	}()

	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server starting on %s", addr)
	if err := fiberApp.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Println("Server stopped")
}

func handleAppError(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"message": message,
		"error":   fmt.Errorf("%+v", err),
	})
}
