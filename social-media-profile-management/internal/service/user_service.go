package service

import (
	"context"
	"social-server/internal/domain"
)

type UserService struct {
	*GenericService[domain.User]
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{
		GenericService: NewGenericService[domain.User](repo),
	}
}

func (s *UserService) Create(ctx context.Context, user domain.User) (domain.User, error) {
	return s.GenericService.Create(ctx, user)
}

func (s *UserService) GetById(ctx context.Context, id int) (domain.User, error) {
	return s.GenericService.GetById(ctx, id)
}

func (s *UserService) Update(ctx context.Context, user domain.User) (domain.User, error) {
	return s.GenericService.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, id int) error {
	return s.GenericService.Delete(ctx, id)
}
