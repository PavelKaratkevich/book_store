package handler

import (
	"book_store/internal/domain"
	"book_store/internal/domain/mocks"
	"book_store/internal/middleware"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func Test_if_GetAllBook_Handler_Returns_Error_400_StatusBadRequest(t *testing.T) {

	// Arrange
	mock := new(mocks.Service)
	router := gin.Default()

	rr := httptest.NewRecorder()

	router.Use(middleware.Logger())
	router.Use(middleware.CORS())
	router.Use(middleware.CheckToken())

	mock.On("GetAllBooks").Return(nil, gin.H{"error": "Please provide a valid token"})

	// Act
	request, err := http.NewRequest(http.MethodGet, "/books/", nil)
	assert.NoError(t, err)

	router.ServeHTTP(rr, request)

	// Assert
	assert.Equal(t, 400, rr.Code)
}

func Test_if_GetAllBook_Handler_Returns_Books(t *testing.T) {

	// Arrange
	books := []domain.Book{
		{
			Authors: pq.StringArray{"Mr Sean"},
			Title:   "Test",
			Year:    "2022-11-01T00:00:00Z",
		},
		{
			Authors: pq.StringArray{"Mr Paul"},
			Title:   "Test2",
			Year:    "2022-11-01T00:00:00Z",
		},
	}

	expected, err := json.Marshal(books)
	assert.NoError(t, err)

	mock := new(mocks.Service)
	mock.On("GetAllBooks").Return(books, nil)

	rr := httptest.NewRecorder()
	router := gin.Default()
	
	request, err := http.NewRequest(http.MethodPost, "/books/", nil)
	assert.NoError(t, err)

	router.ServeHTTP(rr, request)
	assert.JSONEq(t, string(expected), rr.Body.String())
}
