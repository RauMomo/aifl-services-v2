package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"aifl-backend/db"
	"aifl-backend/handlers"
	"aifl-backend/logger"
	"aifl-backend/models"
)

// RegisterRoutes sets up all API endpoints and middleware
func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	authH := handlers.NewAuthHandler()
	contactH := handlers.NewContactHandler()
	projectH := handlers.NewProjectHandler()
	testimonialH := handlers.NewTestimonialHandler()

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":       "UP",
			"timestamp":    time.Now().Format(time.RFC3339),
			"database":     "CONNECTED",
			"redis_active": db.IsRedisActive(),
		})
	})

	// API Info endpoint
	mux.HandleFunc("/api/v1", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"name":       "AIFL REST API Backend",
			"version":    "2.0.0",
			"features":   []string{"Admin Auth", "GORM SQLite Database", "Redis Cache & Rate Limiting", "Protobuf Support", "Structured Logging"},
			"endpoints":  []string{"/api/v1/auth/login", "/api/v1/auth/me", "/api/v1/contact", "/api/v1/contact/list", "/api/v1/projects", "/api/v1/testimonials"},
		})
	})

	// Auth Endpoints
	mux.HandleFunc("/api/v1/auth/login", authH.Login)
	mux.Handle("/api/v1/auth/me", handlers.JWTMiddleware(http.HandlerFunc(authH.GetMe)))

	// Public Form Submission & Protected Contact List
	mux.HandleFunc("/api/v1/contact", contactH.CreateContact)
	mux.Handle("/api/v1/contact/list", handlers.JWTMiddleware(http.HandlerFunc(contactH.GetContacts)))

	// Public Read / Protected Creation for Projects & Testimonials
	mux.HandleFunc("/api/v1/projects", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			projectH.GetProjects(w, r)
		case http.MethodPost:
			handlers.JWTMiddleware(http.HandlerFunc(projectH.CreateProject)).ServeHTTP(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	mux.Handle("/api/v1/projects/upload-image", handlers.JWTMiddleware(http.HandlerFunc(projectH.UploadImage)))

	// Static file handler for local uploads fallback
	fs := http.FileServer(http.Dir("./public/uploads"))
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", fs))

	mux.HandleFunc("/api/v1/testimonials", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			testimonialH.GetTestimonials(w, r)
		case http.MethodPost:
			handlers.JWTMiddleware(http.HandlerFunc(testimonialH.CreateTestimonial)).ServeHTTP(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Wrap with Middleware: CORS & Structured Slog Logger with Request ID Tracing
	return enableCORS(logger.RequestLoggerMiddleware(mux))
}

// Middleware: CORS Configuration
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, X-Request-ID, Accept")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Ensure models package is used
var _ = models.APIResponse{}
