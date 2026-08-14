package repository

import (
	"context"

	"github.com/LarsSeverson/charter/services/user/internal/domain/user"
)

type UserRepository interface {
	FindByID(ctx context.Context, id user.ID) (*user.User, error)
	Create(ctx context.Context, value *user.User) error
	Update(ctx context.Context, value *user.User) error
}
