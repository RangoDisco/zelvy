package types

type Chart struct {
	Labels   []string  `json:"labels"`
	Type     string    `json:"type"`
	Datasets []Dataset `json:"datasets"`
}

type Dataset struct {
	Label           string `json:"label"`
	Data            []int  `json:"data"`
	BackgroundColor string `json:"backgroundColor"`
	BorderColor     string `json:"borderColor"`
}

type ChartColors struct {
	Background string `json:"background"`
	Border     string `json:"border"`
}
