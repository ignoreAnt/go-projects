package memory

import (
	"social-server/internal/domain"
)

type EmploymentDetailRepository struct {
	*GenericRepository[domain.EmploymentDetail]
}

func (e *EmploymentDetailRepository) NewRepository() *EmploymentDetailRepository {
	return &EmploymentDetailRepository{
		GenericRepository: NewRepository[domain.EmploymentDetail](),
	}
}
