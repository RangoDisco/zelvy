package types

type MetricResponse struct {
	Name         string `json:"name"`
	DisplayValue string `json:"displayValue"`
	Threshold    string `json:"threshold"`
	Success      bool   `json:"success"`
	IsOff        bool   `json:"isOff"`
}

type MetricData struct {
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}
