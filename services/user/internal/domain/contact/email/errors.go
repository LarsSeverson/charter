package email

import "errors"

var (
	ErrInvalidID           = errors.New("invalid email ID")
	ErrInvalidUserID       = errors.New("invalid user ID")
	ErrInvalidEmailAddress = errors.New("invalid email address")
	ErrInvalidEmailType    = errors.New("invalid email type")
)
