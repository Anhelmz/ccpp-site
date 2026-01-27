package services

import "ccpp-backend/internal/domain/models"

type EventCalendarService interface {
	CreateEvent(event *models.EventCalendar) error
	GetAllEvents() ([]models.EventCalendar, error)
	GetUpcomingEvents() ([]models.EventCalendar, error)
	GetEventByID(id uint) (*models.EventCalendar, error)
	UpdateEvent(event *models.EventCalendar) error
	DeleteEvent(id uint) error
}
