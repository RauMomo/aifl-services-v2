# ==============================================================================
# Multi-stage Dockerfile for AIFL Enterprise Golang Backend
# Port: 8200
# Features: GORM SQLite, Redis Fallback, Admin JWT Auth, Protobuf, Slog Logger
# ==============================================================================

# ------------------------------------------------------------------------------
# Stage 1: Build Binary
# ------------------------------------------------------------------------------
FROM golang:alpine AS builder

# Set environment variables for CGO-free static compilation
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# Install git and ca-certificates
RUN apk add --no-cache git ca-certificates tzdata

# Copy dependency manifests first for Docker layer caching
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy full application source code
COPY . .

# Compile optimized static Go binary
RUN go build \
    -ldflags="-w -s" \
    -a -installsuffix cgo \
    -o server main.go

# ------------------------------------------------------------------------------
# Stage 2: Minimal Production Runtime
# ------------------------------------------------------------------------------
FROM alpine:latest

# Install runtime dependencies (TLS certs, timezone data, curl/wget for healthchecks)
RUN apk --no-cache add ca-certificates tzdata wget \
    && addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Create persistent database directory
RUN mkdir -p /app/data && chown -R appuser:appgroup /app

# Copy binary from builder stage
COPY --from=builder --chown=appuser:appgroup /app/server .

# Environment defaults
ENV PORT=8200 \
    DATABASE_PATH=/app/data/aifl.db \
    GO_ENV=production

# Switch to non-root user for security
USER appuser

# Expose backend service port
EXPOSE 8200

# Health check directive
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://localhost:8200/health || exit 1

# Run the compiled server
CMD ["./server"]
