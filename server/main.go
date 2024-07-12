package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"rangodisco.eu/zelby-server/models"
)

func setupDb() *gorm.DB {
	// Open database connection
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Migrate all models
	err = db.AutoMigrate(&models.Metrics{}, &models.Workout{})
	if err != nil {
		log.Fatal("Failed to migrate models")
	}

	return db
}

func main() {

	// Setup database
	_ = setupDb()

	// Start gin server
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err := r.Run()
	if err != nil {
		return
	}
}
