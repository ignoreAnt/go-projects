package memory

import (
	"social-server/internal/domain"
)

type WorkDetailRepository struct {
	*GenericRepository[domain.WorkDetail]
}

func (r *WorkDetailRepository) NewRepository() *WorkDetailRepository {
	return &WorkDetailRepository{
		GenericRepository: NewRepository[domain.WorkDetail](),
	}
}
