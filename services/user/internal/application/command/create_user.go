package command

import (
	"context"
	"errors"

	"github.com/LarsSeverson/charter/services/user/internal/application/port"
	"github.com/LarsSeverson/charter/services/user/internal/domain/user"
)

type CreateUser struct {
	users port.UserRepository
}

type CreateUserInput struct{} // Populate when needed

type CreateUserResult struct {
	ID     string
	Status string
}

func NewCreateUser(users port.UserRepository) *CreateUser {
	return &CreateUser{users: users}
}

func (h *CreateUser) Handle(ctx context.Context, input CreateUserInput) (CreateUserResult, error) {
	id, err := user.NewID()
	if err != nil {
		return CreateUserResult{}, errors.Join(ErrCreateUser, err)
	}

	value, err := user.New(id)
	if err != nil {
		return CreateUserResult{}, errors.Join(ErrCreateUser, err)
	}

	if err := h.users.Create(ctx, value); err != nil {
		return CreateUserResult{}, errors.Join(ErrCreateUser, err)
	}

	return CreateUserResult{
		ID:     value.ID().String(),
		Status: value.Status().String(),
	}, nil
}
