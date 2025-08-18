package summary

import (
	"context"
	"github.com/rangodisco/zelvy/gen/zelvy/summary"
	"github.com/rangodisco/zelvy/server/internal/services"
	"google.golang.org/grpc"
)

type server struct {
	summary.UnimplementedSummaryServiceServer
}

func RegisterServer(s *grpc.Server) {
	summary.RegisterSummaryServiceServer(s, &server{})
}

func (s *server) GetSummary(_ context.Context, request *summary.GetSummaryResquest) (*summary.GetSummaryResponse, error) {
	// Fetch summary
	sum, err := services.FetchSummaryByDate(request.Day)
	if err != nil {
		return &summary.GetSummaryResponse{}, err
	}

	// Format data to fit fields in the view
	res, err := services.CreateSummaryViewModel(&sum)
	if err != nil {
		return &summary.GetSummaryResponse{}, err
	}

	return res, nil
}
