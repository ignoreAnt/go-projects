package service

import (
	"social-server/internal/domain"
)

// LifeEventService is a service for life events
type LifeEventService struct {
	*GenericService[domain.LifeEvent]
}

// NewService creates a new life event service
func (s *LifeEventService) NewService(repo domain.Repository[domain.LifeEvent]) *LifeEventService {
	return &LifeEventService{
		GenericService: NewService[domain.LifeEvent](repo),
	}
}
