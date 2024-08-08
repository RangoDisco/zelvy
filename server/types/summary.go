package types

import "github.com/rangodisco/zelby/server/models"

type SummaryResponse struct {
	ID       string            `json:"id"`
	Date     string            `json:"date"`
	Steps    int               `json:"steps"`
	Metrics  []MetricResponse  `json:"metrics"`
	Workouts []WorkoutResponse `json:"workouts"`
	Winner   models.User       `json:"winner"`
}

type RequestBody struct {
	Metrics  []MetricData  `json:"metrics"`
	Workouts []WorkoutData `json:"workouts"`
}
