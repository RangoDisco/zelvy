package tests

import (
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/models"
)

func init() {
	// Load test environment variables
	godotenv.Load("../.env.test")

	// Set test environment
	os.Setenv("GIN_MODE", "test")

	database.SetupDatabase()

	seedDatabase()
}

func seedDatabase() {
	// Seed database
	testUser := models.User{
		ID:          uuid.New(),
		Username:    "test123",
		DiscordID:   "123456789",
		PaypalEmail: "testEmail@gmail.com",
		CreatedAt:   time.Now(),
	}
	res := database.DB.Create(&testUser)

	if res.Error != nil {
		panic(res.Error)
	}
}
