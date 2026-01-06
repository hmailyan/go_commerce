package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Uid string
	jwt.StandardClaims
}

type TokenUtils struct {
	SECRET_KEY string
}

func NewTokenUtils() TokenUtils {
	return TokenUtils{
		SECRET_KEY: os.Getenv("SECRET_KEY"),
	}
}

func (t TokenUtils) GenerateUserTokens(ID string) (signedToken string, err error) {
	claims := &SignedDetails{
		Uid: ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(t.SECRET_KEY))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (t TokenUtils) ValidateToken(signedToken string) (Uid string, msg error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(t.SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = errors.New("the token is invalid")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = errors.New("Token is expired")
		return
	}
	return claims.Uid, msg
}

func (t TokenUtils) GenerateRandomToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
