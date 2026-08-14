package routes

import (
	"log"
	"net/http"
	"time"

	"aifl-backend/handlers"
	"aifl-backend/models"
)

// RegisterRoutes sets up all API endpoints and middleware
func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	contactH := handlers.NewContactHandler()
	projectH := handlers.NewProjectHandler()
	testimonialH := handlers.NewTestimonialHandler()

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"UP","timestamp":"` + time.Now().Format(time.RFC3339) + `"}`))
	})

	// API Info endpoint
	mux.HandleFunc("/api/v1", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"name":"AIFL REST API Backend","version":"1.0.0","entities":["contact","projects","testimonials"]}`))
	})

	// Entity Endpoints
	mux.HandleFunc("/api/v1/contact", contactH.CreateContact)
	mux.HandleFunc("/api/v1/contact/list", contactH.GetContacts)

	mux.HandleFunc("/api/v1/projects", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			projectH.GetProjects(w, r)
		case http.MethodPost:
			projectH.CreateProject(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/v1/testimonials", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			testimonialH.GetTestimonials(w, r)
		case http.MethodPost:
			testimonialH.CreateTestimonial(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Wrap with Middleware (CORS & Logging)
	return enableCORS(requestLogger(mux))
}

// Middleware: CORS Configuration
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Middleware: Request Logger
func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s %s - %v", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}

// Ensure models package is used
var _ = models.APIResponse{}
