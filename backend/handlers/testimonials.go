package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"aifl-backend/db"
	"aifl-backend/logger"
	"aifl-backend/models"
)

type TestimonialHandler struct{}

func NewTestimonialHandler() *TestimonialHandler {
	h := &TestimonialHandler{}
	h.seedDefaultTestimonials()
	return h
}

func (h *TestimonialHandler) seedDefaultTestimonials() {
	var count int64
	db.DB.Model(&models.Testimonial{}).Count(&count)
	if count > 0 {
		return
	}

	samples := []models.Testimonial{
		{
			ID:         "testi-1",
			AuthorName: "Hendry Prasetya",
			Position:   "Head of Operations",
			Company:    "PT Nusantara Synergy",
			Quote:      "Solusi digital AIFL mengubah alur kerja manual kami menjadi terotomatisasi secara presisi dan efisien.",
			Rating:     5,
			CreatedAt:  time.Now(),
		},
		{
			ID:         "testi-2",
			AuthorName: "Budi Santoso",
			Position:   "General Manager",
			Company:    "PT Corporate Indonesia",
			Quote:      "Dashboard monitoring dan aplikasi kustom AIFL memberikan kontrol penuh atas visibilitas tim kami.",
			Rating:     5,
			CreatedAt:  time.Now(),
		},
	}

	for _, t := range samples {
		db.DB.Create(&t)
	}
	logger.Log.Info("✅ Sample testimonials seeded into database")
}

// GetTestimonials handles GET /api/v1/testimonials
func (h *TestimonialHandler) GetTestimonials(w http.ResponseWriter, r *http.Request) {
	cacheKey := "cache:testimonials:list"
	var testimonials []models.Testimonial

	// Check Redis Cache
	if db.CacheGet(r.Context(), cacheKey, &testimonials) {
		logger.Log.Debug("⚡ Serving testimonials list from Redis cache")
		respondJSON(w, http.StatusOK, models.APIResponse{
			Success: true,
			Message: fmt.Sprintf("Berhasil mengambil %d testimoni (Cached)", len(testimonials)),
			Data:    testimonials,
		})
		return
	}

	if err := db.DB.Order("created_at desc").Find(&testimonials).Error; err != nil {
		logger.Log.Error("Failed to fetch testimonials from database", "error", err.Error())
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Gagal mengambil data testimoni",
		})
		return
	}

	// Cache result in Redis for 10 minutes
	_ = db.CacheSet(r.Context(), cacheKey, testimonials, 10*time.Minute)

	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: fmt.Sprintf("Berhasil mengambil %d testimoni", len(testimonials)),
		Data:    testimonials,
	})
}

// CreateTestimonial handles POST /api/v1/testimonials
func (h *TestimonialHandler) CreateTestimonial(w http.ResponseWriter, r *http.Request) {
	var req models.Testimonial
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Format JSON tidak valid",
		})
		return
	}

	req.AuthorName = strings.TrimSpace(req.AuthorName)
	req.Quote = strings.TrimSpace(req.Quote)
	if req.AuthorName == "" || req.Quote == "" {
		respondJSON(w, http.StatusUnprocessableEntity, models.APIResponse{
			Success: false,
			Error:   "Nama penulis dan kutipan testimoni wajib diisi",
		})
		return
	}

	if req.Rating < 1 || req.Rating > 5 {
		req.Rating = 5
	}

	req.ID = fmt.Sprintf("testi-%d", time.Now().UnixNano())
	req.CreatedAt = time.Now()

	if err := db.DB.Create(&req).Error; err != nil {
		logger.Log.Error("Failed to save testimonial to database", "error", err.Error())
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Gagal menyimpan data testimoni",
		})
		return
	}

	// Invalidate Redis cache
	db.CacheDelete(r.Context(), "cache:testimonials:list")

	logger.Log.Info("⭐ New testimonial created", "id", req.ID, "author", req.AuthorName)

	respondJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Testimoni berhasil ditambahkan",
		Data:    req,
	})
}
