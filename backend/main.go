package main

import (
	"net/http"
	"os"

	"aifl-backend/config"
	"aifl-backend/db"
	"aifl-backend/logger"
	"aifl-backend/routes"
)

func main() {
	// Load configuration using SPF13 Viper (config.yaml + Environment variables)
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Log.Warn("Failed to load Viper configuration", "error", err.Error())
	}

	port := cfg.Server.Port
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
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
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	logger.Log.Info("🚀 AIFL Golang Enterprise Backend Server started",
		"port", port,
		"env", cfg.Server.Env,
		"config", "config.yaml",
		"health", "http://localhost:"+port+"/health",
		"auth_login", "http://localhost:"+port+"/api/v1/auth/login",
	)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Log.Error("❌ Server failed to start", "error", err.Error())
		os.Exit(1)
	}
}
