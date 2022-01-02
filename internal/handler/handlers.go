package handler

import (
	"book_store/internal/domain"
	"book_store/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CustomerHandler incorporates service struct
type BookHandler struct {
	Service service.BookService
}

func (bh BookHandler) GetAllBook(ctx *gin.Context) {
	res, err := bh.Service.GetAllBooks()
	if err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (bh BookHandler) GetBookbyIdNumber(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	res, err := bh.Service.GetBookById(id)
	if err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (bh BookHandler) UploadNewBook(ctx *gin.Context) {
	var newBook domain.Book
	// Validating if all the fields are filled in
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.JSON(http.StatusBadRequest, "All fields should be filled in")
		return
	}
	res, err := bh.Service.PostNewBook(newBook)
	if err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (bh BookHandler) DeleteBookByItsIdNumber(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	rowsDeleted, err := bh.Service.DeleteBookById(id)
	if err != nil {
		ctx.JSON(err.Code, "Server error.")
		return
	}
	switch rowsDeleted {
	case 0:
		ctx.JSON(http.StatusNotFound, "ID not found")
	case 1:
		ctx.JSON(http.StatusOK, "Book has been deleted successfully")
	}
}
