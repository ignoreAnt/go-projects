package service

import (
	"context"
	"social-server/internal/domain"
)

type LifeEventService struct {
	*GenericService[domain.LifeEvent]
}

func NewLifeEventService(repo domain.LifeEventRepository) *LifeEventService {
	return &LifeEventService{
		GenericService: NewGenericService[domain.LifeEvent](repo),
	}
}

func (s *LifeEventService) Create(ctx context.Context, event domain.LifeEvent) (domain.LifeEvent, error) {
	// Add any additional logic for life event creation if needed
	return s.GenericService.Create(ctx, event)
}

func (s *LifeEventService) GetById(ctx context.Context, id int) (domain.LifeEvent, error) {
	// Add any additional logic for life event retrieval if needed
	return s.GenericService.GetById(ctx, id)
}

func (s *LifeEventService) Update(ctx context.Context, event domain.LifeEvent) (domain.LifeEvent, error) {
	// Add any additional logic for life event update if needed
	return s.GenericService.Update(ctx, event)
}

func (s *LifeEventService) Delete(ctx context.Context, id int) error {
	// Add any additional logic for life event deletion if needed
	return s.GenericService.Delete(ctx, id)
}
