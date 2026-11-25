package interceptor

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const requestIDMetadataKey = "x-request-id"

func requestIDFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	values := md.Get(requestIDMetadataKey)
	if len(values) == 0 {
		return ""
	}

	return values[0]
}
