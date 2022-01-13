package domain

import (
	"github.com/gin-gonic/gin"
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
	GetAllBooks(ctx *gin.Context) ([]Book, error)
	GetBookById(ctx *gin.Context, id int) (*Book, error)
	PostNewBook(ctx *gin.Context, req Book) (int, error)
	DeleteBookById(ctx *gin.Context, id int) (int, error)
	UpdateBookById(ctx *gin.Context, req Book) (int, error)
}

// Secondary Port for the database implementation
type BookRepository interface {
	GetBooks(ctx *gin.Context) ([]Book, error)
	GetBook(ctx *gin.Context, id int) (*Book, error)
	NewBook(ctx *gin.Context, req Book) (int, error)
	DeleteBook(ctx *gin.Context, id int) (int, error)
	UpdateBook(ctx *gin.Context, req Book) (int, error)
}
