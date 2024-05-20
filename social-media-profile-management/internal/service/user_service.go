package service

import (
	"social-server/internal/domain"
)

// UserService is a service for users
type UserService struct {
	*GenericService[domain.User]
}

// NewService creates a new user service
func (s *UserService) NewService(repo domain.Repository[domain.User]) *UserService {
	return &UserService{
		GenericService: NewService[domain.User](repo),
	}
}
