package user

// Presentation layer. Can be more composed than just this file

import (
	userv1 "github.com/LarsSeverson/charter/gen/grpc/user/v1"
	"github.com/LarsSeverson/charter/services/user/internal/domain/user"
)

func encodeCreateUser(
	// result command.CreateUserResult, // TODO: Need a command result
	result *user.User,
) *userv1.CreateUserResponse {
	return &userv1.CreateUserResponse{
		Id:     result.ID().String(),
		Status: result.Status().String(),
	}
}
