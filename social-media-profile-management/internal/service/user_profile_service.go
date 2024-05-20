package service

import (
	"social-server/internal/domain"
)

// UserProfileService is a service for User Profiles
type UserProfileService struct {
	*GenericService[domain.UserProfile]
}

// NewService creates a new user profile service
func (s *UserProfileService) NewService(repo domain.Repository[domain.UserProfile]) *UserProfileService {
	return &UserProfileService{
		GenericService: NewService[domain.UserProfile](repo),
	}
}
