package memory

import (
	"go-projects/social-media-profile-management/internal/domain"
	appErrors "go-projects/social-media-profile-management/internal/errors"
	"sync"
)

// UserRepository is a struct that defines the repository for the user
type UserRepository struct {
	users  map[int]domain.User
	mu     sync.RWMutex // RWMutex is a reader/writer mutual exclusion lock to ensure that multiple goroutines can read the map concurrently
	nextID int
}

// UpdateUserName is a method that updates a user's username
func (r *UserRepository) UpdateUserName(userID int, newUserName string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.users[userID]; !exists {
		return appErrors.ErrNotFound
	}
	r.users[userID] = domain.User{
		UserID:   userID,
		UserName: newUserName,
	}
	return nil
}

// NewUserRepository is a function that returns a new UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  make(map[int]domain.User),
		nextID: 1,
	}
}

// Create is a method that creates a new user
func (r *UserRepository) Create(user domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	user.UserID = r.nextID
	r.nextID++
	r.users[user.UserID] = user
	return nil
}

// Update is a method that updates a user
func (r *UserRepository) Update(user domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.users[user.UserID]; !exists {
		return appErrors.ErrNotFound
	}
	r.users[user.UserID] = user
	return nil
}

// GetById is a method that gets a user by id
func (r *UserRepository) GetById(id int) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if user, exists := r.users[id]; exists {
		return &user, nil
	}
	return nil, appErrors.ErrNotFound
}

// Delete is a method that deletes a user by id
func (r *UserRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.users[id]; !exists {
		return appErrors.ErrNotFound
	}
	delete(r.users, id)
	return nil
}
