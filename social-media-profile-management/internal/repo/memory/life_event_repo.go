package memory

import (
	"context"
	"errors"
	"social-server/internal/domain"
	"social-server/pkg/util"
	"sync/atomic"
)

type LifeEventRepository struct {
	*util.SyncMap[int, domain.LifeEvent]
	nextID int32
}

func NewLifeEventRepository() *LifeEventRepository {
	repo := &LifeEventRepository{
		nextID: 1,
	}
	generateID := func() int {
		return int(atomic.AddInt32(&repo.nextID, 1))
	}
	repo.SyncMap = util.NewSyncMap[int, domain.LifeEvent](generateID)
	return repo
}

func (r *LifeEventRepository) Create(ctx context.Context, event domain.LifeEvent) (domain.LifeEvent, error) {
	id, _ := r.SyncMap.Create(event)
	event.LifeEventID = id
	return event, nil
}

func (r *LifeEventRepository) GetById(ctx context.Context, id int) (domain.LifeEvent, error) {
	event, exists := r.SyncMap.Retrieve(id)
	if !exists {
		return domain.LifeEvent{}, errors.New("life event not found")
	}
	return event, nil
}

func (r *LifeEventRepository) Update(ctx context.Context, event domain.LifeEvent) (domain.LifeEvent, error) {
	_, exists := r.SyncMap.Update(event.LifeEventID, event)
	if !exists {
		return domain.LifeEvent{}, errors.New("life event not found")
	}
	return event, nil
}

func (r *LifeEventRepository) Delete(ctx context.Context, id int) error {
	if !r.SyncMap.Delete(id) {
		return errors.New("life event not found")
	}
	return nil
}
