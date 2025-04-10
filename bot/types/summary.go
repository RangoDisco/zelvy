package types

type WorkoutData struct {
	KcalBurned   int    `json:"kcalBurned"`
	ActivityType string `json:"activityType"`
	Name         string `json:"name"`
	Duration     string `json:"duration"`
}

type Metric struct {
	Name             string  `json:"name"`
	Value            float64 `json:"value"`
	DisplayValue     string  `json:"displayValue"`
	Threshold        float64 `json:"threshold"`
	DisplayThreshold string  `json:"displayThreshold"`
	Success          bool    `json:"success"`
	IsOff            bool    `json:"isOff"`
	Progression      int     `json:"difference"`
	Picto            string  `json:"picto"`
}

type Summary struct {
	ID       string        `json:"id"`
	Date     string        `json:"date"`
	Steps    int           `json:"steps"`
	Metrics  []Metric      `json:"metrics"`
	Workouts []WorkoutData `json:"workouts"`
	Winner   Winner        `json:"winner"`
}

type Winner struct {
	DiscordID string `json:"discordID"`
}
