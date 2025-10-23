package summary

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	pb_sum "github.com/rangodisco/zelvy/gen/zelvy/summary"
	"github.com/rangodisco/zelvy/server/config/database"
	"github.com/rangodisco/zelvy/server/internal/models"
	"github.com/rangodisco/zelvy/server/internal/services"
	"google.golang.org/grpc"
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
		return &pb_sum.GetSummaryResponse{}, errors.New(fmt.Sprint("error fetching summary for day: ", request.Day))
	}

	// Format data to fit fields in the view
	res, err := services.CreateSummaryViewModel(&sum)
	if err != nil {
		return &pb_sum.GetSummaryResponse{}, errors.New("error creating summary view model")
	}

	return res, nil
}

func (s *server) AddSummary(_ context.Context, request *pb_sum.AddSummaryRequest) (*pb_sum.AddSummaryResponse, error) {
	// Convert to models
	summary := models.Summary{
		ID:   uuid.New(),
		Date: time.Now(),
	}

	// Build and add workouts to the summary object
	for _, w := range request.Workouts {
		workout := services.ConvertToWorkoutModel(w, summary.ID)
		summary.Workouts = append(summary.Workouts, workout)
	}

	goals, err := services.FindAllActiveGoals()
	if err != nil {
		return &pb_sum.AddSummaryResponse{}, errors.New("unable to find actives goals")
	}

	// Build and add metrics to the summary object
	for _, g := range *goals {
		relatedMetric := services.GetMetricFromGoalID(g, request.Metrics)
		metric, err := services.ConvertToMetricModel(summary.ID, g, relatedMetric, summary.Workouts)
		if err != nil {
			continue
		}
		summary.Metrics = append(summary.Metrics, metric)
	}

	// Set summary's success prop if all goals are met
	success := true
	for _, m := range summary.Metrics {
		if m.Success == false {
			success = false
			break
		}
	}
	summary.Success = success

	// Pick winner
	w, err := services.PickWinner()
	if err != nil {
		return &pb_sum.AddSummaryResponse{}, errors.New("unable to pick winner")
	}
	summary.WinnerID = w

	// Save summary
	if err := database.GetDB().Create(&summary).Error; err != nil {
		return &pb_sum.AddSummaryResponse{}, errors.New("unable to add summary to database")
	}

	return &pb_sum.AddSummaryResponse{
		Message: "Summary saved successfully!",
	}, nil
}

func (s *server) GetSummaryHeatmap(_ context.Context, request *pb_sum.GetSummaryHeatmapRequest) (*pb_sum.GetSummaryHeatmapResponse, error) {
	items, err := services.FindHeatmapResults(request.StartDate, request.EndDate)

	if err != nil {
		return &pb_sum.GetSummaryHeatmapResponse{}, errors.New("unable to generate heatmap data")
	}

	return &pb_sum.GetSummaryHeatmapResponse{Items: items}, nil
}
