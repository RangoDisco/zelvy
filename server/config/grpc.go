package config

import (
	"github.com/rangodisco/zelvy/server/internal/api/grpc/summary"
	"github.com/rangodisco/zelvy/server/internal/api/grpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func SetupGRpc() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	s := grpc.NewServer()

	summary.RegisterServer(s)
	user.RegisterServer(s)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
