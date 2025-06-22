package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte(os.Getenv("JWT_SECRET")) // move to env later

func GenerateToken(email string, expiry time.Duration) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   email,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
