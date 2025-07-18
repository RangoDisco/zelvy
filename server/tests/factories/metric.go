package factories

import (
	"server/internal/enums"
	"server/internal/models"
	"server/pkg/types"

	"github.com/google/uuid"
)

func CreateMetricModels(surmmaryId uuid.UUID) []models.Metric {
	return []models.Metric{
		{
			ID:        uuid.New(),
			Type:      enums.KcalBurned,
			Value:     1091.9,
			SummaryID: surmmaryId,
		},
		{
			ID:        uuid.New(),
			Type:      enums.KcalConsumed,
			Value:     2083,
			SummaryID: surmmaryId,
		},
		{
			ID:        uuid.New(),
			Type:      enums.MilliliterDrank,
			Value:     2200,
			SummaryID: surmmaryId,
		},
	}
}

func CreateMetricInputModels() []types.MetricInputModel {
	return []types.MetricInputModel{
		{
			Value: 1091.9,
			Type:  enums.KcalBurned,
		},
		{
			Value: 2083,
			Type:  enums.KcalConsumed,
		},
		{
			Value: 2200,
			Type:  enums.MilliliterDrank,
		},
	}
}
