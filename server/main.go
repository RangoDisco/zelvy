package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rangodisco/zelby/server/config"
	"github.com/rangodisco/zelby/server/database"
)

func main() {

	// Load environment variables
	config.LoadEnv()

	// Setup database
	database.SetupDatabase()

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Setup router
	r := config.SetupRouter()

	// Run server
	err := r.Run()
	if err != nil {
		return
	}
}
