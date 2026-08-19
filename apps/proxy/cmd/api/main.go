package main

import (
	"fmt"
	"log"
	"proxy-dashboard/internal/config"
	"proxy-dashboard/internal/database"
	"proxy-dashboard/internal/server"
)

func main () {
	// load config
	config.LoadEnv()

	// init DB
	dbConn := config.InitDB()
	defer dbConn.Close()
	queries := database.New(dbConn)

	// server build
	frontendURL := config.GetEnv("FRONTEND_URL", "*")
	srv := server.NewServer(queries, frontendURL)

	// turnOn Server
	port := config.GetEnv("PORT", "3000")
	serverAddr := fmt. Sprintf(":%s", port)

	log.Printf("API Proxy Running on localhost:%s\n", port)

	if err := srv.Start(serverAddr); err != nil {
		log.Fatalf("Server Crashed: %v", err)
	}
}
