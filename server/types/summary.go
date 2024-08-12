package types

type SummaryViewModel struct {
	ID       string             `json:"id"`
	Date     string             `json:"date"`
	Steps    int                `json:"steps"`
	Metrics  []MetricViewModel  `json:"metrics"`
	Workouts []WorkoutViewModel `json:"workouts"`
	Winner   Winner             `json:"winner"`
}

type RequestBody struct {
	Metrics  []MetricData  `json:"metrics"`
	Workouts []WorkoutData `json:"workouts"`
}
