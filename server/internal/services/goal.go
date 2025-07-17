package services

import (
	"server/internal/enums"
	"server/pkg/types"
	"strconv"
	"time"

	"server/config/database"
	"server/internal/models"

	"github.com/google/uuid"
)

// convertToGoalViewModel Check if a goal is achieved, off or failed for each metric
func convertToGoalViewModel(m *models.Metric, g *models.Goal, workouts *[]models.Workout) (types.GoalViewModel, error) {
	var displayValue string
	var displayThreshold string
	var value float64

	// Populate goal based on its type
	if g.Type == enums.MainWorkoutDuration || g.Type == enums.ExtraWorkoutDuration {
		value = calculateWorkoutDuration(workouts, g.Type)
		displayValue = convertMsToHour(value)
		displayThreshold = convertMsToHour(g.Value)
	} else {
		switch g.Unit {
		case "L":
			displayValue = strconv.FormatFloat(m.Value, 'f', 2, 64) + "L"
			displayThreshold = strconv.FormatFloat(g.Value, 'f', 2, 64) + "L"

		default:
			displayValue = strconv.Itoa(int(m.Value))
			displayThreshold = strconv.Itoa(int(g.Value))
		}
	}

	isOff := isOff(g.ID)
	return types.GoalViewModel{
		Value:            m.Value,
		DisplayValue:     displayValue,
		Threshold:        g.Value,
		DisplayThreshold: displayThreshold,
		Name:             g.Name,
		Success:          isAchieved(m.Value, g.Value, g.Comparison, isOff),
		IsOff:            isOff,
		Progression:      getProgression(m.Value, g.Value),
		Picto:            getMetricPicto(g.Type),
	}, nil
}

// isAchieved, used when creating a view model, determines if the goal is achieved, based on the threshold
func isAchieved(value, goalValue float64, comparison string, isOffDay bool) bool {
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

// isOff used when creating a view model, check if a goal has been disabled for a given day
func isOff(goalId uuid.UUID) bool {
	var offDay models.Offday
	sod := time.Now().Truncate(24 * time.Hour)
	eod := sod.Add(24 * time.Hour)

	if err := database.GetDB().Where("goal_id = ? AND day >= ? AND day < ?", goalId, sod, eod).Find(&offDay).Error; err != nil {
		return false
	}

	if offDay.ID != uuid.Nil {
		return true
	}

	return false
}
