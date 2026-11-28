package grpc

import (
	"google.golang.org/grpc"
)

type Config struct {
	Address           string
	Options           []grpc.ServerOption
	UnaryInterceptors []grpc.UnaryServerInterceptor
}
