package logger

import (
	"crypto/rand"
	"encoding/hex"
	"log/slog"
	"os"
)

type contextKey string

const RequestIDKey contextKey = "request_id"

var Log *slog.Logger

func init() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	Log = slog.New(handler)
	slog.SetDefault(Log)
}

// GenerateRequestID returns a unique hexadecimal string
func GenerateRequestID() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return hex.EncodeToString(b)[:12]
	}
	return hex.EncodeToString(b)
}
