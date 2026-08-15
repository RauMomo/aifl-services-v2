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

type ContactService interface {
	CreateContact(ctx context.Context, req *models.ContactUs) (*models.ContactUs, error)
	GetAllContacts(ctx context.Context) ([]models.ContactUs, error)
}

type contactService struct {
	contactRepo repository.ContactRepository
}

func NewContactService(contactRepo repository.ContactRepository) ContactService {
	return &contactService{contactRepo: contactRepo}
}

func (s *contactService) CreateContact(ctx context.Context, req *models.ContactUs) (*models.ContactUs, error) {
	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)
	req.Message = strings.TrimSpace(req.Message)

	if req.Name == "" || req.Email == "" || req.Message == "" {
		return nil, errors.New("Nama, email, dan pesan wajib diisi")
	}

	req.ID = fmt.Sprintf("ct-%d", time.Now().UnixNano())
	req.Status = "unread"
	req.CreatedAt = time.Now()

	if err := s.contactRepo.Create(req); err != nil {
		logger.Log.Error("Failed to insert contact in database", "error", err.Error())
		return nil, errors.New("Gagal menyimpan data pesan")
	}

	logger.Log.Info("📩 New contact submission created", "id", req.ID, "name", req.Name, "email", req.Email)
	return req, nil
}

func (s *contactService) GetAllContacts(ctx context.Context) ([]models.ContactUs, error) {
	contacts, err := s.contactRepo.FindAll()
	if err != nil {
		logger.Log.Error("Failed to query contacts from database", "error", err.Error())
		return nil, errors.New("Gagal mengambil data kontak")
	}
	return contacts, nil
}
