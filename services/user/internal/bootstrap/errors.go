package bootstrap

import "errors"

var (
	ErrOpenDatabase   = errors.New("open database")
	ErrCreateListener = errors.New("create listener")
)
