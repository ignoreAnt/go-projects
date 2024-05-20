package memory

import (
	"context"
	"errors"
	"social-server/internal/domain"
	"social-server/pkg/util"
	"sync/atomic"
)

type UserRepository struct {
	*util.SyncMap[int, domain.User]
	nextID int32
}

func NewUserRepository() *UserRepository {
	repo := &UserRepository{
		nextID: 1,
	}
	generateID := func() int {
		return int(atomic.AddInt32(&repo.nextID, 1))
	}
	repo.SyncMap = util.NewSyncMap[int, domain.User](generateID)
	return repo
}

func (r *UserRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	id, _ := r.SyncMap.Create(user)
	user.ID = id
	return user, nil
}

func (r *UserRepository) GetById(ctx context.Context, id int) (domain.User, error) {
	user, exists := r.SyncMap.Retrieve(id)
	if !exists {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, user domain.User) (domain.User, error) {
	_, exists := r.SyncMap.Update(user.ID, user)
	if !exists {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	if !r.SyncMap.Delete(id) {
		return errors.New("user not found")
	}
	return nil
}
