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

func (h *Handler) GetCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := context.GetUserID(c)

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Interval Server Error"})
		}

		cart, err := h.service.GetCart(c, userId)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"cart": ToCartResponse(cart)})
	}
}

func (h *Handler) RemoveItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RemoveItemRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userId, ok := context.GetUserID(c)

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Interval Server Error"})
			return
		}

		err := h.service.RemoveItem(c, req, userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"message": "Item successfully modifyed"})

	}
}
