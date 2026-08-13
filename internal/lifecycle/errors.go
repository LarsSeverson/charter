package lifecycle

import "errors"

var (
	ErrAlreadyDeleted   = errors.New("entity is already deleted")
	ErrInvalidLifecycle = errors.New("invalid lifecycle")
)
