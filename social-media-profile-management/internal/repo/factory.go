package repo

import (
	"database/sql"
	"go-projects/social-media-profile-management/internal/domain"
	"go-projects/social-media-profile-management/internal/repo/database"
	"go-projects/social-media-profile-management/internal/repo/memory"
)

type RepoType string

const (
	// MemoryRepo represents memory repository
	MemoryRepo RepoType = "memory"
	// DbRepo represents database repository
	DbRepo RepoType = "db"
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

// GetUserRepository returns a new instance of DataManager for User
func GetUserRepository(repoType RepoType) domain.DataManager[domain.User] {
	switch repoType {
	case MemoryRepo:
		return memory.NewBackend[domain.User]()
	case DbRepo:
		return nil
	}
	return nil
}

// GetUserProfileRepository returns a new instance of DataManager for UserProfile
func GetUserProfileRepository(repoType RepoType) domain.DataManager[domain.UserProfile] {
	switch repoType {
	case MemoryRepo:
		return memory.NewBackend[domain.UserProfile]()
	case DbRepo:
		return nil
	}
	return nil
}

// GetWorkDetailsRepository returns a new instance of DataManager for WorkDetail
func GetWorkDetailsRepository(repoType RepoType) domain.DataManager[domain.WorkDetail] {
	switch repoType {
	case MemoryRepo:
		return memory.NewBackend[domain.WorkDetail]()
	case DbRepo:
		return nil
	}
	return nil
}

// GetPictureRepository returns a new instance of DataManager for Picture
func GetPictureRepository(repoType RepoType) domain.DataManager[domain.Picture] {
	switch repoType {
	case MemoryRepo:
		return memory.NewBackend[domain.Picture]()
	case DbRepo:
		return nil
	}
	return nil
}

// GetEducationDetailsRepository returns a new instance of DataManager for EducationDetail
func GetEducationDetailsRepository(repoType RepoType) domain.DataManager[domain.EducationDetail] {
	switch repoType {
	case MemoryRepo:
		return memory.NewBackend[domain.EducationDetail]()
	case DbRepo:
		return nil
	}
	return nil
}

// GetEducationDetailsRepository returns a new instance of DataManager for EducationDetail
func GetEmploymentDetailsRepository(repoType RepoType) domain.DataManager[domain.EmploymentDetail] {
	switch repoType {
	case MemoryRepo:
		return memory.NewBackend[domain.EmploymentDetail]()
	case DbRepo:
		return nil
	}
	return nil
}

// GetLifeEventsRepository returns a new instance of DataManager for Life Event
func GetLifeEventsRepository(repoType RepoType) domain.DataManager[domain.LifeEvent] {
	switch repoType {
	case MemoryRepo:
		return memory.NewBackend[domain.LifeEvent]()
	case DbRepo:
		return nil
	}
	return nil
}

// GetPrivacyRepository returns a new instance of DataManager for Privacy
func GetPrivacyRepository(repoType RepoType) domain.DataManager[domain.Privacy] {
	switch repoType {
	case MemoryRepo:
		return memory.NewBackend[domain.Privacy]()
	case DbRepo:
		return nil
	}
	return nil
}
