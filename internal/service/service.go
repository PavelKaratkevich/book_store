package service

import (
	"book_store/internal/domain"
	err "book_store/internal/response"
)

// BooksService is the implementation (adapter) of Client interface
type BookService struct {
	repo domain.BookRepository
}

// helper function
func NewBookService(repository domain.BookRepository) BookService {
	return BookService{
		repository,
	}
}

func(b BookService) GetAllBooks() ([]domain.Book, *err.AppError) {
	return b.repo.GetBooks()
}

func(b BookService) GetBookById(id int) (*domain.Book, *err.AppError) {
	return b.repo.GetBook(id)
}