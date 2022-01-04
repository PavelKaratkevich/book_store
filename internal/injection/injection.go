package injection

import (
	"book_store/internal/handler"
	"book_store/internal/middleware"
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

	// Enabling middleware
	g.Use(middleware.CORS())
	g.Use(middleware.Logger())

	// Declaring routes and handlers
	clientRoutes := g.Group("/books")
	{
		clientRoutes.GET("/", bh.GetAllBook)
		clientRoutes.GET("/:id/", bh.GetBookbyId)
		clientRoutes.POST("/", bh.UploadNewBook)
		clientRoutes.DELETE("/:id", bh.DeleteBook)
		clientRoutes.PUT("/:id", bh.UpdateBook)
	}

	return g
}
