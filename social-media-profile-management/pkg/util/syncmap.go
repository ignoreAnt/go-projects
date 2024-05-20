package util

import (
	"sync"
)

// SyncMap is a generic thread-safe map with basic CRUD operations.
type SyncMap[K comparable, V any] struct {
	sync.Mutex
	items      map[K]V
	generateID func() K
}

// NewSyncMap creates a new instance of SyncMap with a given ID generation function.
func NewSyncMap[K comparable, V any](generateID func() K) *SyncMap[K, V] {
	return &SyncMap[K, V]{
		items:      make(map[K]V),
		generateID: generateID,
	}
}

// Create inserts a new item and generates an ID.
func (m *SyncMap[K, V]) Create(value V) (K, V) {
	m.Lock()
	defer m.Unlock()
	id := m.generateID()
	m.items[id] = value
	return id, value
}

// Retrieve returns an item by key.
func (m *SyncMap[K, V]) Retrieve(key K) (V, bool) {
	m.Lock()
	defer m.Unlock()
	item, exists := m.items[key]
	return item, exists
}

// Update modifies an item at a given key.
func (m *SyncMap[K, V]) Update(key K, value V) (V, bool) {
	m.Lock()
	defer m.Unlock()
	if _, exists := m.items[key]; !exists {
		var zero V
		return zero, false
	}
	m.items[key] = value
	return value, true
}

// Delete removes an item by key.
func (m *SyncMap[K, V]) Delete(key K) bool {
	m.Lock()
	defer m.Unlock()
	if _, exists := m.items[key]; !exists {
		return false
	}
	delete(m.items, key)
	return true
}
