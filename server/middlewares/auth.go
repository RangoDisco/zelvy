package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func CheckKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientAPIKey := c.GetHeader("X-API-KEY")
		serverAPIKey := os.Getenv("API_KEY")

		if clientAPIKey == "" || clientAPIKey != serverAPIKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
