package user

import "github.com/LarsSeverson/charter/internal/identifier"

type userIDKind struct{}

type UserID = identifier.ID[userIDKind]

func NewUserID() (UserID, error) {
	return identifier.New[userIDKind]()
}

func ParseUserID(input string) (UserID, error) {
	return identifier.Parse[userIDKind](input)
}
