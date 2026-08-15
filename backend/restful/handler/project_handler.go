package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"aifl-backend/common/models"
	"aifl-backend/restful/service"
)

type ProjectHandler struct {
	projectService service.ProjectService
}

func NewProjectHandler(projectService service.ProjectService) *ProjectHandler {
	return &ProjectHandler{projectService: projectService}
}

func (h *ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
	projects, cached, err := h.projectService.GetProjects(r.Context())
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	msg := fmt.Sprintf("Berhasil mengambil %d proyek", len(projects))
	if cached {
		msg += " (Cached)"
	}

	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: msg,
		Data:    projects,
	})
}

func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var req models.Project
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Format JSON tidak valid",
		})
		return
	}

	created, err := h.projectService.CreateProject(r.Context(), &req)
	if err != nil {
		respondJSON(w, http.StatusUnprocessableEntity, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	respondJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Proyek berhasil ditambahkan",
		Data:    created,
	})
}

func (h *ProjectHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondJSON(w, http.StatusMethodNotAllowed, models.APIResponse{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}

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

	result, err := h.projectService.UploadImage(r.Context(), file, header)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	respondJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Gambar berhasil diunggah",
		Data:    result,
	})
}
