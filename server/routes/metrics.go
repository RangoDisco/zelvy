package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/helpers"
	"github.com/rangodisco/zelby/server/models"
	"github.com/rangodisco/zelby/server/types"
	"net/http"
	"time"
)

func RegisterMetricsRoutes(r *gin.Engine) {
	r.GET("/api/metrics/today", getTodayMetrics)
	r.POST("/api/metrics", addMetrics)
}

// ROUTES
func getTodayMetrics(c *gin.Context) {

	var metrics models.Metrics

	// Get today's date
	sod := time.Now().Truncate(24 * time.Hour)
	eod := sod.Add(24 * time.Hour)

	// Query routes from today
	if err := database.DB.Where("date >= ? AND date < ?", sod, eod).Preload("Workouts").Find(&metrics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch current goals
	var goals []models.Goal
	if err := database.DB.Find(&goals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Build metrics response
	var metricsResponse types.MetricsResponse
	metricsResponse.ID = metrics.ID.String()
	metricsResponse.Date = metrics.Date.Format(time.RFC3339)
	metricsResponse.Steps = metrics.Steps
	metricsResponse.Metrics = helpers.CompareMetricsWithGoals(metrics, goals)

	// Add workouts to metrics object
	for _, w := range metrics.Workouts {
		workout := helpers.ConvertToWorkoutResponse(w)
		metricsResponse.Workouts = append(metricsResponse.Workouts, workout)
	}

	c.JSON(http.StatusOK, metricsResponse)
}

func addMetrics(c *gin.Context) {
	// Parse body
	var body types.RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert to models
	metrics := models.Metrics{
		ID:              uuid.New(),
		Date:            time.Now(),
		Steps:           body.Steps,
		KcalBurned:      body.KcalBurned,
		KcalConsumed:    body.KcalConsumed,
		MilliliterDrank: body.MilliliterDrank,
	}

	// Build and add workouts to the metrics object
	for _, w := range body.Workouts {
		workout := helpers.ConvertToWorkoutModel(w, metrics.ID)
		metrics.Workouts = append(metrics.Workouts, workout)
	}

	// Save routes and workouts
	if err := database.DB.Create(&metrics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Metrics saved successfully!"})

}
