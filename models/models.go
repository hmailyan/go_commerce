package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Use UUIDs generated in application code for primary keys
type User struct {
	ID string `gorm:"type:uuid;primaryKey" json:"id"`
	// legacy compatibility field (was used as hex of ObjectID)
	User_ID         string        `gorm:"-:all" json:"user_id"`
	First_Name      string        `gorm:"size:30;not null" json:"first_name" validate:"required,min=2,max=30"`
	Last_Name       string        `gorm:"size:30;not null" json:"last_name" validate:"required,min=2,max=30"`
	Password        string        `gorm:"not null" json:"password" validate:"required,min=6"`
	Email           string        `gorm:"size:100;uniqueIndex;not null" json:"email" validate:"email,required"`
	Phone           string        `gorm:"size:30;uniqueIndex;not null" json:"phone" validate:"required"`
	Token           string        `json:"token"`
	Refreshtoken    string        `json:"refresh_token"`
	Created_At      time.Time     `json:"created_at"`
	Updated_At      time.Time     `json:"updated_at"`
	UserCart        []ProductUser `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;" json:"usercart"`
	Address_Details []Address     `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;" json:"address"`
	Order_Status    []Order       `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;" json:"order_status"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	// keep legacy User_ID for compatibility with existing code
	u.User_ID = u.ID
	return
}

type Product struct {
	ID           string `gorm:"type:uuid;primaryKey" json:"id"`
	Product_ID   string `gorm:"-:all" json:"product_id"`
	Product_Name string `gorm:"size:200;not null" json:"product_name"`
	Price        uint64 `json:"price"`
	Rating       uint8  `json:"rating"`
	Image        string `json:"image"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		p.ID = uuid.NewString()
	}
	p.Product_ID = p.ID
	return
}

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
