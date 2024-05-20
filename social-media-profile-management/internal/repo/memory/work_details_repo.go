package memory

import (
	"context"
	"errors"
	"social-server/internal/domain"
	"social-server/pkg/util"
	"sync/atomic"
)

type WorkDetailRepository struct {
	*util.SyncMap[int, domain.WorkDetail]
	nextID int32
}

func NewWorkDetailRepository() *WorkDetailRepository {
	repo := &WorkDetailRepository{
		nextID: 1,
	}
	generateID := func() int {
		return int(atomic.AddInt32(&repo.nextID, 1))
	}
	repo.SyncMap = util.NewSyncMap[int, domain.WorkDetail](generateID)
	return repo
}

func (r *WorkDetailRepository) Create(ctx context.Context, detail domain.WorkDetail) (domain.WorkDetail, error) {
	id, _ := r.SyncMap.Create(detail)
	detail.WorkDetailID = id
	return detail, nil
}

func (r *WorkDetailRepository) GetById(ctx context.Context, id int) (domain.WorkDetail, error) {
	detail, exists := r.SyncMap.Retrieve(id)
	if !exists {
		return domain.WorkDetail{}, errors.New("work detail not found")
	}
	return detail, nil
}

func (r *WorkDetailRepository) Update(ctx context.Context, detail domain.WorkDetail) (domain.WorkDetail, error) {
	_, exists := r.SyncMap.Update(detail.WorkDetailID, detail)
	if !exists {
		return domain.WorkDetail{}, errors.New("work detail not found")
	}
	return detail, nil
}

func (r *WorkDetailRepository) Delete(ctx context.Context, id int) error {
	if !r.SyncMap.Delete(id) {
		return errors.New("work detail not found")
	}
	return nil
}
