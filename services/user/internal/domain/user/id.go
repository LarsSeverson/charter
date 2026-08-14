package user

import "github.com/LarsSeverson/charter/internal/identifier"

type userIDKind struct{}

type ID = identifier.ID[userIDKind]

func NewID() (ID, error) {
	return identifier.New[userIDKind]()
}

func ParseID(input string) (ID, error) {
	return identifier.Parse[userIDKind](input)
}
