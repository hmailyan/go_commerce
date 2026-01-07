package products

import (
	"context"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, product *Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *Repository) List(ctx context.Context) ([]*Product, error) {
	var products []*Product
	if err := r.db.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
func (r *Repository) GetByID(ctx context.Context, id string) (*Product, error) {
	var product Product
	if err := r.db.WithContext(ctx).First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
func (r *Repository) SearchByQuery(ctx context.Context, query string) ([]*Product, error) {
	var products []*Product
	like := "%" + query + "%"
	if err := r.db.WithContext(ctx).Where("name ILIKE ?", like).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
