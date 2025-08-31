package middlewares

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"os"
)

func AuthInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "UNAUTHORIZED")
	}
	reqKey := meta["authorization"]
	if len(reqKey) == 0 {
		return nil, status.Error(codes.Unauthenticated, "UNAUTHORIZED")
	}

	serverAPIKey := os.Getenv("API_KEY")
	if reqKey[0] != serverAPIKey {
		return nil, status.Error(codes.Unauthenticated, "UNAUTHORIZED")
	}

	return handler(ctx, req)
}
