package profile

import "github.com/LarsSeverson/charter/services/user/internal/domain/user"

type UserProfile struct {
	userID    user.ID
	name      string
	avatarKey string
	bio       string
}
