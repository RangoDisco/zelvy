package tests

import (
	"os"

	"server/config"
	"server/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var Router *gin.Engine

func init() {
	// Load main environment variables
	godotenv.Load("../.env")

	// Ensure that the environment is set to test
	os.Setenv("GIN_MODE", "test")

	// Load environment variables
	config.LoadEnv()

	database.SetupDatabase()

	// Setup router
	Router = config.SetupRouter()
}
