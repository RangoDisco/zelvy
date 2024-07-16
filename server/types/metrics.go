package types

type Metric struct {
	Name         string `json:"name"`
	Value        int    `json:"value"`
	DisplayValue string `json:"displayValue"`
	Threshold    string `json:"threshold"`
	Success      bool   `json:"success"`
}

type MetricsResponse struct {
	ID       string            `json:"id"`
	Date     string            `json:"date"`
	Steps    int               `json:"steps"`
	Metrics  []Metric          `json:"metrics"`
	Workouts []WorkoutResponse `json:"workouts"`
}

type RequestBody struct {
	KcalBurned      int           `json:"kcalBurned"`
	Steps           int           `json:"steps"`
	Workouts        []WorkoutData `json:"workouts"`
	KcalConsumed    int           `json:"kcalConsumed"`
	MilliliterDrank int           `json:"milliliterDrank"`
}
