package tests

import (
	"net/http"
	"os"
	"strings"
)

func SendRequest(method string, path string, body []byte) *http.Request {
	req, _ := http.NewRequest(method, path, strings.NewReader(string(body)))

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-API-KEY", os.Getenv("API_TEST_KEY"))

	return req

}
