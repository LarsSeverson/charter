package port

import (
	"context"

	"github.com/LarsSeverson/charter/services/user/internal/domain/user"
)

type Repository interface {
	FindByID(ctx context.Context, id user.ID) (*user.User, error)
	Create(ctx context.Context, u *user.User) error
	Update(ctx context.Context, u *user.User) error
}
