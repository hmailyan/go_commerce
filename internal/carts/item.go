package carts

import (
	"time"

	"github.com/google/uuid"
	"github.com/hmailyan/go_ecommerce/internal/products"
)

type CartItem struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CartID    uuid.UUID `gorm:"type:uuid;index;not null;uniqueIndex:idx_cart_product" json:"cart_id"`
	ProductID uuid.UUID `gorm:"type:uuid;index;not null;uniqueIndex:idx_cart_product" json:"product_id"`
	Quantity  int       `gorm:"not null;check:quantity > 0" json:"quantity"`

	Product products.Product `gorm:"foreignKey:ProductID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
