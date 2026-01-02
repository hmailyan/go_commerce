package database

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/hmailyan/go_ecommerce/models"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

func GetAllProducts() ([]models.Product, error) {
	if DB == nil {
		SetupGORM()
	}
	var products []models.Product
	if err := DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func SearchProductsByName(name string) ([]models.Product, error) {
	if DB == nil {
		SetupGORM()
	}
	var products []models.Product
	like := "%" + strings.ToLower(name) + "%"
	// Use LOWER(...) to be compatible with SQLite and Postgres for case-insensitive search
	if err := DB.Where("LOWER(product_name) LIKE ?", like).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func CreateProduct(p *models.Product) error {
	if DB == nil {
		SetupGORM()
	}
	if p.ID == "" {
		p.ID = uuid.NewString()
	}
	return DB.Create(p).Error
}
