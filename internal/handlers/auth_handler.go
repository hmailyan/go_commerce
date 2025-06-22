package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"go_ecommerce/internal/db"
	"go_ecommerce/internal/helpers"
	"go_ecommerce/internal/models"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET")) // move to env later

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	json.NewDecoder(r.Body).Decode(&creds)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)

	user := models.User{Email: creds.Email, Password: string(hashedPassword)}
	result := db.DB.Create(&user)

	if result.Error != nil {
		http.Error(w, "User already exists or bad input", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	json.NewDecoder(r.Body).Decode(&creds)
	fmt.Printf("LOGIN")
	var user models.User
	if err := db.DB.Where("email = ?", creds.Email).First(&user).Error; err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	accessToken, err := helpers.GenerateToken(user.Email, 15*time.Minute)
	if err != nil {
		http.Error(w, "Token generation failed", http.StatusInternalServerError)
		return
	}

	refreshToken, err := helpers.GenerateToken(user.Email, 7*24*time.Hour)
	if err != nil {
		http.Error(w, "Token generation failed", http.StatusInternalServerError)
		return
	}

	resp := TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	json.NewEncoder(w).Encode(resp)
}
func RefreshToken(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("refresh")

	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(req.RefreshToken, claims, func(t *jwt.Token) (interface{}, error) {
		return helpers.JwtKey, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid or expired refresh token", http.StatusUnauthorized)
		return
	}

	newAccessToken, _ := helpers.GenerateToken(claims.Subject, 15*time.Minute)
	json.NewEncoder(w).Encode(map[string]string{"access_token": newAccessToken})
}
