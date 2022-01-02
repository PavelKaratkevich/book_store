package domain

import (
	err "book_store/internal/response"

	"github.com/lib/pq"
)

// Main model
type Book struct {
	ID      int            `json:"ID,omitempty" db:"id"`
	Title   string         `json:"Title" db:"title" binding:"required"`
	Authors pq.StringArray `json:"Authors" db:"authors" binding:"required"`
	Year    string         `json:"Year" db:"year" binding:"required"`
}

// Primary Port for the Service implementation
type Service interface {
	GetAllBooks() ([]Book, *err.AppError)
	GetBookById(id int) (*Book, *err.AppError)
	PostNewBook(req Book) (int, *err.AppError)
}

// Secondary Port for the database implementation
type BookRepository interface {
	GetBooks() ([]Book, *err.AppError)
	GetBook(id int) (*Book, *err.AppError)
	NewBook(req Book) (int, *err.AppError)
	DeleteBook(id int) (int, *err.AppError)
}
