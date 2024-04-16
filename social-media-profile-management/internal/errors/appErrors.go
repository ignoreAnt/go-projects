package appErrors

import "errors"

// ErrNotFound Define custom errors as needed
var ErrNotFound = errors.New("not found")

// ErrConflict Define custom errors as needed
var ErrConflict = errors.New("conflict")

// ErrInvalidInput ErrInternal Define custom errors as needed
var ErrInvalidInput = errors.New("invalid input")

// NewValidationError Define custom errors as needed
func NewValidationError(message string) error {
	return errors.New("validation error: " + message)
}
