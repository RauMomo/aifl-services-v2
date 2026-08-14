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

type ContactHandler struct {
	mu       sync.RWMutex
	contacts map[string]models.ContactUs
}

func NewContactHandler() *ContactHandler {
	h := &ContactHandler{
		contacts: make(map[string]models.ContactUs),
	}
	
	// Seed sample initial data
	seedID := "ct-001"
	h.contacts[seedID] = models.ContactUs{
		ID:        seedID,
		Name:      "Budi Santoso",
		Email:     "budi@example.com",
		Company:   "PT Maju Sejahtera",
		Subject:   "Diskusi Digitalisasi Operasional",
		Message:   "Kami tertarik untuk mengotomatisasi sistem pendaftaran dan audit internal perusahaan.",
		Status:    "unread",
		CreatedAt: time.Now().Add(-24 * time.Hour),
	}
	
	return h
}

// CreateContact handles POST /api/v1/contact
func (h *ContactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondJSON(w, http.StatusMethodNotAllowed, models.APIResponse{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}

	var req struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Company string `json:"company"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Invalid JSON request body",
		})
		return
	}

	if strings.TrimSpace(req.Name) == "" || strings.TrimSpace(req.Email) == "" || strings.TrimSpace(req.Message) == "" {
		respondJSON(w, http.StatusUnprocessableEntity, models.APIResponse{
			Success: false,
			Error:   "Name, email, and message are required fields",
		})
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	id := fmt.Sprintf("ct-%d", time.Now().UnixNano())
	contact := models.ContactUs{
		ID:        id,
		Name:      req.Name,
		Email:     req.Email,
		Company:   req.Company,
		Subject:   req.Subject,
		Message:   req.Message,
		Status:    "unread",
		CreatedAt: time.Now(),
	}

	h.contacts[id] = contact

	respondJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Pesan Anda telah berhasil dikirim. Tim AIFL akan menghubungi Anda segera.",
		Data:    contact,
	})
}

// GetContacts handles GET /api/v1/contact
func (h *ContactHandler) GetContacts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondJSON(w, http.StatusMethodNotAllowed, models.APIResponse{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}

	h.mu.RLock()
	defer h.mu.RUnlock()

	list := make([]models.ContactUs, 0, len(h.contacts))
	for _, c := range h.contacts {
		list = append(list, c)
	}

	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    list,
	})
}

// Helper function to send JSON response
func respondJSON(w http.ResponseWriter, statusCode int, data models.APIResponse) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}
