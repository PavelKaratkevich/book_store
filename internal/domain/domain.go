package domain

import (
	// err "book_store/internal/response"

	"github.com/lib/pq"
)

// Main model
type Book struct {
	ID      int            `json:"ID,omitempty" db:"id"`
	Title   string         `json:"Title" db:"title" binding:"required"`
	Authors pq.StringArray `json:"Authors" db:"authors" binding:"required"`
	Year    string         `json:"Year" db:"year" binding:"required"`
}

// Primary Port for the Service/Use Case implementation
type Service interface {
	GetAllBooks() ([]Book, *error)
	GetBookById(id int) (*Book, *error)
	PostNewBook(req Book) (int, *error)
	DeleteBookById(id int) (int, *error)
	UpdateBookById(req Book) (int, *error)
}

// Secondary Port for the database implementation
type BookRepository interface {
	GetBooks() ([]Book, *error)
	GetBook(id int) (*Book, *error)
	NewBook(req Book) (int, *error)
	DeleteBook(id int) (int, *error)
	UpdateBook(req Book) (int, *error)
}
