package memory

import (
	"social-server/internal/domain"
	appErrors "social-server/internal/errors"
	"sync"
)

// WorkDetailsRepository is a struct that defines the repository for the work details
type WorkDetailsRepository struct {
	workDetails map[int]domain.WorkDetail
	mu          sync.RWMutex
	nextID      int
}

// NewWorkDetailsRepository Creates new domain.WorkDetail repository
func NewWorkDetailsRepository() *WorkDetailsRepository {
	return &WorkDetailsRepository{
		workDetails: make(map[int]domain.WorkDetail),
		nextID:      1,
	}
}

// Create Work Details Repository
func (w *WorkDetailsRepository) Create(workDetail domain.WorkDetail) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	id := w.nextID
	workDetail.WorkDetailID = id
	w.nextID++
	w.workDetails[workDetail.WorkDetailID] = workDetail

	return nil
}

// Update the workdetail repository
func (w *WorkDetailsRepository) Update(workDetail domain.WorkDetail) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if _, exists := w.workDetails[workDetail.WorkDetailID]; !exists {
		return appErrors.ErrNotFound
	}

	w.workDetails[workDetail.WorkDetailID] = workDetail

	return nil
}

// Delete the workdetail from repository
func (w *WorkDetailsRepository) Delete(workDetailsID int) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if _, exists := w.workDetails[workDetailsID]; !exists {
		return appErrors.ErrNotFound
	}

	delete(w.workDetails, workDetailsID)

	return nil
}

func (w *WorkDetailsRepository) GetWorkDetailByid(workDetailsID int) (*domain.WorkDetail, error) {
	w.mu.RLock()
	defer w.mu.Unlock()

	if _, exists := w.workDetails[workDetailsID]; !exists {
		return nil, appErrors.ErrNotFound
	}

	wd := w.workDetails[workDetailsID]
	return &wd, nil
}
