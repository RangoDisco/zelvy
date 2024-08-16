package factories

import (
	"github.com/rangodisco/zelby/server/enums"
	"github.com/rangodisco/zelby/server/types"
)

func CreateMetricViewModels() []types.MetricViewModel {
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
