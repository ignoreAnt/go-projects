package service

import "errors"

type Borrower struct {
	ID      int
	Name    string
	Email   string
	Phone   string
	Address string
}

// AddBorrower adds a new borrower to the library
func (b *Borrower) AddBorrower(name, email, phone, address string) error {

	newBorrower := Borrower{
		ID:      generateUniqueID(BorrowerConstant),
		Name:    name,
		Email:   email,
		Phone:   phone,
		Address: address,
	}

	library.Borrowers = append(library.Borrowers, newBorrower)

	return nil
}

// UpdateBorrower updates the details of an existing borrower
func (b *Borrower) UpdateBorrower(id int, name, email, phone, address string) error {
	for i, borrower := range library.Borrowers {
		if borrower.ID == id {
			library.Borrowers[i].Name = name
			library.Borrowers[i].Email = email
			library.Borrowers[i].Phone = phone
			library.Borrowers[i].Address = address

			return nil
		}
	}

	return errors.New("borrower not found")
}

// DeleteBorrower deletes a borrower from the library
func (b *Borrower) DeleteBorrower(id int) error {
	for i, borrower := range library.Borrowers {
		if borrower.ID == id {
			library.Borrowers = append(library.Borrowers[:i], library.Borrowers[i+1:]...)

			return nil
		}
	}

	return errors.New("borrower not found")
}

// FindBorrower finds a borrower from the library
func (b *Borrower) FindBorrower(id int) (Borrower, error) {
	for _, borrower := range library.Borrowers {
		if borrower.ID == id {
			return borrower, nil
		}
	}

	return Borrower{}, errors.New("borrower not found")
}

// ListBorrowers lists all borrowers in the library
func (b *Borrower) ListBorrowers() []Borrower {
	return library.Borrowers
}

//

// ReturnBook returns a book to the library
func (b *Borrower) ReturnBook(borrowerID, bookID int) error {
	borrower, err := b.FindBorrower(borrowerID)
	if err != nil {
		return err
	}
	library.BorrowedBooks[borrower] = library.BorrowedBooks[borrower][:0]

	b2 := Book{}
	book, err := b2.GetBookById(bookID)
	if err != nil {
		return err
	}

	book.Quantity++

	return nil

}

// BorrowBook borrows a book from the library
func (b *Borrower) BorrowBook(borrowerID, bookID int) error {
	borrower, err := b.FindBorrower(borrowerID)
	if err != nil {
		return err
	}

	b2 := Book{}
	book, err := b2.GetBookById(bookID)
	if err != nil {
		return err
	}

	if book.Quantity == 0 {
		return errors.New("book not available")
	}
	library.BorrowedBooks[borrower] = append(library.BorrowedBooks[borrower], *book)

	book.Quantity--

	return nil
}
