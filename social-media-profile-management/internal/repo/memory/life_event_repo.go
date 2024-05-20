package memory

import (
	"social-server/internal/domain"
)

type LifeEventRepository struct {
	*GenericRepository[domain.LifeEvent]
}

func (r *LifeEventRepository) NewRepository() *LifeEventRepository {
	return &LifeEventRepository{
		GenericRepository: NewRepository[domain.LifeEvent](),
	}
}
