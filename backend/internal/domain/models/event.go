package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	Title            string         `gorm:"not null" json:"title" binding:"required"`
	Description      string         `gorm:"type:text" json:"description"`
	Summary          string         `gorm:"type:text" json:"summary"`
	StartTime        time.Time      `gorm:"not null" json:"startTime" binding:"required"`
	EndTime          time.Time      `gorm:"not null" json:"endTime" binding:"required"`
	Location         string         `json:"location"`
	Category         string         `json:"category"`
	Color            string         `json:"color"`
	Recurrence       string         `gorm:"type:text;default:none" json:"recurrence"`
	RecurrenceEndsAt *time.Time     `json:"recurrenceEndsAt"`
	Timezone         string         `gorm:"type:text;default:UTC" json:"timezone"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}
