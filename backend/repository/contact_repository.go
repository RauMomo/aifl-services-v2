package repository

import (
	"aifl-backend/common/models"

	"gorm.io/gorm"
)

type ContactRepository interface {
	Create(contact *models.ContactUs) error
	FindAll() ([]models.ContactUs, error)
}

type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{db: db}
}

func (r *contactRepository) Create(contact *models.ContactUs) error {
	return r.db.Create(contact).Error
}

func (r *contactRepository) FindAll() ([]models.ContactUs, error) {
	var contacts []models.ContactUs
	err := r.db.Order("created_at desc").Find(&contacts).Error
	return contacts, err
}
