package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/models"
	"github.com/rangodisco/zelby/server/types"
	"github.com/rangodisco/zelby/server/utils"
	"net/http"
	"time"
)

func RegisterSummaryRoutes(r *gin.Engine) {
	r.GET("/api/summaries/today", getTodaySummary)
	r.POST("/api/summaries", addSummary)
}

// ROUTES
func getTodaySummary(c *gin.Context) {

	var summary models.Summary

	// Get today's date
	sod := time.Now().Truncate(24 * time.Hour)
	eod := sod.Add(24 * time.Hour)

	// Query routes from today
	if err := database.DB.Where("date >= ? AND date < ?", sod, eod).Preload("Workouts").Preload("Metrics").Preload("Winner").Find(&summary).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch current goals
	var goals []models.Goal
	if err := database.DB.Find(&goals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Build summary response
	var res types.SummaryResponse
	res.ID = summary.ID.String()
	res.Date = summary.Date.Format(time.RFC3339)
	res.Winner.DiscordID = summary.Winner.DiscordID
	metrics, err := utils.CompareMetricsWithGoals(summary, goals)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res.Metrics = metrics

	// Add workouts to metrics object
	for _, w := range summary.Workouts {
		workout := utils.ConvertToWorkoutResponse(w)
		res.Workouts = append(res.Workouts, workout)
	}

	c.JSON(http.StatusOK, res)
}

func addSummary(c *gin.Context) {
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
		summary.Metrics = append(summary.Metrics, utils.ConvertToMetricModel(m, summary.ID))
	}

	// Build and add workouts to the summary object
	for _, w := range body.Workouts {
		workout := utils.ConvertToWorkoutModel(w, summary.ID)
		summary.Workouts = append(summary.Workouts, workout)
	}

	// Pick winner
	summary.WinnerID = utils.PickWinner()

	// Save summary
	if err := database.DB.Create(&summary).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Summary saved successfully!"})

}
