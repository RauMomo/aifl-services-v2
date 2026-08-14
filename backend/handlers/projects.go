package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"aifl-backend/db"
	"aifl-backend/logger"
	"aifl-backend/models"
	"aifl-backend/storage"
)

type ProjectHandler struct {
	ikService *storage.ImageKitService
}

func NewProjectHandler() *ProjectHandler {
	h := &ProjectHandler{
		ikService: storage.NewImageKitService(),
	}
	h.seedDefaultProjects()
	return h
}

func (h *ProjectHandler) seedDefaultProjects() {
	var count int64
	db.DB.Model(&models.Project{}).Count(&count)
	if count > 0 {
		return
	}

	samples := []models.Project{
		{
			ID:          "proj-1",
			Title:       "DriveMaster Driving Education System",
			Slug:        "drive-master-system",
			Description: "Platform edukasi dan manajemen kursus mengemudi terpusat dengan penugasan instruktur, jadwal presisi, dan tracking modul.",
			Category:    "Course & Training Platform",
			ClientName:  "DriveMaster Indonesia",
			ImageURL:    "/drive-master-mockup.png",
			Tags:        []string{"Ember.js", "Golang", "Operational Control"},
			ImpactMetrics: []string{
				"Efisiensi Operasional +85%",
				"Otomatisasi Penjadwalan 100%",
			},
			IsFeatured: true,
			CreatedAt:  time.Now(),
		},
		{
			ID:          "proj-2",
			Title:       "Enterprise Risk-Based Audit Control",
			Slug:        "risk-based-audit-control",
			Description: "Sistem pengawasan audit berbasis risiko perusahaan untuk mengontrol kepatuhan, temuan audit, dan pelaporan eksekutif secara real-time.",
			Category:    "Operational Control",
			ClientName:  "AIFL Enterprise Client",
			ImageURL:    "/aifl-software-mockup.png",
			Tags:        []string{"Process Discovery", "Workflow Automation", "Audit System"},
			ImpactMetrics: []string{
				"Zero Data Loss Audit Trail",
				"Kecepatan Laporan 5x Lebih Cepat",
			},
			IsFeatured: true,
			CreatedAt:  time.Now(),
		},
	}

	for _, p := range samples {
		db.DB.Create(&p)
	}
	logger.Log.Info("✅ Sample projects seeded into database")
}

// GetProjects handles GET /api/v1/projects
func (h *ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
	cacheKey := "cache:projects:list"
	var projects []models.Project

	// Check Redis Cache
	if db.CacheGet(r.Context(), cacheKey, &projects) {
		logger.Log.Debug("⚡ Serving projects list from Redis cache")
		respondJSON(w, http.StatusOK, models.APIResponse{
			Success: true,
			Message: fmt.Sprintf("Berhasil mengambil %d proyek (Cached)", len(projects)),
			Data:    projects,
		})
		return
	}

	if err := db.DB.Order("is_featured desc, created_at desc").Find(&projects).Error; err != nil {
		logger.Log.Error("Failed to fetch projects from database", "error", err.Error())
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Gagal mengambil data proyek",
		})
		return
	}

	// Cache result in Redis for 10 minutes
	_ = db.CacheSet(r.Context(), cacheKey, projects, 10*time.Minute)

	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: fmt.Sprintf("Berhasil mengambil %d proyek", len(projects)),
		Data:    projects,
	})
}

// CreateProject handles POST /api/v1/projects
func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var req models.Project
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Format JSON tidak valid",
		})
		return
	}

	req.Title = strings.TrimSpace(req.Title)
	if req.Title == "" {
		respondJSON(w, http.StatusUnprocessableEntity, models.APIResponse{
			Success: false,
			Error:   "Judul proyek wajib diisi",
		})
		return
	}

	req.ID = fmt.Sprintf("proj-%d", time.Now().UnixNano())
	if req.Slug == "" {
		req.Slug = strings.ToLower(strings.ReplaceAll(req.Title, " ", "-"))
	}
	req.CreatedAt = time.Now()

	if err := db.DB.Create(&req).Error; err != nil {
		logger.Log.Error("Failed to save project to database", "error", err.Error())
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Gagal menyimpan data proyek",
		})
		return
	}

	// Invalidate Redis cache
	db.CacheDelete(r.Context(), "cache:projects:list")

	logger.Log.Info("🚀 New project created", "id", req.ID, "title", req.Title)

	respondJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Proyek berhasil ditambahkan",
		Data:    req,
	})
}

// UploadImage handles POST /api/v1/projects/upload-image
func (h *ProjectHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondJSON(w, http.StatusMethodNotAllowed, models.APIResponse{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}

	// 10MB max upload memory
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		respondJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Gagal memproses file upload (Maksimal 10MB)",
		})
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		file, header, err = r.FormFile("file")
		if err != nil {
			respondJSON(w, http.StatusBadRequest, models.APIResponse{
				Success: false,
				Error:   "File gambar tidak ditemukan dalam form (Gunakan field 'image' atau 'file')",
			})
			return
		}
	}
	defer file.Close()

	// If ImageKit is configured, upload to cloud
	if h.ikService.IsConfigured() {
		result, err := h.ikService.UploadFile(r.Context(), file, header.Filename, "/projects")
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, models.APIResponse{
				Success: false,
				Error:   "Gagal mengunggah gambar ke ImageKit: " + err.Error(),
			})
			return
		}

		respondJSON(w, http.StatusCreated, models.APIResponse{
			Success: true,
			Message: "Gambar proyek berhasil diunggah ke ImageKit",
			Data:    result,
		})
		return
	}

	// Fallback Local Disk Storage if ImageKit env variables are not set
	uploadDir := "./public/uploads"
	_ = os.MkdirAll(uploadDir, 0755)

	ext := filepath.Ext(header.Filename)
	uniqueFileName := fmt.Sprintf("proj-img-%d%s", time.Now().UnixNano(), ext)
	dstPath := filepath.Join(uploadDir, uniqueFileName)

	dst, err := os.Create(dstPath)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Gagal menyimpan file lokal",
		})
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Gagal menyalin data file",
		})
		return
	}

	localURL := fmt.Sprintf("/uploads/%s", uniqueFileName)
	logger.Log.Info("📁 File saved to local storage (ImageKit unconfigured)", "url", localURL)

	respondJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Gambar berhasil disimpan di penyimpanan lokal (ImageKit belum dikonfigurasi)",
		Data: storage.UploadResult{
			URL:    localURL,
			FileID: uniqueFileName,
			Name:   header.Filename,
		},
	})
}
