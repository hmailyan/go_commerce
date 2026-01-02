package database

import (
	"errors"

	"github.com/google/uuid"
	"github.com/hmailyan/go_ecommerce/models"
)

var (
	ErrAddressNotFound = errors.New("address not found")
)

func AddAddressGORM(userID string, addr *models.Address) error {
	if DB == nil {
		SetupGORM()
	}
	var count int64
	if err := DB.Model(&models.Address{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return err
	}
	if count >= 2 {
		return errors.New("maximum 2 addresses allowed")
	}
	if addr.ID == "" {
		addr.ID = uuid.NewString()
	}
	addr.UserID = userID
	return DB.Create(addr).Error
}

func EditHomeAddressGORM(userID string, edit models.Address) error {
	if DB == nil {
		SetupGORM()
	}
	var addr models.Address
	if err := DB.Where("user_id = ?", userID).Order("id asc").Offset(0).Limit(1).First(&addr).Error; err != nil {
		return ErrAddressNotFound
	}
	return DB.Model(&addr).Updates(models.Address{House: edit.House, Street: edit.Street, City: edit.City, Pincode: edit.Pincode}).Error
}

func EditWorkAddressGORM(userID string, edit models.Address) error {
	if DB == nil {
		SetupGORM()
	}
	var addr models.Address
	if err := DB.Where("user_id = ?", userID).Order("id asc").Offset(1).Limit(1).First(&addr).Error; err != nil {
		return ErrAddressNotFound
	}
	return DB.Model(&addr).Updates(models.Address{House: edit.House, Street: edit.Street, City: edit.City, Pincode: edit.Pincode}).Error
}

func DeleteAddressesGORM(userID string) error {
	if DB == nil {
		SetupGORM()
	}
	return DB.Where("user_id = ?", userID).Delete(&models.Address{}).Error
}
