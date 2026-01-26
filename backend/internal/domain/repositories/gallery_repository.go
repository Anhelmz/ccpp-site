package repositories

import "ccpp-backend/internal/domain/models"

type GalleryRepository interface {
	Create(gallery *models.Gallery) error
	FindAll() ([]models.Gallery, error)
	FindByCategory(category string) ([]models.Gallery, error)
	FindByID(id uint) (*models.Gallery, error)
	Delete(id uint) error
	DeleteByCategory(category string) error
	// FindOldPhotos finds photos that are soft-deleted or don't have ImageData (old file-based photos)
	FindOldPhotos() ([]models.Gallery, error)
	// DeleteOldPhotos permanently deletes old photos (soft-deleted or without ImageData)
	DeleteOldPhotos() (int64, error)
}

