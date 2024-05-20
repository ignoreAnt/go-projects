package memory

import (
	"social-server/internal/domain"
)

type UserRepository struct {
	*GenericRepository[domain.User]
}

func (r *UserRepository) NewRepository() *UserRepository {
	return &UserRepository{
		GenericRepository: NewRepository[domain.User](),
	}
}
