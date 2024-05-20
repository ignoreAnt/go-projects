package service

import (
	"context"
	"social-server/internal/domain"
)

type UserProfileService struct {
	*GenericService[domain.UserProfile]
}

func NewUserProfileService(repo domain.UserProfileRepository) *UserProfileService {
	return &UserProfileService{
		GenericService: NewGenericService[domain.UserProfile](repo),
	}
}

func (s *UserProfileService) Create(ctx context.Context, profile domain.UserProfile) (domain.UserProfile, error) {
	return s.GenericService.Create(ctx, profile)
}

func (s *UserProfileService) GetById(ctx context.Context, id int) (domain.UserProfile, error) {
	return s.GenericService.GetById(ctx, id)
}

func (s *UserProfileService) Update(ctx context.Context, profile domain.UserProfile) (domain.UserProfile, error) {
	return s.GenericService.Update(ctx, profile)
}

func (s *UserProfileService) Delete(ctx context.Context, id int) error {
	return s.GenericService.Delete(ctx, id)
}
