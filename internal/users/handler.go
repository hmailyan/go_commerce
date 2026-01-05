package users

import (
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

		out, err := h.service.SignUp(c.Request.Context(), req)

		if err != nil {
			if err == ErrEmailAlreadyExists {
				c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		c.JSON(http.StatusCreated, out)

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
