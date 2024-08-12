package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rangodisco/zelby/server/components"
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/gintemplrenderer"
	"github.com/rangodisco/zelby/server/models"
	"github.com/rangodisco/zelby/server/services"
	"github.com/rangodisco/zelby/server/types"
)

func RegisterSummaryRoutes(r *gin.Engine) {
	r.GET("/summaries", getTodaySummary)
	r.POST("/api/summaries", addSummary)
}

// ROUTES
func getTodaySummary(c *gin.Context) {
	// Get date from params
	date := c.Param("date")

	// Fetch summary
	summary, err := services.FetchSummaryByDate(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Format data to fit fields in the view
	res, err := services.CreateSummaryViewModel(summary)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Temp fix to handle both HTML and JSON responses
	accept := c.GetHeader("Accept")

	if accept == "application/json" {
		c.JSON(http.StatusOK, res)
		return
	}

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, components.Home(res))

	c.Render(http.StatusOK, r)

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
		summary.Metrics = append(summary.Metrics, services.ConvertToMetricModel(m, summary.ID))
	}

	// Build and add workouts to the summary object
	for _, w := range body.Workouts {
		workout := services.ConvertToWorkoutModel(w, summary.ID)
		summary.Workouts = append(summary.Workouts, workout)
	}

	// Pick winner
	summary.WinnerID = services.PickWinner()

	// Save summary
	if err := database.DB.Create(&summary).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Summary saved successfully!"})

}
