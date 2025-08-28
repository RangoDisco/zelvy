package factories

import (
	"github.com/google/uuid"
	pb_goa "github.com/rangodisco/zelvy/gen/zelvy/goal"
	pb_met "github.com/rangodisco/zelvy/gen/zelvy/metric"
	"github.com/rangodisco/zelvy/server/internal/models"
)

func CreateMetricModels(surmmaryId uuid.UUID) []models.Metric {
	return []models.Metric{
		{
			ID:        uuid.New(),
			Type:      pb_goa.GoalType_KCAL_BURNED.String(),
			Value:     1091.9,
			SummaryID: surmmaryId,
		},
		{
			ID:        uuid.New(),
			Type:      pb_goa.GoalType_KCAL_CONSUMED.String(),
			Value:     2083,
			SummaryID: surmmaryId,
		},
		{
			ID:        uuid.New(),
			Type:      pb_goa.GoalType_MILLILITER_DRANK.String(),
			Value:     2200,
			SummaryID: surmmaryId,
		},
	}
}

func CreateMetricInputModels() []*pb_met.AddSummaryMetricRequest {
	return []*pb_met.AddSummaryMetricRequest{
		{
			Value: 1091.9,
			Type:  pb_goa.GoalType_KCAL_BURNED,
		},
		{
			Value: 2083,
			Type:  pb_goa.GoalType_KCAL_CONSUMED,
		},
		{
			Value: 2200,
			Type:  pb_goa.GoalType_MILLILITER_DRANK,
		},
	}
}
