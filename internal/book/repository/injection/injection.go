package injection

import (
	"book_store/internal/book/delivery/http"
	"book_store/internal/book/delivery/middleware"
	postgresdb "book_store/internal/book/repository/postgresDB"
	"book_store/internal/book/service"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{DisableHTMLEscape: true})
	log.SetOutput(os.Stdout)
}

// checkEnvVars checks if all required envs are set
func checkEnvVars() {
	requiredEnvs := []string{"POSTGRES_PASSWORD", "POSTGRES_USER"}
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
	db, err := postgresdb.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	postgresBookRepositoryDB := postgresdb.NewBookRepositoryDb(db)
	bookService := service.NewBookService(postgresBookRepositoryDB)
	router := gin.Default()
	private := router.Group("/api")
	public := router.Group("")
	if gin.Mode() == gin.ReleaseMode {
		private.Use(middleware.CORS())
		private.Use(middleware.Logger())
		private.Use(middleware.CheckToken())
	}
	bookHTTP.RegisterBooksEndpoints(private, bookService)
	public.GET("/metrics/", func(ctx *gin.Context) {
		promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
	})
	public.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	return router
}
