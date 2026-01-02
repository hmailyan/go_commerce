package database

import (
	"testing"

	"github.com/hmailyan/go_ecommerce/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupInMemoryDBForProduct(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory sqlite: %v", err)
	}
	if err := db.AutoMigrate(&models.User{}, &models.Product{}, &models.ProductUser{}, &models.Address{}, &models.Order{}); err != nil {
		t.Fatalf("automigrate failed: %v", err)
	}
	return db
}

func TestProductCreateAndSearch(t *testing.T) {
	db := setupInMemoryDBForProduct(t)
	DB = db

	p := models.Product{Product_Name: "unique-product", Price: 250}
	if err := CreateProduct(&p); err != nil {
		t.Fatalf("CreateProduct failed: %v", err)
	}

	all, err := GetAllProducts()
	if err != nil {
		t.Fatalf("GetAllProducts failed: %v", err)
	}
	if len(all) == 0 {
		t.Fatalf("expected at least one product, got %d", len(all))
	}

	res, err := SearchProductsByName("unique")
	if err != nil {
		t.Fatalf("SearchProductsByName failed: %v", err)
	}
	if len(res) == 0 {
		t.Fatalf("expected search to find product, got %d", len(res))
	}
}
