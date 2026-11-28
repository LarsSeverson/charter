package users

import (
	userv1 "github.com/LarsSeverson/charter/gen/charter/user/v1"
	appuser "github.com/LarsSeverson/charter/services/user/internal/app/user"
)

type Handler struct {
	userv1.UnimplementedUserServiceServer

	users appuser.UseCases
}

func NewHandler(users appuser.UseCases) *Handler {
	handler := &Handler{
		users: users,
	}
	return handler
}

var _ userv1.UserServiceServer = (*Handler)(nil)
