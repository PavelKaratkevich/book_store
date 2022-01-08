package service

import (
	"book_store/internal/domain"
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

func(b BookService) GetAllBooks() ([]domain.Book, *error) {
	return b.repo.GetBooks()
}

func(b BookService) GetBookById(id int) (*domain.Book, *error) {
	return b.repo.GetBook(id)
}

func(b BookService) PostNewBook(req domain.Book) (int, *error) {
	return b.repo.NewBook(req)
}

func(b BookService) DeleteBookById(id int) (int, *error) {
	return b.repo.DeleteBook(id)
}

func(b BookService) UpdateBookById(req domain.Book) (int, *error) {
	return b.repo.UpdateBook(req)
}