package service

import (
	"context"
	"social-server/internal/domain"
)

type PictureService struct {
	*GenericService[domain.Picture]
}

func NewPictureService(repo domain.PictureRepository) *PictureService {
	return &PictureService{
		GenericService: NewGenericService[domain.Picture](repo),
	}
}

func (s *PictureService) Create(ctx context.Context, picture domain.Picture) (domain.Picture, error) {
	return s.GenericService.Create(ctx, picture)
}

func (s *PictureService) GetById(ctx context.Context, id int) (domain.Picture, error) {
	return s.GenericService.GetById(ctx, id)
}

func (s *PictureService) Update(ctx context.Context, picture domain.Picture) (domain.Picture, error) {
	return s.GenericService.Update(ctx, picture)
}

func (s *PictureService) Delete(ctx context.Context, id int) error {
	return s.GenericService.Delete(ctx, id)
}
