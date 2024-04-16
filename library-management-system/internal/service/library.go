package service

var library Library

type Library struct {
	Books     []Book
	Borrowers []Borrower
	BorrowedBooks map[Borrower][]Book
}

func init() {
	library = Library{
		Books:     []Book{},
		Borrowers: []Borrower{},
	}
}

// GetLibrary returns the library instance
func GetLibrary() Library {
	return library
}

// ListBooks lists all books in the library
func ListBooks() []Book {
	return library.Books
}

// ListBorrowers lists all borrowers in the library
func ListBorrowers() []Borrower {
	return library.Borrowers
}

// ListBorrowedBooks lists all books borrowed by a borrower
func ListBorrowedBooks(borrower Borrower) []Book {
	return library.BorrowedBooks[borrower]
}


