package repositories

import "ccpp-backend/internal/domain/models"

type EventCalendarRepository interface {
	Create(event *models.EventCalendar) error
	FindAll() ([]models.EventCalendar, error)
	FindUpcoming() ([]models.EventCalendar, error)
	FindByID(id uint) (*models.EventCalendar, error)
	Update(event *models.EventCalendar) error
	Delete(id uint) error
}
