package utils

import (
	"encoding/json"
	"fmt"
)

// TYPES
type WorkoutData struct {
	KcalBurned   int    `json:"kcalBurned"`
	ActivityType string `json:"activityType"`
	Name         string `json:"name"`
	Duration     string `json:"duration"`
}

type Metric struct {
	Name             string  `json:"name"`
	Value            float64 `json:"value"`
	DisplayValue     string  `json:"displayValue"`
	Threshold        float64 `json:"threshold"`
	DisplayThreshold string  `json:"displayThreshold"`
	Success          bool    `json:"success"`
	IsOff            bool    `json:"isOff"`
	Progression      int     `json:"difference"`
	Picto            string  `json:"picto"`
}

type Summary struct {
	ID       string        `json:"id"`
	Date     string        `json:"date"`
	Steps    int           `json:"steps"`
	Metrics  []Metric      `json:"metrics"`
	Workouts []WorkoutData `json:"workouts"`
	Winner   Winner        `json:"winner"`
}

type Winner struct {
	DiscordID string `json:"discordID"`
}

/**
 * Fetch today's  summary from the API
 */
func FetchSummary() (Summary, error) {

	resp, err := Request("GET", "/summaries", nil)
	if err != nil || resp.StatusCode() != 200 {
		return Summary{}, fmt.Errorf("error fetching summary: %d %v", resp.StatusCode(), err)
	}

	// Unmarshal response body to Summary struct
	var summary Summary
	if err := json.Unmarshal(resp.Body(), &summary); err != nil {
		return Summary{}, fmt.Errorf("error unmarshalling response body: %v", err)
	}

	return summary, nil
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
