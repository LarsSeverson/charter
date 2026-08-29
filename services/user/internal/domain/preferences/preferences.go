package preferences

import "github.com/LarsSeverson/charter/services/user/internal/domain/user"

type UserPreferences struct {
	userID   user.ID
	language string
	theme    string
}
