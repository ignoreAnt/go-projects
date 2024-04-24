package service

import (
	"go-projects/social-media-profile-management/internal/domain"
)

// SocialMediaProfileService is a service for managing social media profiles.
type SocialMediaProfileService[T any] struct {
	repo domain.DataManager[T]
}

// NewSocialMediaProfileService creates a new SocialMediaProfileService.
func NewSocialMediaProfileService[T any](repo domain.DataManager[T]) *SocialMediaProfileService[T] {
	return &SocialMediaProfileService[T]{repo: repo}
}

// Create adds a new item to the backend and returns error.
func (s *SocialMediaProfileService[T]) Create(item T) error {
	return s.repo.Create(item)
}

// Get retrieves an item by ID. It returns the item or an error if not found.
func (s *SocialMediaProfileService[T]) Get(id int) (*T, error) {
	return s.repo.Get(id)
}

// Update modifies an item with the given ID. Returns an error if the item is not found.
func (s *SocialMediaProfileService[T]) Update(id int, item T) error {
	return s.repo.Update(id, item)
}

// Delete removes an item by ID. It returns an error if the item does not exist.
func (s *SocialMediaProfileService[T]) Delete(id int) error {
	return s.repo.Delete(id)
}
