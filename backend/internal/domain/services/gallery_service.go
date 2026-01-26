package services

import "ccpp-backend/internal/domain/models"

type GalleryService interface {
	CreateGallery(gallery *models.Gallery) error
	GetAllGalleries() ([]models.Gallery, error)
	GetGalleriesByCategory(category string) ([]models.Gallery, error)
	GetGalleryByID(id uint) (*models.Gallery, error)
	DeleteGallery(id uint) error
	DeleteGalleriesByCategory(category string) error
	GetOldPhotos() ([]models.Gallery, error)
	DeleteOldPhotos() (int64, error)
}

