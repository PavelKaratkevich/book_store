package injection

import (
	"book_store/internal/handler"
	postgresdb "book_store/internal/repositoryDB/postgresDB"
	"book_store/internal/service"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	// Create DB client for PostgreSQL
	db := postgresdb.ConnectDB()

	// Wiring
	postgresBookRepositoryDB := postgresdb.NewBookRepositoryDb(db)
	bookService := service.NewBookService(postgresBookRepositoryDB)
	bh := handler.BookHandler{Service: bookService}

	// Creating router and defining routes and handlers
	g := gin.Default()
	g.GET("/books/", bh.GetAllBook)
	g.GET("/books/:id/", bh.GetBookbyIdNumber)
	g.POST("/books/", bh.UploadNewBook)
	return g
}