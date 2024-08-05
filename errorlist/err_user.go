package errorlist

import "errors"

var (
	UserConflict   = errors.New("User already exists")
	SignUpfail     = errors.New("SignUp failed")
	UserNotFound   = errors.New("User not found")
	UserUpdateFail = errors.New("User update failed")
	UserNotUpdated = errors.New("User not updated")
)
