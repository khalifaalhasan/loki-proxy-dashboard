package server

import (
	"proxy-dashboard/internal/database"
	"proxy-dashboard/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)


type Server struct {
	App *fiber.App
	DB *database.Queries
}

func NewServer(db *database.Queries, frontendURL string) *Server {
	// init fiber
	app := fiber.New(fiber.Config{
		AppName: "Dashboard API Proxy V1",
	})
		// global middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: frontendURL,
		AllowHeaders: "origin, Content-Type, Accept, Authorization",
	}))

	routes.Setup(app, db)
	return &Server{
		App: app,
		DB: db,
	}

}

func (s * Server) Start(addr string) error {
	return s.App.Listen(addr)
}