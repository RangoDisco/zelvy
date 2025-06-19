package services

import (
	"server/config/database"
	"strconv"

	"server/internal/enums"
	"server/internal/models"
	"server/pkg/types"

	"github.com/google/uuid"
)

// ConvertToMetricModel Converts a metric input to a db model
func ConvertToMetricModel(m *types.MetricInputModel, summaryId uuid.UUID) (models.Metric, bool) {
	// Fetch linked goal
	var goal models.Goal
	if database.GetDB().Where("type = ?", m.Type).First(&goal).Error == nil {
		return models.Metric{}, false
	}
	return models.Metric{
		ID:        uuid.New(),
		SummaryID: summaryId,
		Type:      m.Type,
		Value:     m.Value,
		Goal:      goal,
	}, true
}

// Determine if the metric is successful based on the threshold
func IsMetricSuccessful(value, goalValue float64, comparison string, isOffDay bool) bool {
	// In case the day is off, goal is always achieved
	if isOffDay {
		return true
	}
	switch comparison {
	case "greater":
		return value >= goalValue
	case "less":
		return value <= goalValue
	default:
		return false
	}
}

// ConvertToMetricViewModel Converts a model to ViewModel that matches fields needed by the frontend
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

// ConvertToWorkoutMetricViewModel Converts a workout related metric to ViewModel that matches fields needed by the frontend
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

// getMetricPicto Used to display picto in dashboard
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

// getProgression Used to display progress bar in dashboard
func getProgression(value float64, threshold float64) int {

	progression := int(value / threshold * 100)

	if progression > 100 {
		progression = 100
	}

	return progression
}

// CompareMetricsWithGoals Check if goal is achieved, off or failed for each metric
func CompareMetricsWithGoals(metrics *[]models.Metric, workouts *[]models.Workout) ([]types.MetricViewModel, error) {
	var comparedMetrics []types.MetricViewModel
	for _, m := range *metrics {
		var result types.MetricViewModel
		g := m.Goal
		isOff := IsOff(m.GoalID)

		// Populate metric based on its type
		switch g.Type {
		case enums.MainWorkoutDuration:
			duration := CalculateMainWorkoutDuration(workouts)
			result = ConvertToWorkoutMetricViewModel(g.Type, duration, g.Value, "Durée séance", g.Comparison, isOff)
		case enums.ExtraWorkoutDuration:
			duration := CalculateExtraWorkoutDuration(workouts)
			result = ConvertToWorkoutMetricViewModel(g.Type, duration, g.Value, "Durée supplémentaire", g.Comparison, isOff)
		default:
			result = ConvertToMetricViewModel(g.Type, m.Value, g.Value, g.Name, g.Comparison, g.Unit, isOff)
		}

		comparedMetrics = append(comparedMetrics, result)
	}

	return comparedMetrics, nil
}
