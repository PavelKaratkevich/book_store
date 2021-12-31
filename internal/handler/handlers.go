package handler

import (
	"book_store/internal/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CustomerHandler incorporates service struct
type BookHandler struct {
	Service service.BookService
}

func(bh BookHandler) GetAllBook(ctx *gin.Context) {
	res, err := bh.Service.GetAllBooks()
		if err != nil {
			log.Printf("Error while calling GetAllBooks func: %v", err)
			ctx.JSON(http.StatusInternalServerError, "Unknown error occured. Please try again later")
			return
			} 
	ctx.JSON(http.StatusOK, res)
}

func(bh BookHandler) GetBookbyIdNumber(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	res, err := bh.Service.GetBookById(id)
		if err != nil {
			log.Printf("Error while calling GetBookById func: %v", err)
			ctx.JSON(http.StatusInternalServerError, "Unknown error occured. Please try again later")
			return
			} 
	ctx.JSON(http.StatusOK, res)
}