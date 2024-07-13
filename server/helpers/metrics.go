package helpers

import (
	"github.com/rangodisco/zelby/server/models"
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
