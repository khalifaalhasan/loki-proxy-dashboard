package routes

import (
	"proxy-dashboard/internal/database"
	"proxy-dashboard/internal/handlers"

	"github.com/gofiber/fiber/v2"
)


func Setup(app *fiber.App, db *database.Queries) {
	authhandler := handlers.NewAuthHandler(db)


	// endpoint prefix
	api := app.Group("/api")

	// Heatlh Check Endpoint
	api.Get("/health", func(c *fiber.Ctx) error{
		return c.Status(200).JSON(fiber.Map{
			"status": "UP",
			"message": "Fiber Backend is running",
		})
	})

	// registered endpoint
	v1 := api.Group("/v1")
	v1.Post("/register", authhandler.Register)
}