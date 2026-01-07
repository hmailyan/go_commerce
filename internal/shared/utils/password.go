package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordUtils struct{}

func NewPasswordUtils() PasswordUtils {
	return PasswordUtils{}
}
func (p PasswordUtils) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (p PasswordUtils) VerifyPassword(password, givenPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(givenPassword))
	if err != nil {
		return err
	}
	return nil
}
