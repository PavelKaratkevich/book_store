package service

import (
	"book_store/internal/domain"
	"book_store/internal/dto"
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

func(b BookService) GetAllBooks() ([]dto.Book, error) {
	var dto []dto.Book
	books, err := b.repo.GetBooks()
		if err != nil {
			log.Printf("Error while getting books from DB repo: %v", err)
			return nil, err
		}
	for _, j := range books {
		dto = append(dto, j.ToDto())
	}
	return dto, nil	
}

func(b BookService) GetBookById(id int) (*dto.Book, error) {
	var dto dto.Book
	book, err := b.repo.GetBook(id)
		if err != nil {
			log.Printf("Error while getting book by ID from DB repo: %v", err)
			return nil, err
		}
		dto = book.ToDto()
	return &dto, nil	
}

