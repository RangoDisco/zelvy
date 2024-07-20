package types

type WorkoutResponse struct {
	ID           string `json:"id"`
	KcalBurned   int    `json:"kcalBurned"`
	ActivityType string `json:"activityType"`
	Name         string `json:"name"`
	Duration     string `json:"duration"`
}

type WorkoutData struct {
	KcalBurned   int     `json:"kcalBurned"`
	ActivityType string  `json:"activityType"`
	Name         string  `json:"name"`
	Duration     float64 `json:"duration"`
}
