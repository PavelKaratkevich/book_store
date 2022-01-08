package injection

import (
	"book_store/internal/handler"
	"book_store/internal/middleware"
	jwtAuth "book_store/internal/middleware/jwt"
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
	g.POST("/login", jwtAuth.Login()) // JWT authentication

	// Declaring routes and handlers
	routes := g.Group("/books")
	{
		routes.GET("/", bh.GetAllBook)
		routes.GET("/:id/", bh.GetBookbyId)
		routes.POST("/", bh.UploadNewBook)
		routes.DELETE("/:id", bh.DeleteBook)
		routes.PUT("/:id", bh.UpdateBook)
	}

	return g
}
