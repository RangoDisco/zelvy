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

func RegisterSummaryRoutes(r *gin.Engine) {
	r.GET("/api/summaries/today", getTodaySummary)
	r.POST("/api/summaries", addMetrics)
}

// ROUTES
func getTodaySummary(c *gin.Context) {

	var summary models.Summary

	// Get today's date
	sod := time.Now().Truncate(24 * time.Hour)
	eod := sod.Add(24 * time.Hour)

	// Query routes from today
	if err := database.DB.Where("date >= ? AND date < ?", sod, eod).Preload("Workouts").Preload("Metrics").Find(&summary).Error; err != nil {
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
	var metricsResponse types.SummaryResponse
	metricsResponse.ID = summary.ID.String()
	metricsResponse.Date = summary.Date.Format(time.RFC3339)
	metricsResponse.Metrics = helpers.CompareMetricsWithGoals(summary, goals)

	// Add workouts to metrics object
	for _, w := range summary.Workouts {
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
	summary := models.Summary{
		ID:   uuid.New(),
		Date: time.Now(),
	}

	// Build and add metrics to the summary object
	for _, m := range body.Metrics {
		summary.Metrics = append(summary.Metrics, helpers.ConvertToMetricModel(m, summary.ID))
	}

	// Build and add workouts to the metrics object
	for _, w := range body.Workouts {
		workout := helpers.ConvertToWorkoutModel(w, summary.ID)
		summary.Workouts = append(summary.Workouts, workout)
	}

	// Save routes and workouts
	if err := database.DB.Create(&summary).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Metrics saved successfully!"})

}
