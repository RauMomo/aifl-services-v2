package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"aifl-backend/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handler := routes.RegisterRoutes()

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("🚀 AIFL Golang Backend Server running on port %s", port)
	log.Printf("📌 Health Check endpoint: http://localhost:%s/health", port)
	log.Printf("📌 Contact API: http://localhost:%s/api/v1/contact", port)
	log.Printf("📌 Projects API: http://localhost:%s/api/v1/projects", port)
	log.Printf("📌 Testimonials API: http://localhost:%s/api/v1/testimonials", port)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
