package grpc

import (
	"context"

	userv1 "github.com/LarsSeverson/charter/gen/grpc/user/v1"
	"github.com/LarsSeverson/charter/services/user/internal/application/command"
)

type UserHandler struct {
	userv1.UnimplementedUserServiceServer

	createUser *command.CreateUser
}

func NewUserHandler(createUser *command.CreateUser) *UserHandler {
	return &UserHandler{createUser: createUser}
}

func (h *UserHandler) CreateUser(
	ctx context.Context,
	_ *userv1.CreateUserRequest,
) (*userv1.CreateUserResponse, error) {
	value, err := h.createUser.Handle(ctx)
	if err != nil {
		return nil, err // encodeError
	}

	return &userv1.CreateUserResponse{
		Id:     value.ID().String(),
		Status: value.Status().String(),
	}, nil
}

var _ userv1.UserServiceServer = (*UserHandler)(nil)
