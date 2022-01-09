package injection

import (
	"book_store/internal/handler"
	"book_store/internal/middleware"
	jwtAuth "book_store/internal/middleware/jwt"
	postgresdb "book_store/internal/repositoryDB/postgresDB"
	"book_store/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

// StartApp function creates and returns a gin router
func StartApp() *gin.Engine {

	// Create DB client for PostgreSQL
	db := postgresdb.ConnectDB()

	// Wiring
	postgresBookRepositoryDB := postgresdb.NewBookRepositoryDb(db)
	bookService := service.NewBookService(postgresBookRepositoryDB)
	bh := handler.BookHandler{Service: bookService}

	// Creating router and defining routes and handlers
	g := gin.Default()

	// Enable Prometheus metrics for the gin router
	DoPrometheusMetrics(g)

	// Enabling middleware
	g.Use(middleware.CORS())
	g.Use(middleware.Logger())

	// Providing endpoint for JWT authentication
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

func DoPrometheusMetrics(router *gin.Engine) {
	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(router)
}