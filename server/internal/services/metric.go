package services

import (
	"github.com/rangodisco/zelvy/gen/zelvy/metric"
	"github.com/rangodisco/zelvy/server/config/database"
	"github.com/rangodisco/zelvy/server/internal/enums"
	"github.com/rangodisco/zelvy/server/internal/models"
	"slices"

	"github.com/google/uuid"
)

// FindAllActiveGoals queries the database to find all goals that aren't deleted
func FindAllActiveGoals() (*[]models.Goal, error) {
	// Fetch linked goal
	var goals []models.Goal
	err := database.GetDB().Where("active = ?", true).Find(&goals).Error
	if err != nil {
		return nil, err
	}

	return &goals, nil
}

// ConvertToMetricModel Compares a metric input to its goal and converts it to a db model
func ConvertToMetricModel(sID uuid.UUID, g models.Goal, m *metric.AddSummaryMetricRequest, workouts []models.Workout) (models.Metric, error) {
	value := getValue(m, &g, &workouts)
	isOff := isOff(g.ID)
	success := isAchieved(value, g.Value, g.Comparison, isOff)

	return models.Metric{
		ID:        uuid.New(),
		SummaryID: sID,
		Type:      g.Type,
		Value:     value,
		GoalID:    g.ID,
		Success:   success,
		Disabled:  isOff,
	}, nil
}

// GetMetricFromGoalID finds a metric in a slice by its goal ID
func GetMetricFromGoalID(g models.Goal, metrics []*metric.AddSummaryMetricRequest) *metric.AddSummaryMetricRequest {
	idx := slices.IndexFunc(metrics, func(m *metric.AddSummaryMetricRequest) bool {
		return m.Type.String() == g.Type
	})

	// Workout goals don't have a related metric
	if idx == -1 {
		return nil
	} else {
		return metrics[idx]
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
func getProgression(value float64, threshold float64) int64 {
	progression := int64(value / threshold * 100)
	if progression > 100 {
		progression = 100
	}
	return progression
}
