package helpers

import (
	"github.com/rangodisco/zelby/server/enums"
	"github.com/rangodisco/zelby/server/models"
	"github.com/rangodisco/zelby/server/types"
	"strconv"
	"time"
)

// Convert ms to hour and minute format
func ConvertMsToHour(ms int) string {
	duration := time.Duration(ms) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	return strconv.Itoa(hours) + "h" + strconv.Itoa(minutes) + "m"
}

// Calculate main workout duration
func CalculateMainWorkoutDuration(workouts []models.Workout) int {
	var duration int
	for _, w := range workouts {
		if w.ActivityType == "strength" {
			duration = duration + w.Duration
		}
	}
	return duration
}

// Calculate extra workout duration
func CalculateExtraWorkoutDuration(workouts []models.Workout) int {
	var duration int
	for _, w := range workouts {
		if w.ActivityType != "strength" {
			duration = duration + w.Duration
		}
	}
	return duration
}

// Determine if the metric is successful based on the threshold
func IsMetricSuccessful(value int, goalValue int, shouldBeSmaller bool) bool {
	if shouldBeSmaller {
		return value >= goalValue
	}
	return value <= goalValue
}

func PopulateMetric(value int, threshold int, name string, shouldThresholdBeSmaller bool, unit string) types.Metric {
	return types.Metric{
		Value:        value,
		DisplayValue: strconv.Itoa(value) + unit,
		Threshold:    strconv.Itoa(threshold) + unit,
		Name:         name,
		Success:      IsMetricSuccessful(value, threshold, shouldThresholdBeSmaller),
	}
}

func CompareMetricsWithGoals(metrics models.Metrics, goals []models.Goal) []types.Metric {
	var comparedMetrics []types.Metric

	for _, g := range goals {
		var metric types.Metric

		// Add threshold to metric
		metric.Threshold = strconv.Itoa(g.Value)

		switch g.Type {
		case enums.KcalBurned:
			metric = PopulateMetric(metrics.KcalBurned, g.Value, "Calories brulées", true, "")

		case enums.KcalConsumed:
			metric = PopulateMetric(metrics.KcalConsumed, g.Value, "Calories consommées", false, "")

		case enums.MilliliterDrank:
			metric = PopulateMetric(metrics.MilliliterDrank, g.Value, "Eau", true, "ml")

		case enums.MainWorkoutDuration:
			duration := CalculateMainWorkoutDuration(metrics.Workouts)
			metric = PopulateWorkoutMetric(duration, g.Value, "Durée séance", true)

		case enums.ExtraWorkoutDuration:
			duration := CalculateExtraWorkoutDuration(metrics.Workouts)
			metric = PopulateWorkoutMetric(duration, g.Value, "Durée supplémentaire", true)
		}

		comparedMetrics = append(comparedMetrics, metric)
	}

	return comparedMetrics
}
