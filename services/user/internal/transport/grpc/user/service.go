package user

import (
	userv1 "github.com/LarsSeverson/charter/gen/grpc/user/v1"
	"github.com/LarsSeverson/charter/services/user/internal/application/command"
)

type Service struct {
	userv1.UnimplementedUserServiceServer

	createUser *command.CreateUser
}

func NewService(
	createUser *command.CreateUser,
) *Service {
	return &Service{
		createUser: createUser,
	}
}

var _ userv1.UserServiceServer = (*Service)(nil)
