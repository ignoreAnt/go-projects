package memory

import (
	"go-projects/social-media-profile-management/internal/domain"
	appErrors "go-projects/social-media-profile-management/internal/errors"
	"sync"
)

// LifeEventRepository is a struct that defines the repository for the life event
type LifeEventRepository struct {
	lifeEvents map[int]domain.LifeEvent
	mu         sync.RWMutex
	nextID     int
}

// NewLifeEventRepository is a function that returns a new LifeEventRepository
func NewLifeEventRepository() *LifeEventRepository {
	return &LifeEventRepository{
		lifeEvents: make(map[int]domain.LifeEvent),
		nextID:     1,
	}
}

// Create is a method that creates a new life event
func (r *LifeEventRepository) Create(lifeEvent domain.LifeEvent) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	lifeEvent.LifeEventID = r.nextID
	r.nextID++
	r.lifeEvents[lifeEvent.LifeEventID] = lifeEvent

	return nil

}

// Update is a method that updates a life event
func (r *LifeEventRepository) Update(lifeEvent domain.LifeEvent) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.lifeEvents[lifeEvent.LifeEventID]; !exists {
		return appErrors.ErrNotFound
	}
	r.lifeEvents[lifeEvent.LifeEventID] = lifeEvent
	return nil
}

// Delete is a method that deletes a life event
func (r *LifeEventRepository) Delete(lifeEventID int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.lifeEvents[lifeEventID]; !exists {
		return appErrors.ErrNotFound
	}
	delete(r.lifeEvents, lifeEventID)
	return nil
}
