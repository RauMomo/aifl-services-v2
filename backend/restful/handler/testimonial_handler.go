package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"aifl-backend/common/models"
	"aifl-backend/restful/service"
)

type TestimonialHandler struct {
	testimonialService service.TestimonialService
}

func NewTestimonialHandler(testimonialService service.TestimonialService) *TestimonialHandler {
	return &TestimonialHandler{testimonialService: testimonialService}
}

func (h *TestimonialHandler) GetTestimonials(w http.ResponseWriter, r *http.Request) {
	testimonials, cached, err := h.testimonialService.GetTestimonials(r.Context())
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	msg := fmt.Sprintf("Berhasil mengambil %d testimoni", len(testimonials))
	if cached {
		msg += " (Cached)"
	}

	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: msg,
		Data:    testimonials,
	})
}

func (h *TestimonialHandler) CreateTestimonial(w http.ResponseWriter, r *http.Request) {
	var req models.Testimonial
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Format JSON tidak valid",
		})
		return
	}

	created, err := h.testimonialService.CreateTestimonial(r.Context(), &req)
	if err != nil {
		respondJSON(w, http.StatusUnprocessableEntity, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	respondJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Testimoni berhasil ditambahkan",
		Data:    created,
	})
}
