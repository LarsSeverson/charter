package lifecycle

import "errors"

var (
	ErrComponentStart = errors.New("hook start")
	ErrComponentStop  = errors.New("hook stop")
)
