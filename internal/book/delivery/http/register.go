package bookHTTP

import (
	"book_store/internal/book/service"

	"github.com/gin-gonic/gin"
)

func RegisterBooksEndpoints(g *gin.RouterGroup, h service.BookService) {
	bh := BookHandler{Service: h}

	books := g.Group("")
	{
		books.GET("/", bh.GetAllBook)
		books.GET("/:id/", bh.GetBookbyId)
		books.POST("/", bh.UploadNewBook)
		books.DELETE("/:id", bh.DeleteBook)
		books.PUT("/:id", bh.UpdateBook)
	}
}