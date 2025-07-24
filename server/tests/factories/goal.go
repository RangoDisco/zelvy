package factories

import "server/pkg/types"

func CreateGoalViewModels() []types.GoalViewModel {
	return []types.GoalViewModel{
		{
			Name:             "Calories brulées",
			Value:            1500,
			DisplayValue:     "1500",
			Threshold:        1500,
			DisplayThreshold: "1500",
			IsSuccessful:     true,
			IsOff:            false,
			Progression:      100,
			Picto:            "picto",
		}, {
			Name:             "Calories consommées",
			Value:            3600,
			DisplayValue:     "1h00",
			Threshold:        7200,
			DisplayThreshold: "2h00",
			IsSuccessful:     false,
			IsOff:            false,
			Progression:      50,
			Picto:            "picto",
		},
	}
}
