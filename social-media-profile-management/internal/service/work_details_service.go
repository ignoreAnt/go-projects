package service

import (
	"social-server/internal/domain"
)

// WorkDetailService is a service for work details
type WorkDetailService struct {
	*GenericService[domain.WorkDetail]
}

// NewService creates a new work detail service
func (s *WorkDetailService) NewService(repo domain.Repository[domain.WorkDetail]) *WorkDetailService {
	return &WorkDetailService{
		GenericService: NewService[domain.WorkDetail](repo),
	}
}
