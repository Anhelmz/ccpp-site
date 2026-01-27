package models

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Title       string         `gorm:"column:title;not null" json:"title" binding:"required"`
	Description string         `gorm:"column:description;type:text" json:"description"`
	YouTubeID   string         `gorm:"column:youtube_id;type:varchar(11);not null" json:"youtubeId" binding:"required"`
	YouTubeURL  string         `gorm:"column:youtube_url;type:text;not null" json:"youtubeUrl" binding:"required"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

// TableName specifies the table name for Video
func (Video) TableName() string {
	return "videos"
}
