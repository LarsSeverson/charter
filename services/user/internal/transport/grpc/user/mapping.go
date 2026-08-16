package user

import (
	userv1 "github.com/LarsSeverson/charter/gen/grpc/user/v1"
	"github.com/LarsSeverson/charter/services/user/internal/application/command"
)

func decodeCreateUser(
	request *userv1.CreateUserRequest,
) command.CreateUserInput {
	return command.CreateUserInput{}
}

func encodeCreateUser(
	result command.CreateUserResult,
) *userv1.CreateUserResponse {
	return &userv1.CreateUserResponse{
		Id:     result.ID,
		Status: result.Status,
	}
}
