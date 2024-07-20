package types

type SummaryResponse struct {
	ID       string            `json:"id"`
	Date     string            `json:"date"`
	Steps    int               `json:"steps"`
	Metrics  []MetricResponse  `json:"metrics"`
	Workouts []WorkoutResponse `json:"workouts"`
}

type RequestBody struct {
	Metrics  []MetricData  `json:"metrics"`
	Workouts []WorkoutData `json:"workouts"`
}
