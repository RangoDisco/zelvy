package types

type Chart struct {
	Labels   []string  `json:"labels"`
	Type     string    `json:"type"`
	Datasets []Dataset `json:"datasets"`
}

type Dataset struct {
	Label string `json:"label"`
	Data  []int  `json:"data"`
}
