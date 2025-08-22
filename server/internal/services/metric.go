package services

import (
	pb_met "github.com/rangodisco/zelvy/gen/zelvy/metric"
	"github.com/rangodisco/zelvy/server/config/database"
	"github.com/rangodisco/zelvy/server/internal/enums"
	"github.com/rangodisco/zelvy/server/internal/models"

	"github.com/google/uuid"
)

// ConvertToMetricModel Converts a metric input to a db model
func ConvertToMetricModel(m *pb_met.AddSummaryMetricRequest, summaryId uuid.UUID) (models.Metric, bool) {
	// Fetch linked goal
	var goal models.Goal
	if database.GetDB().Where("type = ?", m.Type).First(&goal).Error != nil {
		return models.Metric{}, false
	}
	return models.Metric{
		ID:        uuid.New(),
		SummaryID: summaryId,
		// TODO: handle enum
		Type:   string(m.Type),
		Value:  m.Value,
		GoalID: goal.ID,
	}, true
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
