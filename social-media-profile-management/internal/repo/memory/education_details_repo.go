package memory

import (
	"social-server/internal/domain"
)

type EducationDetailRepository struct {
	*GenericRepository[domain.EducationDetail]
}

func (r *EducationDetailRepository) NewRepository() *EducationDetailRepository {
	return &EducationDetailRepository{
		GenericRepository: NewRepository[domain.EducationDetail](),
	}
}
