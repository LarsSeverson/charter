package identifier

import (
	"go.jetify.com/typeid"
)

type ID[Kind any] struct {
	value typeid.AnyID
}

func New[Kind any]() (ID[Kind], error) {
	value, err := typeid.WithPrefix("")
	if err != nil {
		return ID[Kind]{}, err
	}

	return ID[Kind]{value: value}, nil
}

func NewWithPrefix[Kind any](prefix string) (ID[Kind], error) {
	value, err := typeid.WithPrefix(prefix)
	if err != nil {
		return ID[Kind]{}, err
	}

	return ID[Kind]{value: value}, nil
}

func Parse[Kind any](input string) (ID[Kind], error) {
	value, err := typeid.FromString(input)
	if err != nil {
		return ID[Kind]{}, err
	}

	return ID[Kind]{value: value}, nil
}

func (id ID[Kind]) Prefix() string {
	return id.value.Prefix()
}

func (id ID[Kind]) String() string {
	return id.value.String()
}
