package lookup

import (
	"context"

	"github.com/LarsSeverson/charter/services/user/internal/app/user/port"
	"github.com/LarsSeverson/charter/services/user/internal/domain/user"
)

type GetQuery struct {
	UserID user.ID
}

type Get struct {
	users port.Repository
}

func (uc Get) Execute(ctx context.Context, query GetQuery) (*user.User, error) {
	u, err := uc.users.FindByID(ctx, query.UserID)
	return u, err
}
