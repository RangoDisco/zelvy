package factories

import pb_goa "github.com/rangodisco/zelvy/gen/zelvy/goal"

func CreateGoalViewModels() []*pb_goa.GoalViewModel {
	return []*pb_goa.GoalViewModel{
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
		},
		{
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
