package bookHTTP

import (
	"book_store/internal/domain/mocks"
	"book_store/internal/domain"
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func Test_if_GetAllBooks_Handler_Returns_Books(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	books := []domain.Book{
		{
			ID:      1,
			Title:   "Check",
			Authors: pq.StringArray{"Mr Bean"},
			Year:    "1999-07-25T00:00:00Z",
		},
		{
			ID:      2,
			Title:   "Check",
			Authors: pq.StringArray{"Mr Bean"},
			Year:    "1999-07-25T00:00:00Z",
		},
	}
	expected, err := json.Marshal(books)
	assert.NoError(t, err)

	router := gin.Default()

	mock := new(mocks.Service)
	routes := router.Group("/api")

	RegisterBooksEndpoints(routes, mock)

	mock.On("GetAllBooks", context.Background()).Return(books, nil)

	rr := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/api/books", nil)
	assert.NoError(t, err)

	router.ServeHTTP(rr, request)

	log.Printf("Output of the test: %v %v", rr.Code, rr.Body)

	assert.JSONEq(t, string(expected), string(rr.Body.Bytes()))
	assert.Equal(t, 200, rr.Code)
	mock.AssertExpectations(t)
}

func Test_if_GetBookbyId_Handler_Returns_Book(t *testing.T) {
	gin.SetMode(gin.TestMode)

	book := &domain.Book{
		ID:      1,
		Title:   "Test",
		Authors: pq.StringArray{"Test"},
		Year:    "2020-01-2T00:00:00Z",
	}

	expected, err := json.Marshal(book)
	assert.NoError(t, err)

	mock := new(mocks.Service)
	r := gin.Default()

	api := r.Group("/api")
	RegisterBooksEndpoints(api, mock)

	mock.On("GetBookById", context.Background(), book.ID).Return(book, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/books/1", nil)
	r.ServeHTTP(w, req)
	actual := w.Body.Bytes()

	log.Printf("Output of the test: %v %v", w.Code, w.Body)

	assert.Equal(t, expected, actual)
	assert.Equal(t, 200, w.Code)
	mock.AssertExpectations(t)
}

func Test_if_PostNewBook_Handler_Returns_int(t *testing.T) {
	gin.SetMode(gin.TestMode)

	book := domain.Book{
		Title:   "Test",
		Authors: pq.StringArray{"Test"},
		Year:    "2020-01-21",
	}

	output, err := json.Marshal(book)
	assert.NoError(t, err)

	mock := new(mocks.Service)
	r := gin.Default()

	api := r.Group("/api")
	RegisterBooksEndpoints(api, mock)

	expected, err2 := json.Marshal(1)
	assert.NoError(t, err2)

	mock.On("PostNewBook", context.Background(), book).Return(1, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/books", bytes.NewBuffer(output))

	r.ServeHTTP(w, req)
	actual := w.Body.Bytes()

	log.Printf("Output of the test: %v %v", w.Code, w.Body)

	assert.Equal(t, expected, actual)
	assert.Equal(t, 200, w.Code)
	mock.AssertExpectations(t)
}


func Test_if_DeleteBookById_Handler_Returns_Book(t *testing.T) {
	gin.SetMode(gin.TestMode)

	bookID := 1

	mock := new(mocks.Service)
	r := gin.Default()

	api := r.Group("/api")
	RegisterBooksEndpoints(api, mock)

	mock.On("DeleteBookById", context.Background(), bookID).Return(1, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/books/1", nil)
	
	r.ServeHTTP(w, req)

	log.Printf("Output of the test: %v %v %T", w.Code, w.Body, w.Body)

	assert.Equal(t, 200, w.Code)
	mock.AssertExpectations(t)
}

func Test_if_UpdateBookById_Handler_Returns_int(t *testing.T) {
	gin.SetMode(gin.TestMode)

	book := domain.Book{
		ID: 1,
		Title:   "Test",
		Authors: pq.StringArray{"Test"},
		Year:    "2020-01-21",
	}

	output, err := json.Marshal(book)
	assert.NoError(t, err)

	mock := new(mocks.Service)
	r := gin.Default()

	api := r.Group("/api")
	RegisterBooksEndpoints(api, mock)
	
	mock.On("UpdateBookById", context.Background(), book).Return(1, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/books/1", bytes.NewBuffer(output))

	r.ServeHTTP(w, req)

	log.Printf("Output of the test: %v %v", w.Code, w.Body)

	assert.Equal(t, 200, w.Code)
	mock.AssertExpectations(t)
}