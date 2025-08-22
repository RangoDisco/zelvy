package utils

import (
	"context"
	"fmt"
	pb_goa "github.com/rangodisco/zelvy/gen/zelvy/goal"
	pb_sum "github.com/rangodisco/zelvy/gen/zelvy/summary"
	"time"
)

// FetchSummary fetches today's summary from the API
func FetchSummary() (*pb_sum.GetSummaryResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := Client.GetSummary(ctx, &pb_sum.GetSummaryResquest{})
	if err != nil {
		return &pb_sum.GetSummaryResponse{}, fmt.Errorf("error fetching summary: %v", err)
	}
	return resp, nil
}

// IsSuccessful determines if the summary is successful based on each metric success
func IsSuccessful(goals []*pb_goa.GoalViewModel) bool {
	// For each metric, check if it's a success
	for _, g := range goals {
		if !g.IsSuccessful && !g.IsOff {
			return false
		}
	}
	return true
}
