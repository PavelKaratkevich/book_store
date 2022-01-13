package injection

import (
	"book_store/internal/handler"
	"book_store/internal/middleware"
	jwtAuth "book_store/internal/middleware/jwt"
	postgresdb "book_store/internal/repositoryDB/postgresDB"
	"book_store/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// StartApp function creates and returns a gin router
func StartApp() *gin.Engine {

	var public *gin.RouterGroup

	gin.SetMode(gin.TestMode)

	// Create DB client for PostgreSQL
	db := postgresdb.ConnectDB()

	// Wiring
	postgresBookRepositoryDB := postgresdb.NewBookRepositoryDb(db)
	bookService := service.NewBookService(postgresBookRepositoryDB)
	bh := handler.BookHandler{Service: bookService}

	// Creating routers and defining routes and handlers
	appRouter := gin.Default()
	metricRouter := gin.Default()

	m := ginmetrics.GetMonitor()
	m.UseWithoutExposingEndpoint(appRouter)
	m.Expose(metricRouter)

	// switching on/off JWT Authentication depending on a Release or Test mode
	switch gin.Mode() {
	case gin.ReleaseMode:
		// Enabling middleware
		appRouter.Use(middleware.CORS())
		appRouter.Use(middleware.Logger())
		appRouter.POST("/login", jwtAuth.Login())
		public = appRouter.Group("/books", middleware.CheckToken())
	case gin.TestMode:
		public = appRouter.Group("/books")
	}
	{
		public.GET("/", bh.GetAllBook)
		public.GET("/:id/", bh.GetBookbyId)
		public.POST("/", bh.UploadNewBook)
		public.DELETE("/:id", bh.DeleteBook)
		public.PUT("/:id", bh.UpdateBook)
	}

	private := appRouter.Group("/api")
	{
		private.GET("/metrics/", func(ctx *gin.Context) {
			promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
		})
	}
	go func() {
		_ = metricRouter.Run(":8081")
	}()

	return appRouter
}
