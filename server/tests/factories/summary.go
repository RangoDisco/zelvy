package factories

import (
	"time"

	"server/internal/models"
	"server/pkg/types"

	"github.com/google/uuid"
)

func CreateSummaryModel() models.Summary {
	id := uuid.New()

	return models.Summary{
		ID:       id,
		Date:     time.Now(),
		Metrics:  CreateMetricModels(id),
		Workouts: CreateWorkoutModels(id),
		Winner:   CreateWinner(),
	}
}

func CreateSummaryViewModel() types.SummaryViewModel {
	return types.SummaryViewModel{
		ID:       "id",
		Date:     "2024-01-01",
		Steps:    1000,
		Metrics:  CreateMetricViewModels(),
		Workouts: CreateWorkoutViewModels(),
		Winner:   CreateWinnerViewModel(),
	}
}

func CreateSummaryInputModel() types.SummaryInputModel {
	return types.SummaryInputModel{
		Metrics:  CreateMetricInputModels(),
		Workouts: CreateWorkoutInputModels(),
	}

}

func CreateWinnerViewModel() types.Winner {
	return types.Winner{
		DiscordID: "123456789",
	}
}

func CreateWinner() models.User {
	return models.User{
		DiscordID:   "123456789",
		Username:    "Test User",
		PaypalEmail: "dummy@test.com",
		CreatedAt:   time.Now(),
	}
}
