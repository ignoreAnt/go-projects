package repo

import (
	"social-server/internal/repo/memory"
)

func GetRepository[T any](repoType string) *memory.GenericRepository[T] {
	switch repoType {
	case "memory":
		return memory.NewRepository[T]()
	default:
		return nil
	}
}
