// filepath: /home/hrant/Desktop/go_commerce/controllers/auth.go
package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"github.com/hmailyan/go_ecommerce/models"
)

var UserCollection *mongo.Collection

// SetUserCollection allows wiring the MongoDB collection from main
func SetUserCollection(c *mongo.Collection) {
	UserCollection = c
}

// SignUp registers a new user. Expects JSON matching models.User (password in plain text).
func SignUp(c *gin.Context) {
	if UserCollection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user collection not initialized"})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
		return
	}

	// basic required checks
	if input.Email == "" || input.Password == "" || input.First_Name == "" || input.Last_Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "first_name, last_name, email and password are required"})
		return
	}

	// check existing user by email
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	count, err := UserCollection.CountDocuments(ctx, bson.M{"email": input.Email})
	if err == nil && count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
		return
	}

	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	now := time.Now()
	newID := primitive.NewObjectID()
	input.ID = newID
	input.Password = string(hashed)
	input.Created_At = now
	input.Updated_At = now

	accessToken, refreshToken, err := generateUserTokens(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate tokens"})
		return
	}
	input.Token = accessToken
	input.Refreshtoken = refreshToken

	_, err = UserCollection.InsertOne(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user", "details": err.Error()})
		return
	}

	// do not return password
	input.Password = ""
	c.JSON(http.StatusCreated, gin.H{
		"message":       "user created",
		"user":          input,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// Login authenticates a user by email and password.
func Login(c *gin.Context) {
	if UserCollection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user collection not initialized"})
		return
	}

	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if creds.Email == "" || creds.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password required"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := UserCollection.FindOne(ctx, bson.M{"email": creds.Email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query user"})
		return
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// generate new tokens and update
	accessToken, refreshToken, err := generateUserTokens(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate tokens"})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"token":         accessToken,
			"refresh_token": refreshToken,
			"updated_at":    time.Now(),
		},
	}
	_, err = UserCollection.UpdateByID(ctx, user.ID, update)
	if err != nil {
		// not fatal for login, but report
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user tokens"})
		return
	}

	user.Password = ""
	user.Token = accessToken
	user.Refreshtoken = refreshToken

	c.JSON(http.StatusOK, gin.H{
		"message":       "login successful",
		"user":          user,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
