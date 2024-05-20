package service

import (
	"context"
	"social-server/internal/domain"
)

type GenericService[T any] struct {
	repo domain.Repository[T]
}

func NewGenericService[T any](repo domain.Repository[T]) *GenericService[T] {
	return &GenericService[T]{repo: repo}
}

func (s *GenericService[T]) Create(ctx context.Context, t T) (T, error) {
	return s.repo.Create(ctx, t)
}

func (s *GenericService[T]) GetById(ctx context.Context, id int) (T, error) {
	return s.repo.GetById(ctx, id)
}

func (s *GenericService[T]) Update(ctx context.Context, t T) (T, error) {
	return s.repo.Update(ctx, t)
}

func (s *GenericService[T]) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
