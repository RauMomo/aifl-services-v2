package repository

import (
	"time"

	"aifl-backend/common/logger"
	"aifl-backend/common/models"

	"gorm.io/gorm"
)

type TestimonialRepository interface {
	Create(testimonial *models.Testimonial) error
	FindAll() ([]models.Testimonial, error)
	SeedDefaultTestimonials() error
}

type testimonialRepository struct {
	db *gorm.DB
}

func NewTestimonialRepository(db *gorm.DB) TestimonialRepository {
	r := &testimonialRepository{db: db}
	_ = r.SeedDefaultTestimonials()
	return r
}

func (r *testimonialRepository) Create(testimonial *models.Testimonial) error {
	return r.db.Create(testimonial).Error
}

func (r *testimonialRepository) FindAll() ([]models.Testimonial, error) {
	var testimonials []models.Testimonial
	err := r.db.Order("created_at desc").Find(&testimonials).Error
	return testimonials, err
}

func (r *testimonialRepository) SeedDefaultTestimonials() error {
	var count int64
	r.db.Model(&models.Testimonial{}).Count(&count)
	if count > 0 {
		return nil
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
		r.db.Create(&t)
	}
	logger.Log.Info("✅ Sample testimonials seeded into database")
	return nil
}
