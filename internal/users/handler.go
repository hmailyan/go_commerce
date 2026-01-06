package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) SignUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		var req SignUpRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := h.service.SignUp(c.Request.Context(), req)

		if err != nil {
			if err == ErrEmailAlreadyExists {
				c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

	}
}

func (h *Handler) VerifyEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")

		if token == "" {
			fmt.Printf("test")
			fmt.Printf(token)
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidVerificationToken})
			return
		}

		err := h.service.VerifyEmail(c.Request.Context(), token)
		if err != nil {
			fmt.Printf("test2")

			if err == ErrInvalidVerificationToken {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "email verified successfully"})
	}
}

func (h *Handler) Me() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientToken := c.GetHeader("Authorization")

		if ClientToken == "" {
			c.JSON(http.StatusUnauthorized, ErrInvalidToken)
			return
		}

		out, err := h.service.Me(c.Request.Context(), ClientToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		c.JSON(http.StatusOK, out)
	}
}
