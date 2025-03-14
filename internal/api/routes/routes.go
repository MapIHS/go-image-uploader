package routes

import (
	"github.com/MapIhs/go-image-uploader/internal/api/handlers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes mengatur semua rute API
func SetupRoutes(app *fiber.App, imageHandler *handlers.ImageHandler, diskHandler *handlers.DiskHandler) {
	// Halaman UI dengan template
	app.Get("/", diskHandler.RenderHomePage)
	app.Get("/upload", imageHandler.RenderUploadForm)

	// API Endpoints
	api := app.Group("/api")
	api.Post("/upload", imageHandler.UploadImage)
	api.Get("/images", imageHandler.GetImageList)
	api.Get("/disk-usage", diskHandler.GetDiskUsage)

	// Serve file statis
	app.Static("/images", "./public/images")
}
