package topic

import "errors"

var (
	ErrInvalidCode   = errors.New("invalid topic code")
	ErrInvalidStatus = errors.New("invalid topic status")
)
