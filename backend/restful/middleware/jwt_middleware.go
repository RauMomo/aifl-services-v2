package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"aifl-backend/common/models"
	"aifl-backend/proto/pb"
	"aifl-backend/restful/service"

	"google.golang.org/protobuf/proto"
)

type authContextKey string

const ClaimsContextKey authContextKey = "jwt_claims"

func JWTMiddleware(authService service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				respondAuthError(w, r, http.StatusUnauthorized, "Akses ditolak. Token autentikasi diperlukan.")
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := authService.ValidateToken(tokenString)
			if err != nil {
				respondAuthError(w, r, http.StatusUnauthorized, err.Error())
				return
			}

			ctx := context.WithValue(r.Context(), ClaimsContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
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
