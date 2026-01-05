package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordUtils struct{}

func NewPasswordUtils() PasswordUtils {
	return PasswordUtils{}
}
func (p PasswordUtils) HashPassword(password string) (string, error) {
	// implementation for hashing password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (p PasswordUtils) VerifyPassword(password, givenPassword string) (bool, string) {
	// implementation for checking password hash
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(password))
	valid := true
	msg := ""

	if err != nil {
		msg = "login or password is incorrect"
		valid = false
	}
	return valid, msg
}
