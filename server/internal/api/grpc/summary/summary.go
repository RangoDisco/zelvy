package summary

import (
	"context"
	"github.com/google/uuid"
	pb_sum "github.com/rangodisco/zelvy/gen/zelvy/summary"
	"github.com/rangodisco/zelvy/server/config/database"
	"github.com/rangodisco/zelvy/server/internal/models"
	"github.com/rangodisco/zelvy/server/internal/services"
	"google.golang.org/grpc"
	"time"
)

type server struct {
	pb_sum.UnimplementedSummaryServiceServer
}

func RegisterServer(s *grpc.Server) {
	pb_sum.RegisterSummaryServiceServer(s, &server{})
}

func (s *server) GetSummary(_ context.Context, request *pb_sum.GetSummaryRequest) (*pb_sum.GetSummaryResponse, error) {
	// Fetch summary
	sum, err := services.FetchSummaryByDate(request.Day)
	if err != nil {
		return &pb_sum.GetSummaryResponse{}, err
	}

	// Format data to fit fields in the view
	res, err := services.CreateSummaryViewModel(&sum)
	if err != nil {
		return &pb_sum.GetSummaryResponse{}, err
	}

	return res, nil
}

func (s *server) AddSummary(_ context.Context, request *pb_sum.AddSummaryRequest) (*pb_sum.AddSummaryResponse, error) {
	// Convert to models
	summary := models.Summary{
		ID:   uuid.New(),
		Date: time.Now(),
	}

	// Build and add metrics to the summary object
	for _, m := range request.Metrics {
		mm, success := services.ConvertToMetricModel(m, summary.ID)
		if !success {
			continue
		}
		summary.Metrics = append(summary.Metrics, mm)
	}

	// Build and add workouts to the summary object
	for _, w := range request.Workouts {
		workout := services.ConvertToWorkoutModel(w, summary.ID)
		summary.Workouts = append(summary.Workouts, workout)
	}

	// Pick winner
	w, err := services.PickWinner()
	if err != nil {
		return &pb_sum.AddSummaryResponse{}, err
	}
	summary.WinnerID = w

	// Save summary
	if err := database.GetDB().Create(&summary).Error; err != nil {
		return &pb_sum.AddSummaryResponse{}, err
	}

	return &pb_sum.AddSummaryResponse{
		Message: "Summary saved successfully!",
	}, nil
}
