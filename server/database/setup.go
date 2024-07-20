package database

import (
	"fmt"
	"github.com/rangodisco/zelby/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func SetupDatabase() {
	var err error

	// Open a database connection
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	// Use psql in production, sqlite in development
	if os.Getenv("GIN_MODE") == "release" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host, user, password, name, port)

		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	}

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Auto-migrate your models
	err = DB.AutoMigrate(&models.Summary{}, &models.Metric{}, &models.Workout{}, &models.Goal{}, &models.Offday{})
	if err != nil {
		log.Fatal("Failed to migrate models")
	}
}
