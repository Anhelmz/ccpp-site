package services

import (
	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/repositories"
	"ccpp-backend/internal/domain/services"
)

type galleryServiceImpl struct {
	galleryRepo repositories.GalleryRepository
}

func NewGalleryService(galleryRepo repositories.GalleryRepository) services.GalleryService {
	return &galleryServiceImpl{galleryRepo: galleryRepo}
}

func (s *galleryServiceImpl) CreateGallery(gallery *models.Gallery) error {
	return s.galleryRepo.Create(gallery)
}

func (s *galleryServiceImpl) GetAllGalleries() ([]models.Gallery, error) {
	return s.galleryRepo.FindAll()
}

func (s *galleryServiceImpl) GetGalleriesByCategory(category string) ([]models.Gallery, error) {
	return s.galleryRepo.FindByCategory(category)
}

func (s *galleryServiceImpl) GetGalleryByID(id uint) (*models.Gallery, error) {
	return s.galleryRepo.FindByID(id)
}

func (s *galleryServiceImpl) DeleteGallery(id uint) error {
	return s.galleryRepo.Delete(id)
}

func (s *galleryServiceImpl) DeleteGalleriesByCategory(category string) error {
	return s.galleryRepo.DeleteByCategory(category)
}

func (s *galleryServiceImpl) GetOldPhotos() ([]models.Gallery, error) {
	return s.galleryRepo.FindOldPhotos()
}

func (s *galleryServiceImpl) DeleteOldPhotos() (int64, error) {
	return s.galleryRepo.DeleteOldPhotos()
}
