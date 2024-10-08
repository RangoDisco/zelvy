package handlers

import (
	"net/http"
	"time"

	"server/components"
	"server/database"
	"server/gintemplrenderer"
	"server/models"
	"server/services"
	"server/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterSummaryRoutes(r *gin.Engine) {
	r.GET("/summaries", getTodaySummary)
	r.POST("/api/summaries", AddSummary)
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

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, components.Summary(res))

	c.Render(http.StatusOK, r)

}

func AddSummary(c *gin.Context) {
	// Parse body
	var body types.SummaryInputModel
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
	if err := database.GetDB().Create(&summary).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Summary saved successfully!"})

}
