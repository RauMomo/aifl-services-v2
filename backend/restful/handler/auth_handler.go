package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"aifl-backend/common/models"
	"aifl-backend/proto/pb"
	"aifl-backend/restful/middleware"
	"aifl-backend/restful/service"

	"google.golang.org/protobuf/proto"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondAuthError(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	clientIP := getClientIP(r)
	var req models.LoginRequest

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

	user, token, err := h.authService.Login(r.Context(), req.UsernameOrEmail, req.Password, clientIP)
	if err != nil {
		status := http.StatusUnauthorized
		if strings.Contains(err.Error(), "Terlalu banyak") {
			status = http.StatusTooManyRequests
		}
		respondAuthError(w, r, status, err.Error())
		return
	}

	respondAuthSuccess(w, r, "Login berhasil", token, user)
}

func (h *AuthHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondAuthError(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	claims, ok := r.Context().Value(middleware.ClaimsContextKey).(*models.JWTClaims)
	if !ok || claims == nil {
		respondAuthError(w, r, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user, err := h.authService.GetUserByID(r.Context(), claims.UserID)
	if err != nil {
		respondAuthError(w, r, http.StatusNotFound, "User not found")
		return
	}

	respondAuthSuccess(w, r, "User profile retrieved", "", user)
}

func respondAuthSuccess(w http.ResponseWriter, r *http.Request, message string, token string, user *models.AdminUser) {
	acceptHeader := r.Header.Get("Accept")

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
