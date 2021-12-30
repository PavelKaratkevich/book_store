package handler

import (
	"book_store/internal/service"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CustomerHandler incorporates service struct
type BookHandler struct {
	Service service.BookService
}

func(bh BookHandler) GetAllBook(ctx *gin.Context) {
	res, err := bh.Service.GetAllBooks(context.Background())
		if err != nil {
			log.Printf("Error while calling GetAllBooks func: %v", err)
			ctx.JSON(http.StatusInternalServerError, "Unknown error occured. Please try again later")
			return
			} 
	ctx.JSON(http.StatusOK, res)
}