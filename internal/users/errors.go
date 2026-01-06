package users

import "errors"

var (
	ErrEmailAlreadyExists       = errors.New("email already exists")
	ErrCantFindUser             = errors.New("Cant find user")
	ErrInvalidToken             = errors.New("invalid token")
	ErrInvalidVerificationToken = errors.New("invalid verification token")
)
