package services

import (
	"errors"
	"strconv"
	"time"

	"github.com/rangodisco/zelvy/gen/zelvy/metric"
	"github.com/rangodisco/zelvy/server/internal/enums"

	pb_goa "github.com/rangodisco/zelvy/gen/zelvy/goal"
	"github.com/rangodisco/zelvy/server/config/database"
	"github.com/rangodisco/zelvy/server/internal/models"

	"github.com/google/uuid"
)

// DisableGoal fetches a goals by its type, creates an Offday and inserts in db, disabling the goal for the day
func DisableGoal(goalType pb_goa.GoalType) error {
	var goal models.Goal
	if err := database.GetDB().Where("type = ? AND active = true", goalType.String()).First(&goal).Error; err != nil {
		return err
	}
	err := createOffDay(goal)
	if err != nil {
		return err
	}

	return nil
}

func createOffDay(goal models.Goal) error {
	offDay := models.Offday{
		GoalID: goal.ID,
		Day:    time.Now(),
	}

	// See if goal has already been disabled for today
	isAlreadyOff := isOff(goal.ID)
	if isAlreadyOff {
		return nil
	}

	if err := database.GetDB().Save(&offDay).Error; err != nil {
		return err
	}
	return nil
}

// convertToGoalViewModel Check if a goal is achieved, off or failed for each metric
func convertToGoalViewModel(m *models.Metric, g *models.Goal) (pb_goa.GoalViewModel, error) {

	if m == nil {
		return pb_goa.GoalViewModel{}, errors.New("metric is nil")
	}

	displayValue, displayThreshold := formatDisplayValue(m.Value, g)

	return pb_goa.GoalViewModel{
		Value:            m.Value,
		DisplayValue:     displayValue,
		Threshold:        g.Value,
		DisplayThreshold: displayThreshold,
		Name:             g.Name,
		IsSuccessful:     m.Success,
		IsOff:            m.Disabled,
		Progression:      getProgression(m.Value, g.Value),
		Picto:            getMetricPicto(g.Type),
		Type:             g.Type,
	}, nil
}

func getValue(m *metric.AddSummaryMetricRequest, g *models.Goal, w *[]models.Workout) float64 {
	if g.Type == pb_goa.GoalType_MAIN_WORKOUT_DURATION.String() || g.Type == pb_goa.GoalType_EXTRA_WORKOUT_DURATION.String() {
		return calculateWorkoutDuration(w, g.Type)
	}

	return m.Value
}

func formatDisplayValue(value float64, g *models.Goal) (string, string) {
	if g.Type == enums.MainWorkoutDuration || g.Type == enums.ExtraWorkoutDuration {
		return convertMsToHour(value), convertMsToHour(g.Value)
	}

	switch g.Unit {
	case "L":

		return strconv.FormatFloat(value, 'f', 2, 64) + "L", strconv.FormatFloat(g.Value, 'f', 2, 64) + "L"

	default:
		return strconv.Itoa(int(value)), strconv.Itoa(int(g.Value))
	}
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
