package database

import (
	"log"

	// "github.com/hmailyan/go_ecommerce/internal/addresses"
	// "github.com/hmailyan/go_ecommerce/internal/orders"
	"github.com/hmailyan/go_ecommerce/internal/carts"
	"github.com/hmailyan/go_ecommerce/internal/products"
	"github.com/hmailyan/go_ecommerce/internal/users"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	log.Println("Running AutoMigrate...")

	err := db.AutoMigrate(
		&users.User{},
		&products.Product{},
		&carts.Cart{},
		&carts.CartItem{},
		// &addresses.Address{},
	)

	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
}
