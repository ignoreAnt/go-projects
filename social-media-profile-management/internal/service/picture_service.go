package service

import (
	"go-projects/social-media-profile-management/internal/domain"
)

// PictureService is the struct for service for picture
type PictureService struct {
	repository domain.PictureRepository
}

// NewPictureService creates a new Pictureservice
func NewPictureService(picrepo domain.PictureRepository) *PictureService {
	return &PictureService{picrepo}
}

// CreatePicture creates a new picture
func (ps *PictureService) CreatePicture(pic domain.Picture) error {
	ps.repository.Create(pic)
	return nil
}

// UpdatePicture updates a picture
func (ps *PictureService) UpdatePicture(pic domain.Picture) error {
	ps.repository.Update(pic)
	return nil
}

// DeletePicture deletes a picture
func (ps *PictureService) DeletePicture(picID int) error {
	ps.repository.Delete(picID)
	return nil
}

// GetPictureByID gets a picture by id
func (ps *PictureService) GetPictureByID(picID int) (*domain.Picture, error) {
	return ps.repository.GetById(picID)
}
