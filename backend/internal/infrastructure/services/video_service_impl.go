package services

import (
	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/repositories"
	"ccpp-backend/internal/domain/services"
)

type videoServiceImpl struct {
	repo repositories.VideoRepository
}

func NewVideoService(repo repositories.VideoRepository) services.VideoService {
	return &videoServiceImpl{
		repo: repo,
	}
}

func (s *videoServiceImpl) CreateVideo(video *models.Video) error {
	return s.repo.Create(video)
}

func (s *videoServiceImpl) GetAllVideos() ([]models.Video, error) {
	return s.repo.FindAll()
}

func (s *videoServiceImpl) GetVideoByID(id uint) (*models.Video, error) {
	return s.repo.FindByID(id)
}

func (s *videoServiceImpl) UpdateVideo(video *models.Video) error {
	return s.repo.Update(video)
}

func (s *videoServiceImpl) DeleteVideo(id uint) error {
	return s.repo.Delete(id)
}
