package users

import "errors"

var (
	ErrFindUser         = errors.New("find user")
	ErrCreateUser       = errors.New("create user")
	ErrUpdateUser       = errors.New("update user")
	ErrParseUserID      = errors.New("parse persisted user ID")
	ErrReconstituteUser = errors.New("reconstitute persisted user")
)
