package repositories

import (
	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/repositories"
	"time"

	"gorm.io/gorm"
)

type eventRepositoryImpl struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) repositories.EventRepository {
	return &eventRepositoryImpl{db: db}
}

func (r *eventRepositoryImpl) Create(event *models.Event) error {
	return r.db.Create(event).Error
}

func (r *eventRepositoryImpl) FindAll() ([]models.Event, error) {
	var events []models.Event
	err := r.db.Order("start_time asc, end_time asc").Find(&events).Error
	return events, err
}

func (r *eventRepositoryImpl) FindUpcoming() ([]models.Event, error) {
	var events []models.Event
	now := time.Now()
	err := r.db.
		Where("start_time >= ?", now).
		Order("start_time asc, end_time asc").
		Find(&events).Error
	return events, err
}

func (r *eventRepositoryImpl) FindByID(id uint) (*models.Event, error) {
	var event models.Event
	err := r.db.First(&event, id).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepositoryImpl) Update(event *models.Event) error {
	return r.db.Save(event).Error
}

func (r *eventRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Event{}, id).Error
}




