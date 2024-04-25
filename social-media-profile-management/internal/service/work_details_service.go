package service

import "social-server/internal/domain"

// WorkDetailsService is a struct that defines the service for the work details
type WorkDetailsService struct {
	repository domain.WorkDetailsRepository
}

// NewWorkDetailsService creates new WorkDetailsService
func NewWorkDetailsService(repo domain.WorkDetailsRepository) *WorkDetailsService {
	return &WorkDetailsService{
		repository: repo,
	}
}

// CreateWorkDetail creates work detail
func (w *WorkDetailsService) CreateWorkDetails(workDetail domain.WorkDetail) error {
	return w.repository.Create(workDetail)
}
