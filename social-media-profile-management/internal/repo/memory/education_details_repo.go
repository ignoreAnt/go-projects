package memory

import (
	"context"
	"errors"
	"social-server/internal/domain"
	"social-server/pkg/util"
	"sync/atomic"
)

type EducationDetailRepository struct {
	*util.SyncMap[int, domain.EducationDetail]
	nextID int32
}

func NewEducationDetailRepository() *EducationDetailRepository {
	repo := &EducationDetailRepository{
		nextID: 1,
	}
	generateID := func() int {
		return int(atomic.AddInt32(&repo.nextID, 1))
	}
	repo.SyncMap = util.NewSyncMap[int, domain.EducationDetail](generateID)
	return repo
}

func (r *EducationDetailRepository) Create(ctx context.Context, detail domain.EducationDetail) (domain.EducationDetail, error) {
	id, _ := r.SyncMap.Create(detail)
	detail.EducationDetailID = id
	return detail, nil
}

func (r *EducationDetailRepository) GetById(ctx context.Context, id int) (domain.EducationDetail, error) {
	detail, exists := r.SyncMap.Retrieve(id)
	if !exists {
		return domain.EducationDetail{}, errors.New("education detail not found")
	}
	return detail, nil
}

func (r *EducationDetailRepository) Update(ctx context.Context, detail domain.EducationDetail) (domain.EducationDetail, error) {
	_, exists := r.SyncMap.Update(detail.EducationDetailID, detail)
	if !exists {
		return domain.EducationDetail{}, errors.New("education detail not found")
	}
	return detail, nil
}

func (r *EducationDetailRepository) Delete(ctx context.Context, id int) error {
	if !r.SyncMap.Delete(id) {
		return errors.New("education detail not found")
	}
	return nil
}
