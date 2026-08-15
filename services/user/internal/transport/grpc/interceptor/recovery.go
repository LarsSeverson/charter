package interceptor

import (
	"context"
	"log/slog"
	"runtime/debug"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Recovery(logger *slog.Logger) grpc.UnaryServerInterceptor {
	if logger == nil {
		logger = slog.Default()
	}

	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (response any, err error) {
		defer func() {
			recovered := recover()
			if recovered == nil {
				return
			}

			attributes := []slog.Attr{
				slog.String("rpc.system.name", "grpc"),
				slog.String("rpc.method", info.FullMethod),

				slog.Any("panic.value", recovered),
				slog.String("panic.stack", string(debug.Stack())),
			}

			logger.LogAttrs(
				ctx,
				slog.LevelError,
				"recovered from panic during gRPC request",
				attributes...,
			)

			response = nil
			err = status.Error(
				codes.Internal,
				"internal server error",
			)
		}()

		return handler(ctx, req)
	}
}
