package server

import (
	"time"

	"google.golang.org/grpc"
)

type Config struct {
	ShutdownTimeout   time.Duration
	Options           []grpc.ServerOption
	UnaryInterceptors []grpc.UnaryServerInterceptor
}
