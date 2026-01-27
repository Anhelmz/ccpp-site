package services

import (
	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/repositories"
	"ccpp-backend/internal/domain/services"
)

type eventCalendarServiceImpl struct {
	repo repositories.EventCalendarRepository
}

func NewEventCalendarService(repo repositories.EventCalendarRepository) services.EventCalendarService {
	return &eventCalendarServiceImpl{
		repo: repo,
	}
}

func (s *eventCalendarServiceImpl) CreateEvent(event *models.EventCalendar) error {
	return s.repo.Create(event)
}

func (s *eventCalendarServiceImpl) GetAllEvents() ([]models.EventCalendar, error) {
	return s.repo.FindAll()
}

func (s *eventCalendarServiceImpl) GetUpcomingEvents() ([]models.EventCalendar, error) {
	return s.repo.FindUpcoming()
}

func (s *eventCalendarServiceImpl) GetEventByID(id uint) (*models.EventCalendar, error) {
	return s.repo.FindByID(id)
}

func (s *eventCalendarServiceImpl) UpdateEvent(event *models.EventCalendar) error {
	return s.repo.Update(event)
}

func (s *eventCalendarServiceImpl) DeleteEvent(id uint) error {
	return s.repo.Delete(id)
}
