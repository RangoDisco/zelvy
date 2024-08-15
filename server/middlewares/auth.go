package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CheckKey(publicRoutes []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientAPIKey := c.GetHeader("X-API-KEY")
		serverAPIKey := os.Getenv("API_KEY")

		if isProtected(c.Request.URL.Path, publicRoutes, c.Request.Method) && (clientAPIKey == "" || clientAPIKey != serverAPIKey) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func isProtected(route string, publicRoutes []string, method string) bool {
	for _, publicRoute := range publicRoutes {
		if publicRoute == route && method == http.MethodGet {
			return false
		}
	}
	return true
}
