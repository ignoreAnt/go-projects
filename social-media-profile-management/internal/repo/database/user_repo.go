package database

import (
	"database/sql"
	"errors"
	"go-projects/social-media-profile-management/internal/domain"
	app_errors "go-projects/social-media-profile-management/internal/errors"
)

// UserRepository is a struct that defines the repository for the user
type UserRepository struct {
	db *sql.DB
}

// UpdateUserName updates the username for the userID
func (r *UserRepository) UpdateUserName(userID int, newUserName string) error {
	query := `UPDATE users SET username = ? WHERE id = ?`
	_, err := r.db.Exec(query, newUserName, userID)
	return err
}

// NewUserRepository is a function that returns a new UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create is a method that creates a new user
func (r *UserRepository) Create(user domain.User) error {
	query := `INSERT INTO users (username) VALUES (?) RETURNING id`
	return r.db.QueryRow(query, user.UserName).Scan(&user.UserID)
}

// Update is a method that updates a user
func (r *UserRepository) Update(user domain.User) error {
	query := `UPDATE users SET username = ? WHERE id = ?`
	_, err := r.db.Exec(query, user.UserName, user.UserID)
	return err
}

// GetById is a method that gets a user by id
func (r *UserRepository) GetById(id int) (*domain.User, error) {
	user := domain.User{}
	query := `SELECT id, username FROM users WHERE id = ?`
	err := r.db.QueryRow(query, id).Scan(&user.UserID, &user.UserName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, app_errors.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

// Delete is a method that deletes a user by id
func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
