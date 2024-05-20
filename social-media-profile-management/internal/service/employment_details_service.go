package service

import (
	"social-server/internal/domain"
)

// EmploymentDetailService is a service for employment details
type EmploymentDetailService struct {
	*GenericService[domain.EmploymentDetail]
}

// NewService creates a new employment detail service
func (s *EmploymentDetailService) NewService(repo domain.Repository[domain.EmploymentDetail]) *EmploymentDetailService {
	return &EmploymentDetailService{
		GenericService: NewService[domain.EmploymentDetail](repo),
	}
}
