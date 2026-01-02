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

		if database.DB == nil {
			database.SetupGORM()
		}
		var count int64
		if err := database.DB.Model(&models.Address{}).Where("user_id = ?", uid).Count(&count).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count >= 2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "You can add maximum 2 addresses"})
			return
		}
		if err := database.DB.Create(&addr).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		if database.DB == nil {
			database.SetupGORM()
		}
		var addr models.Address
		if err := database.DB.Where("user_id = ?", uid).Order("id asc").Offset(0).Limit(1).First(&addr).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "home address not found"})
			return
		}
		if err := database.DB.Model(&addr).Updates(models.Address{House: editAddress.House, Street: editAddress.Street, City: editAddress.City, Pincode: editAddress.Pincode}).Error; err != nil {
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
		if database.DB == nil {
			database.SetupGORM()
		}
		var addr models.Address
		if err := database.DB.Where("user_id = ?", uid).Order("id asc").Offset(1).Limit(1).First(&addr).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "work address not found"})
			return
		}
		if err := database.DB.Model(&addr).Updates(models.Address{House: editAddress.House, Street: editAddress.Street, City: editAddress.City, Pincode: editAddress.Pincode}).Error; err != nil {
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
		if database.DB == nil {
			database.SetupGORM()
		}
		if err := database.DB.Where("user_id = ?", uid).Delete(&models.Address{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"message": "deleted"})
	}
}
