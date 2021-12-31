package service

import (
	"book_store/internal/domain"
	"log"
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

func(b BookService) GetAllBooks() ([]domain.Book, error) {
	books, err := b.repo.GetBooks()
		if err != nil {
			log.Printf("Error while getting books from DB repo: %v", err)
			return nil, err
		}
	
	return books, nil	
}

func(b BookService) GetBookById(id int) (*domain.Book, error) {
	book, err := b.repo.GetBook(id)
		if err != nil {
			log.Printf("Error while getting book by ID from DB repo: %v", err)
			return nil, err
		}
	return book, nil	
}