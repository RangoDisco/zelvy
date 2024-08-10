package services

import (
	"strconv"

	"github.com/google/uuid"
	"github.com/rangodisco/zelby/server/enums"
	"github.com/rangodisco/zelby/server/models"
	"github.com/rangodisco/zelby/server/types"
)

func ConvertToMetricModel(m types.MetricData, summaryId uuid.UUID) models.Metric {
	return models.Metric{
		ID:        uuid.New(),
		SummaryID: summaryId,
		Type:      m.Type,
		Value:     m.Value,
	}
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

	default:
		displayValue = strconv.Itoa(int(value))
		displayThreshold = strconv.Itoa(int(threshold))
	}

	return types.MetricResponse{
		Value:            value,
		DisplayValue:     displayValue,
		Threshold:        threshold,
		DisplayThreshold: displayThreshold,
		Name:             name,
		Success:          IsMetricSuccessful(value, threshold, comparison, isOffDay),
		IsOff:            isOffDay,
		Progression:      getProgression(value, threshold),
	}
}

func PopulateWorkoutMetric(duration float64, goalValue float64, name string, comparison string, isOffDay bool) types.MetricResponse {
	return types.MetricResponse{
		Value:            duration,
		DisplayValue:     ConvertMsToHour(duration),
		Threshold:        goalValue,
		DisplayThreshold: ConvertMsToHour(goalValue),
		Name:             name,
		Success:          IsMetricSuccessful(duration, goalValue, comparison, isOffDay),
		IsOff:            isOffDay,
		Progression:      getProgression(duration, goalValue),
	}
}

// Used to display progress bar in dashboard

func getProgression(value float64, threshold float64) int {

	progression := int(value / threshold * 100)

	if progression > 100 {
		progression = 100
	}

	return progression
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
		offDay := FetchByGoalAndDate(g.ID)
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
