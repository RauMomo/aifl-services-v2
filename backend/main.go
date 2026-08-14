package main

import (
	"net/http"
	"os"
	"time"

	"aifl-backend/db"
	"aifl-backend/logger"
	"aifl-backend/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize Database (GORM SQLite) & Seed Default Admin Account
	db.InitDB()

	// Initialize Redis Client & Rate Limiter
	db.InitRedis()

	// Register API Routes & Middleware
	handler := routes.RegisterRoutes()

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.Log.Info("🚀 AIFL Golang Enterprise Backend Server started",
		"port", port,
		"health", "http://localhost:"+port+"/health",
		"auth_login", "http://localhost:"+port+"/api/v1/auth/login",
	)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Log.Error("❌ Server failed to start", "error", err.Error())
		os.Exit(1)
	}
}
