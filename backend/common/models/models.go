package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// AdminUser represents an authenticated administrator account
type AdminUser struct {
	ID           string    `gorm:"primaryKey;type:varchar(64)" json:"id"`
	Username     string    `gorm:"uniqueIndex;not null;type:varchar(64)" json:"username"`
	Email        string    `gorm:"uniqueIndex;not null;type:varchar(128)" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Role         string    `gorm:"default:'admin';type:varchar(32)" json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}

// LoginRequest represents payload for admin authentication
type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}

// LoginResponse represents response for admin authentication
type LoginResponse struct {
	Success bool       `json:"success"`
	Message string     `json:"message,omitempty"`
	Token   string     `json:"token,omitempty"`
	User    *AdminUser `json:"user,omitempty"`
	Error   string     `json:"error,omitempty"`
}

// JWTClaims defines payload stored inside signed JWT tokens
type JWTClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// ContactUs represents a consultation request from the website contact form
type ContactUs struct {
	ID        string    `gorm:"primaryKey;type:varchar(64)" json:"id"`
	Name      string    `gorm:"not null;type:varchar(128)" json:"name"`
	Email     string    `gorm:"not null;type:varchar(128)" json:"email"`
	Company   string    `gorm:"type:varchar(128)" json:"company,omitempty"`
	Subject   string    `gorm:"type:varchar(256)" json:"subject,omitempty"`
	Message   string    `gorm:"not null;type:text" json:"message"`
	Status    string    `gorm:"default:'unread';type:varchar(32)" json:"status"` // "unread", "read", "archived"
	CreatedAt time.Time `json:"created_at"`
}

// Project represents a portfolio case study or digital solution project
type Project struct {
	ID            string    `gorm:"primaryKey;type:varchar(64)" json:"id"`
	Title         string    `gorm:"not null;type:varchar(256)" json:"title"`
	Slug          string    `gorm:"uniqueIndex;type:varchar(256)" json:"slug"`
	Description   string    `gorm:"type:text" json:"description"`
	Category      string    `gorm:"type:varchar(128)" json:"category"`
	ClientName    string    `gorm:"type:varchar(128)" json:"client_name"`
	ImageURL      string    `gorm:"type:text" json:"image_url"`
	Tags          []string  `gorm:"serializer:json" json:"tags"`
	ImpactMetrics []string  `gorm:"serializer:json" json:"impact_metrics"`
	IsFeatured    bool      `gorm:"default:false" json:"is_featured"`
	CreatedAt     time.Time `json:"created_at"`
}

// Testimonial represents a client review or endorsement
type Testimonial struct {
	ID         string    `gorm:"primaryKey;type:varchar(64)" json:"id"`
	AuthorName string    `gorm:"not null;type:varchar(128)" json:"author_name"`
	Position   string    `gorm:"type:varchar(128)" json:"position"`
	Company    string    `gorm:"type:varchar(128)" json:"company"`
	Quote      string    `gorm:"not null;type:text" json:"quote"`
	Rating     int       `gorm:"default:5" json:"rating"` // 1 - 5
	AvatarURL  string    `gorm:"type:text" json:"avatar_url,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

// APIResponse standardizes JSON output for all endpoints
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
