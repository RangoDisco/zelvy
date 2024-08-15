package factories

import (
	"github.com/rangodisco/zelby/server/enums"
	"github.com/rangodisco/zelby/server/types"
)

func CreateSummaryViewModel() types.SummaryViewModel {
	return types.SummaryViewModel{
		ID:       "id",
		Date:     "2024-01-01",
		Steps:    1000,
		Metrics:  CreateMetrics(),
		Workouts: CreateWorkout(),
		Winner:   CreateWinner(),
	}
}

func CreateMetrics() []types.MetricViewModel {
	return []types.MetricViewModel{
		{
			Name:             "Calories brulées",
			Value:            1500,
			DisplayValue:     "1500",
			Threshold:        1500,
			DisplayThreshold: "1500",
			Success:          true,
			IsOff:            false,
			Progression:      100,
			Picto:            "picto",
		}, {
			Name:             "Calories consommées",
			Value:            3600,
			DisplayValue:     "1h00",
			Threshold:        7200,
			DisplayThreshold: "2h00",
			Success:          false,
			IsOff:            false,
			Progression:      50,
			Picto:            "picto",
		},
	}
}

func CreateWorkout() []types.WorkoutViewModel {
	return []types.WorkoutViewModel{
		{
			ID:           "id",
			KcalBurned:   320,
			ActivityType: enums.WorkoutTypeStrength,
			Name:         "Push 1",
			Duration:     "1h00",
			Picto:        "picto",
		},
		{
			ID:           "id",
			KcalBurned:   300,
			ActivityType: enums.WorkoutTypeWalking,
			Name:         "Marche",
			Duration:     "1h00",
			Picto:        "picto",
		},
	}
}

func CreateWinner() types.Winner {
	return types.Winner{
		DiscordID: "123456789",
	}
}
