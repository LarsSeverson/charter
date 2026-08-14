package repository

import (
	"context"

	"github.com/LarsSeverson/charter/services/user/internal/domain/user"
)

type UserRepository interface {
	FindByID(ctx context.Context, id user.ID) (*user.User, error)
	Save(ctx context.Context, value *user.User) (*user.User, error)
	Delete(ctx context.Context, id user.ID) (*user.User, error)
}
