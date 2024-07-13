package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"rangodisco.eu/zelby-server/models"
)

var DB *gorm.DB

func SetupDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Auto-migrate your models
	err = DB.AutoMigrate(&models.Metrics{}, &models.Workout{}, &models.Goal{})
	if err != nil {
		log.Fatal("Failed to migrate models")
	}
}
