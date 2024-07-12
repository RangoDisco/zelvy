package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"rangodisco.eu/zelby-server/models"
	"time"
)

type WorkoutData struct {
	KcalBurned   int    `json:"kcalBurned"`
	ActivityType string `json:"activityType"`
	Name         string `json:"name"`
	Duration     int    `json:"duration"`
}

type RequestBody struct {
	KcalBurned   int           `json:"kcalBurned"`
	Steps        int           `json:"steps"`
	Workouts     []WorkoutData `json:"workouts"`
	KcalConsumed int           `json:"kcalConsumed"`
}

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
	db := setupDb()

	// Start gin server
	r := gin.Default()

	r.GET("/api/metrics/today", func(c *gin.Context) {
		var metrics models.Metrics

		// Get today's date
		sod := time.Now().Truncate(24 * time.Hour)
		eod := sod.Add(24 * time.Hour)

		// Query metrics from today
		if err := db.Where("date >= ? AND date < ?", sod, eod).Preload("Workouts").Find(&metrics).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, metrics)

	})

	r.POST("/api/metrics", func(c *gin.Context) {
		// Parse body
		var body RequestBody
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Convert to models
		metrics := models.Metrics{
			ID:           uuid.New(),
			Date:         time.Now(),
			Steps:        body.Steps,
			KcalBurned:   body.KcalBurned,
			KcalConsumed: body.KcalConsumed,
		}

		for _, w := range body.Workouts {
			workout := models.Workout{
				ID:           uuid.New(),
				MetricsRefer: metrics.ID,
				KcalBurned:   w.KcalBurned,
				ActivityType: w.ActivityType,
				Name:         w.Name,
				Duration:     w.Duration,
			}
			metrics.Workouts = append(metrics.Workouts, workout)
		}

		// Save metrics and workouts
		if err := db.Create(&metrics).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Metrics saved successfully!"})

	})

	// Run server
	err := r.Run()
	if err != nil {
		return
	}
}
