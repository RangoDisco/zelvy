package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/middlewares"
	"github.com/rangodisco/zelby/server/routes"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Setup database
	database.SetupDatabase()

	// Start gin server
	r := gin.Default()

	// Middleware to check API key in header
	r.Use(middlewares.CheckKey())

	// Register routes from routes package
	routes.RegisterMetricsRoutes(r)
	routes.RegisterGoalRoutes(r)

	// Run server
	err = r.Run()
	if err != nil {
		return
	}
}
