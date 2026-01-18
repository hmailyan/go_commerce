package products

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"size:200;not null" json:"name"`
	Price     uint64    `json:"price"`
	Rating    uint8     `json:"rating"`
	Image     string    `json:"image"`
	MasterID  uuid.UUID `gorm:"type:uuid;index;" json:"master_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Master     *Product  `gorm:"foreignKey:MasterID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Variations []Product `gorm:"foreignKey:MasterID"`
}
