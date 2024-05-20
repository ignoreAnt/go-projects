package memory

import (
	"context"
	"errors"
	"social-server/internal/domain"
	"social-server/pkg/util"
	"sync/atomic"
)

type EmploymentDetailRepository struct {
	*util.SyncMap[int, domain.EmploymentDetail]
	nextID int32
}

func NewEmploymentDetailRepository() *EmploymentDetailRepository {
	repo := &EmploymentDetailRepository{
		nextID: 1,
	}
	generateID := func() int {
		return int(atomic.AddInt32(&repo.nextID, 1))
	}
	repo.SyncMap = util.NewSyncMap[int, domain.EmploymentDetail](generateID)
	return repo
}

func (r *EmploymentDetailRepository) Create(ctx context.Context, detail domain.EmploymentDetail) (domain.EmploymentDetail, error) {
	id, _ := r.SyncMap.Create(detail)
	detail.EmploymentDetailID = id
	return detail, nil
}

func (r *EmploymentDetailRepository) GetById(ctx context.Context, id int) (domain.EmploymentDetail, error) {
	detail, exists := r.SyncMap.Retrieve(id)
	if !exists {
		return domain.EmploymentDetail{}, errors.New("employment detail not found")
	}
	return detail, nil
}

func (r *EmploymentDetailRepository) Update(ctx context.Context, detail domain.EmploymentDetail) (domain.EmploymentDetail, error) {
	_, exists := r.SyncMap.Update(detail.EmploymentDetailID, detail)
	if !exists {
		return domain.EmploymentDetail{}, errors.New("employment detail not found")
	}
	return detail, nil
}

func (r *EmploymentDetailRepository) Delete(ctx context.Context, id int) error {
	if !r.SyncMap.Delete(id) {
		return errors.New("employment detail not found")
	}
	return nil
}
