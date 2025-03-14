package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/MapIhs/go-image-uploader/config"
	"github.com/MapIhs/go-image-uploader/internal/domain/models"
)

// ImageStorage mendefinisikan operasi penyimpanan gambar
type ImageStorage interface {
	Save(file *multipart.FileHeader) (*models.Image, error)
	GetDiskUsage() (int64, int64, float64, error) // used, available, percentage
}

// LocalImageStorage mengimplementasikan ImageStorage dengan menyimpan ke filesystem lokal
type LocalImageStorage struct {
	config *config.Config
}

// NewLocalImageStorage membuat instance baru dari LocalImageStorage
func NewLocalImageStorage(cfg *config.Config) *LocalImageStorage {
	return &LocalImageStorage{
		config: cfg,
	}
}

// Save menyimpan gambar ke filesystem lokal
func (s *LocalImageStorage) Save(fileHeader *multipart.FileHeader) (*models.Image, error) {
	// Validasi ukuran file
	if fileHeader.Size > s.config.ImageSizeLimit {
		return nil, fmt.Errorf("ukuran file terlalu besar (maksimum: %d bytes)", s.config.ImageSizeLimit)
	}

	// Buka file yang diunggah
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Buat nama file unik
	ext := filepath.Ext(fileHeader.Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	// Path tujuan
	dstPath := filepath.Join(s.config.UploadPath, fileName)

	// Buat file tujuan
	dst, err := os.Create(dstPath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// Salin konten file
	if _, err = io.Copy(dst, file); err != nil {
		return nil, err
	}

	// Buat dan kembalikan objek Image
	return &models.Image{
		Filename:  fileHeader.Filename,
		Size:      fileHeader.Size,
		MimeType:  fileHeader.Header.Get("Content-Type"),
		Path:      fileName,
		CreatedAt: time.Now(),
	}, nil
}

// GetDiskUsage mendapatkan informasi penggunaan disk
func (s *LocalImageStorage) GetDiskUsage() (int64, int64, float64, error) {
	var stat syscall.Statfs_t
	err := syscall.Statfs(s.config.UploadPath, &stat)
	if err != nil {
		return 0, 0, 0, err
	}

	// Hitung total, available, dan used space
	totalSpace := stat.Blocks * uint64(stat.Bsize)
	availSpace := stat.Bavail * uint64(stat.Bsize)
	usedSpace := totalSpace - availSpace
	usedPercent := float64(usedSpace) / float64(totalSpace) * 100.0

	return int64(usedSpace), int64(availSpace), usedPercent, nil
}
