package service

import "go-projects/social-media-profile-management/internal/domain"

type LifeEventService struct {
	repository domain.LifeEventsRepository
}

// NewLifeEventService creates a new LifeEventService
func NewLifeEventService(repository domain.LifeEventsRepository) *LifeEventService {
	return &LifeEventService{repository: repository}
}

// Create is a method that creates a new life event
func (s *LifeEventService) Create(lifeEvent domain.LifeEvent) error {
	return s.repository.Create(lifeEvent)
}

// Update is a method that updates a life event
func (s *LifeEventService) Update(lifeEvent domain.LifeEvent) error {
	return s.repository.Update(lifeEvent)
}

// Delete is a method that deletes a life event
func (s *LifeEventService) Delete(lifeEventID int) error {
	return s.repository.Delete(lifeEventID)
}
