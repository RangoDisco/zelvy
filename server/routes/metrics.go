package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/helpers"
	"github.com/rangodisco/zelby/server/models"
	"net/http"
	"strconv"
	"time"
)

// TYPES
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

type Metric struct {
	Name         string `json:"name"`
	Value        int    `json:"value"`
	DisplayValue string `json:"displayValue"`
	Threshold    string `json:"threshold"`
	Success      bool   `json:"success"`
}

type MetricsResponse struct {
	ID       string           `json:"id"`
	Date     string           `json:"date"`
	Steps    int              `json:"steps"`
	Metrics  []Metric         `json:"metrics"`
	Workouts []models.Workout `json:"workouts"`
}

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
	var metricsResponse MetricsResponse
	metricsResponse.ID = metrics.ID.String()
	metricsResponse.Date = metrics.Date.Format(time.RFC3339)
	metricsResponse.Steps = metrics.Steps
	metricsResponse.Workouts = metrics.Workouts
	metricsResponse.Metrics = compareMetricsWithGoals(metrics, goals)

	c.JSON(http.StatusOK, metricsResponse)
}

func addMetrics(c *gin.Context) {
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
		// Handle name based on activity type in case null
		if w.Name == "" {
			switch w.ActivityType {
			case "strength":
				w.Name = "Séance de musculation"
				break
			case "running":
				w.Name = "Footing"
				break
			case "cycling":
				w.Name = "Vélo"
				break
			case "walk":
				w.Name = "Marche"
				break
			}
		}

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

	// Save routes and workouts
	if err := database.DB.Create(&metrics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Metrics saved successfully!"})

}

// Helpers
func compareMetricsWithGoals(metrics models.Metrics, goals []models.Goal) []Metric {
	var comparedMetrics []Metric

	for _, g := range goals {
		var metric Metric
		// Add threshold and type to metric
		metric.Threshold = strconv.Itoa(g.Value)

		switch g.Type {
		case "kcalBurned":
			metric = populateMetric(metrics.KcalBurned, g.Value, "Calories brulées", true)

		case "kcalConsumed":
			metric = populateMetric(metrics.KcalConsumed, g.Value, "Calories consommées", false)

		case "litterDrank":
			metric = populateMetric(metrics.CentiliterDrank, g.Value, "Litres bus", true)

		case "mainWorkoutDuration":
			duration := helpers.CalculateMainWorkoutDuration(metrics.Workouts)
			metric = populateWorkoutMetric(duration, g.Value, "Durée séance", true)

		case "extraWorkoutDuration":
			duration := helpers.CalculateExtraWorkoutDuration(metrics.Workouts)
			metric = populateWorkoutMetric(duration, g.Value, "Durée supplémentaire", true)
		}

		comparedMetrics = append(comparedMetrics, metric)
	}

	return comparedMetrics

}

// Helper function to populate metric
func populateMetric(value int, threshold int, name string, shouldThresholdBeSmaller bool) Metric {
	return Metric{
		Value:        value,
		DisplayValue: strconv.Itoa(value),
		Threshold:    strconv.Itoa(threshold),
		Name:         name,
		Success:      helpers.IsMetricSuccessful(value, threshold, shouldThresholdBeSmaller),
	}
}

// Helper function to populate workout metric
func populateWorkoutMetric(duration int, goalValue int, name string, shouldThresholdBeSmaller bool) Metric {
	return Metric{
		Value:        duration,
		DisplayValue: helpers.ConvertMsToHour(duration),
		Threshold:    helpers.ConvertMsToHour(goalValue),
		Name:         name,
		Success:      helpers.IsMetricSuccessful(duration, goalValue, shouldThresholdBeSmaller),
	}
}
