package repositories

import "ccpp-backend/internal/domain/models"

type VideoRepository interface {
	Create(video *models.Video) error
	FindAll() ([]models.Video, error)
	FindByID(id uint) (*models.Video, error)
	Update(video *models.Video) error
	Delete(id uint) error
}
