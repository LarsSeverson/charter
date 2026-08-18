package user

import (
	"fmt"

	userv1 "github.com/LarsSeverson/charter/gen/charter/user/v1"
	"github.com/LarsSeverson/charter/services/user/internal/application/command"
)

func decodeCreateUser(
	request *userv1.CreateUserRequest,
) command.CreateUserInput {
	return command.CreateUserInput{}
}

func encodeCreateUser(result command.CreateUserResult) (*userv1.CreateUserResponse, error) {
	status, err := encodeUserStatus(result.Status)
	if err != nil {
		return nil, err
	}

	user := &userv1.User{
		Id:     result.ID,
		Status: status,
	}

	return &userv1.CreateUserResponse{
		User: user,
	}, nil
}

func encodeUserStatus(status string) (userv1.UserStatus, error) {
	switch status {
	case "pending":
		return userv1.UserStatus_USER_STATUS_PENDING, nil
	case "active":
		return userv1.UserStatus_USER_STATUS_ACTIVE, nil
	case "suspended":
		return userv1.UserStatus_USER_STATUS_SUSPENDED, nil
	case "closed":
		return userv1.UserStatus_USER_STATUS_CLOSED, nil
	default:
		return 0, fmt.Errorf("invalid status: %s", status)
	}
}
