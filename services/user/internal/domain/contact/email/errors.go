package email

import "errors"

var (
	ErrInvalidEmailAddress = errors.New("invalid email address")
	ErrInvalidEmailType    = errors.New("invalid email type")
)
