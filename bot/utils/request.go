package utils

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
)

func Request(m string, e string, b interface{}) (*resty.Response, error) {
	baseUrl := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")

	var resp *resty.Response
	var err error

	client := resty.New()
	req := client.R().SetHeader("X-API-KEY", apiKey)

	if b != nil {
		req.SetBody(b)
	}

	switch m {
	case "GET":
		resp, err = req.Get(baseUrl + e)
	case "POST":
		resp, err = req.Post(baseUrl + e)
	default:
		return nil, fmt.Errorf("invalid method")
	}

	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	return resp, nil
}
