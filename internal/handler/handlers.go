package handler

import (
	"book_store/internal/domain"
	"book_store/internal/service"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CustomerHandler incorporates service adapter (struct)
type BookHandler struct {
	Service service.BookService
}

// GetAllBook sends JSON with all books listed in the database
func (bh BookHandler) GetAllBook(ctx *gin.Context) {

	res, err := bh.Service.GetAllBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown error"})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetBookbyIdNumber returns JSON with a particular book depending on its ID
func (bh BookHandler) GetBookbyId(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown error"})
		return
	}

	res, err2 := bh.Service.GetBookById(id)
	if err2 != nil {
		if *err2 == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "ID not found"})
			return
		} else {
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validating date
	validate := validator.New()
	if err := validate.Struct(newBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := bh.Service.PostNewBook(newBook)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while making a book record"})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// DeleteBookByItsIdNumber takes ID of a book from URL and sends back JSON with error or success
func (bh BookHandler) DeleteBook(ctx *gin.Context) {

	id, err1 := strconv.Atoi(ctx.Param("id"))
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown error"})
		return
	}

	rowsDeleted, err := bh.Service.DeleteBookById(id)
	if err != nil {
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown error"})
		return
	}
	// Creating object which will be used for decoding fields of JSON request body (Title, Authors, Year)
	updateBookRequest := domain.Book{ID: id}

	// Validating if all the fields were filled in
	if err := ctx.ShouldBindJSON(&updateBookRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validating date
	validate := validator.New()
	if err := validate.Struct(updateBookRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Invoking service function
	res, err := bh.Service.UpdateBookById(updateBookRequest)
	if err != nil {
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
