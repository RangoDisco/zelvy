package config

import (
	pb_sum "github.com/rangodisco/zelvy/server/internal/api/grpc/summary"
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

	pb_sum.RegisterServer(s)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
