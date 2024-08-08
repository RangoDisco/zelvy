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
	client.R().SetHeader("X-API-KEY", apiKey)

	if b != nil {
		client.R().SetBody(b)
	}

	switch m {
	case "GET":
		resp, err = client.R().Get(baseUrl + e)
	case "POST":
		resp, err = client.R().Post(baseUrl + e)
	default:
		return nil, fmt.Errorf("invalid method")
	}

	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	return resp, nil
}
