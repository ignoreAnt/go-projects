package repo

import (
	"social-server/internal/domain"
	"social-server/internal/repo/memory"
)

func GetRepository[T domain.Repository[T]](repoType string) *memory.GenericRepository[T] {
	switch repoType {
	case "memory":
		return memory.NewRepository[T]()
	default:
		return nil
	}
}
