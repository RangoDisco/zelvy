package factories

import (
	"github.com/rangodisco/zelby/server/enums"
	"github.com/rangodisco/zelby/server/types"
)

func CreateWorkoutViewModels() []types.WorkoutViewModel {
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

func CreateWorkoutInputModels() []types.WorkoutInputModel {
	return []types.WorkoutInputModel{
		{
			Name:         "Push 1",
			Duration:     3600,
			ActivityType: enums.WorkoutTypeStrength,
			KcalBurned:   320,
		},
		{
			Name:         "Marche",
			Duration:     3600,
			ActivityType: enums.WorkoutTypeWalking,
			KcalBurned:   300,
		},
	}
}
