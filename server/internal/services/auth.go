package services

import (
	"context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"os"
	"strings"
)

// IsWhitelisted compares the user's IP to the trusted ips list
func IsWhitelisted(ctx *context.Context) bool {
	trustedIps := strings.Split(os.Getenv("TRUSTED_IPS"), ",")
	peer, ok := peer.FromContext(*ctx)

	if !ok {
		return false
	}

	peerIp := strings.Split(peer.Addr.String(), ":")[0]

	for _, trustedIp := range trustedIps {
		if peerIp == trustedIp {
			return true
		}
	}

	return false
}

// IsAuthorized checks for the request's authorization metadata
func IsAuthorized(ctx *context.Context) bool {
	meta, ok := metadata.FromIncomingContext(*ctx)
	if !ok {
		return false
	}
	reqKey := meta["authorization"]
	if len(reqKey) == 0 {
		return false
	}

	serverAPIKey := os.Getenv("API_KEY")
	return reqKey[0] == serverAPIKey
}
