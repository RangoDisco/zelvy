package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckKey(publicRoutes []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientAPIKey := c.GetHeader("X-API-KEY")
		serverAPIKey := os.Getenv("API_KEY")

		// Bypass auth in test context
		if os.Getenv("APP_ENV") == "test" {
			c.Next()
			return
		}

		if isProtected(c.Request.URL.Path, publicRoutes, c.Request.Method) && (clientAPIKey == "" || clientAPIKey != serverAPIKey) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func isProtected(route string, publicRoutes []string, method string) bool {

	// Check if the route starts with "/assets/"
	if strings.HasPrefix(route, "/assets/") && method == http.MethodGet {
		return false
	}

	for _, publicRoute := range publicRoutes {
		if publicRoute == route && method == http.MethodGet {
			return false
		}
	}
	return true
}
