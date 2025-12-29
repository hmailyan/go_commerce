package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id" json:"_id"`
	First_Name      string             `bson:"first_name" json:"first_name" validate:"required,min=2,max=30"`
	Last_Name       string             `bson:"last_name" json:"last_name" validate:"required,min=2,max=30"`
	Password        string             `bson:"password" json:"password" validate:"required,min=6"`
	Email           string             `bson:"email" json:"email" validate:"email,required"`
	Phone           string             `bson:"phone" json:"phone" validate:"required"`
	Token           string             `bson:"token" json:"token"`
	Refreshtoken    string             `bson:"refresh_token" json:"refresh_token"`
	Created_At      time.Time          `bson:"created_at" json:"created_at"`
	Updated_At      time.Time          `bson:"updated_at" json:"updated_at"`
	UserCart        []Product          `bson:"usercart" json:"usercart"`
	User_ID         string             `bson:"user_id" json:"user_id"`
	Address_Details []Address          `bson:"address" json:"address"`
	Order_Status    []Order            `bson:"order_status" json:"order_status"`
}

type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id" json:"_id"`
	Product_Name string             `bson:"product_name" json:"product_name"`
	Price        uint64             `bson:"price" json:"price"`
	Rating       uint8              `bson:"rating" json:"rating"`
	Image        string             `bson:"image" json:"image"`
}
type ProductUser struct {
	Product_ID   primitive.ObjectID `bson:"_id" json:"_id"`
	Product_Name string             `bson:"product_name" json:"product_name"`
	Price        uint64             `bson:"price" json:"price"`
	Rating       uint8              `bson:"rating" json:"rating"`
	Image        string             `bson:"image" json:"image"`
	Quantity     int                `bson:"quantity" json:"quantity"`
}
type Address struct {
	Address_ID primitive.ObjectID `bson:"_id" json:"_id"`
	House      string             `bson:"house" json:"house"`
	Street     string             `bson:"street" json:"street"`
	City       string             `bson:"city" json:"city"`
	Pincode    string             `bson:"pincode" json:"pincode"`
}
type Order struct {
	Order_ID       primitive.ObjectID `bson:"_id" json:"_id"`
	Order_Cart     []ProductUser      `bson:"order_cart" json:"order_cart"`
	Ordered_At     time.Time          `bson:"ordered_at" json:"ordered_at"`
	Price          float64            `bson:"price" json:"price"`
	Discount       float64            `bson:"discount" json:"discount"`
	Payment_Method string             `bson:"payment_method" json:"payment_method"`
	Order_Status   string             `bson:"order_status" json:"order_status"`
}
type Payment struct {
	Digital string `bson:"digital" json:"digital"`
	COD     string `bson:"cod" json:"cod"`
}
