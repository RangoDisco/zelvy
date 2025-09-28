package middlewares

import (
	"context"
	"github.com/rangodisco/zelvy/server/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	whitelisted := services.IsWhitelisted(&ctx)
	if !whitelisted {
		return nil, status.Error(codes.PermissionDenied, "BLACKLISTED")
	}

	isAuthorized := services.IsAuthorized(&ctx)
	if !isAuthorized {
		return nil, status.Error(codes.PermissionDenied, "UNAUTHORIZED")
	}

	return handler(ctx, req)
}
