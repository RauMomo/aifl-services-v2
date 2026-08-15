package service

import (
	"context"
	"errors"
	"os"
	"time"

	"aifl-backend/common/config"
	"aifl-backend/common/logger"
	"aifl-backend/common/models"
	"aifl-backend/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, usernameOrEmail, password, clientIP string) (*models.AdminUser, string, error)
	GetUserByID(ctx context.Context, id string) (*models.AdminUser, error)
	ValidateToken(tokenStr string) (*models.JWTClaims, error)
}

type authService struct {
	adminRepo repository.AdminRepository
	redisRepo repository.RedisRepository
	cfg       *config.Config
}

func NewAuthService(adminRepo repository.AdminRepository, redisRepo repository.RedisRepository, cfg *config.Config) AuthService {
	return &authService{
		adminRepo: adminRepo,
		redisRepo: redisRepo,
		cfg:       cfg,
	}
}

func (s *authService) getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if s.cfg != nil && s.cfg.JWT.Secret != "" {
		secret = s.cfg.JWT.Secret
	}
	if secret == "" {
		secret = "aifl-jwt-secret-key-production-change"
	}
	return []byte(secret)
}

func (s *authService) Login(ctx context.Context, usernameOrEmail, password, clientIP string) (*models.AdminUser, string, error) {
	rateKey := "ratelimit:login:" + clientIP
	allowed, err := s.redisRepo.RateLimit(ctx, rateKey, 5, 5*time.Minute)
	if err == nil && !allowed {
		logger.Log.Warn("Login rate limit exceeded", "ip", clientIP)
		return nil, "", errors.New("Terlalu banyak percobaan login. Silakan tunggu 5 menit.")
	}

	user, err := s.adminRepo.FindByUsernameOrEmail(usernameOrEmail)
	if err != nil {
		return nil, "", errors.New("Username/Email atau Password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, "", errors.New("Username/Email atau Password salah")
	}

	expHours := 24
	if s.cfg != nil && s.cfg.JWT.ExpirationHours > 0 {
		expHours = s.cfg.JWT.ExpirationHours
	}

	expirationTime := time.Now().Add(time.Duration(expHours) * time.Hour)
	claims := &models.JWTClaims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   user.ID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.getJWTSecret())
	if err != nil {
		logger.Log.Error("Failed to sign JWT token", "error", err.Error())
		return nil, "", errors.New("Gagal membuat sesi login")
	}

	logger.Log.Info("🔐 Admin login successful", "username", user.Username, "ip", clientIP)
	return user, tokenString, nil
}

func (s *authService) GetUserByID(ctx context.Context, id string) (*models.AdminUser, error) {
	return s.adminRepo.FindByID(id)
}

func (s *authService) ValidateToken(tokenStr string) (*models.JWTClaims, error) {
	claims := &models.JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return s.getJWTSecret(), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("Token autentikasi tidak valid atau sudah kedaluwarsa.")
	}
	return claims, nil
}
