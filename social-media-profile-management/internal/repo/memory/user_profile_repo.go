package memory

import (
	"social-server/internal/domain"
)

type UserProfileRepository struct {
	*GenericRepository[domain.UserProfile]
}

func (r *UserProfileRepository) NewRepository() *UserProfileRepository {
	return &UserProfileRepository{
		GenericRepository: NewRepository[domain.UserProfile](),
	}
}
