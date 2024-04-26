package memory

import (
	"social-server/internal/domain"
	appErrors "social-server/internal/errors"
	"sync"
)

// PictureRepository is a struct that defines the repository for the picture
type PictureRepository struct {
	pictures map[int]domain.Picture
	mu       sync.RWMutex // RWMutex is a reader/writer mutual exclusion lock to ensure that multiple goroutines can read the map concurrently
	nextID   int
}

// NewPictureRepository is a function that returns a new PictureRepository
func NewPictureRepository() *PictureRepository {
	return &PictureRepository{
		pictures: make(map[int]domain.Picture),
		nextID:   1,
	}
}

// Create is a method that creates a new picture
func (r *PictureRepository) Create(picture domain.Picture) error {
	// Acquire the lock
	r.mu.Lock()
	defer r.mu.Unlock()

	picture.PictureID = r.nextID
	r.nextID++
	r.pictures[picture.PictureID] = picture

	return nil

}

// Update is a method that updates a picture
func (r *PictureRepository) Update(picture domain.Picture) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.pictures[picture.PictureID]; !exists {
		return appErrors.ErrNotFound
	}
	r.pictures[picture.PictureID] = picture
	return nil
}

// Delete is a method that deletes a picture
func (r *PictureRepository) Delete(pictureID int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.pictures[pictureID]; !exists {
		return appErrors.ErrNotFound
	}
	delete(r.pictures, pictureID)
	return nil
}

// GetById is a method that gets a picture by its ID
func (r *PictureRepository) GetById(pictureID int) (*domain.Picture, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if picture, exists := r.pictures[pictureID]; exists {
		return &picture, nil
	}
	return nil, appErrors.ErrNotFound
}
