package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main () {
	app := fiber.New(fiber.Config{
		AppName: "Dashboard API Proxy V1",
	})

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "origin, Content-Type, Accept, Authorization",
	}))

	api := app.Group("/api")
	api.Get("/health", func(c *fiber.Ctx) error{
		return c.Status(200).JSON(fiber.Map{
			"status": "UP",
			"message": "Fiber Backend is running",
		})
	})

	log.Println("API Proxy Running on localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
