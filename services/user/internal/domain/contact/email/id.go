package email

import "github.com/LarsSeverson/charter/internal/identifier"

type emailIDKind struct{}

type ID = identifier.ID[emailIDKind]

func NewID() (ID, error) {
	return identifier.New[emailIDKind]()
}

func ParseID(input string) (ID, error) {
	return identifier.Parse[emailIDKind](input)
}
