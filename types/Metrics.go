package types

import "github.com/rangodisco/zelby/server/models"

type WorkoutData struct {
	KcalBurned   int    `json:"kcalBurned"`
	ActivityType string `json:"activityType"`
	Name         string `json:"name"`
	Duration     int    `json:"duration"`
}

type RequestBody struct {
	KcalBurned   int           `json:"kcalBurned"`
	Steps        int           `json:"steps"`
	Workouts     []WorkoutData `json:"workouts"`
	KcalConsumed int           `json:"kcalConsumed"`
}

type Metric struct {
	Type      string `json:"type"`
	Value     int    `json:"value"`
	Threshold int    `json:"threshold"`
	Success   bool   `json:"success"`
}

type MetricsResponse struct {
	ID       string           `json:"id"`
	Date     string           `json:"date"`
	Steps    int              `json:"steps"`
	Metrics  []Metric         `json:"metrics"`
	Workouts []models.Workout `json:"workouts"`
}
