package channel

import "errors"

var (
	ErrInvalidCode   = errors.New("invalid channel code")
	ErrInvalidStatus = errors.New("invalid channel status")
)
