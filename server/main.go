package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/routes"
)

func main() {

	// Setup database
	database.SetupDatabase()

	// Start gin server
	r := gin.Default()

	// Register routes from routes package
	routes.RegisterMetricsRoutes(r)
	routes.RegisterGoalRoutes(r)

	// Run server
	err := r.Run()
	if err != nil {
		return
	}
}
