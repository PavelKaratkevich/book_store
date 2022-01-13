package handler

import (
	"book_store/internal/domain"
	"book_store/internal/domain/mocks"
	// "book_store/internal/injection"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func Test_if_GetAllBook_Handler_Returns_Error_400_StatusBadRequest(t *testing.T) {

	gin.SetMode(gin.TestMode)

	var err error

	// Arrange
	mock := new(mocks.Service)
	router := gin.Default()

	rr := httptest.NewRecorder()

	mock.On("GetAllBooks").Return(nil, err)

	// Act	
	request, err := http.NewRequest(http.MethodGet, "/books", nil)
	assert.NoError(t, err)

	router.ServeHTTP(rr, request)

	// Assert
	assert.Equal(t, 404, rr.Code)

	log.Printf("Output of the test: %v %v", rr.Code, rr.Body)
}

func Test_if_GetAllBook_Handler_Returns_Books(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	books := []domain.Book{
		{
			ID:      1,
			Authors: pq.StringArray{"Mr Sean"},
			Title:   "Test",
			Year:    "2022-11-01T00:00:00Z",
		},
		{
			ID:      2,
			Authors: pq.StringArray{"Mr Paul"},
			Title:   "Test2",
			Year:    "2022-11-01T00:00:00Z",
		},
	}

	expected, err := json.Marshal(gin.H{
		"books": books,
	})
	assert.NoError(t, err)

	log.Printf("Expected output: %v", expected)


	rr := httptest.NewRecorder()
	router := gin.Default()

	mock := new(mocks.Service)

	mock.On("GetAllBooks").Return(books, nil)

	request, err := http.NewRequest(http.MethodGet, "/books", nil)
	assert.NoError(t, err)

	router.ServeHTTP(rr, request)

	log.Printf("Output of the test: %v %v", rr.Code, rr.Body)

	assert.JSONEq(t, string(expected), string(rr.Body.Bytes()))
	assert.Equal(t, 200, rr.Code)
	mock.AssertExpectations(t)
}
