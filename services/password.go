package services

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// implementation for hashing password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func VerifyPassword(password, givenPassword string) (bool, string) {
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
