package memory

import (
	"context"
	"errors"
	"reflect"
	"social-server/internal/domain"
	"social-server/pkg/util"
	"sync/atomic"
)

// RepositoryFactory is a generic factory for creating repositories.
type RepositoryFactory[T domain.Repository[U], U any] struct {
	nextID int32
}

// GenericRepository is a generic implementation of the Repository interface.
type GenericRepository[U any] struct {
	*util.SyncMap[int, U]
	nextID int32
}

// NewRepository creates a new GenericRepository.
func NewRepository[U any]() *GenericRepository[U] {
	repo := &GenericRepository[U]{
		nextID: 1,
	}
	generateID := func() int {
		return int(atomic.AddInt32(&repo.nextID, 1))
	}
	repo.SyncMap = util.NewSyncMap[int, U](generateID)
	return repo
}

func (r *GenericRepository[U]) Create(ctx context.Context, entity U) (U, error) {
	_, _ = r.SyncMap.Create(entity)
	return entity, nil
}

func (r *GenericRepository[U]) GetById(ctx context.Context, id int) (U, error) {
	entity, exists := r.SyncMap.Retrieve(id)
	if !exists {
		var zero U
		return zero, errors.New("entity not found")
	}
	return entity, nil
}

func (r *GenericRepository[U]) Update(ctx context.Context, entity U) (U, error) {
	id := reflect.ValueOf(entity).FieldByName("ID").Interface().(int)
	_, exists := r.SyncMap.Update(id, entity)
	if !exists {
		var zero U
		return zero, errors.New("entity not found")
	}
	return entity, nil
}

func (r *GenericRepository[U]) Delete(ctx context.Context, id int) error {
	if !r.SyncMap.Delete(id) {
		return errors.New("entity not found")
	}
	return nil
}
