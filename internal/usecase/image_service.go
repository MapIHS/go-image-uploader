package usecase

import (
	"mime/multipart"

	"github.com/MapIhs/go-image-uploader/internal/domain/models"
	"github.com/MapIhs/go-image-uploader/internal/infrastructure/database"
	"github.com/MapIhs/go-image-uploader/internal/infrastructure/storage"
)

// ImageService mendefinisikan layanan untuk operasi terkait gambar
type ImageService interface {
	UploadImage(file *multipart.FileHeader) (*models.Image, error)
	GetDiskUsage() (map[string]interface{}, error)
	GetImageList() ([]models.Image, error)
}

type imageService struct {
	storage  storage.ImageStorage
	database *database.Database
}

// NewImageService membuat instance baru dari ImageService
func NewImageService(storage storage.ImageStorage, db *database.Database) ImageService {
	return &imageService{
		storage:  storage,
		database: db,
	}
}

// UploadImage menangani proses unggah gambar
func (s *imageService) UploadImage(fileHeader *multipart.FileHeader) (*models.Image, error) {
	// Simpan file ke penyimpanan
	image, err := s.storage.Save(fileHeader)
	if err != nil {
		return nil, err
	}

	// Simpan metadata ke database
	result := s.database.DB.Create(image)
	if result.Error != nil {
		return nil, result.Error
	}

	return image, nil
}

// GetDiskUsage mendapatkan informasi penggunaan disk
func (s *imageService) GetDiskUsage() (map[string]interface{}, error) {
	used, available, percentage, err := s.storage.GetDiskUsage()
	if err != nil {
		return nil, err
	}

	// Format ke MB untuk tampilan yang lebih baik
	return map[string]interface{}{
		"used_bytes":      used,
		"available_bytes": available,
		"used_mb":         float64(used) / (1024 * 1024),
		"available_mb":    float64(available) / (1024 * 1024),
		"percentage_used": percentage,
	}, nil
}

// GetImageList mendapatkan daftar gambar yang telah diunggah
func (s *imageService) GetImageList() ([]models.Image, error) {
	var images []models.Image
	result := s.database.DB.Order("created_at desc").Find(&images)
	if result.Error != nil {
		return nil, result.Error
	}

	return images, nil
}
