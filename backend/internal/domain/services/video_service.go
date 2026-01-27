package services

import "ccpp-backend/internal/domain/models"

type VideoService interface {
	CreateVideo(video *models.Video) error
	GetAllVideos() ([]models.Video, error)
	GetVideoByID(id uint) (*models.Video, error)
	UpdateVideo(video *models.Video) error
	DeleteVideo(id uint) error
}
