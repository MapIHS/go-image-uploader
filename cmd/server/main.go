package main

import (
	"log"

	"github.com/MapIhs/go-image-uploader/config"
	"github.com/MapIhs/go-image-uploader/internal/api/handlers"
	"github.com/MapIhs/go-image-uploader/internal/api/routes"
	"github.com/MapIhs/go-image-uploader/internal/infrastructure/database"
	"github.com/MapIhs/go-image-uploader/internal/infrastructure/storage"
	"github.com/MapIhs/go-image-uploader/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Load konfigurasi
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Gagal memuat konfigurasi: %v", err)
	}

	// Inisialisasi database
	db, err := database.NewDatabase(cfg.PostgresConnStr)
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	// Inisialisasi template engine
	engine := html.New("./templates", ".html")
	engine.AddFunc("div", func(a, b int64) float64 {
		return float64(a) / float64(b)
	})

	// Inisialisasi aplikasi Fiber
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Daftarkan middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// Inisialisasi komponen
	imgStorage := storage.NewLocalImageStorage(cfg)
	imgService := usecase.NewImageService(imgStorage, db)
	imgHandler := handlers.NewImageHandler(imgService, cfg)
	diskHandler := handlers.NewDiskHandler(imgService)

	// Setup routes
	routes.SetupRoutes(app, imgHandler, diskHandler)

	// Mulai server
	log.Printf("Server berjalan pada http://localhost:%s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Gagal memulai server: %v", err)
	}
}
