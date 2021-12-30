package domain

import (
	"book_store/internal/dto"

	"github.com/lib/pq"
)

type Book struct {
	ID      int            `json:"ID" db:"id"`
	Title   string         `json:"Title" db:"title"`
	Authors pq.StringArray `json:"Authors" db:"authors"`
	Year    string         `json:"Year" db:"year"`
}

// Primary Port for the Service implementation
type Service interface {
	GetAllBooks() ([]Book, error)
}

// Secondary Port for the database implementation
type BookRepository interface {
	GetBooks() ([]Book, error)
}

// Move data to DTO
func (b Book) ToDto() dto.Book {
	return dto.Book{
		ID: b.ID,
		Title: b.Title,
		Authors: b.Authors,
		Year: b.Year,
	}
}