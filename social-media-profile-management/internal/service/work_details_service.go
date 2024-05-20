package service

import (
	"context"
	"social-server/internal/domain"
)

type WorkDetailService struct {
	*GenericService[domain.WorkDetail]
}

func NewWorkDetailService(repo domain.WorkDetailRepository) *WorkDetailService {
	return &WorkDetailService{
		GenericService: NewGenericService[domain.WorkDetail](repo),
	}
}

func (s *WorkDetailService) Create(ctx context.Context, detail domain.WorkDetail) (domain.WorkDetail, error) {
	return s.GenericService.Create(ctx, detail)
}

func (s *WorkDetailService) GetById(ctx context.Context, id int) (domain.WorkDetail, error) {
	return s.GenericService.GetById(ctx, id)
}

func (s *WorkDetailService) Update(ctx context.Context, detail domain.WorkDetail) (domain.WorkDetail, error) {
	return s.GenericService.Update(ctx, detail)
}

func (s *WorkDetailService) Delete(ctx context.Context, id int) error {
	return s.GenericService.Delete(ctx, id)
}
