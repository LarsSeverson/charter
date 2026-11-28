package grpc

import "errors"

var (
	ErrListen    = errors.New("listen gRPC")
	ErrServe     = errors.New("serve gRPC")
	ErrForceStop = errors.New("force stop gRPC: graceful stop timed out")
)
