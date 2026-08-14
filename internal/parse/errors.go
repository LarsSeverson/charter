package parse

import "errors"

var (
	ErrValueRequired = errors.New("value is required")
	ErrParseInt32    = errors.New("parse int32")
)
