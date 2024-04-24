package memory

import (
	"fmt"
	"go-projects/social-media-profile-management/internal/domain"
	"sync"
)

// Backend is a generic in-memory storage for any type.
type Backend[T any] struct {
	mu      sync.RWMutex
	backend map[int]T
	seq     int // for generating unique IDs
}

// NewBackend creates a new Backend
func NewBackend[T any]() *Backend[T] {
	return &Backend[T]{
		backend: make(map[int]T),
	}
}

// Create adds a new item to the backend and returns its ID.
func (b *Backend[T]) Create(item T) error {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.seq++     // increment sequence number to get a unique ID
	id := b.seq // use the incremented sequence number as the ID
	b.backend[id] = item
	return nil
}

// Get retrieves an item by ID. It returns the item or an error if not found.
func (b *Backend[T]) Get(id int) (*T, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if item, ok := b.backend[id]; ok {
		return &item, nil
	}
	return nil, fmt.Errorf("item not found")
}

// Update modifies an item with the given ID. Returns an error if the item is not found.
func (b *Backend[T]) Update(id int, item T) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.backend[id]; !exists {
		return fmt.Errorf("item not found")
	}
	b.backend[id] = item
	return nil
}

// Delete removes an item by ID. It returns an error if the item does not exist.
func (b *Backend[T]) Delete(id int) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.backend[id]; !exists {
		return fmt.Errorf("item not found")
	}
	delete(b.backend, id)
	return nil
}

// Verify Backend implements the DataManager interface
var _ domain.DataManager[domain.UserProfile] = &Backend[domain.UserProfile]{}
var _ domain.DataManager[domain.User] = &Backend[domain.User]{}
var _ domain.DataManager[domain.Picture] = &Backend[domain.Picture]{}
var _ domain.DataManager[domain.Privacy] = &Backend[domain.Privacy]{}
var _ domain.DataManager[domain.WorkDetail] = &Backend[domain.WorkDetail]{}
var _ domain.DataManager[domain.EmploymentDetail] = &Backend[domain.EmploymentDetail]{}
