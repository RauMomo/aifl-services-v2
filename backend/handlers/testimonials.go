package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"aifl-backend/models"
)

type TestimonialHandler struct {
	mu           sync.RWMutex
	testimonials map[string]models.Testimonial
}

func NewTestimonialHandler() *TestimonialHandler {
	h := &TestimonialHandler{
		testimonials: make(map[string]models.Testimonial),
	}

	// Seed realistic initial testimonials
	t1 := models.Testimonial{
		ID:         "testi-01",
		AuthorName: "Hendry Prasetya",
		Position:   "Head of Operations",
		Company:    "PT Nusantara Synergy",
		Quote:      "Solusi digital dari AIFL mengubah cara kami bekerja secara fundamental. Sistem penjadwalan dan tracking data kini berjalan otomatis tanpa error manual.",
		Rating:     5,
		AvatarURL:  "https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&w=150&q=80",
		CreatedAt:  time.Now().Add(-45 * 24 * time.Hour),
	}

	t2 := models.Testimonial{
		ID:         "testi-02",
		AuthorName: "Siti Rahmawati",
		Position:   "Managing Director",
		Company:    "Drive Master Academy",
		Quote:      "Platform kustom AIFL sangat intuitif dan mudah digunakan oleh seluruh tim. Registrasi peserta meningkat drastis dan administrasi jauh lebih efisien.",
		Rating:     5,
		AvatarURL:  "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?auto=format&fit=crop&w=150&q=80",
		CreatedAt:  time.Now().Add(-20 * 24 * time.Hour),
	}

	t3 := models.Testimonial{
		ID:         "testi-03",
		AuthorName: "Bambang Wijaya",
		Position:   "Internal Audit Manager",
		Company:    "Global Manufacturing Corp",
		Quote:      "Proses audit yang sebelumnya rumit kini dapat terpantau secara real-time. Transparansi dan akurasi data kami meningkat signifikan.",
		Rating:     5,
		AvatarURL:  "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?auto=format&fit=crop&w=150&q=80",
		CreatedAt:  time.Now().Add(-10 * 24 * time.Hour),
	}

	h.testimonials[t1.ID] = t1
	h.testimonials[t2.ID] = t2
	h.testimonials[t3.ID] = t3

	return h
}

// GetTestimonials handles GET /api/v1/testimonials
func (h *TestimonialHandler) GetTestimonials(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondJSON(w, http.StatusMethodNotAllowed, models.APIResponse{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}

	h.mu.RLock()
	defer h.mu.RUnlock()

	list := make([]models.Testimonial, 0, len(h.testimonials))
	for _, t := range h.testimonials {
		list = append(list, t)
	}

	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    list,
	})
}

// CreateTestimonial handles POST /api/v1/testimonials
func (h *TestimonialHandler) CreateTestimonial(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondJSON(w, http.StatusMethodNotAllowed, models.APIResponse{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}

	var req models.Testimonial
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Invalid JSON payload",
		})
		return
	}

	if strings.TrimSpace(req.AuthorName) == "" || strings.TrimSpace(req.Quote) == "" {
		respondJSON(w, http.StatusUnprocessableEntity, models.APIResponse{
			Success: false,
			Error:   "Author name and quote are required",
		})
		return
	}

	if req.Rating <= 0 || req.Rating > 5 {
		req.Rating = 5
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	id := fmt.Sprintf("testi-%d", time.Now().UnixNano())
	req.ID = id
	req.CreatedAt = time.Now()

	h.testimonials[id] = req

	respondJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Testimonial added successfully",
		Data:    req,
	})
}
