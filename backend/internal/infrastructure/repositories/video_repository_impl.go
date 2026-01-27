package repositories

import (
	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/repositories"
	"errors"

	"gorm.io/gorm"
)

type videoRepositoryImpl struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) repositories.VideoRepository {
	return &videoRepositoryImpl{db: db}
}

func (r *videoRepositoryImpl) Create(video *models.Video) error {
	return r.db.Create(video).Error
}

func (r *videoRepositoryImpl) FindAll() ([]models.Video, error) {
	var videos []models.Video
	err := r.db.Order("created_at desc").Find(&videos).Error
	return videos, err
}

func (r *videoRepositoryImpl) FindByID(id uint) (*models.Video, error) {
	var video models.Video
	err := r.db.First(&video, id).Error
	if err != nil {
		return nil, err
	}
	return &video, nil
}

func (r *videoRepositoryImpl) Update(video *models.Video) error {
	return r.db.Save(video).Error
}

func (r *videoRepositoryImpl) Delete(id uint) error {
	// Perform hard delete using raw SQL to ensure permanent removal from database
	result := r.db.Exec("DELETE FROM videos WHERE id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	// Verify that a row was actually deleted
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return nil
}
