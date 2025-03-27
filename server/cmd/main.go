package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"server/config"
	"server/config/database"
)

func main() {

	// Load environment variables
	err := config.LoadEnv()
	if err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	// Setup database
	err = database.SetupDatabase()
	if err != nil {
		log.Fatalf("failed to setup database: %v", err)
	}

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Setup router
	r := config.SetupRouter()

	// Run server
	err = r.Run()
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
