package repository

import (
	"time"

	"aifl-backend/common/logger"
	"aifl-backend/common/models"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	Create(project *models.Project) error
	FindAll() ([]models.Project, error)
	SeedDefaultProjects() error
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	r := &projectRepository{db: db}
	_ = r.SeedDefaultProjects()
	return r
}

func (r *projectRepository) Create(project *models.Project) error {
	return r.db.Create(project).Error
}

func (r *projectRepository) FindAll() ([]models.Project, error) {
	var projects []models.Project
	err := r.db.Order("is_featured desc, created_at desc").Find(&projects).Error
	return projects, err
}

func (r *projectRepository) SeedDefaultProjects() error {
	var count int64
	r.db.Model(&models.Project{}).Count(&count)
	if count > 0 {
		return nil
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
		r.db.Create(&p)
	}
	logger.Log.Info("✅ Sample projects seeded into database")
	return nil
}
