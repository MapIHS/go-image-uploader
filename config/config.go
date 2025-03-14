package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config menyimpan semua konfigurasi aplikasi
type Config struct {
	Port            string
	ImageSizeLimit  int64 // dalam bytes
	UploadPath      string
	PostgresConnStr string
}

// LoadConfig memuat konfigurasi dari file .env
func LoadConfig() (*Config, error) {
	// Muat file .env jika ada
	godotenv.Load()

	// Nilai default
	cfg := &Config{
		Port:            "8080",
		ImageSizeLimit:  5 * 1024 * 1024, // 5MB default
		UploadPath:      "./public/images",
		PostgresConnStr: "postgresql://postgres:postgres@localhost:5432/imagedb?sslmode=disable",
	}

	// Override dengan nilai dari environment jika tersedia
	if port := os.Getenv("PORT"); port != "" {
		cfg.Port = port
	}

	if sizeLimit := os.Getenv("IMAGE_SIZE_LIMIT"); sizeLimit != "" {
		if size, err := strconv.ParseInt(sizeLimit, 10, 64); err == nil {
			cfg.ImageSizeLimit = size
		}
	}

	if uploadPath := os.Getenv("UPLOAD_PATH"); uploadPath != "" {
		cfg.UploadPath = uploadPath
	}

	if dbConn := os.Getenv("POSTGRES_CONN"); dbConn != "" {
		cfg.PostgresConnStr = dbConn
	}

	// Pastikan direktori upload ada
	if err := os.MkdirAll(cfg.UploadPath, os.ModePerm); err != nil {
		return nil, err
	}

	return cfg, nil
}
