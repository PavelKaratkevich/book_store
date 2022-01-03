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
	
    // Enabling CORS
    g.Use(CORS)

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

func CORS(c *gin.Context) {

    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Methods", "*")
    c.Header("Access-Control-Allow-Headers", "*")
    c.Header("Content-Type", "application/json")

    if c.Request.Method != "OPTIONS" {      
        c.Next()
    } else {
        c.AbortWithStatus(http.StatusOK)
    }
}
