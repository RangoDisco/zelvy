package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rangodisco/zelvy/server/config"
	"github.com/rangodisco/zelvy/server/config/database"
	"log"
	"os"
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

	err = config.SetupGRpc()
	if err != nil {
		log.Fatalf("failed to setup gRpc: %v", err)
	}
}
