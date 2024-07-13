package main

import (
	"github.com/gin-gonic/gin"
	"rangodisco.eu/zelby-server/database"
	"rangodisco.eu/zelby-server/routes"
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
