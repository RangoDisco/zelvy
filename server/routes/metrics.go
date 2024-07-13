package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rangodisco.eu/zelby-server/database"
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

type Metric struct {
	Type      string `json:"type"`
	Value     int    `json:"value"`
	Threshold int    `json:"threshold"`
	Success   bool   `json:"success"`
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

func compareMetricsWithGoals(metrics models.Metrics, goals []models.Goal) []Metric {
	var comparedMetrics []Metric
	shouldThresholdBeBigger := true

	for _, g := range goals {
		var metric Metric

		// Add threshold and type to metric
		metric.Threshold = g.Value
		metric.Type = g.Type

		switch g.Type {
		case "kcalBurned":
			metric.Value = metrics.KcalBurned
			break

		case "kcalConsumed":
			metric.Value = metrics.KcalConsumed
			shouldThresholdBeBigger = false
			break

		case "litterDrank":
			metric.Value = metrics.CentiliterDrank
			break

		case "mainWorkoutDuration":
			// Find workouts with type "strength" and sum duration
			for _, w := range metrics.Workouts {
				if w.ActivityType == "strength" {
					metric.Value = metric.Value + w.Duration
				}
			}
			break

		case "extraWorkoutDuration":
			// Sum all workouts duration EXCEPT strength
			for _, w := range metrics.Workouts {
				if w.ActivityType != "strength" {
					metric.Value = metric.Value + w.Duration
				}
			}
			break
		}

		// Compare metric with threshold. Note that only kcalConsumed should be lower than threshold
		if shouldThresholdBeBigger {
			metric.Success = metric.Value >= metric.Threshold
		} else {
			metric.Success = metric.Value <= metric.Threshold
		}

		// Finally push in comparedMetrics
		comparedMetrics = append(comparedMetrics, metric)
	}
	return comparedMetrics

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
