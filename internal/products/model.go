package products

import (
	"github.com/google/uuid"
)

type Product struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name   string    `gorm:"size:200;not null" json:"product_name"`
	Price  uint64    `json:"price"`
	Rating uint8     `json:"rating"`
	Image  string    `json:"image"`
}
