package service

import (
	"social-server/internal/domain"
)

type EducationDetailService struct {
	*GenericService[domain.EducationDetail]
}

func (s *EducationDetailService) NewService(repo domain.Repository[domain.EducationDetail]) *EducationDetailService {
	return &EducationDetailService{
		GenericService: NewService[domain.EducationDetail](repo),
	}
}
