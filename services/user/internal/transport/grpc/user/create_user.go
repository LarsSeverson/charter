package user

import (
	"context"

	userv1 "github.com/LarsSeverson/charter/gen/grpc/user/v1"
	"github.com/LarsSeverson/charter/services/user/internal/transport/grpc/grpcerror"
)

func (s *Service) CreateUser(
	ctx context.Context,
	req *userv1.CreateUserRequest,
) (*userv1.CreateUserResponse, error) {
	// TOOD: input := GetSomeInput...

	result, err := s.createUser.Handle(ctx)
	if err != nil {
		return nil, grpcerror.Encode(err)
	}

	return encodeCreateUser(result), nil
}
