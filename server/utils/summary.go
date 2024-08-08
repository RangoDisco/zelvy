package utils

import (
	"github.com/google/uuid"
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/enums"
	"github.com/rangodisco/zelby/server/models"
	"github.com/rangodisco/zelby/server/types"
	"strconv"
	"time"
)

// Convert ms to hour and minute format
func ConvertMsToHour(ms float64) string {
	duration := time.Duration(ms) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	return strconv.Itoa(hours) + "h" + strconv.Itoa(minutes) + "m"
}

func ConvertToMetricModel(m types.MetricData, summaryId uuid.UUID) models.Metric {
	return models.Metric{
		ID:        uuid.New(),
		SummaryID: summaryId,
		Type:      m.Type,
		Value:     m.Value,
	}

}

// Calculate extra workout duration
func CalculateExtraWorkoutDuration(workouts []models.Workout) float64 {
	var duration float64
	for _, w := range workouts {
		if w.ActivityType != "strength" {
			duration = duration + w.Duration
		}
	}
	return duration
}

// Determine if the metric is successful based on the threshold
func IsMetricSuccessful(value float64, goalValue float64, comparison string, isOffDay bool) bool {
	if comparison == "greater" {
		return value >= goalValue || isOffDay
	}
	return value <= goalValue || isOffDay
}

func PopulateMetric(value float64, threshold float64, name string, comparison string, unit string, isOffDay bool) types.MetricResponse {
	var displayValue string
	var displayThreshold string
	// Handle weird float/int diff between goals
	switch unit {
	case "L":
		displayValue = strconv.FormatFloat(value, 'f', 2, 64) + "L"
		displayThreshold = strconv.FormatFloat(threshold, 'f', 2, 64) + "L"
		break

	default:
		displayValue = strconv.Itoa(int(value))
		displayThreshold = strconv.Itoa(int(threshold))
	}

	return types.MetricResponse{
		DisplayValue: displayValue,
		Threshold:    displayThreshold,
		Name:         name,
		Success:      IsMetricSuccessful(value, threshold, comparison, isOffDay),
	}
}

func CompareMetricsWithGoals(summary models.Summary, goals []models.Goal) ([]types.MetricResponse, error) {
	var comparedMetrics []types.MetricResponse

	// Create a map of metrics to values and then iterate over the goals
	metricMap := make(map[string]float64)
	for _, m := range summary.Metrics {
		metricMap[m.Type] = m.Value
	}

	for _, g := range goals {
		var result types.MetricResponse
		var isOffDay bool
		// Find metric by goal type
		value := metricMap[g.Type]

		// Search if goal is off for today
		offDay, err := FetchByGoalAndDate(g.ID)
		if err != nil {
			return nil, err
		}
		if offDay != nil {
			isOffDay = true
		}

		// Populate metric based on goal type
		if g.Type == enums.MainWorkoutDuration {
			duration := CalculateMainWorkoutDuration(summary.Workouts)
			result = PopulateWorkoutMetric(duration, g.Value, "Durée séance", g.Comparison, isOffDay)
		} else if g.Type == enums.ExtraWorkoutDuration {
			duration := CalculateExtraWorkoutDuration(summary.Workouts)
			result = PopulateWorkoutMetric(duration, g.Value, "Durée supplémentaire", g.Comparison, isOffDay)
		} else {
			result = PopulateMetric(value, g.Value, g.Name, g.Comparison, g.Unit, isOffDay)
		}

		// Add threshold to metric
		comparedMetrics = append(comparedMetrics, result)
	}

	return comparedMetrics, nil
}

func PickWinner() uuid.UUID {
	var u models.User

	// Get user from db randomly
	database.DB.Order("RANDOM()").First(&u)

	return u.ID
}
