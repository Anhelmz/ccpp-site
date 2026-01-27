package repositories

import (
	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/repositories"
	"errors"
	"time"

	"gorm.io/gorm"
)

type eventCalendarRepositoryImpl struct {
	db *gorm.DB
}

func NewEventCalendarRepository(db *gorm.DB) repositories.EventCalendarRepository {
	return &eventCalendarRepositoryImpl{db: db}
}

func (r *eventCalendarRepositoryImpl) Create(event *models.EventCalendar) error {
	return r.db.Create(event).Error
}

func (r *eventCalendarRepositoryImpl) FindAll() ([]models.EventCalendar, error) {
	var events []models.EventCalendar
	err := r.db.Order("start_date asc").Find(&events).Error
	return events, err
}

func (r *eventCalendarRepositoryImpl) FindUpcoming() ([]models.EventCalendar, error) {
	var events []models.EventCalendar
	now := time.Now()
	err := r.db.Where("end_date >= ?", now).Order("start_date asc").Find(&events).Error
	return events, err
}

func (r *eventCalendarRepositoryImpl) FindByID(id uint) (*models.EventCalendar, error) {
	var event models.EventCalendar
	err := r.db.First(&event, id).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventCalendarRepositoryImpl) Update(event *models.EventCalendar) error {
	return r.db.Save(event).Error
}

func (r *eventCalendarRepositoryImpl) Delete(id uint) error {
	// Perform hard delete using raw SQL to ensure permanent removal from database
	result := r.db.Exec("DELETE FROM events_calendar WHERE id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	// Verify that a row was actually deleted
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return nil
}
