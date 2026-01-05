package utils

import (
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

func (t TokenUtils) UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {

	// if database.DB == nil {
	// 	database.SetupGORM()
	// }

	// updates := map[string]interface{}{
	// 	"token":        signedToken,
	// 	"refreshtoken": signedRefreshToken,
	// 	"updated_at":   time.Now(),
	// }

	// if err := database.DB.Model(&models.User{}).Where("id = ?", userId).Updates(updates).Error; err != nil {
	// 	log.Printf("failed to update tokens: %v", err)
	// }

}
