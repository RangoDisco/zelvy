package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/routes"
	"log"
	"net/http"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Setup database
	database.SetupDatabase()

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Start gin server
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	// Middleware to check API key in header
	//r.Use(middlewares.CheckKey())

	// Serve static files
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/clicked", func(c *gin.Context) {
		c.HTML(http.StatusOK, "clicked.html", nil)
	})

	// Register routes from routes package
	routes.RegisterSummaryRoutes(r)
	routes.RegisterGoalRoutes(r)
	routes.RegisterOffDayRoutes(r)
	routes.RegisterUserRoutes(r)

	// Run server
	err = r.Run()
	if err != nil {
		return
	}
}
