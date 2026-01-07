package users

import "errors"

var (
	ErrEmailAlreadyExists       = errors.New("email already exists")
	ErrCantFindUser             = errors.New("Cant find user")
	ErrInvalidToken             = errors.New("invalid token")
	ErrInvalidVerificationToken = errors.New("invalid verification token")
	ErrInvalidLogin             = errors.New("Email or Password are incorrect")
	ErrVerificationRequired     = errors.New("User not verifyed")
)
