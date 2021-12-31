package domain

import (
	"book_store/internal/dto"

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
	GetAllBooks() ([]dto.Book, error)
	GetBookById(id int) (*dto.Book, error)
}

// Secondary Port for the database implementation
type BookRepository interface {
	GetBooks() ([]Book, error)
	GetBook(id int) (*Book, error)
}

// ToDto is func that moves data to a DTO
func (b Book) ToDto() dto.Book {
	return dto.Book{
		ID: b.ID,
		Title: b.Title,
		Authors: b.Authors,
		Year: b.Year,
	}
}