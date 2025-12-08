package config

import (
	"log"
	"net"

	"github.com/rangodisco/zelvy/server/internal/api/grpc/goal"
	"github.com/rangodisco/zelvy/server/internal/api/grpc/summary"
	"github.com/rangodisco/zelvy/server/internal/api/grpc/user"
	"github.com/rangodisco/zelvy/server/internal/api/grpc/workout"
	"github.com/rangodisco/zelvy/server/internal/api/middlewares"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetupGRpc(errChan chan<- error, stopChan <-chan struct{}) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(middlewares.AuthInterceptor))

	summary.RegisterServer(s)
	user.RegisterServer(s)
	goal.RegisterServer(s)
	workout.RegisterServer(s)

	reflection.Register(s)

	go func() {
		<-stopChan
		s.GracefulStop()
	}()
	go func() {
		if err := s.Serve(lis); err != nil {
			errChan <- err
		}
	}()
}
