package database

import (
	"testing"
	"time"

	"github.com/hmailyan/go_ecommerce/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupInMemoryDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory sqlite: %v", err)
	}
	if err := db.AutoMigrate(&models.User{}, &models.Product{}, &models.ProductUser{}, &models.Address{}, &models.Order{}); err != nil {
		t.Fatalf("automigrate failed: %v", err)
	}
	return db
}

func TestAddRemoveAndBuyCartFlow(t *testing.T) {
	db := setupInMemoryDB(t)
	DB = db // point package DB to in-memory DB

	// create a product and a user
	p := models.Product{Product_Name: "test", Price: 100}
	if err := DB.Create(&p).Error; err != nil {
		t.Fatalf("create product failed: %v", err)
	}

	u := models.User{First_Name: "A", Last_Name: "B", Email: "x@example.com", Phone: "123", Password: "secret"}
	if err := DB.Create(&u).Error; err != nil {
		t.Fatalf("create user failed: %v", err)
	}

	// Add to cart
	if err := AddProductToCartGORM(u.ID, p.ID); err != nil {
		t.Fatalf("AddProductToCartGORM failed: %v", err)
	}

	items, total, err := GetCartItemsGORM(u.ID)
	if err != nil {
		t.Fatalf("GetCartItemsGORM failed: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("expected 1 cart item, got %d", len(items))
	}
	if total != int64(p.Price) {
		t.Fatalf("expected total %d, got %d", p.Price, total)
	}

	// Buy from cart
	orderID, err := BuyFromCartGORM(u.ID)
	if err != nil {
		t.Fatalf("BuyFromCartGORM failed: %v", err)
	}
	if orderID == "" {
		t.Fatalf("expected non-empty order id")
	}

	// Cart should be empty
	items, total, err = GetCartItemsGORM(u.ID)
	if err != nil {
		t.Fatalf("GetCartItemsGORM failed: %v", err)
	}
	if len(items) != 0 {
		t.Fatalf("expected 0 cart items after buy, got %d", len(items))
	}
	if total != 0 {
		t.Fatalf("expected total 0 after buy, got %d", total)
	}

	// Instant buy
	orderID2, err := InstantBuyGORM(u.ID, p.ID, 2)
	if err != nil {
		t.Fatalf("InstantBuyGORM failed: %v", err)
	}
	if orderID2 == "" {
		t.Fatalf("expected non-empty order id for instant buy")
	}

	// ensure order created with appropriate price
	var ord models.Order
	if err := DB.First(&ord, "id = ?", orderID2).Error; err != nil {
		t.Fatalf("failed to find order: %v", err)
	}
	if ord.Price != int(p.Price)*2 {
		t.Fatalf("expected order price %d, got %d", int(p.Price)*2, ord.Price)
	}

	// short wait to ensure any DB tasks settle
	time.Sleep(10 * time.Millisecond)
}
