package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type RequestBody struct {
	Goals []string `json:"goals"`
}

/**
 * Send an array of goal type to be disabled for today
 */
func SetOffDay(goals []string) {
	baseUrl := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")

	var b = RequestBody{
		goals,
	}

	j, err := json.Marshal(b)
	checkErr(err)

	r, err := http.NewRequest("POST", baseUrl+"/api/offdays", bytes.NewBuffer(j))
	checkErr(err)

	r.Header.Add("X-API-KEY", apiKey)

	client := &http.Client{}
	resp, err := client.Do(r)
	checkErr(err)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		log.Fatalf("error sending goals to disable: %v", resp.Status)
	}
}
