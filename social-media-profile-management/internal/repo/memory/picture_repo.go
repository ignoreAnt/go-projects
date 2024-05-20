package memory

import (
	"social-server/internal/domain"
)

type PictureRepository struct {
	*GenericRepository[domain.Picture]
}

func (r *PictureRepository) NewRepository() *PictureRepository {
	return &PictureRepository{
		GenericRepository: NewRepository[domain.Picture](),
	}
}
