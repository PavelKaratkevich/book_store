package handler

import (
	"book_store/internal/domain"
	"book_store/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CustomerHandler incorporates service adapter (struct)
type BookHandler struct {
	Service service.BookService
}

// GetAllBook sends JSON with all books listed in the database
func (bh BookHandler) GetAllBook(ctx *gin.Context) {

	res, err := bh.Service.GetAllBooks()
	if err != nil {
		ctx.JSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetBookbyIdNumber returns JSON with a particular book depending on its ID
func (bh BookHandler) GetBookbyIdNumber(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	res, err := bh.Service.GetBookById(id)
	if err != nil {
		ctx.JSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// UploadNewBook requires JSON fields (Title, Authors, Year) to be filled in, and sends back the number of books added 
func (bh BookHandler) UploadNewBook(ctx *gin.Context) {

	var newBook domain.Book
	
	// Validating if all the fields are filled in
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "All fields should be filled in"} )
		return
	}
	
	res, err := bh.Service.PostNewBook(newBook)
	if err != nil {
		ctx.JSON(err.Code, gin.H{"error": err.Message})
		return
	}
	
	ctx.JSON(http.StatusOK, res)
}

// DeleteBookByItsIdNumber takes ID of a book from URL and sends back JSON with error or success
func (bh BookHandler) DeleteBookByItsIdNumber(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))
	
	rowsDeleted, err := bh.Service.DeleteBookById(id)
	if err != nil {
		ctx.JSON(err.Code, gin.H{"error": gin.H{"error":err.Message}})
		return
	}
	
	switch rowsDeleted {
	case 0:
		ctx.JSON(http.StatusNotFound, gin.H{"error":"ID not found"})
	case 1:
		ctx.JSON(http.StatusOK, gin.H{"error":"Book has been deleted successfully"})
	}
}

/* UpdateBookByItsId takes ID of the book from URL, and takes Title, Authors, Year fields from request body,
and sends back status code and status message */
func (bh BookHandler) UpdateBookByItsId(ctx *gin.Context) {
	// Retrieving ID from URL
	id, _ := strconv.Atoi(ctx.Param("id"))
	
	// Creating object which will be used for decoding fields of JSON request body (Title, Authors, Year)
	updateBookRequest := domain.Book{ID: id,}
	
	// Validating if all the fields were filled in
	if err := ctx.ShouldBindJSON(&updateBookRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "All fields should be filled in"} )
		return
	}
	// Invoking service function
	res, err := bh.Service.UpdateBookById(updateBookRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Internal error"})
	}
	// Switching over RowsAffected: if 1 - status 200, if 0 - status 404
	switch res {
	case 0:
		ctx.JSON(http.StatusNotFound, gin.H{"error": "ID not found"})
	case 1:
		ctx.JSON(http.StatusOK, gin.H{"error":"Book has been updated successfully"})
	}
}
