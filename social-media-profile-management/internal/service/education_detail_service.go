package service

import (
	"context"
	"social-server/internal/domain"
)

type EducationDetailService struct {
	*GenericService[domain.EducationDetail]
}

func NewEducationDetailService(repo domain.EducationDetailRepository) *EducationDetailService {
	return &EducationDetailService{
		GenericService: NewGenericService[domain.EducationDetail](repo),
	}
}

func (s *EducationDetailService) Create(ctx context.Context, detail domain.EducationDetail) (domain.EducationDetail, error) {
	return s.GenericService.Create(ctx, detail)
}

func (s *EducationDetailService) GetById(ctx context.Context, id int) (domain.EducationDetail, error) {
	return s.GenericService.GetById(ctx, id)
}

func (s *EducationDetailService) Update(ctx context.Context, detail domain.EducationDetail) (domain.EducationDetail, error) {
	return s.GenericService.Update(ctx, detail)
}

func (s *EducationDetailService) Delete(ctx context.Context, id int) error {
	return s.GenericService.Delete(ctx, id)
}
