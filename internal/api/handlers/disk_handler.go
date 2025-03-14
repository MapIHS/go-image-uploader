package handlers

import (
	"github.com/MapIhs/go-image-uploader/internal/domain/models"
	"github.com/MapIhs/go-image-uploader/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

// DiskHandler menangani request terkait informasi disk
type DiskHandler struct {
	imageService usecase.ImageService
}

// NewDiskHandler membuat instance baru dari DiskHandler
func NewDiskHandler(imageService usecase.ImageService) *DiskHandler {
	return &DiskHandler{
		imageService: imageService,
	}
}

// GetDiskUsage menangani API untuk informasi penggunaan disk
func (h *DiskHandler) GetDiskUsage(c *fiber.Ctx) error {
	diskInfo, err := h.imageService.GetDiskUsage()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mendapatkan informasi penggunaan disk: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   diskInfo,
	})
}

// RenderHomePage menampilkan halaman beranda dengan informasi penggunaan disk
func (h *DiskHandler) RenderHomePage(c *fiber.Ctx) error {
	diskInfo, err := h.imageService.GetDiskUsage()
	if err != nil {
		diskInfo = map[string]interface{}{
			"error": "Gagal mendapatkan informasi penggunaan disk",
		}
	}

	// Dapatkan daftar gambar untuk ditampilkan di beranda
	images, err := h.imageService.GetImageList()
	if err != nil {
		images = []models.Image{}
	}

	return c.Render("home", fiber.Map{
		"title":     "Beranda - Layanan Unggah Gambar",
		"disk_info": diskInfo,
		"images":    images,
	})
}
