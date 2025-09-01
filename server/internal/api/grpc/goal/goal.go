package goal

import (
	"context"
	pb_goa "github.com/rangodisco/zelvy/gen/zelvy/goal"
	"github.com/rangodisco/zelvy/server/internal/services"
	"google.golang.org/grpc"
)

type server struct {
	pb_goa.UnimplementedGoalServiceServer
}

func RegisterServer(s *grpc.Server) {
	pb_goa.RegisterGoalServiceServer(s, &server{})
}

func (s *server) DisableGoals(_ context.Context, request *pb_goa.DisableGoalsRequest) (*pb_goa.DisableGoalsResponse, error) {
	var success []pb_goa.GoalType
	var errors []pb_goa.GoalType

	for _, g := range request.Goals {
		err := services.DisableGoal(g)
		if err != nil {
			errors = append(errors, g)
			continue
		}
		success = append(success, g)
	}

	return &pb_goa.DisableGoalsResponse{DisabledGoals: success, ErrorGoals: errors}, nil
}
