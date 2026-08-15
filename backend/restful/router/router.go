package router

import (
	"encoding/json"
	"net/http"
	"time"

	"aifl-backend/repository"
	"aifl-backend/restful/handler"
	"aifl-backend/restful/middleware"
	"aifl-backend/restful/service"
)

type RouterDependencies struct {
	AuthHandler        *handler.AuthHandler
	ContactHandler     *handler.ContactHandler
	ProjectHandler     *handler.ProjectHandler
	TestimonialHandler *handler.TestimonialHandler
	AuthService        service.AuthService
	RedisRepo          repository.RedisRepository
}

func RegisterRoutes(deps *RouterDependencies) http.Handler {
	mux := http.NewServeMux()

	jwtMW := middleware.JWTMiddleware(deps.AuthService)

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":       "UP",
			"timestamp":    time.Now().Format(time.RFC3339),
			"database":     "CONNECTED",
			"redis_active": deps.RedisRepo.IsActive(),
		})
	})

	// API Info endpoint
	mux.HandleFunc("/api/v1", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"name":      "AIFL REST API Backend",
			"version":   "2.0.0",
			"features":  []string{"Admin Auth", "GORM SQLite Database", "Redis Cache & Rate Limiting", "Protobuf Support", "Structured Logging"},
			"endpoints": []string{"/api/v1/auth/login", "/api/v1/auth/me", "/api/v1/contact", "/api/v1/contact/list", "/api/v1/projects", "/api/v1/testimonials"},
		})
	})

	// Auth Endpoints
	mux.HandleFunc("/api/v1/auth/login", deps.AuthHandler.Login)
	mux.Handle("/api/v1/auth/me", jwtMW(http.HandlerFunc(deps.AuthHandler.GetMe)))

	// Public Form Submission & Protected Contact List
	mux.HandleFunc("/api/v1/contact", deps.ContactHandler.CreateContact)
	mux.Handle("/api/v1/contact/list", jwtMW(http.HandlerFunc(deps.ContactHandler.GetContacts)))

	// Public Read / Protected Creation for Projects & Testimonials
	mux.HandleFunc("/api/v1/projects", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			deps.ProjectHandler.GetProjects(w, r)
		case http.MethodPost:
			jwtMW(http.HandlerFunc(deps.ProjectHandler.CreateProject)).ServeHTTP(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	mux.Handle("/api/v1/projects/upload-image", jwtMW(http.HandlerFunc(deps.ProjectHandler.UploadImage)))

	// Static file handler for local uploads fallback
	fs := http.FileServer(http.Dir("./public/uploads"))
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", fs))

	mux.HandleFunc("/api/v1/testimonials", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			deps.TestimonialHandler.GetTestimonials(w, r)
		case http.MethodPost:
			jwtMW(http.HandlerFunc(deps.TestimonialHandler.CreateTestimonial)).ServeHTTP(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	return middleware.CORSMiddleware(middleware.LoggerMiddleware(mux))
}
