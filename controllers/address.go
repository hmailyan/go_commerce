package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hmailyan/go_ecommerce/database"
	"github.com/hmailyan/go_ecommerce/models"
)

func AddAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetString("uid")
		if uid == "" {
			uid = c.Query("id")
			if uid == "" {
				c.JSON(http.StatusNotFound, gin.H{"error": "User ID is required"})
				return
			}
		}

		// Validate UUID format
		if _, err := uuid.Parse(uid); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			return
		}

		var addr models.Address
		if err := c.BindJSON(&addr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		addr.UserID = uid

		if err := database.AddAddressGORM(uid, &addr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Successfully added address", "address": addr})
	}
}

func EditHomeAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetString("uid")
		if uid == "" {
			uid = c.Query("id")
			if uid == "" {
				c.JSON(http.StatusNotFound, gin.H{"error": "User ID is required"})
				return
			}
		}

		var editAddress models.Address
		if err := c.BindJSON(&editAddress); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := database.EditHomeAddressGORM(uid, editAddress); err != nil {
			if err == database.ErrAddressNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "home address not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"message": "updated"})
	}
}

func EditWorkAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetString("uid")
		if uid == "" {
			uid = c.Query("id")
			if uid == "" {
				c.JSON(http.StatusNotFound, gin.H{"error": "User ID is required"})
				return
			}
		}

		var editAddress models.Address
		if err := c.BindJSON(&editAddress); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := database.EditWorkAddressGORM(uid, editAddress); err != nil {
			if err == database.ErrAddressNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "work address not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"message": "updated"})
	}
}

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetString("uid")
		if uid == "" {
			uid = c.Query("id")
			if uid == "" {
				c.JSON(http.StatusNotFound, gin.H{"error": "User ID is required"})
				return
			}
		}

		if err := database.DeleteAddressesGORM(uid); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"message": "deleted"})
	}
}
