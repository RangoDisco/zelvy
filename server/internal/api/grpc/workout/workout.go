package workout

import (
	"context"

	pb_wkr "github.com/rangodisco/zelvy/gen/zelvy/workout"
	"github.com/rangodisco/zelvy/server/internal/services"
	"google.golang.org/grpc"
)

type server struct {
	pb_wkr.UnimplementedWorkoutServiceServer
}

func RegisterServer(s *grpc.Server) {
	pb_wkr.RegisterWorkoutServiceServer(s, &server{})
}

func (s *server) GetWorkouts(_ context.Context, request *pb_wkr.GetWorkoutsRequest) (*pb_wkr.GetWorkoutsResponse, error) {
	wkrs, err := services.FetchWorkoutsByDateRange(request.StartDate, request.EndDate)
	if err != nil {
		return nil, err
	}
	var wkrViewModels []*pb_wkr.WorkoutViewModel

	for _, w := range wkrs {
		wkrViewModels = append(wkrViewModels, services.ConvertToWorkoutViewModel(&w))
	}

	return &pb_wkr.GetWorkoutsResponse{Workouts: wkrViewModels}, err
}
