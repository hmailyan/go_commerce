package users

import (
	"time"

	"github.com/google/uuid"
	"github.com/hmailyan/go_ecommerce/internal/carts"
)

type User struct {
	ID                uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FirstName         string     `gorm:"size:30;not null" json:"first_name" validate:"required,min=2,max=30"`
	LastName          string     `gorm:"size:30;not null" json:"last_name" validate:"required,min=2,max=30"`
	Password          string     `gorm:"not null" json:"password" validate:"required,min=6"`
	Email             string     `gorm:"size:100;uniqueIndex;not null" json:"email" validate:"email,required"`
	Phone             string     `gorm:"size:30;;not null" json:"phone" validate:"required"`
	VerificationAt    *time.Time `gorm:"default:null"; json:"verification_at"`
	VerificationToken string     `json:"verification_token"`

	Cart carts.Cart `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`

	// Token        string    `json:"token"`
	// Refreshtoken string    `json:"refresh_token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
