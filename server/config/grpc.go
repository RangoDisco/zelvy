package config

import (
	"github.com/rangodisco/zelvy/server/internal/api/grpc/summary"
	"github.com/rangodisco/zelvy/server/internal/api/grpc/user"
	"github.com/rangodisco/zelvy/server/internal/api/middlewares"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func SetupGRpc() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(middlewares.AuthInterceptor))

	summary.RegisterServer(s)
	user.RegisterServer(s)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
