package grpcerror

// "Client" facing errors live here

import (
	"errors"

	"github.com/LarsSeverson/charter/services/user/internal/application/port"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Encode(err error) error {
	switch {
	case errors.Is(err, port.ErrUserNotFound):
		return status.Error(codes.NotFound, "user not found")
	default:
		return status.Error(codes.Internal, "internal error")
	}
}
