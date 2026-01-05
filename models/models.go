package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductUser struct {
	ID           string `gorm:"type:uuid;primaryKey" json:"id"`
	UserID       string `gorm:"type:uuid;index" json:"user_id"`
	OrderID      string `gorm:"type:uuid;index" json:"order_id"`
	ProductID    string `gorm:"type:uuid;index" json:"product_id"`
	Product_Name string `json:"product_name"`
	Price        uint64 `json:"price"`
	Rating       uint8  `json:"rating"`
	Image        string `json:"image"`
	Quantity     int    `json:"quantity"`
}

func (pu *ProductUser) BeforeCreate(tx *gorm.DB) (err error) {
	if pu.ID == "" {
		pu.ID = uuid.NewString()
	}
	return
}

type Address struct {
	ID         string `gorm:"type:uuid;primaryKey" json:"id"`
	Address_ID string `gorm:"-:all" json:"address_id"`
	UserID     string `gorm:"type:uuid;index" json:"user_id"`
	House      string `json:"house"`
	Street     string `json:"street"`
	City       string `json:"city"`
	Pincode    string `json:"pincode"`
}

func (a *Address) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID == "" {
		a.ID = uuid.NewString()
	}
	a.Address_ID = a.ID
	return
}

type Order struct {
	ID             string        `gorm:"type:uuid;primaryKey" json:"id"`
	UserID         string        `gorm:"type:uuid;index" json:"user_id"`
	Order_Cart     []ProductUser `gorm:"foreignKey:OrderID;references:ID;constraint:OnDelete:CASCADE;" json:"order_cart"`
	Ordered_At     time.Time     `json:"ordered_at"`
	Price          int           `json:"price"`
	Discount       float64       `json:"discount"`
	Payment_Method string        `json:"payment_method"`
	Order_Status   string        `json:"order_status"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID == "" {
		o.ID = uuid.NewString()
	}
	return
}

type Payment struct {
	Digital string `json:"digital"`
	COD     string `json:"cod"`
}
