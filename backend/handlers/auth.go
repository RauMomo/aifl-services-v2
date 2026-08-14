package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"aifl-backend/config"
	"aifl-backend/db"
	"aifl-backend/logger"
	"aifl-backend/models"
	"aifl-backend/proto/pb"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/proto"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if config.AppConfig != nil && config.AppConfig.JWT.Secret != "" {
		secret = config.AppConfig.JWT.Secret
	}
	if secret == "" {
		secret = "aifl-jwt-secret-key-production-change"
	}
	return []byte(secret)
}

type authContextKey string

const ClaimsContextKey authContextKey = "jwt_claims"

// Login handles POST /api/v1/auth/login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondAuthError(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Redis Rate Limiting (5 login attempts per 5 minutes per remote IP)
	clientIP := getClientIP(r)
	rateKey := "ratelimit:login:" + clientIP
	allowed, err := db.RateLimit(r.Context(), rateKey, 5, 5*time.Minute)
	if err == nil && !allowed {
		logger.Log.Warn("Login rate limit exceeded", "ip", clientIP)
		respondAuthError(w, r, http.StatusTooManyRequests, "Terlalu banyak percobaan login. Silakan tunggu 5 menit.")
		return
	}

	var req models.LoginRequest

	// Support both JSON & Protobuf payloads
	if strings.Contains(r.Header.Get("Content-Type"), "application/x-protobuf") {
		body, err := io.ReadAll(r.Body)
		if err == nil {
			var pbReq pb.AdminLoginRequestProto
			if proto.Unmarshal(body, &pbReq) == nil {
				req.UsernameOrEmail = pbReq.GetUsernameOrEmail()
				req.Password = pbReq.GetPassword()
			}
		}
	} else {
		_ = json.NewDecoder(r.Body).Decode(&req)
	}

	req.UsernameOrEmail = strings.TrimSpace(req.UsernameOrEmail)
	req.Password = strings.TrimSpace(req.Password)

	if req.UsernameOrEmail == "" || req.Password == "" {
		respondAuthError(w, r, http.StatusBadRequest, "Username/Email and Password are required")
		return
	}

	var user models.AdminUser
	err = db.DB.Where("username = ? OR email = ?", req.UsernameOrEmail, req.UsernameOrEmail).First(&user).Error
	if err != nil {
		respondAuthError(w, r, http.StatusUnauthorized, "Username/Email atau Password salah")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		respondAuthError(w, r, http.StatusUnauthorized, "Username/Email atau Password salah")
		return
	}

	expHours := 24
	if config.AppConfig != nil && config.AppConfig.JWT.ExpirationHours > 0 {
		expHours = config.AppConfig.JWT.ExpirationHours
	}

	// Generate JWT Token
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
	tokenString, err := token.SignedString(getJWTSecret())
	if err != nil {
		logger.Log.Error("Failed to sign JWT token", "error", err.Error())
		respondAuthError(w, r, http.StatusInternalServerError, "Gagal membuat sesi login")
		return
	}

	logger.Log.Info("🔐 Admin login successful", "username", user.Username, "ip", clientIP)

	respondAuthSuccess(w, r, "Login berhasil", tokenString, &user)
}

// GetMe handles GET /api/v1/auth/me (Protected Endpoint)
func (h *AuthHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondAuthError(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	claims, ok := r.Context().Value(ClaimsContextKey).(*models.JWTClaims)
	if !ok || claims == nil {
		respondAuthError(w, r, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var user models.AdminUser
	if err := db.DB.Where("id = ?", claims.UserID).First(&user).Error; err != nil {
		respondAuthError(w, r, http.StatusNotFound, "User not found")
		return
	}

	respondAuthSuccess(w, r, "User profile retrieved", "", &user)
}

// JWTMiddleware validates Bearer tokens for protected endpoints
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			respondAuthError(w, r, http.StatusUnauthorized, "Akses ditolak. Token autentikasi diperlukan.")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &models.JWTClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return getJWTSecret(), nil
		})

		if err != nil || !token.Valid {
			respondAuthError(w, r, http.StatusUnauthorized, "Token autentikasi tidak valid atau sudah kedaluwarsa.")
			return
		}

		ctx := context.WithValue(r.Context(), ClaimsContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func respondAuthSuccess(w http.ResponseWriter, r *http.Request, message string, token string, user *models.AdminUser) {
	acceptHeader := r.Header.Get("Accept")

	// Protobuf response support
	if strings.Contains(acceptHeader, "application/x-protobuf") {
		pbUser := &pb.AdminUserProto{
			Id:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
		}
		pbResp := &pb.AdminLoginResponseProto{
			Success: true,
			Message: message,
			Token:   token,
			User:    pbUser,
		}
		data, err := proto.Marshal(pbResp)
		if err == nil {
			w.Header().Set("Content-Type", "application/x-protobuf")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(data)
			return
		}
	}

	// JSON response fallback
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.LoginResponse{
		Success: true,
		Message: message,
		Token:   token,
		User:    user,
	})
}

func respondAuthError(w http.ResponseWriter, r *http.Request, status int, errMsg string) {
	acceptHeader := r.Header.Get("Accept")

	if strings.Contains(acceptHeader, "application/x-protobuf") {
		pbResp := &pb.AdminLoginResponseProto{
			Success: false,
			Error:   errMsg,
		}
		data, err := proto.Marshal(pbResp)
		if err == nil {
			w.Header().Set("Content-Type", "application/x-protobuf")
			w.WriteHeader(status)
			_, _ = w.Write(data)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(models.LoginResponse{
		Success: false,
		Error:   errMsg,
	})
}

func getClientIP(r *http.Request) string {
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		parts := strings.Split(xff, ",")
		return strings.TrimSpace(parts[0])
	}
	return r.RemoteAddr
}
