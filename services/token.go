package tokens

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hmailyan/go_ecommerce/models"
)

func GenerateUserTokens(u models.User) (string, string, error) {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		secret = "secret" // fallback, replace in production
	}

	// access token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = u.ID.Hex()
	atClaims["email"] = u.Email
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", "", err
	}

	// refresh token
	rtClaims := jwt.MapClaims{}
	rtClaims["user_id"] = u.ID.Hex()
	rtClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshToken, err := rt.SignedString([]byte(secret))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
