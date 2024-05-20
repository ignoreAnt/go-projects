package memory

import (
	"context"
	"errors"
	"social-server/internal/domain"
	"social-server/pkg/util"
	"sync/atomic"
)

type PictureRepository struct {
	*util.SyncMap[int, domain.Picture]
	nextID int32
}

func NewPictureRepository() *PictureRepository {
	repo := &PictureRepository{
		nextID: 1,
	}
	generateID := func() int {
		return int(atomic.AddInt32(&repo.nextID, 1))
	}
	repo.SyncMap = util.NewSyncMap[int, domain.Picture](generateID)
	return repo
}

func (r *PictureRepository) Create(ctx context.Context, picture domain.Picture) (domain.Picture, error) {
	id, _ := r.SyncMap.Create(picture)
	picture.PictureID = id
	return picture, nil
}

func (r *PictureRepository) GetById(ctx context.Context, id int) (domain.Picture, error) {
	picture, exists := r.SyncMap.Retrieve(id)
	if !exists {
		return domain.Picture{}, errors.New("picture not found")
	}
	return picture, nil
}

func (r *PictureRepository) Update(ctx context.Context, picture domain.Picture) (domain.Picture, error) {
	_, exists := r.SyncMap.Update(picture.PictureID, picture)
	if !exists {
		return domain.Picture{}, errors.New("picture not found")
	}
	return picture, nil
}

func (r *PictureRepository) Delete(ctx context.Context, id int) error {
	if !r.SyncMap.Delete(id) {
		return errors.New("picture not found")
	}
	return nil
}
