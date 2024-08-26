package main

import (
	"os"

	"server/config"
	"server/config/database"

	"github.com/gin-gonic/gin"
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
