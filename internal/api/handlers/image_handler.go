package handlers

import (
	"github.com/MapIhs/go-image-uploader/config"
	"github.com/MapIhs/go-image-uploader/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

// ImageHandler menangani request terkait gambar
type ImageHandler struct {
	imageService usecase.ImageService
	config       *config.Config
}

// NewImageHandler membuat instance baru dari ImageHandler
func NewImageHandler(imageService usecase.ImageService, cfg *config.Config) *ImageHandler {
	return &ImageHandler{
		imageService: imageService,
		config:       cfg,
	}
}

// UploadImage menangani API untuk unggah gambar
func (h *ImageHandler) UploadImage(c *fiber.Ctx) error {
	// Ambil file dari request
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Gagal mengambil file gambar",
		})
	}

	// Validasi ukuran file
	if file.Size > h.config.ImageSizeLimit {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":          "Ukuran file terlalu besar",
			"max_size_bytes": h.config.ImageSizeLimit,
			"max_size_mb":    float64(h.config.ImageSizeLimit) / (1024 * 1024),
		})
	}

	// Proses unggah gambar
	image, err := h.imageService.UploadImage(file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengunggah gambar: " + err.Error(),
		})
	}

	// Kembalikan respons sukses
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   image,
	})
}

// RenderUploadForm menampilkan halaman form unggah
func (h *ImageHandler) RenderUploadForm(c *fiber.Ctx) error {
	return c.Render("upload", fiber.Map{
		"title":       "Unggah Gambar",
		"max_size_mb": float64(h.config.ImageSizeLimit) / (1024 * 1024),
	})
}

// GetImageList menangani API untuk mendapatkan daftar gambar
func (h *ImageHandler) GetImageList(c *fiber.Ctx) error {
	images, err := h.imageService.GetImageList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mendapatkan daftar gambar: " + err.Error(),
		})
	}

	// Melanjutkan kode dari GetImageList
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   images,
	})
}
