package types

type MetricViewModel struct {
	Name             string  `json:"name"`
	Value            float64 `json:"value"`
	DisplayValue     string  `json:"displayValue"`
	Threshold        float64 `json:"threshold"`
	DisplayThreshold string  `json:"displayThreshold"`
	Success          bool    `json:"success"`
	IsOff            bool    `json:"isOff"`
	Progression      int     `json:"difference"`
}

type MetricData struct {
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}
