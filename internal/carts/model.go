package carts

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID  `gorm:"type:uuid;uniqueIndex;not null" json:"user_id"`
	Items     []CartItem `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
