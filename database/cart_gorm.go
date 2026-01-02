package database

import (
	"errors"
	"time"

	"github.com/hmailyan/go_ecommerce/models"
)

var (
	ErrNotFound = errors.New("not found")
)

// AddProductToCartGORM adds or increments a ProductUser for the given user and product
func AddProductToCartGORM(userID, productID string) error {
	if DB == nil {
		SetupGORM()
	}

	var pu models.ProductUser
	if err := DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&pu).Error; err == nil {
		pu.Quantity += 1
		pu.Price = pu.Price // keep same
		if err := DB.Save(&pu).Error; err != nil {
			return err
		}
		return nil
	}

	// otherwise create new ProductUser
	var prod models.Product
	if err := DB.First(&prod, "id = ?", productID).Error; err != nil {
		return ErrNotFound
	}

	newPU := models.ProductUser{
		UserID:       userID,
		ProductID:    prod.ID,
		Product_Name: prod.Product_Name,
		Price:        prod.Price,
		Rating:       prod.Rating,
		Image:        prod.Image,
		Quantity:     1,
	}
	if err := DB.Create(&newPU).Error; err != nil {
		return err
	}
	return nil
}

// RemoveProductFromCartGORM removes a product from a user's cart
func RemoveProductFromCartGORM(userID, productID string) error {
	if DB == nil {
		SetupGORM()
	}
	if err := DB.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.ProductUser{}).Error; err != nil {
		return err
	}
	return nil
}

// GetCartItemsGORM returns cart items and total price for a user
func GetCartItemsGORM(userID string) (items []models.ProductUser, total int64, err error) {
	if DB == nil {
		SetupGORM()
	}
	if err = DB.Where("user_id = ?", userID).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	var sum int64
	for _, it := range items {
		sum += int64(it.Price) * int64(it.Quantity)
	}
	return items, sum, nil
}

// BuyFromCartGORM creates an order from user's cart and clears cart
func BuyFromCartGORM(userID string) (orderID string, err error) {
	if DB == nil {
		SetupGORM()
	}
	items, total, err := GetCartItemsGORM(userID)
	if err != nil {
		return "", err
	}
	if len(items) == 0 {
		return "", errors.New("cart empty")
	}

	order := models.Order{
		UserID:         userID,
		Order_Cart:     items,
		Ordered_At:     time.Now(),
		Price:          int(total),
		Payment_Method: "cod",
		Order_Status:   "created",
	}
	if err := DB.Create(&order).Error; err != nil {
		return "", err
	}
	// create order items and remove cart items
	var ids []string
	for _, it := range items {
		ids = append(ids, it.ID)
		newPU := models.ProductUser{
			UserID:       "",
			OrderID:      order.ID,
			ProductID:    it.ProductID,
			Product_Name: it.Product_Name,
			Price:        it.Price,
			Rating:       it.Rating,
			Image:        it.Image,
			Quantity:     it.Quantity,
		}
		if err := DB.Create(&newPU).Error; err != nil {
			return "", err
		}
	}
	// delete original cart items
	if err := DB.Where("id IN ?", ids).Delete(&models.ProductUser{}).Error; err != nil {
		return order.ID, err
	}
	return order.ID, nil
}

// InstantBuyGORM buys a single product for a user
func InstantBuyGORM(userID, productID string, qty int) (orderID string, err error) {
	if DB == nil {
		SetupGORM()
	}
	var prod models.Product
	if err := DB.First(&prod, "id = ?", productID).Error; err != nil {
		return "", ErrNotFound
	}
	pu := models.ProductUser{
		UserID:       userID,
		ProductID:    prod.ID,
		Product_Name: prod.Product_Name,
		Price:        prod.Price,
		Rating:       prod.Rating,
		Image:        prod.Image,
		Quantity:     qty,
	}
	order := models.Order{
		UserID:         userID,
		Order_Cart:     []models.ProductUser{pu},
		Ordered_At:     time.Now(),
		Price:          int(prod.Price) * qty,
		Payment_Method: "cod",
		Order_Status:   "created",
	}
	if err := DB.Create(&order).Error; err != nil {
		return "", err
	}
	return order.ID, nil
}
