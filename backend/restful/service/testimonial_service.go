package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"aifl-backend/common/logger"
	"aifl-backend/common/models"
	"aifl-backend/repository"
)

type TestimonialService interface {
	GetTestimonials(ctx context.Context) ([]models.Testimonial, bool, error)
	CreateTestimonial(ctx context.Context, req *models.Testimonial) (*models.Testimonial, error)
}

type testimonialService struct {
	testimonialRepo repository.TestimonialRepository
	redisRepo       repository.RedisRepository
}

func NewTestimonialService(testimonialRepo repository.TestimonialRepository, redisRepo repository.RedisRepository) TestimonialService {
	return &testimonialService{
		testimonialRepo: testimonialRepo,
		redisRepo:       redisRepo,
	}
}

func (s *testimonialService) GetTestimonials(ctx context.Context) ([]models.Testimonial, bool, error) {
	cacheKey := "cache:testimonials:list"
	var testimonials []models.Testimonial

	if s.redisRepo.CacheGet(ctx, cacheKey, &testimonials) {
		logger.Log.Debug("⚡ Serving testimonials list from Redis cache")
		return testimonials, true, nil
	}

	testimonials, err := s.testimonialRepo.FindAll()
	if err != nil {
		logger.Log.Error("Failed to fetch testimonials from database", "error", err.Error())
		return nil, false, errors.New("Gagal mengambil data testimoni")
	}

	_ = s.redisRepo.CacheSet(ctx, cacheKey, testimonials, 10*time.Minute)
	return testimonials, false, nil
}

func (s *testimonialService) CreateTestimonial(ctx context.Context, req *models.Testimonial) (*models.Testimonial, error) {
	req.AuthorName = strings.TrimSpace(req.AuthorName)
	req.Quote = strings.TrimSpace(req.Quote)
	if req.AuthorName == "" || req.Quote == "" {
		return nil, errors.New("Nama penulis dan kutipan testimoni wajib diisi")
	}

	if req.Rating < 1 || req.Rating > 5 {
		req.Rating = 5
	}

	req.ID = fmt.Sprintf("testi-%d", time.Now().UnixNano())
	req.CreatedAt = time.Now()

	if err := s.testimonialRepo.Create(req); err != nil {
		logger.Log.Error("Failed to save testimonial to database", "error", err.Error())
		return nil, errors.New("Gagal menyimpan data testimoni")
	}

	s.redisRepo.CacheDelete(ctx, "cache:testimonials:list")
	logger.Log.Info("⭐ New testimonial created", "id", req.ID, "author", req.AuthorName)

	return req, nil
}
