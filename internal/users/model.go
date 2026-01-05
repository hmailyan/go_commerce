package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName     string    `gorm:"size:30;not null" json:"first_name" validate:"required,min=2,max=30"`
	LastName      string    `gorm:"size:30;not null" json:"last_name" validate:"required,min=2,max=30"`
	Password      string    `gorm:"not null" json:"password" validate:"required,min=6"`
	Email         string    `gorm:"size:100;uniqueIndex;not null" json:"email" validate:"email,required"`
	Phone         string    `gorm:"size:30;uniqueIndex;not null" json:"phone" validate:"required"`
	EmailVerified bool      `gorm:"default:false" json:"email_verified"`

	Token        string    `json:"token"`
	Refreshtoken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
