package service

import "social-server/internal/domain"

// UserService is a struct that defines the service for the user
type UserService struct {
	repository domain.UserRepository
}

// NewUserService is a function that returns a new UserService
func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repository: repo}
}

// CreateUser is a method that creates a new user
func (u *UserService) CreateUser(user domain.User) (int, error) {
	return u.repository.Create(user)
}

// GetUserByID is a method that gets a user by id
func (u *UserService) GetUserByID(id int) (*domain.User, error) {
	return u.repository.GetById(id)
}
