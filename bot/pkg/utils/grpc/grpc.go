package grpc

import (
	pb_sum "github.com/rangodisco/zelvy/gen/zelvy/summary"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	Client pb_sum.SummaryServiceClient
)

func SetupClient() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Cannot connect to Discord: %v", err)
	}

	Client = pb_sum.NewSummaryServiceClient(conn)
}
