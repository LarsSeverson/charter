package user

import (
	"context"

	userv1 "github.com/LarsSeverson/charter/gen/charter/user/v1"
	"github.com/LarsSeverson/charter/services/user/internal/transport/grpc/grpcerror"
)

func (s *Service) CreateUser(
	ctx context.Context,
	req *userv1.CreateUserRequest,
) (*userv1.CreateUserResponse, error) {
	input := decodeCreateUser(req)

	result, err := s.createUser.Handle(ctx, input)
	if err != nil {
		return nil, grpcerror.Encode(err)
	}

	output, err := encodeCreateUser(result)
	if err != nil {
		return nil, grpcerror.Encode(err)
	}

	return output, nil
}
