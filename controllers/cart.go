package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hmailyan/go_ecommerce/database"
)

type Application struct{}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		if database.DB == nil {
			c.JSON(http.StatusNotImplemented, gin.H{"error": "cart endpoints not implemented for Postgres yet"})
			return
		}
		productID := c.Query("id")
		if productID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing product id"})
			return
		}
		uid := c.GetString("uid")
		if uid == "" {
			uid = c.Query("UserID")
			if uid == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found in token or query"})
				return
			}
		}
		if err := database.AddProductToCartGORM(uid, productID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "product added to cart"})
	}
}

func (app *Application) RemoveFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		if database.DB == nil {
			c.JSON(http.StatusNotImplemented, gin.H{"error": "cart endpoints not implemented for Postgres yet"})
			return
		}
		productID := c.Query("id")
		if productID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing product id"})
			return
		}
		uid := c.GetString("uid")
		if uid == "" {
			uid = c.Query("UserID")
			if uid == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found in token or query"})
				return
			}
		}
		if err := database.RemoveProductFromCartGORM(uid, productID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "product removed from cart"})
	}
}

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		if database.DB == nil {
			c.JSON(http.StatusNotImplemented, gin.H{"error": "cart endpoints not implemented for Postgres yet"})
			return
		}
		uid := c.GetString("uid")
		if uid == "" {
			uid = c.Query("UserID")
			if uid == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found in token or query"})
				return
			}
		}
		orderID, err := database.BuyFromCartGORM(uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"order_id": orderID})
	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		if database.DB == nil {
			c.JSON(http.StatusNotImplemented, gin.H{"error": "cart endpoints not implemented for Postgres yet"})
			return
		}
		productID := c.Query("id")
		if productID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing product id"})
			return
		}
		qtyStr := c.Query("qty")
		if qtyStr == "" {
			qtyStr = "1"
		}
		qty, err := strconv.Atoi(qtyStr)
		if err != nil || qty <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid qty"})
			return
		}
		uid := c.GetString("uid")
		if uid == "" {
			uid = c.Query("UserID")
			if uid == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found in token or query"})
				return
			}
		}
		orderID, err := database.InstantBuyGORM(uid, productID, qty)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"order_id": orderID})
	}
}

func GetItemFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		if database.DB == nil {
			c.JSON(http.StatusNotImplemented, gin.H{"error": "cart endpoints not implemented for Postgres yet"})
			return
		}
		uid := c.GetString("uid")
		if uid == "" {
			uid = c.Query("UserID")
			if uid == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found in token or query"})
				return
			}
		}
		items, total, err := database.GetCartItemsGORM(uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"items": items, "total": total})
	}
}
