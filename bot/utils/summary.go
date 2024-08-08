package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
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

/**
 * Pick a winner from the list of user ids
 * Happens every day, even if the goals are met
 */
func PickWinner(s *discordgo.Session, summary Summary) (*discordgo.User, error) {
	baseUrl := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")

	fmt.Printf("Picking winner for summary ID: %s\n", summary.ID)

	body := map[string]string{
		"summaryId": summary.ID,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %v", err)
	}

	req, err := http.NewRequest("POST", baseUrl+"/api/users/pick-winner", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Add api key and content type to headers
	req.Header.Add("X-API-KEY", apiKey)
	req.Header.Add("Content-Type", "application/json")

	// Create client & send request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer res.Body.Close()

	fmt.Printf("Response status: %s\n", res.StatusCode)

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Check if the response is JSON
	var response struct {
		Winner struct {
			DiscordID string `json:"discordID"`
		} `json:"winner"`
	}
	if err := json.Unmarshal(respBody, &response); err != nil {
		// If it's not JSON, assume the response is just the winner ID
		response.Winner.DiscordID = string(respBody)
	}

	if response.Winner.DiscordID == "" {
		return nil, fmt.Errorf("no winner ID received")
	}

	fmt.Printf("Winner ID: %s\n", response.Winner.DiscordID)

	winner, err := s.User(response.Winner.DiscordID)
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %v", err)
	}

	return winner, nil
}
