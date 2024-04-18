package database

import (
	"database/sql"
)

// PictureRepository is a struct that defines the repository for the picture
type PictureDbRepository struct {
	db *sql.DB
}

// NewPictureRepository is a function that returns a new PictureRepository
func NewPictureRepository(db *sql.DB) *PictureDbRepository {
	return &PictureDbRepository{
		db: db,
	}
}
