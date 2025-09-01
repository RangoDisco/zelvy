package utils

import (
	"context"
	"github.com/rangodisco/zelvy/server/internal/api/grpc/goal"
	"github.com/rangodisco/zelvy/server/internal/api/grpc/summary"
	"github.com/rangodisco/zelvy/server/internal/api/grpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
)

var (
	Conn *grpc.ClientConn
	Lis  *bufconn.Listener
)

func bufDialer(context.Context, string) (net.Conn, error) {
	return Lis.Dial()
}

func SetupGrpc() {
	bufferInsecure := 1024 * 1024
	Lis = bufconn.Listen(bufferInsecure)

	s := grpc.NewServer()
	summary.RegisterServer(s)
	user.RegisterServer(s)
	goal.RegisterServer(s)

	go func() {
		if err := s.Serve(Lis); err != nil {
			log.Fatal(err)
		}
	}()

	Conn, _ = grpc.NewClient("passthrough:///bufnet", grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
}
