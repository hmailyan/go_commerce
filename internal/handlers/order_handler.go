package handlers

import (
	"encoding/json"
	"go_ecommerce/internal/db"
	"go_ecommerce/internal/models"
	"net/http"

	"gorm.io/gorm"
)

type CreateOrderRequest struct {
	Items []struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	} `json:"items"`
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value("userEmail")

	var orders []models.Order
	err := db.DB.Preload("OrderItems.Product").
		Where("user_email = ?", userEmail).
		Find(&orders).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		http.Error(w, "Error fetching orders", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	userEmail := r.Context().Value("userEmail")

	var total float64
	var orderItems []models.OrderItem

	for _, item := range req.Items {
		var product models.Product
		if err := db.DB.First(&product, item.ProductID).Error; err != nil {
			http.Error(w, "Product not found", http.StatusBadRequest)
			return
		}

		subtotal := product.Price * float64(item.Quantity)
		total += subtotal

		orderItems = append(orderItems, models.OrderItem{
			ProductID: product.ID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})
	}

	order := models.Order{
		UserEmail:  userEmail.(string),
		TotalPrice: total,
		OrderItems: orderItems,
	}

	if err := db.DB.Create(&order).Error; err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}
