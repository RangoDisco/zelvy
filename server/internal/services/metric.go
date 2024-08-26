package services

import (
	"strconv"

	"server/internal/enums"
	"server/internal/models"
	"server/pkg/types"

	"github.com/google/uuid"
)

func ConvertToMetricModel(m types.MetricInputModel, summaryId uuid.UUID) models.Metric {
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

func ConvertToMetricViewModel(goalType string, value float64, threshold float64, name string, comparison string, unit string, isOffDay bool) types.MetricViewModel {
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

	return types.MetricViewModel{
		Value:            value,
		DisplayValue:     displayValue,
		Threshold:        threshold,
		DisplayThreshold: displayThreshold,
		Name:             name,
		Success:          IsMetricSuccessful(value, threshold, comparison, isOffDay),
		IsOff:            isOffDay,
		Progression:      getProgression(value, threshold),
		Picto:            getMetricPicto(goalType),
	}
}

func ConvertToWorkoutMetricViewModel(goalType string, duration float64, goalValue float64, name string, comparison string, isOffDay bool) types.MetricViewModel {
	return types.MetricViewModel{
		Value:            duration,
		DisplayValue:     ConvertMsToHour(duration),
		Threshold:        goalValue,
		DisplayThreshold: ConvertMsToHour(goalValue),
		Name:             name,
		Success:          IsMetricSuccessful(duration, goalValue, comparison, isOffDay),
		IsOff:            isOffDay,
		Progression:      getProgression(duration, goalValue),
		Picto:            getMetricPicto(goalType),
	}
}

func getMetricPicto(goalType string) string {
	switch goalType {
	case enums.MainWorkoutDuration:
		return "🏋️"
	case enums.ExtraWorkoutDuration:
		return "👟"
	case enums.KcalBurned:
		return "🔥"
	case enums.KcalConsumed:
		return "🍛"
	case enums.MilliliterDrank:
		return "💧"
	default:
		return "📊"
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

func CompareMetricsWithGoals(metrics []models.Metric, workouts []models.Workout) ([]types.MetricViewModel, error) {

	// Firt fetch all goals
	goals, err := FetchGoals()
	if err != nil {
		return []types.MetricViewModel{}, err
	}

	var comparedMetrics []types.MetricViewModel

	// Create a map of metrics to values and then iterate over the goals
	metricMap := make(map[string]float64)
	for _, m := range metrics {
		metricMap[m.Type] = m.Value
	}

	for _, g := range goals {
		var result types.MetricViewModel
		// Find metric by goal type
		value := metricMap[g.Type]

		// Search if goal is off for today
		isOffDay := IsOff(g.ID)

		// Populate metric based on goal type
		switch g.Type {
		case enums.MainWorkoutDuration:
			duration := CalculateMainWorkoutDuration(workouts)
			result = ConvertToWorkoutMetricViewModel(g.Type, duration, g.Value, "Durée séance", g.Comparison, isOffDay)
		case enums.ExtraWorkoutDuration:
			duration := CalculateExtraWorkoutDuration(workouts)
			result = ConvertToWorkoutMetricViewModel(g.Type, duration, g.Value, "Durée supplémentaire", g.Comparison, isOffDay)
		default:
			result = ConvertToMetricViewModel(g.Type, value, g.Value, g.Name, g.Comparison, g.Unit, isOffDay)
		}

		// Add threshold to metric
		comparedMetrics = append(comparedMetrics, result)
	}

	return comparedMetrics, nil
}
