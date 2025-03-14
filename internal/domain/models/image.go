package models

import (
	"time"
)

// Image menyimpan metadata gambar yang diunggah
type Image struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Filename  string    `json:"filename"`
	Size      int64     `json:"size"` // dalam bytes
	MimeType  string    `json:"mime_type"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at"`
}
