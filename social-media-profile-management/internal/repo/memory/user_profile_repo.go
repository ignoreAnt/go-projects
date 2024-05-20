package memory

import (
	"context"
	"errors"
	"social-server/internal/domain"
	"social-server/pkg/util"
	"sync/atomic"
)

type UserProfileRepository struct {
	*util.SyncMap[int, domain.UserProfile]
	nextID int32
}

func NewUserProfileRepository() *UserProfileRepository {
	repo := &UserProfileRepository{
		nextID: 1,
	}
	generateID := func() int {
		return int(atomic.AddInt32(&repo.nextID, 1))
	}
	repo.SyncMap = util.NewSyncMap[int, domain.UserProfile](generateID)
	return repo
}

func (r *UserProfileRepository) Create(ctx context.Context, profile domain.UserProfile) (domain.UserProfile, error) {
	id, _ := r.SyncMap.Create(profile)
	profile.ProfileID = id
	return profile, nil
}

func (r *UserProfileRepository) GetById(ctx context.Context, id int) (domain.UserProfile, error) {
	profile, exists := r.SyncMap.Retrieve(id)
	if !exists {
		return domain.UserProfile{}, errors.New("profile not found")
	}
	return profile, nil
}

func (r *UserProfileRepository) Update(ctx context.Context, profile domain.UserProfile) (domain.UserProfile, error) {
	_, exists := r.SyncMap.Update(profile.ProfileID, profile)
	if !exists {
		return domain.UserProfile{}, errors.New("profile not found")
	}
	return profile, nil
}

func (r *UserProfileRepository) Delete(ctx context.Context, id int) error {
	if !r.SyncMap.Delete(id) {
		return errors.New("profile not found")
	}
	return nil
}
