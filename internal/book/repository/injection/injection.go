package injection

import (
	bookHTTP "book_store/internal/book/delivery/http"
	"book_store/internal/book/delivery/middleware"
	jwtAuth "book_store/internal/book/delivery/middleware/jwt"
	postgresdb "book_store/internal/book/repository/postgresDB"
	"book_store/internal/book/service"
	"os"
	"strings"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{DisableHTMLEscape: true})
	log.SetOutput(os.Stdout)
}

func checkEnvVars() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	requiredEnvs := []string{"POSTGRES_USER", "POSTGRES_PASSWORD", "DB_ADDRESS", "DB_PORT", "POSTGRES_DB"}
	var msg []string
	for _, el := range requiredEnvs {
		val, exists := os.LookupEnv(el)
		if !exists || len(val) == 0 {
			msg = append(msg, el)
		}
	}
	if len(msg) > 0 {
		log.Fatal(strings.Join(msg, ", "), " env(s) not set")
	}
}

// StartApp function creates and returns a gin router
func StartApp() *gin.Engine {
	initLogger()
	checkEnvVars()

	gin.SetMode(gin.ReleaseMode)

	var public *gin.RouterGroup

	// Create DB client for PostgreSQL
	db, err := postgresdb.ConnectDB()
	if err != nil {
		log.Fatalf("Could not establish DB connection: %v", err.Error())
	}

	// Wiring
	postgresBookRepositoryDB := postgresdb.NewBookRepositoryDb(db)
	bookService := service.NewBookService(postgresBookRepositoryDB)

	// Creating routers and defining routes and handlers
	appRouter := gin.Default()
	metricRouter := gin.Default()

	m := ginmetrics.GetMonitor()
	m.UseWithoutExposingEndpoint(appRouter)
	m.Expose(metricRouter)

	// switching on/off JWT Authentication depending on a Release or Test mode
	if gin.Mode() == gin.ReleaseMode {
		appRouter.Use(middleware.CORS())
		appRouter.Use(middleware.Logger())
		appRouter.POST("/login", jwtAuth.Login())
		public = appRouter.Group("/api", middleware.CheckToken())
	} else {
		public = appRouter.Group("/api")
	}

	// Register endpoints for public
	bookHTTP.RegisterBooksEndpoints(public, bookService)

	private := appRouter.Group("")
	{
		private.GET("/metrics", func(ctx *gin.Context) {
			promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
		})
		private.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"status": "healthy"})
		})
	}

	go func() {
		_ = metricRouter.Run(":8081")
	}()

	return appRouter
}
