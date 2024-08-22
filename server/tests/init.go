package tests

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rangodisco/zelvy/server/config"
	"github.com/rangodisco/zelvy/server/database"
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
