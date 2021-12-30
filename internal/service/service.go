package service

import (
	"book_store/internal/domain"
	"book_store/internal/dto"
	"context"
	"log"
)

// BooksService is the implementation (adapter) of UseCase interface
type BookService struct {
	repo domain.BookRepository
}

// helper function
func NewBookService(repository domain.BookRepository) BookService {
	return BookService{
		repository,
	}
}

func(b BookService) GetAllBooks(ctx context.Context) ([]dto.Book, error) {
	var dto []dto.Book

	books, err := b.repo.GetBooks()
		if err != nil {
			log.Printf("Error while getting []books from DB repo: %v", err)
			return nil, err
		}
	for _, j := range books {
		dto = append(dto, j.ToDto())
	}
	return dto, nil	
}

