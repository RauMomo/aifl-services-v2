package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"aifl-backend/common/config"
	"aifl-backend/common/logger"

	"github.com/imagekit-developer/imagekit-go/v2"
	"github.com/imagekit-developer/imagekit-go/v2/option"
)

type ImageKitClient interface {
	IsConfigured() bool
	UploadFile(ctx context.Context, fileReader io.Reader, fileName string, folder string) (*UploadResult, error)
}

type imageKitClient struct {
	ik           *imagekit.Client
	isConfigured bool
	publicKey    string
	privateKey   string
	urlEndpoint  string
}

type UploadResult struct {
	URL          string `json:"url"`
	FileID       string `json:"file_id"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	Name         string `json:"name"`
}

func NewImageKitClient(cfg *config.Config) ImageKitClient {
	publicKey := os.Getenv("IMAGEKIT_PUBLIC_KEY")
	privateKey := os.Getenv("IMAGEKIT_PRIVATE_KEY")
	urlEndpoint := os.Getenv("IMAGEKIT_URL_ENDPOINT")

	if cfg != nil {
		if publicKey == "" {
			publicKey = cfg.ImageKit.PublicKey
		}
		if privateKey == "" {
			privateKey = cfg.ImageKit.PrivateKey
		}
		if urlEndpoint == "" {
			urlEndpoint = cfg.ImageKit.URLEndpoint
		}
	}

	if publicKey == "" || privateKey == "" || urlEndpoint == "" {
		logger.Log.Warn("⚠️ ImageKit credentials incomplete. Set IMAGEKIT_PUBLIC_KEY, IMAGEKIT_PRIVATE_KEY, and IMAGEKIT_URL_ENDPOINT in config.yaml or environment variables to enable cloud uploads.", "url_endpoint", urlEndpoint)
		return &imageKitClient{
			isConfigured: false,
			urlEndpoint:  urlEndpoint,
		}
	}

	opts := []option.RequestOption{
		option.WithPrivateKey(privateKey),
	}
	if urlEndpoint != "" {
		opts = append(opts, option.WithBaseURL(urlEndpoint))
	}

	ikClient := imagekit.NewClient(opts...)

	logger.Log.Info("🟢 ImageKit Client initialized successfully", "url_endpoint", urlEndpoint)

	return &imageKitClient{
		ik:           &ikClient,
		isConfigured: true,
		publicKey:    publicKey,
		privateKey:   privateKey,
		urlEndpoint:  urlEndpoint,
	}
}

func (s *imageKitClient) IsConfigured() bool {
	return s.isConfigured
}

func (s *imageKitClient) UploadFile(ctx context.Context, fileReader io.Reader, fileName string, folder string) (*UploadResult, error) {
	if !s.isConfigured {
		return nil, errors.New("ImageKit Client is not fully configured. Please set IMAGEKIT_PUBLIC_KEY and IMAGEKIT_PRIVATE_KEY in config.yaml or environment variables")
	}

	if folder == "" {
		folder = "/projects"
	}

	resp, err := s.ik.Files.Upload(ctx, imagekit.FileUploadParams{
		File:              fileReader,
		FileName:          fileName,
		Folder:            imagekit.String(folder),
		UseUniqueFileName: imagekit.Bool(true),
	})

	if err != nil {
		logger.Log.Error("Failed to upload file to ImageKit", "error", err.Error(), "file_name", fileName)
		return nil, fmt.Errorf("ImageKit upload error: %w", err)
	}

	logger.Log.Info("📸 Image uploaded to ImageKit successfully", "file_id", resp.FileID, "url", resp.URL)

	return &UploadResult{
		URL:          resp.URL,
		FileID:       resp.FileID,
		ThumbnailURL: resp.ThumbnailURL,
		Name:         resp.Name,
	}, nil
}
