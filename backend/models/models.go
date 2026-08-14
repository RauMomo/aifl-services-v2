package models

import "time"

// ContactUs represents a consultation request from the website contact form
type ContactUs struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Company   string    `json:"company,omitempty"`
	Subject   string    `json:"subject,omitempty"`
	Message   string    `json:"message"`
	Status    string    `json:"status"` // "unread", "read", "archived"
	CreatedAt time.Time `json:"created_at"`
}

// Project represents a portfolio case study or digital solution project
type Project struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Slug          string    `json:"slug"`
	Description   string    `json:"description"`
	Category      string    `json:"category"` // e.g., "Operational Control", "Course Platform", "Audit System"
	ClientName    string    `json:"client_name"`
	ImageURL      string    `json:"image_url"`
	Tags          []string  `json:"tags"`
	ImpactMetrics []string  `json:"impact_metrics"`
	IsFeatured    bool      `json:"is_featured"`
	CreatedAt     time.Time `json:"created_at"`
}

// Testimonial represents a client review or endorsement
type Testimonial struct {
	ID         string    `json:"id"`
	AuthorName string    `json:"author_name"`
	Position   string    `json:"position"`
	Company    string    `json:"company"`
	Quote      string    `json:"quote"`
	Rating     int       `json:"rating"` // 1 - 5
	AvatarURL  string    `json:"avatar_url,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

// APIResponse standardizes JSON output for all endpoints
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
