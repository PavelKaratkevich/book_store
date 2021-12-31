package domain

import (
	err "book_store/internal/response"

	"github.com/lib/pq"
)

// Main model
type Book struct {
	ID      int            `json:"ID" db:"id"`
	Title   string         `json:"Title" db:"title"`
	Authors pq.StringArray `json:"Authors" db:"authors"`
	Year    string         `json:"Year" db:"year"`
}

// Primary Port for the Service implementation
type Service interface {
	GetAllBooks() ([]Book, *err.AppError)
	GetBookById(id int) (*Book, *err.AppError)
}

// Secondary Port for the database implementation
type BookRepository interface {
	GetBooks() ([]Book, *err.AppError)
	GetBook(id int) (*Book, *err.AppError)
}