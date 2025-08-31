package user

import (
	"context"
	pb_usr "github.com/rangodisco/zelvy/gen/zelvy/user"
	"github.com/rangodisco/zelvy/server/internal/services"
	"google.golang.org/grpc"
)

type server struct {
	pb_usr.UnimplementedUserServiceServer
}

func RegisterServer(s *grpc.Server) {
	pb_usr.RegisterUserServiceServer(s, &server{})
}

func (s *server) AddUser(_ context.Context, req *pb_usr.AddUserRequest) (*pb_usr.AddUserResponse, error) {
	err := services.UpsertUser(req)
	if err != nil {
		return nil, err
	}

	return &pb_usr.AddUserResponse{Message: "Upsert successful"}, nil
}
