package command

import (
	"context"
	"errors"

	"github.com/LarsSeverson/charter/services/user/internal/application/repository"
	"github.com/LarsSeverson/charter/services/user/internal/domain/user"
)

type CreateUser struct {
	users repository.UserRepository
}

func NewCreateUser(users repository.UserRepository) *CreateUser {
	return &CreateUser{users: users}
}

func (h *CreateUser) Handle(ctx context.Context) (*user.User, error) {
	id, err := user.NewID()
	if err != nil {
		return nil, errors.Join(ErrCreateUser, err)
	}

	value, err := user.New(id)
	if err != nil {
		return nil, errors.Join(ErrCreateUser, err)
	}

	if err := h.users.Create(ctx, value); err != nil {
		return nil, errors.Join(ErrCreateUser, err)
	}

	return value, nil
}
