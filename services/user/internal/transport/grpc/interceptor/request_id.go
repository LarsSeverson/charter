package interceptor

// incoming metadata is not auto copied to outgoing gRPC calls.
// when a service calls another service, you will want a client
// interceptor that reads the context value and adds x-request-id
// to outgoing metadata. this gives one correlation ID across the 
// complete service call chain

import (
	"context"
	"strings"

	"github.com/LarsSeverson/charter/internal/identifier"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	requestIDMetadataKey = "x-request-id"
	maxRequestIDLength   = 128
)

type requestIDKind struct{}
type requestIDContextKey struct{}

func RequestID() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		_ *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		requestID := requestIDFromMetadata(ctx)
		if requestID == "" {
			generatedId, err := identifier.NewWithPrefix[requestIDKind]("req")
			if err != nil {
				return nil, status.Error(
					codes.Internal,
					"failed to generate request ID",
				)
			}
			requestID = generatedId.String()
		}

		ctx = context.WithValue(
			ctx,
			requestIDContextKey{},
			requestID,
		)

		grpc.SetHeader(
			ctx,
			metadata.Pairs(requestIDMetadataKey, requestID),
		)

		return handler(ctx, req)
	}
}

func requestIDFromMetadata(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	values := md.Get(requestIDMetadataKey)
	if len(values) == 0 {
		return ""
	}

	requestID := strings.TrimSpace(values[0])
	if requestID == "" || len(requestID) > maxRequestIDLength {
		return ""
	}

	return requestID
}
