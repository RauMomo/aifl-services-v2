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

type ProjectHandler struct {
	mu       sync.RWMutex
	projects map[string]models.Project
}

func NewProjectHandler() *ProjectHandler {
	h := &ProjectHandler{
		projects: make(map[string]models.Project),
	}

	// Seed real project case studies from AIFL platform
	p1 := models.Project{
		ID:          "proj-01",
		Title:       "Enterprise Operational Control System",
		Slug:        "enterprise-operational-control",
		Description: "Sistem digital terintegrasi untuk mengelola data peserta, jadwal otomatis, dan operasional harian.",
		Category:    "Operational Control",
		ClientName:  "Lembaga Pelatihan Nasional",
		ImageURL:    "/aifl-software-mockup.png",
		Tags:        []string{"Automated Scheduling", "Real-time Analytics", "Workflow Management"},
		ImpactMetrics: []string{
			"Efisiensi Administrasi +75%",
			"Otomatisasi Penjadwalan 100%",
			"Transparansi Data Real-time",
		},
		IsFeatured: true,
		CreatedAt:  time.Now().Add(-60 * 24 * time.Hour),
	}

	p2 := models.Project{
		ID:          "proj-02",
		Title:       "Drive Master Course Platform",
		Slug:        "drive-master-course-platform",
		Description: "Platform pembelajaran dan manajemen kursus mengemudi dengan tracking progress dan registrasi terpusat.",
		Category:    "Course Platform",
		ClientName:  "Drive Master Academy",
		ImageURL:    "/drive-master-mockup.png",
		Tags:        []string{"Student Portal", "Payment Integration", "Digital Certificates"},
		ImpactMetrics: []string{
			"Pertumbuhan Siswa Baru 3x Lipat",
			"Proses Registrasi Kurang dari 2 Menit",
			"Akses Dashboard 24/7",
		},
		IsFeatured: true,
		CreatedAt:  time.Now().Add(-30 * 24 * time.Hour),
	}

	p3 := models.Project{
		ID:          "proj-03",
		Title:       "Internal Audit & Governance Portal",
		Slug:        "internal-audit-governance-portal",
		Description: "Sistem audit digital terpusat untuk monitoring aktivitas audit, tracking progress, dan dokumentasi risiko.",
		Category:    "Audit System",
		ClientName:  "Perusahaan Manufaktur Multinasional",
		ImageURL:    "/aifl-software-mockup.png",
		Tags:        []string{"Risk Management", "Audit Logging", "Compliance Tracking"},
		ImpactMetrics: []string{
			"Waktu Audit Berkurang 50%",
			"Dokumentasi 100% Terverifikasi",
		},
		IsFeatured: false,
		CreatedAt:  time.Now().Add(-15 * 24 * time.Hour),
	}

	h.projects[p1.ID] = p1
	h.projects[p2.ID] = p2
	h.projects[p3.ID] = p3

	return h
}

// GetProjects handles GET /api/v1/projects
func (h *ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondJSON(w, http.StatusMethodNotAllowed, models.APIResponse{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}

	h.mu.RLock()
	defer h.mu.RUnlock()

	list := make([]models.Project, 0, len(h.projects))
	for _, p := range h.projects {
		list = append(list, p)
	}

	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    list,
	})
}

// CreateProject handles POST /api/v1/projects
func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondJSON(w, http.StatusMethodNotAllowed, models.APIResponse{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}

	var req models.Project
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Invalid JSON payload",
		})
		return
	}

	if strings.TrimSpace(req.Title) == "" || strings.TrimSpace(req.Description) == "" {
		respondJSON(w, http.StatusUnprocessableEntity, models.APIResponse{
			Success: false,
			Error:   "Title and description are required",
		})
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	id := fmt.Sprintf("proj-%d", time.Now().UnixNano())
	req.ID = id
	if req.Slug == "" {
		req.Slug = strings.ToLower(strings.ReplaceAll(req.Title, " ", "-"))
	}
	req.CreatedAt = time.Now()

	h.projects[id] = req

	respondJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Project created successfully",
		Data:    req,
	})
}
