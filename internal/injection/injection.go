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

	// Providing endoint for JWT authentication
	g.POST("/login", jwtAuth.Login())
	
	/* Declaring routes and handlers and enabling jwtAuth middleware 
	(checking if token is provided, verified and valid for all incoming requests)
	*/
	routes := g.Group("/books", middleware.CheckToken())
	{
		routes.GET("/", bh.GetAllBook)
		routes.GET("/:id/", bh.GetBookbyId)
		routes.POST("/", bh.UploadNewBook)
		routes.DELETE("/:id", bh.DeleteBook)
		routes.PUT("/:id", bh.UpdateBook)
	}

	return g
}
