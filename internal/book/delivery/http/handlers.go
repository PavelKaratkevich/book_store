package bookHTTP

import (
	"book_store/internal/domain"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CustomerHandler incorporates service adapter (struct)
type BookHandler struct {
	Service domain.Service
}

func NewBookHandler(service domain.Service) *BookHandler {
	return &BookHandler{Service: service}
}

// GetAllBook sends JSON with all books listed in the database
func (bh BookHandler) GetAllBook(ctx *gin.Context) {

	res, err := bh.Service.GetAllBooks(ctx.Request.Context())
	if err != nil {
		log.Printf("Error: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown error"})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetBookbyIdNumber returns JSON with a particular book depending on its ID
func (bh BookHandler) GetBookbyId(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Printf("Error: %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unknown error"})
		return
	}

	res, err2 := bh.Service.GetBookById(ctx.Request.Context(), id)
	if err2 != nil {
		if err2 == sql.ErrNoRows {
			log.Printf("Error: %v", err2.Error())
			ctx.JSON(http.StatusNotFound, gin.H{"error": "ID not found"})
			return
		} else {
			log.Printf("Error: %v", err2.Error())
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Unknown error"})
			return
		}
	}

	ctx.JSON(http.StatusOK, res)
}

// UploadNewBook requires JSON fields (Title, Authors, Year) to be filled in, and sends back the number of books added
func (bh BookHandler) UploadNewBook(ctx *gin.Context) {

	var newBook domain.Book

	// Validating if all the fields are filled in
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		log.Printf("Error: %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validating date
	validate := validator.New()
	if err := validate.Struct(newBook); err != nil {
		log.Printf("Error: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := bh.Service.PostNewBook(ctx.Request.Context(), newBook)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while making a book record"})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// DeleteBookByItsIdNumber takes ID of a book from URL and sends back JSON with error or success
func (bh BookHandler) DeleteBook(ctx *gin.Context) {

	id, err1 := strconv.Atoi(ctx.Param("id"))
	if err1 != nil {
		log.Printf("Error: %v", err1.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unknown error"})
		return
	}

	rowsDeleted, err := bh.Service.DeleteBookById(ctx.Request.Context(), id)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error while deleting book with %v from DB", id)})
		return
	}

	switch rowsDeleted {
	case 0:
		ctx.JSON(http.StatusNotFound, gin.H{"error": "ID not found"})
	case 1:
		ctx.JSON(http.StatusOK, gin.H{"Message": "Book has been deleted successfully"})
	}
}

/* UpdateBookByItsId takes ID of the book from URL, and takes Title, Authors, Year fields from request body,
and sends back status code and status message */
func (bh BookHandler) UpdateBook(ctx *gin.Context) {

	// Retrieving ID from URL
	id, err1 := strconv.Atoi(ctx.Param("id"))
	if err1 != nil {
		log.Printf("Error: %v", err1.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown error"})
		return
	}
	// Creating object which will be used for decoding fields of JSON request body (Title, Authors, Year)
	updateBookRequest := domain.Book{ID: id}

	// Validating if all the fields were filled in
	if err := ctx.ShouldBindJSON(&updateBookRequest); err != nil {
		log.Printf("Error: %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validating date
	validate := validator.New()
	if err := validate.Struct(updateBookRequest); err != nil {
		log.Printf("Error: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Invoking service function
	res, err := bh.Service.UpdateBookById(ctx.Request.Context(), updateBookRequest)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}
	// Switching over RowsAffected: if 1 - status 200, if 0 - status 404
	switch res {
	case 0:
		ctx.JSON(http.StatusNotFound, gin.H{"error": "ID not found"})
	case 1:
		ctx.JSON(http.StatusOK, gin.H{"Message": "Book has been updated successfully"})
	}
}
