package repositories

import (
	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/repositories"
	"errors"

	"gorm.io/gorm"
)

type galleryRepositoryImpl struct {
	db *gorm.DB
}

func NewGalleryRepository(db *gorm.DB) repositories.GalleryRepository {
	return &galleryRepositoryImpl{db: db}
}

func (r *galleryRepositoryImpl) Create(gallery *models.Gallery) error {
	return r.db.Create(gallery).Error
}

func (r *galleryRepositoryImpl) FindAll() ([]models.Gallery, error) {
	var galleries []models.Gallery
	// GORM automatically adds "deleted_at IS NULL" filter due to DeletedAt field in model
	// Since we use hard deletes (raw SQL), records are permanently removed, so this filter is harmless
	err := r.db.Order("created_at desc").Find(&galleries).Error
	return galleries, err
}

func (r *galleryRepositoryImpl) FindByCategory(category string) ([]models.Gallery, error) {
	var galleries []models.Gallery
	// GORM automatically adds "deleted_at IS NULL" filter due to DeletedAt field in model
	// Since we use hard deletes (raw SQL), records are permanently removed, so this filter is harmless
	err := r.db.Where("category = ?", category).Order("created_at desc").Find(&galleries).Error
	return galleries, err
}

func (r *galleryRepositoryImpl) FindByID(id uint) (*models.Gallery, error) {
	var gallery models.Gallery
	err := r.db.First(&gallery, id).Error
	if err != nil {
		return nil, err
	}
	return &gallery, nil
}

func (r *galleryRepositoryImpl) Delete(id uint) error {
	// Perform hard delete using raw SQL to ensure permanent removal from database
	// This bypasses GORM's soft delete mechanism completely
	result := r.db.Exec("DELETE FROM galleries WHERE id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	// Verify that a row was actually deleted
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return nil
}

func (r *galleryRepositoryImpl) DeleteByCategory(category string) error {
	// Perform hard delete using raw SQL to ensure permanent removal from database
	// This bypasses GORM's soft delete mechanism completely
	result := r.db.Exec("DELETE FROM galleries WHERE category = ?", category)
	if result.Error != nil {
		return result.Error
	}
	// Note: result.RowsAffected will be 0 if no records matched, which is valid
	return nil
}

func (r *galleryRepositoryImpl) FindOldPhotos() ([]models.Gallery, error) {
	var galleries []models.Gallery
	// Find photos that are soft-deleted (deleted_at IS NOT NULL) or don't have ImageData (old file-based photos)
	// Use Unscoped() to include soft-deleted records
	err := r.db.Unscoped().
		Where("deleted_at IS NOT NULL OR image_data IS NULL OR image_data = ''").
		Order("created_at desc").
		Find(&galleries).Error
	return galleries, err
}

func (r *galleryRepositoryImpl) DeleteOldPhotos() (int64, error) {
	// Permanently delete old photos: soft-deleted ones or those without ImageData
	result := r.db.Exec("DELETE FROM galleries WHERE deleted_at IS NOT NULL OR image_data IS NULL OR image_data = ''")
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
