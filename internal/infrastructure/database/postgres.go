package database

import (
	"github.com/MapIhs/go-image-uploader/internal/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database adalah wrapper untuk koneksi database
type Database struct {
	DB *gorm.DB
}

// NewDatabase membuat koneksi database baru
func NewDatabase(connStr string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate model
	err = db.AutoMigrate(&models.Image{})
	if err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}

// Close menutup koneksi database
func (d *Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
