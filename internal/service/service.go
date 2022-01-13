package service

import (
	"book_store/internal/domain"

	"github.com/gin-gonic/gin"
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

func(b BookService) GetAllBooks(ctx *gin.Context) ([]domain.Book, error) {
	return b.repo.GetBooks(ctx)
}

func(b BookService) GetBookById(ctx *gin.Context, id int) (*domain.Book, error) {
	return b.repo.GetBook(ctx, id)
}

func(b BookService) PostNewBook(ctx *gin.Context, req domain.Book) (int, error) {
	return b.repo.NewBook(ctx, req)
}

func(b BookService) DeleteBookById(ctx *gin.Context, id int) (int, error) {
	return b.repo.DeleteBook(ctx, id)
}

func(b BookService) UpdateBookById(ctx *gin.Context, req domain.Book) (int, error) {
	return b.repo.UpdateBook(ctx, req)
}