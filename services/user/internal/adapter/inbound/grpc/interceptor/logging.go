package interceptor

import (
	"context"
	"log/slog"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type peerInfo struct {
	address string
	port    int
}

func Logging(logger *slog.Logger) grpc.UnaryServerInterceptor {
	if logger == nil {
		logger = slog.Default()
	}

	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		startedAt := time.Now()

		response, err := handler(ctx, req)

		code := status.Code(err)
		attributes := []slog.Attr{
			slog.String("rpc.system.name", "grpc"),
			slog.String("rpc.method", info.FullMethod),
			slog.String("rpc.response.status_code", code.String()),

			slog.Duration("rpc.duration", time.Since(startedAt)),
		}

		// Maybe have this be its own interceptor if there are other consumers down the line
		if peer, ok := peerInfoFromContext(ctx); ok {
			attributes = append(attributes, slog.String("network.peer.address", peer.address))
			if peer.port != 0 {
				attributes = append(attributes, slog.Int("network.peer.port", peer.port))
			}
		}

		logger.LogAttrs(
			ctx,
			levelForCode(code),
			"gRPC request completed",
			attributes...,
		)

		return response, err
	}
}

func levelForCode(code codes.Code) slog.Level {
	switch code {
	case codes.OK:
		return slog.LevelInfo

	case codes.Canceled:
		return slog.LevelDebug

	case codes.InvalidArgument,
		codes.NotFound,
		codes.AlreadyExists,
		codes.PermissionDenied,
		codes.Unauthenticated,
		codes.FailedPrecondition,
		codes.OutOfRange:
		return slog.LevelInfo

	case codes.DeadlineExceeded,
		codes.ResourceExhausted,
		codes.Aborted,
		codes.Unavailable:
		return slog.LevelWarn

	case codes.Unknown,
		codes.Unimplemented,
		codes.Internal,
		codes.DataLoss:
		return slog.LevelError

	default:
		return slog.LevelError
	}
}

func peerInfoFromContext(ctx context.Context) (peerInfo, bool) {
	remotePeer, ok := peer.FromContext(ctx)
	if !ok || remotePeer.Addr == nil {
		return peerInfo{}, false
	}

	addr := remotePeer.Addr.String()
	host, portText, err := net.SplitHostPort(addr)
	if err != nil {
		return peerInfo{
			address: addr,
		}, true
	}

	port, err := strconv.Atoi(portText)
	if err != nil {
		return peerInfo{
			address: host,
		}, true
	}

	return peerInfo{
		address: host,
		port:    port,
	}, true
}
