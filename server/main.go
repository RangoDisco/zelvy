package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/gintemplrenderer"
	"github.com/rangodisco/zelby/server/routes"
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

	r.LoadHTMLFiles("./templates/index.html")

	ginHtmlRenderer := r.HTMLRender
	r.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	// Middleware to check API key in header
	//r.Use(middlewares.CheckKey())

	// Register routes from routes package
	routes.RegisterSummaryRoutes(r, ginHtmlRenderer)
	routes.RegisterGoalRoutes(r)
	routes.RegisterOffDayRoutes(r)
	routes.RegisterUserRoutes(r)

	// Serve static files
	r.Static("/assets", "./assets")

	// Run server
	err = r.Run()
	if err != nil {
		return
	}
}
