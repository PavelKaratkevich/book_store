package middleware

import (
	jwtAuth "book_store/internal/middleware/jwt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Host: %v, Request method received: %s, Request path: %s, Status Code: %v", c.Request.Host, c.Request.Method, c.Request.URL.Path, c.Writer.Status())
	}
}

// CheckToken verifies if the token is provided, verified and valid
func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		if token := jwtAuth.ExtractToken(c.Request); token == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Please provide a valid token"})
			return
		}

		if _, err := jwtAuth.VerifyToken(c.Request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Token not verified"})
			return
		}

		if err := jwtAuth.TokenValid(c.Request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Token not valid"})
			return
		}
	}
}