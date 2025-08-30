package utils

import (
	"context"
	pb_sum "github.com/rangodisco/zelvy/gen/zelvy/summary"
	pb_sum_s "github.com/rangodisco/zelvy/server/internal/api/grpc/summary"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
)

var (
	Conn   *grpc.ClientConn
	Client pb_sum.SummaryServiceClient
	Lis    *bufconn.Listener
)

func bufDialer(context.Context, string) (net.Conn, error) {
	return Lis.Dial()
}

func SetupGrpc() {
	bufferInsecure := 1024 * 1024
	Lis = bufconn.Listen(bufferInsecure)

	s := grpc.NewServer()
	pb_sum_s.RegisterServer(s)
	go func() {
		if err := s.Serve(Lis); err != nil {
			log.Fatal(err)
		}
	}()

	Conn, _ = grpc.NewClient("passthrough:///bufnet", grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
}
