package service

import "errors"

type Book struct {
	ID       int
	Title    string
	Author   string
	ISBN     string
	Genre    string
	Quantity int
}

type SearchFunc func(book Book) bool

// AddBook adds a new book to library
func (b *Book) AddBook(title, author, isbn, genre string, quantity int) error {

	// Generate a unique Id for the new book
	// Create a new Book instance with the provided details
	newBook := Book{
		ID:       generateUniqueID("book"),
		Title:    title,
		Author:   author,
		ISBN:     isbn,
		Genre:    genre,
		Quantity: quantity,
	}

	// Add the new book to the library's collection of books
	library.Books = append(library.Books, newBook)

	return nil
}

// UpdateBook updates the details of an existing book
func (b *Book) UpdateBook(id int, title, author, isbn, genre string, quantitity int) error {
	for i, book := range library.Books {
		if book.ID == id {
			// Update the details with provided values
			library.Books[i].Title = title
			library.Books[i].Author = author
			library.Books[i].ISBN = isbn
			library.Books[i].Genre = genre
			library.Books[i].Quantity = quantitity

			return nil
		}
	}

	return errors.New("book not found")
}

// DeleteBook deletes a book from the library
func (b *Book) DeleteBook(id int) error {
	// Find the book with given ID in the library's Collection
	for i, book := range library.Books {
		if book.ID == id {
			// Remove the book from the library's collection
			library.Books = append(library.Books[:i], library.Books[i+1:]...)
			return nil
		}
	}

	return errors.New("book not found")
}

// GetBookById retrieves a book from the library by its ID
func (b *Book) GetBookById(id int) (*Book, error) {
	for _, book := range library.Books {
		if book.ID == id {
			return &book, nil
		}
	}

	return nil, errors.New("book not found")
}

// SearchBooks searches for books in the library based on the provide search function
func (b *Book) SearchBooks(searchFunc SearchFunc) []Book {
	var matchingBooks []Book

	for _, book := range library.Books {
		if searchFunc(book) {
			matchingBooks = append(matchingBooks, book)
		}
	}

	return matchingBooks
}
