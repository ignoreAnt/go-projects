package repo

import (
	"database/sql"
	"go-projects/social-media-profile-management/internal/domain"
	"go-projects/social-media-profile-management/internal/repo/database"
	"go-projects/social-media-profile-management/internal/repo/memory"
)

// UserRepositoryFactory is a function that returns a new UserRepository
func UserRepositoryFactory(useMemory bool, db *sql.DB) domain.UserRepository {
	if useMemory {
		return memory.NewUserRepository()
	}
	return database.NewUserRepository(db)
}

func WorkDetailsRepositoryFactory(useMemory bool, db *sql.DB) domain.WorkDetailsRepository {
	if useMemory {
		return memory.NewWorkDetailsRepository()
	}

	return nil

}
