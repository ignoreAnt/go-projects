package domain

import "context"

// Repository is a generic interface for CRUD operations on an entity type T.
type Repository[T any] interface {
	Create(ctx context.Context, t T) (T, error)
	GetById(ctx context.Context, id int) (T, error)
	Update(ctx context.Context, t T) (T, error)
	Delete(ctx context.Context, id int) error
}

type RepositoryCreator[R any] interface {
	NewRepository() R
}

type ServiceCreator[T any, S any] interface {
	NewService(repository Repository[T]) S
}

// UserRepository User specific repository interface
type UserRepository interface {
	Repository[User]
}

// UserProfileRepository UserProfile specific repository interface
type UserProfileRepository interface {
	Repository[UserProfile]
}

// PictureRepository Picture specific repository interface
type PictureRepository interface {
	Repository[Picture]
}

// WorkDetailRepository WorkDetail specific repository interface
type WorkDetailRepository interface {
	Repository[WorkDetail]
}

// EducationDetailRepository EducationDetail specific repository interface
type EducationDetailRepository interface {
	Repository[EducationDetail]
}

// LifeEventRepository LifeEvent specific repository interface
type LifeEventRepository interface {
	Repository[LifeEvent]
}

type EmploymentDetailRepository interface {
	Repository[EmploymentDetail]
}
