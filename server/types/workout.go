package types

type WorkoutViewModel struct {
	ID           string `json:"id"`
	KcalBurned   int    `json:"kcalBurned"`
	ActivityType string `json:"activityType"`
	Name         string `json:"name"`
	Duration     string `json:"duration"`
	Picto        string `json:"picto"`
}

type WorkoutData struct {
	KcalBurned   int     `json:"kcalBurned"`
	ActivityType string  `json:"activityType"`
	Name         string  `json:"name"`
	Duration     float64 `json:"duration"`
}
