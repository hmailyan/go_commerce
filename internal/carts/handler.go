package carts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hmailyan/go_ecommerce/internal/app/http/context"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) AddItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AddItemRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		userId, ok := context.GetUserID(c)

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Interval Server Error"})
		}

		err := h.service.AddItem(c, req, userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusAccepted, gin.H{"message": "Product added to cart"})
	}
}
