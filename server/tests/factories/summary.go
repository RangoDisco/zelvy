package factories

import (
	"github.com/rangodisco/zelby/server/types"
)

func CreateSummaryViewModel() types.SummaryViewModel {
	return types.SummaryViewModel{
		ID:       "id",
		Date:     "2024-01-01",
		Steps:    1000,
		Metrics:  CreateMetricViewModels(),
		Workouts: CreateWorkoutViewModels(),
		Winner:   CreateWinner(),
	}
}

func CreateSummaryInputModel() types.SummaryInputModel {
	return types.SummaryInputModel{
		Metrics:  CreateMetricInputModels(),
		Workouts: CreateWorkoutInputModels(),
	}

}

func CreateWinner() types.Winner {
	return types.Winner{
		DiscordID: "123456789",
	}
}
