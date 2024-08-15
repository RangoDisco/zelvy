package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rangodisco/zelby/server/components"
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/gintemplrenderer"
	"github.com/rangodisco/zelby/server/handlers"
	"github.com/rangodisco/zelby/server/middlewares"
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

	ginHtmlRenderer := r.HTMLRender
	r.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	// Middleware to check API key in header
	r.Use(middlewares.CheckKey([]string{"/", "/summaries", "/charts"}))

	// Register handlers from handlers package
	handlers.RegisterSummaryRoutes(r)
	handlers.RegisterGoalRoutes(r)
	handlers.RegisterOffDayRoutes(r)
	handlers.RegisterUserRoutes(r)
	handlers.RegisterChartRoutes(r)

	// Serve static files
	r.Static("/assets", "./assets")

	// Handle 404
	r.NoRoute(func(c *gin.Context) {
		r := gintemplrenderer.New(c.Request.Context(), http.StatusNotFound, components.NotFound())
		c.Render(404, r)
	})

	// Run server
	err = r.Run()
	if err != nil {
		return
	}
}
