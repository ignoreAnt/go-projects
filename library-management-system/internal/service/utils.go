package service

import "sync"

const BookConstant = "book"
const BorrowerConstant = "borrower"

var (
	bookIdCounter     int
	borrowerIDCounter int
	idMutex           sync.Mutex
)

func generateUniqueID(entityType string) int {
	idMutex.Lock()
	defer idMutex.Unlock()

	var idCounter *int

	switch entityType {
	case BookConstant:
		idCounter = &bookIdCounter
	case BorrowerConstant:
		idCounter = &borrowerIDCounter
	default:
		idCounter = &bookIdCounter

	}

	*idCounter++

	return *idCounter
}
