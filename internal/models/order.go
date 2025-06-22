package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserEmail  string      `json:"user_email"`
	TotalPrice float64     `json:"total_price"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
