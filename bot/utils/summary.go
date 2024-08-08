package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// TYPES
type WorkoutData struct {
	KcalBurned   int    `json:"kcalBurned"`
	ActivityType string `json:"activityType"`
	Name         string `json:"name"`
	Duration     string `json:"duration"`
}

type Metric struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Value        int    `json:"value"`
	DisplayValue string `json:"displayValue"`
	Threshold    string `json:"threshold"`
	Success      bool   `json:"success"`
}

type Summary struct {
	ID       string        `json:"id"`
	Date     string        `json:"date"`
	Steps    int           `json:"steps"`
	Metrics  []Metric      `json:"metrics"`
	Workouts []WorkoutData `json:"workouts"`
	Winner   User          `json:"winner"`
}

type User struct {
	ID        string `json:"id"`
	DiscordID string `json:"discordID"`
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

/**
 * Fetch today's  summary from the API
 */
func FetchSummary() Summary {
	baseUrl := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")

	req, err := http.NewRequest("GET", baseUrl+"/api/summaries/today", nil)
	checkErr(err)

	// Add api key to headers
	req.Header.Add("X-API-KEY", apiKey)

	// Create client & send request
	client := &http.Client{}
	res, err := client.Do(req)
	checkErr(err)

	// Ensure the response body is closed after reading
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading response body: %v", err)
	}

	// Unmarshal response body to Summary struct
	var summary Summary
	if err := json.Unmarshal(body, &summary); err != nil {
		log.Fatalf("error unmarshalling response body: %v", err)
	}

	return summary
}

/**
 * Determine if the summary is successful based on each metric success
 */
func IsSuccess(metrics []Metric) bool {
	// For each metric, check if it's a success
	for _, metric := range metrics {
		if !metric.Success {
			return false
		}
	}
	return true
}
