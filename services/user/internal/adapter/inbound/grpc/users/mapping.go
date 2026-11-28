package users

import (
	"fmt"

	userv1 "github.com/LarsSeverson/charter/gen/charter/user/v1"
)

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
