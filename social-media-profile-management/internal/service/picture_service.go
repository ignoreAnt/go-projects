package service

import (
	"social-server/internal/domain"
)

// PictureService is a service for pictures
type PictureService struct {
	*GenericService[domain.Picture]
}

// NewService creates a new picture service
func (s *PictureService) NewService(repo domain.Repository[domain.Picture]) *PictureService {
	return &PictureService{
		GenericService: NewService[domain.Picture](repo),
	}
}
