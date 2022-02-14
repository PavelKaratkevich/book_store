package bookHTTP

import (
	"book_store/internal/domain"

	"github.com/gin-gonic/gin"
)

func RegisterBooksEndpoints(g *gin.RouterGroup, h domain.Service) {
	bh := BookHandler{Service: h}

	g.GET("/books", bh.GetAllBook)
	g.GET("/books/:id", bh.GetBookbyId)
	g.POST("/books", bh.UploadNewBook)
	g.DELETE("/books/:id", bh.DeleteBook)
	g.PUT("/books/:id", bh.UpdateBook)
}
