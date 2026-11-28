package email

import (
	"net/mail"
	"strings"
)

type Address struct {
	value string
}

func ParseAddress(input string) (Address, error) {
	value := strings.TrimSpace(input)

	addr, err := mail.ParseAddress(value)
	if err != nil || addr.Address != value {
		return Address{}, ErrInvalidEmailAddress
	}

	return Address{value: addr.Address}, nil
}

func (a Address) String() string {
	return a.value
}

func (a Address) IsZero() bool {
	return a.value == ""
}
