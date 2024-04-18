package service

import (
	"go-projects/social-media-profile-management/internal/domain"
)

// PictureService is the struct for service for picture
type PictureService struct {
	repository domain.PictureRepository
}

// NewPictureService creates a new PictureService
func NewPictureService(picRepo domain.PictureRepository) *PictureService {
	return &PictureService{picRepo}
}

// CreatePicture creates a new picture
func (ps *PictureService) CreatePicture(pic domain.Picture) error {
	return ps.repository.Create(pic)

}

// UpdatePicture updates a picture
func (ps *PictureService) UpdatePicture(pic domain.Picture) error {
	return ps.repository.Update(pic)

}

// DeletePicture deletes a picture
func (ps *PictureService) DeletePicture(picID int) error {
	return ps.repository.Delete(picID)

}

// GetPictureByID gets a picture by id
func (ps *PictureService) GetPictureByID(picID int) (*domain.Picture, error) {
	return ps.repository.GetById(picID)
}
