package domain

import (
	"context"

	"github.com/lib/pq"
)

// Main model
type Book struct {
	ID      int            `json:"ID,omitempty" db:"id"`
	Title   string         `json:"Title" db:"title" binding:"required" validate:"required"`
	Authors pq.StringArray `json:"Authors" db:"authors" binding:"required" validate:"required"`
	Year    string         `json:"Year" db:"year" binding:"required" validate:"required,datetime=2006-01-02"`
}

// Primary Port for the Service/Use Case implementation
type Service interface {
	GetAllBooks(ctx context.Context) ([]Book, error)
	GetBookById(ctx context.Context, id int) (*Book, error)
	PostNewBook(ctx context.Context, req Book) (int, error)
	DeleteBookById(ctx context.Context, id int) (int, error)
	UpdateBookById(ctx context.Context, req Book) (int, error)
}

// Secondary Port for the database implementation
type BookRepository interface {
	GetBooks(ctx context.Context) ([]Book, error)
	GetBook(ctx context.Context, id int) (*Book, error)
	NewBook(ctx context.Context, req Book) (int, error)
	DeleteBook(ctx context.Context, id int) (int, error)
	UpdateBook(ctx context.Context, req Book) (int, error)
}
