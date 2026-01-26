package models

import (
	"time"

	"gorm.io/gorm"
)

type Gallery struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Filename  string         `gorm:"not null" json:"filename"`
	Path      string         `gorm:"not null" json:"path"` // Kept for backward compatibility, but will store data URI
	ImageData string         `gorm:"type:text" json:"-"`  // Store base64 image data
	MimeType  string         `gorm:"type:text" json:"mimeType"` // Store MIME type (image/jpeg, image/png, etc.) - will be set to NOT NULL after backfill
	Category  string         `gorm:"not null" json:"category" binding:"required"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

