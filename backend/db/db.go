package db

import (
	"fmt"
	"os"
	"time"

	"aifl-backend/logger"
	"aifl-backend/models"

	"github.com/glebarez/sqlite"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes GORM database, auto-migrates tables, and seeds initial Admin
func InitDB() {
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "aifl.db"
	}

	gormConfig := &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Silent),
	}

	database, err := gorm.Open(sqlite.Open(dbPath), gormConfig)
	if err != nil {
		logger.Log.Error("❌ Failed to connect to SQLite database", "error", err.Error())
		panic(fmt.Sprintf("Failed to initialize database: %v", err))
	}

	DB = database
	logger.Log.Info("🟢 Database connected successfully", "engine", "SQLite", "file", dbPath)

	// Run Auto Migration
	if err := DB.AutoMigrate(
		&models.AdminUser{},
		&models.ContactUs{},
		&models.Project{},
		&models.Testimonial{},
	); err != nil {
		logger.Log.Error("❌ Database AutoMigrate failed", "error", err.Error())
		panic(err)
	}

	logger.Log.Info("✅ Database schema migrated successfully")

	// Seed default Admin User if table is empty
	SeedDefaultAdmin()
}

// SeedDefaultAdmin creates a default admin account if no admin exists
func SeedDefaultAdmin() {
	var count int64
	DB.Model(&models.AdminUser{}).Count(&count)
	if count > 0 {
		return
	}

	defaultUsername := os.Getenv("ADMIN_INITIAL_USERNAME")
	if defaultUsername == "" {
		defaultUsername = "admin"
	}

	defaultPassword := os.Getenv("ADMIN_INITIAL_PASSWORD")
	if defaultPassword == "" {
		defaultPassword = "admin123"
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.Error("Failed to hash default admin password", "error", err.Error())
		return
	}

	admin := models.AdminUser{
		ID:           fmt.Sprintf("usr-%d", time.Now().UnixNano()),
		Username:     defaultUsername,
		Email:        "admin@aiflservices.com",
		PasswordHash: string(hash),
		Role:         "admin",
		CreatedAt:    time.Now(),
	}

	if err := DB.Create(&admin).Error; err != nil {
		logger.Log.Error("Failed to seed default admin user", "error", err.Error())
		return
	}

	logger.Log.Info("🎉 Default Admin user seeded successfully",
		"username", defaultUsername,
		"password", "[PROTECTED - Default: admin123]",
	)
}
