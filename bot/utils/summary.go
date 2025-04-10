package utils

import (
	"encoding/json"
	"fmt"
	"github.com/rangodisco/zelvy/bot/types"
)

// FetchSummary fetches today's summary from the API
func FetchSummary() (types.Summary, error) {

	resp, err := Request("GET", "/summaries", nil)
	if err != nil || resp.StatusCode() != 200 {
		return types.Summary{}, fmt.Errorf("error fetching summary: %d %v", resp.StatusCode(), err)
	}

	// Unmarshal response body to Summary struct
	var summary types.Summary
	if err := json.Unmarshal(resp.Body(), &summary); err != nil {
		return types.Summary{}, fmt.Errorf("error unmarshalling response body: %v", err)
	}

	return summary, nil
}

// IsSuccess determines if the summary is successful based on each metric success
func IsSuccess(metrics []types.Metric) bool {
	// For each metric, check if it's a success
	for _, metric := range metrics {
		if !metric.Success {
			return false
		}
	}
	return true
}
