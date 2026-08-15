package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"aifl-backend/common/logger"
	"aifl-backend/common/models"
	"aifl-backend/repository"
	"aifl-backend/restful/client"
)

type ProjectService interface {
	GetProjects(ctx context.Context) ([]models.Project, bool, error)
	CreateProject(ctx context.Context, req *models.Project) (*models.Project, error)
	UploadImage(ctx context.Context, file multipart.File, header *multipart.FileHeader) (*client.UploadResult, error)
}

type projectService struct {
	projectRepo repository.ProjectRepository
	redisRepo   repository.RedisRepository
	ikClient    client.ImageKitClient
}

func NewProjectService(projectRepo repository.ProjectRepository, redisRepo repository.RedisRepository, ikClient client.ImageKitClient) ProjectService {
	return &projectService{
		projectRepo: projectRepo,
		redisRepo:   redisRepo,
		ikClient:    ikClient,
	}
}

func (s *projectService) GetProjects(ctx context.Context) ([]models.Project, bool, error) {
	cacheKey := "cache:projects:list"
	var projects []models.Project

	if s.redisRepo.CacheGet(ctx, cacheKey, &projects) {
		logger.Log.Debug("⚡ Serving projects list from Redis cache")
		return projects, true, nil
	}

	projects, err := s.projectRepo.FindAll()
	if err != nil {
		logger.Log.Error("Failed to fetch projects from database", "error", err.Error())
		return nil, false, errors.New("Gagal mengambil data proyek")
	}

	_ = s.redisRepo.CacheSet(ctx, cacheKey, projects, 10*time.Minute)
	return projects, false, nil
}

func (s *projectService) CreateProject(ctx context.Context, req *models.Project) (*models.Project, error) {
	req.Title = strings.TrimSpace(req.Title)
	if req.Title == "" {
		return nil, errors.New("Judul proyek wajib diisi")
	}

	req.ID = fmt.Sprintf("proj-%d", time.Now().UnixNano())
	if req.Slug == "" {
		req.Slug = strings.ToLower(strings.ReplaceAll(req.Title, " ", "-"))
	}
	req.CreatedAt = time.Now()

	if err := s.projectRepo.Create(req); err != nil {
		logger.Log.Error("Failed to save project to database", "error", err.Error())
		return nil, errors.New("Gagal menyimpan data proyek")
	}

	s.redisRepo.CacheDelete(ctx, "cache:projects:list")
	logger.Log.Info("🚀 New project created", "id", req.ID, "title", req.Title)

	return req, nil
}

func (s *projectService) UploadImage(ctx context.Context, file multipart.File, header *multipart.FileHeader) (*client.UploadResult, error) {
	if s.ikClient.IsConfigured() {
		result, err := s.ikClient.UploadFile(ctx, file, header.Filename, "/projects")
		if err != nil {
			return nil, fmt.Errorf("Gagal mengunggah gambar ke ImageKit: %w", err)
		}
		return result, nil
	}

	uploadDir := "./public/uploads"
	_ = os.MkdirAll(uploadDir, 0755)

	ext := filepath.Ext(header.Filename)
	uniqueFileName := fmt.Sprintf("proj-img-%d%s", time.Now().UnixNano(), ext)
	dstPath := filepath.Join(uploadDir, uniqueFileName)

	dst, err := os.Create(dstPath)
	if err != nil {
		return nil, errors.New("Gagal menyimpan file lokal")
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return nil, errors.New("Gagal menyalin data file")
	}

	localURL := fmt.Sprintf("/uploads/%s", uniqueFileName)
	logger.Log.Info("📁 File saved to local storage (ImageKit unconfigured)", "url", localURL)

	return &client.UploadResult{
		URL:    localURL,
		FileID: uniqueFileName,
		Name:   header.Filename,
	}, nil
}
