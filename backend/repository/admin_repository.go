package repository

import (
	"fmt"
	"time"

	"aifl-backend/common/config"
	"aifl-backend/common/logger"
	"aifl-backend/common/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminRepository interface {
	FindByUsernameOrEmail(identifier string) (*models.AdminUser, error)
	FindByID(id string) (*models.AdminUser, error)
	SeedDefaultAdmin(cfg *config.Config) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) FindByUsernameOrEmail(identifier string) (*models.AdminUser, error) {
	var user models.AdminUser
	err := r.db.Where("username = ? OR email = ?", identifier, identifier).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *adminRepository) FindByID(id string) (*models.AdminUser, error) {
	var user models.AdminUser
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *adminRepository) SeedDefaultAdmin(cfg *config.Config) error {
	var count int64
	r.db.Model(&models.AdminUser{}).Count(&count)
	if count > 0 {
		return nil
	}

	defaultUsername := "admin"
	defaultPassword := "admin123"
	defaultEmail := "admin@aiflservices.com"

	if cfg != nil {
		if cfg.Admin.InitialUsername != "" {
			defaultUsername = cfg.Admin.InitialUsername
		}
		if cfg.Admin.InitialPassword != "" {
			defaultPassword = cfg.Admin.InitialPassword
		}
		if cfg.Admin.InitialEmail != "" {
			defaultEmail = cfg.Admin.InitialEmail
		}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.Error("Failed to hash default admin password", "error", err.Error())
		return err
	}

	admin := models.AdminUser{
		ID:           fmt.Sprintf("usr-%d", time.Now().UnixNano()),
		Username:     defaultUsername,
		Email:        defaultEmail,
		PasswordHash: string(hash),
		Role:         "admin",
		CreatedAt:    time.Now(),
	}

	if err := r.db.Create(&admin).Error; err != nil {
		logger.Log.Error("Failed to seed default admin user", "error", err.Error())
		return err
	}

	logger.Log.Info("🎉 Default Admin user seeded successfully",
		"username", defaultUsername,
		"password", "[PROTECTED - Default: admin123]",
	)
	return nil
}
