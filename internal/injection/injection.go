package injection

import (
	"book_store/internal/handler"
	postgresdb "book_store/internal/repositoryDB/postgresDB"
	"book_store/internal/service"
	"net/http"

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

	// Declaring routes and handlers
	clientRoutes := g.Group("/books") 
	{
	clientRoutes.GET("/", bh.GetAllBook)
	clientRoutes.GET("/:id/", bh.GetBookbyIdNumber)
	clientRoutes.POST("/", bh.UploadNewBook)
	clientRoutes.DELETE("/:id", bh.DeleteBookByItsIdNumber)
	clientRoutes.PUT("/:id", bh.UpdateBookByItsId)
	}
	return g
}