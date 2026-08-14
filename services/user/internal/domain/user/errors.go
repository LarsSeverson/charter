package user

import "errors"

var (
	ErrInvalidID           = errors.New("invalid user ID")
	ErrInvalidStatus       = errors.New("invalid user status")
	ErrInvalidStatusChange = errors.New("invalid status change")
)
