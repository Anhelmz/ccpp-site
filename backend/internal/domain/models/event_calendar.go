package models

import (
	"time"

	"gorm.io/gorm"
)

type EventCalendar struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Title     string         `gorm:"not null" json:"title" binding:"required"`
	Details   string         `gorm:"type:text" json:"details"`
	StartDate time.Time      `gorm:"not null" json:"startDate" binding:"required"`
	EndDate   time.Time      `gorm:"not null" json:"endDate" binding:"required"`
	Location  string         `gorm:"type:text" json:"location"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for EventCalendar
func (EventCalendar) TableName() string {
	return "events_calendar"
}

// EventCalendarResponse is used for JSON serialization with UTC dates
type EventCalendarResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Details   string    `json:"details"`
	StartDate string    `json:"startDate"`
	EndDate   string    `json:"endDate"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
