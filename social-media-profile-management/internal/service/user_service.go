package service

import (
	"go-projects/social-media-profile-management/internal/domain"
)

type UserService struct {
	repository domain.UserRepository
}

func NewUserService(repository domain.UserRepository) interface{} {
	return &UserService{
		repository: repository,
	}
}
