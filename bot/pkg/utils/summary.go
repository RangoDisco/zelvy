package utils

import (
	"context"
	"fmt"
	"github.com/rangodisco/zelvy/bot/pkg/config"
	"github.com/rangodisco/zelvy/bot/pkg/utils/grpc"
	pb_goa "github.com/rangodisco/zelvy/gen/zelvy/goal"
	pb_sum "github.com/rangodisco/zelvy/gen/zelvy/summary"
	"google.golang.org/grpc/metadata"
	"time"
)

// FetchSummary fetches today's summary from the API
func FetchSummary() (*pb_sum.GetSummaryResponse, error) {
	ctx, cancel := context.WithTimeout(metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"authorization": config.ApiKey})), 10*time.Second)
	defer cancel()
	resp, err := grpc.Client.GetSummary(ctx, &pb_sum.GetSummaryResquest{})
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
